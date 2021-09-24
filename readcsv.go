package tamboongo

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"os"
)

func DecryptRot(filename string) (b []byte, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}

	reader, err := NewRot128Reader(f)
	if err != nil {
		return
	}
	if reader == nil {
		err = errors.New("returned reader of NewRot128Reader is nil")
		return
	}

	fi, err := f.Stat()
	if err != nil {
		return
	}

	b = make([]byte, fi.Size())
	n, err := reader.Read(b)
	if err != nil {
		return
	}
	if int64(n) != fi.Size() {
		err = errors.New("not all file content is read")
		return
	}

	return
}

func ReadCsv(b []byte) (records [][]string, err error) {
	r := csv.NewReader(bytes.NewReader(b))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return records, err
		}

		records = append(records, record)
	}

	return
}
