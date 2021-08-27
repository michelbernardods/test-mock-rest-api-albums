package handler

import (
	"github.com/stretchr/testify/assert"
	"go-albums/apis"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAlbums(t *testing.T) {
	artist := &apis.MockHttpArtist{}
	artist.On("Get", URL).Return([]byte(`[{"title":"quidem molestiae enim"}]`), nil)
	artist.On("Get", URL).Return([]byte(`[{"id":0}]`), nil)

	req, _ := http.NewRequest("GET", "/albums", nil)
	rec := httptest.NewRecorder()
	get := GetAlbums(artist)
	get.ServeHTTP(rec, req)
	assert.Contains(t, rec.Body.String(), `"title":"quidem molestiae enim"`)
	assert.Contains(t, rec.Body.String(), `"id":0`)
	artist.AssertExpectations(t)
}