package reader

import (
	"errors"
	"io"
)

func ReadChunks(r io.Reader, chunkSize int) ([]byte, error) {
	if chunkSize < 1 {
		return nil, errors.New("chunkSize не может быть меньше 1")
	}
	buffer := make([]byte, chunkSize)
	var result []byte
	for {
		n, err := r.Read(buffer)

		result = append(result, buffer[:n]...)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, err
		}

	}

	return result, nil
}
