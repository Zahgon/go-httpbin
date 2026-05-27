package httpbin

import (
	"time"
)

// OptionFunc uses the "functional options" pattern to customize an HTTPBin
// instance
type OptionFunc func(*HTTPBin)

// WithDefaultParams sets the default params handlers will use
func WithDefaultParams(defaultParams DefaultParams) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// WithMaxBodySize sets the maximum amount of memory
func WithMaxBodySize(m int64) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithMaxDuration sets the maximum amount of time httpbin may take to respond
func WithMaxDuration(d time.Duration) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// WithHostname sets the hostname to return via the /hostname endpoint.
func WithHostname(s string) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithObserver sets the request observer callback
func WithObserver(o Observer) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithEnv sets the HTTPBIN_-prefixed environment variables reported
// by the /env endpoint.
func WithEnv(env map[string]string) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithExcludeHeaders sets the headers to exclude in outgoing responses, to
// prevent possible information leakage.
func WithExcludeHeaders(excludeHeaders string) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// WithPrefix sets the path prefix
func WithPrefix(p string) OptionFunc { _ = "STUB: not implemented"; return *new(OptionFunc) }

// WithAllowedRedirectDomains limits the domains to which the /redirect-to
// endpoint will redirect traffic.
func WithAllowedRedirectDomains(hosts []string) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// WithVersion sets the service name and build metadata to expose via /version.
func WithVersion(service, version, commit, buildDate, goVersion string) OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}

// WithUnsafeAllowDangerousResponses means endpoints that allow clients to
// specify a response Conntent-Type WILL NOT escape HTML entities in the
// response body, which can enable (e.g.) reflected XSS attacks.
//
// This configuration is only supported for backwards compatibility if
// absolutely necessary.
func WithUnsafeAllowDangerousResponses() OptionFunc {
	_ = "STUB: not implemented"
	return *new(OptionFunc)
}
