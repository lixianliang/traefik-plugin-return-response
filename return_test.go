package traefik_plugin_return_response

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnResponse(t *testing.T) {
	cfg := CreateConfig()
	cfg.Code = 404
	cfg.Body = "{\"name\":\"lixianliang\"}"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "return-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	resp := recorder.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	//assert.Equal(t, cfg.Code, recorder.Code)
	if cfg.Code != recorder.Code {
		t.Errorf("%d != %d", cfg.Code, recorder.Code)
	}
	if cfg.Body != string(body) {
		t.Errorf("%s != %s", cfg.Body, body)
	}
}
