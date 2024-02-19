package fipe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kaiquegarcia/gofipe/fipe/entity"
)

func (f *client) doRequest(
	method string,
	path string,
	payload []byte,
	dest interface{},
) (audit entity.Audit, err error) {
	audit.Payload = payload
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s/%s", f.baseURL, path),
		bytes.NewReader(payload),
	)
	if err != nil {
		return
	}

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	audit.Response, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(audit.Response, dest)
	return
}
