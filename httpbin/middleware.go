package httpbin

import (
	"bufio"
	"log/slog"
	"net"
	"net/http"
	"time"
)

func preflight(h http.Handler) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func limitRequestSize(maxSize int64, h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// headResponseWriter implements http.ResponseWriter in order to discard the
// body of the response
type headResponseWriter struct {
	*metaResponseWriter
}

func (hw *headResponseWriter) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// autohead automatically discards the body of responses to HEAD requests
		nil
}

func autohead(h http.Handler) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

// testMode enables additional safety checks to be enabled in the test suite.
var testMode = false

// metaResponseWriter implements http.ResponseWriter and http.Flusher in order
// to record a response's status code and body size for logging purposes.
type metaResponseWriter struct {
	w      http.ResponseWriter
	status int
	size   int64
}

func (mw *metaResponseWriter) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mw *metaResponseWriter) WriteHeader(s int) { _ = "STUB: not implemented"; return }

func (mw *metaResponseWriter) Flush() { _ = "STUB: not implemented"; return }

func (mw *metaResponseWriter) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (mw *metaResponseWriter) Status() int { _ = "STUB: not implemented"; return 0 }

func (mw *metaResponseWriter) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (mw *metaResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

func observe(o Observer, h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Result is the result of handling a request, used for instrumentation
type Result struct {
	Status    int
	Method    string
	URI       string
	Size      int64
	Duration  time.Duration
	UserAgent string
	ClientIP  string
}

// Observer is a function that will be called with the details of a handled
// request, which can be used for logging, instrumentation, etc
type Observer func(result Result)

// StdLogObserver creates an Observer that will log each request in structured
// format using the given stdlib logger
func StdLogObserver(l *slog.Logger) Observer { _ = "STUB: not implemented"; return *new(Observer) }
