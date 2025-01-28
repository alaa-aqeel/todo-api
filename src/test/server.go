package test

import (
	"net/http/httptest"
	"testing"

	"github.com/alaa-aqeel/todo/src/api/http/router"
	"github.com/gavv/httpexpect/v2"
)

type TestServer struct {
	server *httptest.Server
	t      *testing.T
}

func NewServer(t *testing.T) *TestServer {
	return &TestServer{
		server: httptest.NewServer(router.NewChiRouter()),
		t:      t,
	}
}

func (ts *TestServer) Client() *httpexpect.Expect {
	return httpexpect.Default(ts.t, ts.server.URL)
}

func (ts *TestServer) Close() {
	ts.server.Close()
}
