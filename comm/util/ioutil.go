package util

import (
	"io"
	"net"
)

func WriteFull(writer io.Writer, data []byte) error {
	total := len(data)

	for pos := 0; pos < total; {
		n, err := writer.Write(data[pos:])
		if err != nil {
			return err
		}

		pos += n
	}

	return nil
}

func IsEOFOrNetReadError(err error) bool {
	if err == io.EOF {
		return true
	}
	ne, ok := err.(*net.OpError)
	return ok && ne.Op == "read"
}