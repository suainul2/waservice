package helpers

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func DownloadMediaFromURL(url string) (io.Reader, error) {

	client := http.DefaultClient
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	respByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(respByte)

	return r, nil
}
