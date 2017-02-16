package exiv2

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

// Exposer processes data from the output of the exiv2 command with options -P?gnycv.
// When calling exiv2 -P?gnycv [filepath] ? can be [X|E|I] for XMP, EXIF or IPTC.
type Exposer struct {
	in     io.Reader
	fields []*Field
	err    error
	sync.Once
}

// New is the default constructor for handling the output of exiv2 using a reader.
// Parsing occurs when Fields is called.
func New(r io.Reader) *Exposer {
	return &Exposer{in: r}
}

// Fields provides an array of exiv2 records as an array of *Field structs.
func (e *Exposer) Fields() ([]*Field, error) {
	// only parse once, save any errors
	e.Do(func() {
		out, err := ioutil.ReadAll(e.in)
		if err != nil {
			e.err = err
			return
		}
		if err = e.UnmarshalText(out); err != nil {
			e.err = err
			return
		}
		return
	})

	if e.err != nil {
		return nil, e.err
	}
	return e.fields, nil
}

// UnmarshalText parses the output into Fields.
func (e *Exposer) UnmarshalText(text []byte) error {
	s := bufio.NewScanner(bytes.NewReader(text))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()

		record, ok := e.parseLine(line)
		if !ok {
			continue
		}
		e.fields = append(e.fields, record)
	}
	return nil
}

func (e *Exposer) parseLine(s string) (*Field, bool) {
	i := strings.Fields(s)

	// ignore block data
	if len(i) < 4 {
		return nil, false
	}
	c, _ := strconv.Atoi(i[3])
	if c == 0 {
		return nil, false
	}
	record := &Field{
		Group: i[0],
		Name:  i[1],
		Type:  i[2],
		Count: c,
		Value: strings.Join(i[4:], " "),
	}
	if record.Type == "Ascii" {
		record.Count = len(record.Value)
	}
	return record, true
}
