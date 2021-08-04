package httprouter

import (
	"net/http"
	"testing"
)

func TestHttpRouter_Register(t *testing.T) {
	var handler http.Handler

	Default.Register(handler, "test", "GET")

	tests := []struct {
		Input struct {
			Path   string
			Method string
		}

		Expected int
	}{
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "test",
				Method: "GET",
			},
			Expected: 0,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "test",
				Method: "PUT",
			},
			Expected: 405,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "test1",
				Method: "GET",
			},
			Expected: 404,
		},
	}

	for testNumber, test := range tests {
		_, status := Default.handler(test.Input.Method, test.Input.Path)
		if status != test.Expected {
			t.Errorf("Test %d :  %d was expected but got %d", testNumber, test.Expected, status)
		}
	}
}

func TestHttpRouter_DelegatePath(t *testing.T) {
	var handler http.Handler

	Default.DelegatePath(handler, "test", "GET")

	tests := []struct {
		Input struct {
			Path   string
			Method string
		}

		Expected int
	}{
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "test",
				Method: "GET",
			},
			Expected: 0,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "test",
				Method: "PUT",
			},
			Expected: 405,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "test1",
				Method: "GET",
			},
			Expected: 404,
		},
	}

	for testNumber, test := range tests {
		_, status := Default.handler(test.Input.Method, test.Input.Path)
		if status != test.Expected {
			t.Errorf("Test %d :  %d was expected but got %d", testNumber, test.Expected, status)
		}
	}
}
