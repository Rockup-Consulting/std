package wt

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type Headers map[string]string
type Form map[string]string
type JsonBody map[string]any
type Cookies map[string]string

type ReqData interface {
	createReq(t *testing.T, method string, path string) *http.Request
}

// Can only set Form or JsonBody, if both is set the request will panic
type BasicReqData struct {
	Headers Headers
	Cookies Cookies
}

func (b BasicReqData) createReq(
	t *testing.T,
	method string,
	path string,
) *http.Request {
	req := httptest.NewRequest(method, path, nil)

	for k, v := range b.Headers {
		req.Header.Add(k, v)
	}

	for k, v := range b.Cookies {
		req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}

	return req
}

type FormReqData struct {
	Headers Headers
	Cookies Cookies
	Form    Form
}

func (f FormReqData) createReq(
	t *testing.T,
	method string,
	path string,
) *http.Request {
	urlForm := url.Values{}

	for k, v := range f.Form {
		urlForm.Add(k, v)
	}

	req := httptest.NewRequest(method, path, strings.NewReader(urlForm.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	for k, v := range f.Headers {
		req.Header.Add(k, v)
	}

	for k, v := range f.Cookies {
		req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}

	return req
}

type JsonReqData struct {
	Headers  Headers
	Cookies  Cookies
	JsonBody JsonBody
}

func (j JsonReqData) createReq(
	t *testing.T,
	method string,
	path string,
) *http.Request {
	jsonBytes, err := json.Marshal(j.JsonBody)
	if err != nil {
		t.Fatalf("expected nil err, but got %s", err)
	}

	req := httptest.NewRequest(method, path, bytes.NewReader(jsonBytes))

	req.Header.Add("Content-Type", "application/json")

	for k, v := range j.Headers {
		req.Header.Add(k, v)
	}

	for k, v := range j.Cookies {
		req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}

	return req
}

func NewRequest(
	t *testing.T,
	method string,
	path string,
	data ReqData,
) *http.Request {
	t.Helper()

	var req *http.Request
	if data == nil {
		req = httptest.NewRequest(method, path, nil)
	} else {
		req = data.createReq(t, method, path)
	}

	return req
}

// GetCookie is a test helper that gets a cookie set on the Set-Cookie header from a
// httptest.ResponseRecorder. If the Cookie can't be found we panic.
func GetCookie(t *testing.T, res *http.Response, name string) *http.Cookie {
	t.Helper()

	for _, c := range res.Cookies() {
		if c.Name == name {
			return c
		}
	}

	t.Fatalf("cookie %q not found on ResponseRecorder", name)
	return nil
}
