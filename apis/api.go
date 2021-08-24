package apis

import (
	"io/ioutil"
	"net/http"
)

var Artist HttpInterface = &HttpArtist{}

type HttpInterface interface {
	Get(string) ([]byte, error)
}

type HttpArtist struct{}

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (a *HttpArtist) Get(url string) ([]byte, error) {
	return httpGet(url)
}

//mock Test
