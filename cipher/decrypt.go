package cipher

import (
	"fmt"
	"bytes"
	"io/ioutil"
)

func DecryptFile(filePath string) ([]byte, error) {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ioutil ReadFile err: ", err)
		return nil, err
	}
	reader, err := NewRot128Reader(bytes.NewBuffer(dat))
	buf := make([]byte, len(dat))
	_, err = reader.Read(buf)
	if err != nil {
		fmt.Println("rot128Reader err: ", err)
		return nil, err
	}
	return buf, nil
}