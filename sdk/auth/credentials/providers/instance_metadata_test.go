package providers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/stretchr/testify/assert"
)

func TestInstanceMetadataProvider_Retrieve_Success(t *testing.T) {

	// Start a test server locally.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body := "unsupported path: " + r.URL.Path
		status := 500

		switch r.URL.Path {
		case "/latest/meta-data/ram/security-credentials/":
			body = "ELK"
			status = 200
		case "/latest/meta-data/ram/security-credentials/ELK":
			body = ` {
			  "AccessKeyId" : "STS.L4aBSCSJVMuKg5U1vFDw",
			  "AccessKeySecret" : "wyLTSmsyPGP1ohvvw8xYgB29dlGI8KMiH2pKCNZ9",
			  "Expiration" : "2018-08-20T22:30:02Z",
			  "SecurityToken" : "CAESrAIIARKAAShQquMnLIlbvEcIxO6wCoqJufs8sWwieUxu45hS9AvKNEte8KRUWiJWJ6Y+YHAPgNwi7yfRecMFydL2uPOgBI7LDio0RkbYLmJfIxHM2nGBPdml7kYEOXmJp2aDhbvvwVYIyt/8iES/R6N208wQh0Pk2bu+/9dvalp6wOHF4gkFGhhTVFMuTDRhQlNDU0pWTXVLZzVVMXZGRHciBTQzMjc0KgVhbGljZTCpnJjwySk6BlJzYU1ENUJuCgExGmkKBUFsbG93Eh8KDEFjdGlvbkVxdWFscxIGQWN0aW9uGgcKBW9zczoqEj8KDlJlc291cmNlRXF1YWxzEghSZXNvdXJjZRojCiFhY3M6b3NzOio6NDMyNzQ6c2FtcGxlYm94L2FsaWNlLyo=",
			  "LastUpdated" : "2018-08-20T16:30:01Z",
			  "Code" : "Success"
			}`
			status = 200
		case "/latest/api/token":
			body = "token"
			status = 200
		}
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	// Update our securityCredURL and tokenURL to point at our local test server.
	originalSecurityCredURL := securityCredURL
	originalTokenURL := tokenURL
	securityCredURL = strings.Replace(securityCredURL, "http://100.100.100.200", ts.URL, -1)
	tokenURL = strings.Replace(tokenURL, "http://100.100.100.200", ts.URL, -1)
	defer func() {
		securityCredURL = originalSecurityCredURL
		tokenURL = originalTokenURL
	}()

	credential, err := NewInstanceMetadataProvider().Retrieve()
	assert.Nil(t, err)

	stsTokenCredential, ok := credential.(*credentials.StsTokenCredential)
	assert.True(t, ok)

	assert.Equal(t, "STS.L4aBSCSJVMuKg5U1vFDw", stsTokenCredential.AccessKeyId)
	assert.Equal(t, "wyLTSmsyPGP1ohvvw8xYgB29dlGI8KMiH2pKCNZ9", stsTokenCredential.AccessKeySecret)
	assert.True(t, strings.HasPrefix(stsTokenCredential.AccessKeyStsToken, "CAESrAIIARKAA"))
}

func TestInstanceMetadataProvider_Retrieve_Fail1(t *testing.T) {
	// Update our securityCredURL and tokenURL to point at our local test server.
	originalSecurityCredURL := securityCredURL
	originalTokenURL := tokenURL
	securityCredURL = strings.Replace(securityCredURL, "http://100.100.100.200", "http://invalid-domain-xxx", -1)
	tokenURL = strings.Replace(tokenURL, "http://100.100.100.200", "http://invalid-domain-xxx", -1)
	defer func() {
		securityCredURL = originalSecurityCredURL
		tokenURL = originalTokenURL
	}()

	_, err := NewInstanceMetadataProvider().Retrieve()
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "dial tcp: lookup invalid-domain-xxx"))
}

func TestInstanceMetadataProvider_Retrieve_Fail2(t *testing.T) {
	// Start a test server locally.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var body string
		var status int

		switch r.URL.Path {
		case "/latest/meta-data/ram/security-credentials/":
			body = "ELK"
			status = 400
		case "/latest/api/token":
			body = "not found"
			status = 404
		}
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	// Update our securityCredURL and tokenURL to point at our local test server.
	originalSecurityCredURL := securityCredURL
	originalTokenURL := tokenURL
	securityCredURL = strings.Replace(securityCredURL, "http://100.100.100.200", ts.URL, -1)
	tokenURL = strings.Replace(tokenURL, "http://100.100.100.200", ts.URL, -1)
	defer func() {
		securityCredURL = originalSecurityCredURL
		tokenURL = originalTokenURL
	}()

	_, err := NewInstanceMetadataProvider().Retrieve()
	assert.NotNil(t, err)
	assert.Equal(t, "received 400 getting role name: ELK", err.Error())
}

func TestInstanceMetadataProvider_Retrieve_Fail3(t *testing.T) {
	// Start a test server locally.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var body string
		var status int

		switch r.URL.Path {
		case "/latest/meta-data/ram/security-credentials/":
			body = ""
			status = 200
		case "/latest/api/token":
			body = "not found"
			status = 404
		}
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	// Update our securityCredURL and tokenURL to point at our local test server.
	originalSecurityCredURL := securityCredURL
	originalTokenURL := tokenURL
	securityCredURL = strings.Replace(securityCredURL, "http://100.100.100.200", ts.URL, -1)
	tokenURL = strings.Replace(tokenURL, "http://100.100.100.200", ts.URL, -1)
	defer func() {
		securityCredURL = originalSecurityCredURL
		tokenURL = originalTokenURL
	}()

	_, err := NewInstanceMetadataProvider().Retrieve()
	assert.NotNil(t, err)
	assert.Equal(t, "unable to retrieve role name, it may be unset", err.Error())
}

func TestInstanceMetadataProvider_Retrieve_Fail4(t *testing.T) {
	// Start a test server locally.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var body string
		var status int

		switch r.URL.Path {
		case "/latest/meta-data/ram/security-credentials/":
			body = "ELK"
			status = 200
		case "/latest/meta-data/ram/security-credentials/ELK":
			body = ``
			status = 404
		case "/latest/api/token":
			body = "not found"
			status = 404
		}
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	// Update our securityCredURL and tokenURL to point at our local test server.
	originalSecurityCredURL := securityCredURL
	originalTokenURL := tokenURL
	securityCredURL = strings.Replace(securityCredURL, "http://100.100.100.200", ts.URL, -1)
	tokenURL = strings.Replace(tokenURL, "http://100.100.100.200", ts.URL, -1)
	defer func() {
		securityCredURL = originalSecurityCredURL
		tokenURL = originalTokenURL
	}()

	_, err := NewInstanceMetadataProvider().Retrieve()
	assert.NotNil(t, err)
	assert.Equal(t, "received 404 getting security credentials for ELK", err.Error())
}

func TestInstanceMetadataProvider_Retrieve_Fail5(t *testing.T) {
	// Start a test server locally.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var body string
		var status int

		switch r.URL.Path {
		case "/latest/meta-data/ram/security-credentials/":
			body = "ELK"
			status = 200
		case "/latest/meta-data/ram/security-credentials/ELK":
			body = `invalid json`
			status = 200
		case "/latest/api/token":
			body = "not found"
			status = 404
		}
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	defer ts.Close()

	// Update our securityCredURL and tokenURL to point at our local test server.
	originalSecurityCredURL := securityCredURL
	originalTokenURL := tokenURL
	securityCredURL = strings.Replace(securityCredURL, "http://100.100.100.200", ts.URL, -1)
	tokenURL = strings.Replace(tokenURL, "http://100.100.100.200", ts.URL, -1)
	defer func() {
		securityCredURL = originalSecurityCredURL
		tokenURL = originalTokenURL
	}()

	_, err := NewInstanceMetadataProvider().Retrieve()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh metadata err, json.Unmarshal fail: invalid character 'i' looking for beginning of value", err.Error())
}

func mockServer(json string) (server *httptest.Server) {
	// Start a test server locally.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var body string
		var status int

		switch r.URL.Path {
		case "/latest/meta-data/ram/security-credentials/":
			body = "ELK"
			status = 200
		case "/latest/meta-data/ram/security-credentials/ELK":
			body = json
			status = 200
		case "/latest/api/token":
			body = "not found"
			status = 404
		}
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	return ts
}

func test(t *testing.T, input, expected string) {
	// Start a test server locally.
	ts := mockServer(input)
	defer ts.Close()

	// Update our securityCredURL and tokenURL to point at our local test server.
	originalSecurityCredURL := securityCredURL
	originalTokenURL := tokenURL
	securityCredURL = strings.Replace(securityCredURL, "http://100.100.100.200", ts.URL, -1)
	tokenURL = strings.Replace(tokenURL, "http://100.100.100.200", ts.URL, -1)
	defer func() {
		securityCredURL = originalSecurityCredURL
		tokenURL = originalTokenURL
	}()

	_, err := NewInstanceMetadataProvider().Retrieve()
	assert.NotNil(t, err)
	assert.Equal(t, expected, err.Error())
}

func TestInstanceMetadataProvider_Retrieve_Fail6(t *testing.T) {
	test(t, `{}`, "refresh metadata err, fail to get credentials, response: {}")
	test(t, `{"AccessKeyId":true}`,
		"refresh metadata err, json.Unmarshal fail: json: cannot unmarshal bool into Go struct field ecsRAMRoleCredentials.AccessKeyId of type string")
	test(t, `{"AccessKeyId":"access key id"}`,
		"refresh metadata err, fail to get credentials, response: {\"AccessKeyId\":\"access key id\"}")
	test(t, `{"AccessKeyId":"access key id", "AccessKeySecret":true}`,
		"refresh metadata err, json.Unmarshal fail: json: cannot unmarshal bool into Go struct field ecsRAMRoleCredentials.AccessKeySecret of type string")
	test(t, `{"AccessKeyId":"access key id", "AccessKeySecret":"secret"}`,
		"refresh metadata err, fail to get credentials, response: {\"AccessKeyId\":\"access key id\", \"AccessKeySecret\":\"secret\"}")
	test(t, `{"AccessKeyId":"access key id", "AccessKeySecret":"secret","SecurityToken":true}`,
		"refresh metadata err, json.Unmarshal fail: json: cannot unmarshal bool into Go struct field ecsRAMRoleCredentials.SecurityToken of type string")
}
