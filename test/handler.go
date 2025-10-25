package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"pyra/pkg/server"
)

type Handler struct {
	http.Handler
}

func NewMux(method, path string, h http.Handler, logOut io.Writer) Handler {
	h = NewHandler(h, logOut)

	mux := http.NewServeMux()
	mux.Handle(method + " " + path, h)

	return Handler{mux}
}

type HandlerOption func(w *httptest.ResponseRecorder, r *http.Request)

func NewHandler(h http.Handler, logOut io.Writer) Handler {
	logger := NewLogger(logOut)

	h = server.Session(h)
	h = server.Logger(logger, h)

	return Handler{h}
}

func (m *Handler) Handle(method, path string, body io.Reader, opts ...HandlerOption) *http.Response {
	w, r := RR(method, path, body)

	for _, opt := range opts {
		opt(w, r)
	}

	m.ServeHTTP(w, r)

	return w.Result()
}

func SetFormValues(r *http.Request, values map[string]string) {
	r.Form = make(map[string][]string)

	for k, v := range values {
		r.Form.Set(k, v)
	}
}

// RR - returns a pair of ResponseRecorder and Request.
func RR(method string, target string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	return httptest.NewRecorder(), httptest.NewRequest(method, target, body)
}

func WithForm(params map[string]string) HandlerOption {
	return func(w *httptest.ResponseRecorder, r *http.Request) {
		SetFormValues(r, params)
	}
}

func WithPathValues(values map[string]string) HandlerOption {
	return func(w *httptest.ResponseRecorder, r *http.Request) {
		for k, v := range values {
			r.SetPathValue(k, v)
		}
	}
}
