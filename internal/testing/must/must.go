// Package must implements helper functions for testing to eliminate some error
// checking boilerplate.
package must

import (
	"io"
	"net/http"
	"testing"
)

// DoReq makes an HTTP request and fails the test if there is an error.
func DoReq(t *testing.T, client *http.Client, req *http.Request) *http.Response {
	_ = "STUB: not implemented"
	return nil
}

// ReadAll reads all bytes from an io.Reader and fails the test if there is an
// error.
func ReadAll(t *testing.T, r io.Reader) string { _ = "STUB: not implemented"; return "" }

// Unmarshal unmarshals JSON from an io.Reader into a value and fails the test
// if there is an error.
func Unmarshal[T any](t *testing.T, r io.Reader) T { _ = "STUB: not implemented"; return *new(T) }
