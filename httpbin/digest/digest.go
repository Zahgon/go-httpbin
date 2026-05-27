// Package digest provides a limited implementation of HTTP Digest
// Authentication, as defined in RFC 2617.
//
// Only the "auth" QOP directive is handled at this time, and while support for
// the SHA-256 algorithm is implemented here it does not actually work in
// either Chrome or Firefox.
//
// For more info, see:
// https://tools.ietf.org/html/rfc2617
// https://en.wikipedia.org/wiki/Digest_access_authentication
package digest

import (
	"net/http"
)

// digestAlgorithm is an algorithm used to hash digest payloads
type digestAlgorithm int

// Digest algorithms supported by this package
const (
	MD5 digestAlgorithm = iota
	SHA256
)

func (a digestAlgorithm) String() string { _ = "STUB: not implemented"; return "" }

// Check returns a bool indicating whether the request is correctly
// authenticated for the given username and password.
func Check(req *http.Request, username, password string) bool {
	_ = "STUB: not implemented"
	return false
}

// Challenge returns a WWW-Authenticate header value for the given realm and
// algorithm. If an invalid realm or an unsupported algorithm is given
func Challenge(realm string, algorithm digestAlgorithm) string {
	_ = "STUB: not implemented"
	return ""
}

// we use MD5 to hash nonces regardless of hash used for authentication

// sanitizeRealm tries to ensure that a given realm does not include any
// characters that will trip up our extremely simplistic header parser.
func sanitizeRealm(realm string) string { _ = "STUB: not implemented"; return "" }

// authorization is the result of parsing an Authorization header
type authorization struct {
	algorithm digestAlgorithm
	cnonce    string
	nc        string
	nonce     string
	opaque    string
	qop       string
	realm     string
	response  string
	uri       string
	username  string
}

// parseAuthorizationHeader parses an Authorization header into an
// Authorization struct, given a an authorization header like:
//
//	Authorization: Digest username="Mufasa",
//	                     realm="testrealm@host.com",
//	                     nonce="dcd98b7102dd2f0e8b11d0f600bfb0c093",
//	                     uri="/dir/index.html",
//	                     qop=auth,
//	                     nc=00000001,
//	                     cnonce="0a4f113b",
//	                     response="6629fae49393a05397450978507c4ef1",
//	                     opaque="5ccc069c403ebaf9f0171e9517f40e41"
//
// If the given value does not contain a Digest authorization header, or is in
// some other way malformed, nil is returned.
//
// Example from Wikipedia: https://en.wikipedia.org/wiki/Digest_access_authentication#Example_with_explanation
func parseAuthorizationHeader(value string) *authorization { _ = "STUB: not implemented"; return nil }

// parseDictHeader is a simplistic, buggy, and incomplete implementation of
// parsing key-value pairs from a header value into a map.
func parseDictHeader(value string) map[string]string { _ = "STUB: not implemented"; return nil }

// hash generates the hex digest of the given data using the given hashing
// algorithm, which must be one of MD5 or SHA256.
func hash(data []byte, algorithm digestAlgorithm) string { _ = "STUB: not implemented"; return "" }

// makeHA1 returns the HA1 hash, where
//
//	HA1 = H(A1) = H(username:realm:password)
//
// and H is one of MD5 or SHA256.
func makeHA1(realm, username, password string, algorithm digestAlgorithm) string {
	_ = "STUB: not implemented"
	return ""
}

// makeHA2 returns the HA2 hash, where
//
//	HA2 = H(A2) = H(method:digestURI)
//
// and H is one of MD5 or SHA256.
func makeHA2(auth *authorization, method, uri string) string { _ = "STUB: not implemented"; return "" }

// Response calculates the correct digest auth response. If the qop directive's
// value is "auth" or "auth-int" , then compute the response as
//
//	RESPONSE = H(HA1:nonce:nonceCount:clientNonce:qop:HA2)
//
// and if the qop directive is unspecified, then compute the response as
//
//	RESPONSE = H(HA1:nonce:HA2)
//
// where H is one of MD5 or SHA256.
func response(auth *authorization, password, method, uri string) string {
	_ = "STUB: not implemented"
	return ""
}

// compare is a constant-time string comparison
func compare(x, y string) bool { _ = "STUB: not implemented"; return false }
