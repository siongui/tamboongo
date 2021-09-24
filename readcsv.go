package tamboongo

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type CsvRecord struct {
	Name           string
	AmountSubunits int64
	CCNumber       string
	CVV            string
	ExpMonth       time.Month
	ExpYear        int
}

func NewRecord(record []string) (r CsvRecord, err error) {
	amount, err := strconv.ParseInt(record[1], 10, 64)
	if err != nil {
		err = errors.New(strings.Join(record, " ") + " fail to convert donation amount")
		return
	}

	month, err := strconv.Atoi(record[4])
	if err != nil {
		err = errors.New(strings.Join(record, " ") + " fail to convert month")
		return
	}

	year, err := strconv.Atoi(record[5])
	if err != nil {
		err = errors.New(strings.Join(record, " ") + " fail to convert year")
		return
	}

	return CsvRecord{
		Name:           record[0],
		AmountSubunits: amount,
		CCNumber:       record[2],
		CVV:            record[3],
		ExpMonth:       time.Month(month),
		ExpYear:        year,
	}, err
}

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

func ReadCsv(b []byte) (csvRecords []CsvRecord, err error) {
	r := csv.NewReader(bytes.NewReader(b))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return csvRecords, err
		}

		if record[0] == "Name" {
			continue
		}

		csvRecord, err := NewRecord(record)
		if err != nil {
			return csvRecords, err
		}

		csvRecords = append(csvRecords, csvRecord)
	}

	return
}
