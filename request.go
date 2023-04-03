package httpreq

import (
	"io"
	"net/http"
	"strings"
)

type cfg struct {
	body   io.Reader
	header http.Header
}

var Options singleton

type option func(*cfg)

type singleton bool

func (singleton) Body(s string) option            { return func(c *cfg) { c.body = strings.NewReader(s) } }
func (singleton) Header(key, value string) option { return func(c *cfg) { c.header.Add(key, value) } }

func New(method, target string, options ...option) (*http.Request, error) {
	c := newConfig(options)
	request, err := http.NewRequest(method, target, c.body)
	if err != nil {
		return nil, err
	}
	transferHeaders(c, request)
	return request, nil
}
func newConfig(options []option) (c cfg) {
	c.header = make(http.Header)
	for _, callback := range options {
		callback(&c)
	}
	return c
}
func transferHeaders(c cfg, request *http.Request) {
	for key, values := range c.header {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}
}
