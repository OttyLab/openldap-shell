package db

import (
	"bufio"
	"encoding/json"
	"io"
)

type JsonDb struct {
	reader io.Reader
	writer io.Writer
}

func NewJsonDB(reader io.Reader, writer io.Writer) *JsonDb {
	ret := new(JsonDb)
	ret.reader = reader
	ret.writer = writer
	return ret
}

func (jsonDb *JsonDb) Write(dn string, entry Entry) error {
	err := jsonDb.write(dn, entry)
	return err
}

func (jsonDb *JsonDb) Read() (Entries, error) {
	entries, err := jsonDb.read()
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (jsonDb *JsonDb) read() (Entries, error) {
	scaner := bufio.NewScanner(jsonDb.reader)
	bs := make([]byte, 0)
	for scaner.Scan() {
		bs = append(bs, scaner.Bytes()...)
	}

	var entries Entries
	err := json.Unmarshal(bs, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (jsonDb *JsonDb) write(dn string, entry Entry) error {
	entries, err := jsonDb.read()
	entries[dn] = entry

	bs, err := json.Marshal(entries)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(jsonDb.writer)
	_, err = writer.Write(bs)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
