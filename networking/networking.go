package networking

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/juju/errors"
)

// SendSoap send soap message
func SendSoap(httpClient *http.Client, endpoint, message string) (*http.Response, error) {
	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return resp, errors.Annotate(err, "Post")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	return resp, nil
}
