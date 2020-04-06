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
	r.ServeHTTP(w, req)
	return w
}

func TestHandlerPasswordValidCases(t *testing.T) {

	router := SetupRouter()
	okCases := []string{"AbTp9!foo" , "XyzkJ#123", "178Mn@fcb", "SwROTS3!%&*" }

	for _, v := range okCases {
		w := performRequest(router, "POST", "/users/passwords/validations/isValid", strings.NewReader(v))
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `{"IsValid":true,"Errs":null}`, w.Body.String())
	}

}

func TestHandlerPasswordInvalidCases(t *testing.T) {

	router := SetupRouter()
	nokCases := []string{"A", "B" }
	for _, v := range nokCases {
		w := performRequest(router, "POST", "/users/passwords/validations/isValid", strings.NewReader(v))
		assert.Equal(t, http.StatusOK, w.Code)
		//assert.Equal(t, "false", w.Body.String())
	}

}