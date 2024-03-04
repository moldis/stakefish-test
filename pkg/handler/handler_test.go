package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestIPValidate(t *testing.T) {
	handler := NewHandler(nil)
	router := gin.Default()
	router.POST("/v1/validate", func(c *gin.Context) {
		handler.validate(c)
	})

	type want struct {
		code     int
		response string
	}

	type test struct {
		name string
		body string
		want want
	}

	tests := []test{
		{name: "success", body: `{"ip":"192.168.0.1"}`, want: want{code: 200, response: `{"status":true}`}},
		{name: "invalid_body", body: `{ip:"192.168.0.1"}`, want: want{code: 400, response: `{"message":"wrong params"}`}},
		{name: "wrong_ip_format", body: `{"ip":"1bla"}`, want: want{code: 200, response: `{"status":false}`}},
		{name: "valid_ip_2", body: `{"ip":"255.168.0.19"}`, want: want{code: 200, response: `{"status":true}`}},
		{name: "valid_ip_2", body: `{"ip":"259.168.0.19"}`, want: want{code: 200, response: `{"status":false}`}},
		{name: "invalid_ipv6", body: `{"ip":"2001:0db8:85a3:0000:0000:8a2e:0370:7334"}`, want: want{code: 200, response: `{"status":false}`}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/v1/validate", bytes.NewBuffer([]byte(tc.body)))
			router.ServeHTTP(w, req)
			require.Equal(t, w.Code, tc.want.code)
			require.Equal(t, w.Body.String(), tc.want.response)
		})
	}
}

func TestLookupHandler(t *testing.T) {
	handler := NewHandler(nil)
	router := gin.Default()
	router.GET("/v1/lookup", func(c *gin.Context) {
		handler.lookUp(c)
	})

	type want struct {
		code     int
		response string
	}

	type test struct {
		name  string
		query string
		want  want
	}

	tests := []test{
		{name: "invalid_domain", query: `?domain=blabla'`, want: want{code: 400, response: `{"message":"wrong domain"}`}},
		{name: "invalid_domain_2", query: `?domain=https://blabla.com`, want: want{code: 400, response: `{"message":"wrong domain"}`}},
		{name: "invalid_domain_2", query: `?domain=http://blabla.com`, want: want{code: 400, response: `{"message":"wrong domain"}`}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/v1/lookup"+tc.query, nil)
			router.ServeHTTP(w, req)
			require.Equal(t, w.Code, tc.want.code)
			require.Equal(t, w.Body.String(), tc.want.response)
		})
	}
}
