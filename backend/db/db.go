package db

type Attribute []string
type Entry map[string]Attribute
type Entries map[string]Entry

type Db interface {
	Write(dn string, entry Entry) error
	Read() (Entries, error)
}
