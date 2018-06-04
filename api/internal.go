package api

import (
	"io"
	"io/ioutil"
	"log"
)

func read(body io.ReadCloser) ([]byte, error) {
	log.Printf("Reading API response.\n")
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
