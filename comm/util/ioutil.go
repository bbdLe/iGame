package util

import "io"

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
