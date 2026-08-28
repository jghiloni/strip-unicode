package unistripper

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"regexp"
	"strconv"
)

var unicodeRE = regexp.MustCompile(`(?i)\\u[0-9a-f]{2,8}`)

func customScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0 : i+1], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

// StripUnicode takes an input stream, removes all unicode characters from it, and sends it to the given output stream.
// It returns an error if the input stream cannot be read or the output stream cannot be written to, or if the operation
// is cancelled
func StripUnicode(ctx context.Context, in io.Reader, out io.Writer) (err error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(customScanLines)
	for scanner.Scan() {
		// check for whether or not things have been cancelled for each line
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		line := strconv.QuoteToASCII(scanner.Text())
		line = unicodeRE.ReplaceAllString(line, "")
		if line, err = strconv.Unquote(line); err != nil {
			return err
		}
		if _, err = io.WriteString(out, line); err != nil {
			return err
		}
	}

	return scanner.Err()
}
