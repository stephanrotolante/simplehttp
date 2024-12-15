package simplehttp

import (
	"net/http"
	"reflect"
	"testing"
)

func TestMethodAdded(t *testing.T) {

	httpRequest := CreateHttpRequest("")

	setMethodTests := []struct {
		TestDesc   string
		Wants      string
		MethodName string
	}{
		{
			TestDesc:   "set POST",
			Wants:      http.MethodPost,
			MethodName: "Post",
		},
		{
			TestDesc:   "set DELETE",
			Wants:      http.MethodDelete,
			MethodName: "Delete",
		},
		{
			TestDesc:   "set GET",
			Wants:      http.MethodGet,
			MethodName: "Get",
		},
		{
			TestDesc:   "set PUT",
			Wants:      http.MethodPut,
			MethodName: "Put",
		},
	}

	for _, test := range setMethodTests {
		t.Run(test.TestDesc, func(t *testing.T) {

			method := reflect.ValueOf(&httpRequest).MethodByName(test.MethodName)

			if !method.IsValid() {
				t.Errorf("method %s is not valid", test.MethodName)
				return
			}

			method.Call([]reflect.Value{})

			if httpRequest.method != test.Wants {
				t.Error("method not set properly")
				return
			}
		})
	}
}

func TestAddHeader(t *testing.T) {

	httpRequest := CreateHttpRequest("")

	httpRequest.AddHeader("some_header_name", "some_header_value")

	if httpRequest.request.Header.Get("some_header_name") != "some_header_value" {
		t.Error("header not added")
		return
	}
}

func TestUrl(t *testing.T) {

	httpRequest := CreateHttpRequest("https://google.com")

	if httpRequest.url != "https://google.com" {
		t.Error("request didn't initialize url correctly")
		return
	}

	httpRequest.Url("https://weather.com")

	if httpRequest.url != "https://weather.com" {
		t.Error("url not set properly")
		return
	}
}

func TestBody(t *testing.T) {

	httpRequest := CreateHttpRequest("")

	if httpRequest.body != nil {
		t.Error("request didn't initialize body correctly")
		return
	}

	httpRequest.Body([]byte("some_body"))

	buf := make([]byte, len("some_body"))

	httpRequest.body.Read(buf)

	if string(buf) != "some_body" {
		t.Error("body not set properly")
		return
	}
}
