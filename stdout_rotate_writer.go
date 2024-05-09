package filerotatewriter

import "fmt"

type StdoutWriterImpl struct {
}

func (s StdoutWriterImpl) Start() error {
	return nil
}

func (s StdoutWriterImpl) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func StdoutWriter() FileRotateWriter { // TODO add in io streams
	return &StdoutWriterImpl{}
}
