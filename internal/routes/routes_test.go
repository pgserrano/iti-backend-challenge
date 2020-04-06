package routes

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	w.Header().Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	return w
}

func TestHandlerPasswordValidCases(t *testing.T) {

	router := SetupRouter()
	okCases := []string{ `{"password":"AbTp9!foo"}`, `{"password":"XyzkJ#123"}`, `{"password":"178Mn@fcb"}` , `{"password":"SwROTS3!%&*"}` }

	for _, v := range okCases {
		w := performRequest(router, "POST", "/users/passwords/validate", strings.NewReader(v))
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `{"IsValid":true,"Errs":null}`, w.Body.String())
	}

}

func TestHandlerPasswordInvalidCases(t *testing.T) {

	router := SetupRouter()
	nokCases := []string{`{"password":"A"}`, `{"password":"b"}` }
	for _, v := range nokCases {
		w := performRequest(router, "POST", "/users/passwords/validate", strings.NewReader(v))
		assert.Equal(t, http.StatusOK, w.Code)
		//assert.Equal(t, "false", w.Body.String())
	}

}


//TODO TestHandlerPasswordBadRequest