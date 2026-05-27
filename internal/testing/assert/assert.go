// Package assert implements common assertions used in go-httbin's unit tests.
package assert

import (
	"net/http"
	"testing"
	"time"
)

// Equal asserts that two values are equal.
func Equal[T comparable](t *testing.T, got, want T, msg string, arg ...any) {
	_ = "STUB: not implemented"
	return
}

// DeepEqual asserts that two values are deeply equal.
func DeepEqual[T any](t *testing.T, got, want T, msg string, arg ...any) {
	_ = "STUB: not implemented"
	return
}

// NilError asserts that an error is nil.
func NilError(t *testing.T, err error) { _ = "STUB: not implemented"; return }

// Error asserts that an error is not nil.
func Error(t *testing.T, got, expected error) { _ = "STUB: not implemented"; return }

// MinDuration asserts that got >= min.
func MinDuration(t *testing.T, got time.Duration, min time.Duration) {
	_ = "STUB: not implemented"
	return
}

// StatusCode asserts that a response has a specific status code.
func StatusCode(t *testing.T, resp *http.Response, code int) { _ = "STUB: not implemented"; return }

// Ensure our error responses are never served as HTML, so that we do
// not need to worry about XSS or other attacks in error responses.

func isSafeContentType(ct string) bool { _ = "STUB: not implemented"; return false }

// Header asserts that a header key has a specific value in a response.
func Header(t *testing.T, resp *http.Response, key, want string) { _ = "STUB: not implemented"; return }

// ContentType asserts that a response has a specific Content-Type header
// value.
func ContentType(t *testing.T, resp *http.Response, contentType string) {
	_ = "STUB: not implemented"
	return
}

// Contains asserts that needle is found in the given string.
func Contains(t *testing.T, s string, needle string, description string) {
	_ = "STUB: not implemented"
	return
}

// BodyContains asserts that a response body contains a specific substring.
func BodyContains(t *testing.T, resp *http.Response, needle string) {
	_ = "STUB: not implemented"
	return
}

// BodyEquals asserts that a response body is equal to a specific string.
func BodyEquals(t *testing.T, resp *http.Response, want string) { _ = "STUB: not implemented"; return }

// BodySize asserts that a response body is a specific size.
func BodySize(t *testing.T, resp *http.Response, want int) { _ = "STUB: not implemented"; return }
