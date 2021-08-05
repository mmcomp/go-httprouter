package httprouter

import (
	"errors"
	"net/http"
	"testing"
)

func TestHttpRouter_Register(t *testing.T) {
	var handler http.Handler

	Default.Register(handler, "/test", "GET")

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
				Path:   "/test",
				Method: "GET",
			},
			Expected: 0,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/test",
				Method: "PUT",
			},
			Expected: 405,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/test1",
				Method: "GET",
			},
			Expected: 404,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/tes",
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

func TestHttpRouter_RegisterLongPath(t *testing.T) {
	var handler http.Handler

	Default.Register(handler, "/apple/banana/cherry", "GET")

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
				Path:   "/apple/banana/cherry",
				Method: "GET",
			},
			Expected: 0,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/banana/cherry",
				Method: "PUT",
			},
			Expected: 405,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/banana",
				Method: "GET",
			},
			Expected: 404,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/banana/orange",
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

	Default.DelegatePath(handler, "/test", "GET")

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
				Path:   "/test",
				Method: "GET",
			},
			Expected: 0,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/test",
				Method: "PUT",
			},
			Expected: 405,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/tes",
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

func TestHttpRouter_DelegateLongPath(t *testing.T) {
	var handler http.Handler

	Default.DelegatePath(handler, "/apple/banana/cherry", "GET")

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
				Path:   "/apple/banana/cherry",
				Method: "GET",
			},
			Expected: 0,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/banana/cherry",
				Method: "PUT",
			},
			Expected: 405,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/banana",
				Method: "GET",
			},
			Expected: 404,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/banana/orange",
				Method: "GET",
			},
			Expected: 404,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/banana/cherry/dog",
				Method: "GET",
			},
			Expected: 0,
		},
	}

	for testNumber, test := range tests {
		_, status := Default.handler(test.Input.Method, test.Input.Path)
		if status != test.Expected {
			t.Errorf("Test %d :  %d was expected but got %d", testNumber, test.Expected, status)
		}
	}
}

func TestHttpRouter_MixedRegisterLongPath(t *testing.T) {
	var handler http.Handler

	Default.Register(handler, "/apple/banana/cherry", "GET")
	Default.DelegatePath(handler, "/animals", "GET")
	Default.DelegatePath(handler, "/apple/basket", "POST")

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
				Path:   "/apple/banana/cherry",
				Method: "GET",
			},
			Expected: 0,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/banana/cherry",
				Method: "PUT",
			},
			Expected: 405,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/banana",
				Method: "GET",
			},
			Expected: 404,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/banana/orange",
				Method: "GET",
			},
			Expected: 404,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/basket",
				Method: "GET",
			},
			Expected: 405,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/basket",
				Method: "POST",
			},
			Expected: 0,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/basket/banana",
				Method: "POST",
			},
			Expected: 0,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/apple/basket/banana/",
				Method: "POST",
			},
			Expected: 0,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/animals",
				Method: "GET",
			},
			Expected: 0,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/animals/dog",
				Method: "GET",
			},
			Expected: 0,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/animals/dog/cat",
				Method: "GET",
			},
			Expected: 0,
		},
	}

	for testNumber, test := range tests {
		_, status := Default.handler(test.Input.Method, test.Input.Path)
		if status != test.Expected {
			t.Errorf("Test %d :  %d was expected but got %d", testNumber, test.Expected, status)
		}
	}
}

func TestHttpRouter_Reregister(t *testing.T) {
	var handler http.Handler

	Default.Register(handler, "/test", "GET")

	tests := []struct {
		Input struct {
			Path   string
			Method string
		}

		Expected error
	}{
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/test",
				Method: "POST",
			},
			Expected: nil,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/test1",
				Method: "GET",
			},
			Expected: nil,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/test",
				Method: "GET",
			},
			Expected: errors.New("this path and method are already registered"),
		},
	}

	for testNumber, test := range tests {
		err := Default.Register(handler, test.Input.Path, test.Input.Method)
		if test.Expected != nil && (err == nil || err.Error() != test.Expected.Error()) {
			t.Errorf("Test %d :  %q was expected but got %q", testNumber, test.Expected.Error(), err.Error())
		}
	}
}

func TestHttpRouter_RedelegatePath(t *testing.T) {
	var handler http.Handler

	Default.DelegatePath(handler, "/test", "GET")

	tests := []struct {
		Input struct {
			Path   string
			Method string
		}

		Expected error
	}{
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/test",
				Method: "POST",
			},
			Expected: nil,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/test1",
				Method: "GET",
			},
			Expected: nil,
		},
		{
			Input: struct {
				Path   string
				Method string
			}{
				Path:   "/test",
				Method: "GET",
			},
			Expected: errors.New("this path and method are already registered"),
		},
	}

	for testNumber, test := range tests {
		err := Default.DelegatePath(handler, test.Input.Path, test.Input.Method)
		if test.Expected != nil && (err == nil || err.Error() != test.Expected.Error()) {
			t.Errorf("Test %d :  %q was expected but got %q", testNumber, test.Expected.Error(), err.Error())
		}
	}
}
