package io

import "io"

type multiReader struct {
	readers []io.Reader
}

type chanReader struct {
	p   []byte
	n   int
	err error
}

func (m *multiReader) Read(p []byte) (n int, err error) {
	cr := make(chan *chanReader)
	if len(m.readers) == 1 {
		if r, ok := m.readers[0].(*multiReader); ok {
			m.readers = r.readers
		}
	}

	for i := 0; i < len(m.readers); i++ {
		go func(cb chan<- *chanReader, reader io.Reader) {
			p1 := make([]byte, len(p))
			n, err := reader.Read(p1)
			cb <- &chanReader{
				p:   p1,
				n:   n,
				err: err,
			}
		}(cr, m.readers[i])
	}

	ecount := 0
ReadEnd:
	for {
		select {
		case reader := <-cr:
			if reader.err != io.EOF {
				return 0, reader.err
			} else {
				ecount++
			}
			if ecount == len(m.readers) {
				break ReadEnd
			}
		}
	}

	return 0, io.EOF

}

// MultiReader returns a Reader that's the logical concatenation of
// the provided input readers. They're read sequentially. Once all
// inputs have returned EOF, Read will return EOF.  If any of the readers
// return a non-nil, non-EOF error, Read will return that error.
func MultiReader(readers ...io.Reader) io.Reader {
	r := make([]io.Reader, len(readers))
	copy(r, readers)
	return &multiReader{r}
}
