package httpreq_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/mdwhatcott/httpreq"
	"github.com/mdwhatcott/testing/should"
)

func readAll(body io.Reader) string { raw, _ := io.ReadAll(body); return string(raw) }

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct {
	*should.T
}

func (this *Suite) TestAnAwkwardAPI() {
	request, err := http.NewRequest(http.MethodGet, "https://smarty.com", nil)
	this.So(err, should.BeNil)
	this.So(request, should.NOT.BeNil)
	this.So(request.Method, should.Equal, http.MethodGet)
	this.So(request.URL.String(), should.Equal, "https://smarty.com")
	this.So(request.Header, should.BeEmpty)
	this.So(request.Body, should.BeNil)

	request.Header.Add("X-Header", "Value")
	this.So(request.Header.Get("x-header"), should.Equal, "Value")
	request.Close = true
}
func (this *Suite) TestBadURL() {
	request, err := httpreq.New("asdf", "%%%%")
	this.So(err, should.NOT.BeNil)
	this.So(request, should.BeNil)
}
func (this *Suite) TestGET() {
	request, err := httpreq.New(http.MethodGet, "https://smarty.com")
	this.So(err, should.BeNil)
	this.So(request, should.NOT.BeNil)
	this.So(request.Method, should.Equal, http.MethodGet)
	this.So(request.URL.String(), should.Equal, "https://smarty.com")
	this.So(request.Header, should.BeEmpty)
	this.So(request.Body, should.BeNil)
}
func (this *Suite) TestPOST() {
	request, err := httpreq.New(http.MethodPost, "https://smarty.com",
		httpreq.Options.Body("body"),
		httpreq.Options.Header("X-Header", "Value"),
	)
	this.So(err, should.BeNil)
	this.So(request, should.NOT.BeNil)
	this.So(request.Method, should.Equal, http.MethodPost)
	this.So(request.URL.String(), should.Equal, "https://smarty.com")
	this.So(request.Header.Get("x-header"), should.Equal, "Value")
	this.So(request.Body, should.NOT.BeNil)
	this.So(readAll(request.Body), should.Equal, "body")
}
