package httputil

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func ScanBody(r *http.Request) ([]byte, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, err
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}
