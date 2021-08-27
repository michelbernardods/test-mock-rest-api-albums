package apis

import (
	"github.com/stretchr/testify/mock"
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

// MockHttpArtist Mock Test
type MockHttpArtist struct {
	mock.Mock
}

func (m *MockHttpArtist) Get(url string) ([]byte, error) {
	called := m.Called(url)
	return called.Get(0).([]byte), called.Error(1)
}