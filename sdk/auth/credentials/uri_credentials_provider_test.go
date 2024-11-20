package credentials

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewURLCredentialsProvider(t *testing.T) {
	rollback := internal.Memory("ALIBABA_CLOUD_CREDENTIALS_URI")
	defer func() {
		rollback()
	}()
	// case 1: no credentials provider
	_, err := NewURLCredentialsProviderBuilderBuilder().
		Build()
	assert.EqualError(t, err, "[SDK.InvalidParam] The url is empty")

	// case 2: no role arn
	os.Setenv("ALIBABA_CLOUD_CREDENTIALS_URI", "http://localhost:8080")
	p, err := NewURLCredentialsProviderBuilderBuilder().
		Build()
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(p.url, "http://localhost:8080"))

	// case 3: check default role session name
	p, err = NewURLCredentialsProviderBuilderBuilder().
		WithUrl("http://localhost:9090").
		Build()
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(p.url, "http://localhost:9090"))
}

func TestURLCredentialsProvider_getCredentials(t *testing.T) {
	originDo := hookDo
	defer func() { hookDo = originDo }()
	p, err := NewURLCredentialsProviderBuilderBuilder().
		WithUrl("http://localhost:8080").
		Build()
	assert.Nil(t, err)

	// case 1: server error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "mock server error", err.Error())

	// case 2: 4xx error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(400, "4xx error")
			return
		}
	}

	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "get credentials from http://localhost:8080 failed: 4xx error", err.Error())

	// case 3: invalid json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, "invalid json")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "get credentials from http://localhost:8080 failed with error, json unmarshal fail: invalid character 'i' looking for beginning of value", err.Error())

	// case 4: empty response json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, "null")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh credentials from http://localhost:8080 failed: null", err.Error())

	// case 5: empty session ak response json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, "{}")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh credentials from http://localhost:8080 failed: {}", err.Error())

	// case 6: mock ok value
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"AccessKeyId":"saki","AccessKeySecret":"saks","Expiration":"2021-10-20T04:27:09Z","SecurityToken":"token"}`)
			return
		}
	}
	creds, err := p.getCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "saki", creds.AccessKeyId)
	assert.Equal(t, "saks", creds.AccessKeySecret)
	assert.Equal(t, "token", creds.SecurityToken)
	assert.Equal(t, "2021-10-20T04:27:09Z", creds.Expiration)

	// needUpdateCredential
	assert.True(t, p.needUpdateCredential())
	p.expirationTimestamp = time.Now().Unix()
	assert.True(t, p.needUpdateCredential())

	p.expirationTimestamp = time.Now().Unix() + 300
	assert.False(t, p.needUpdateCredential())

	// case 4: mock read response error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = &http.Response{
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				ProtoMinor: 1,
				Header:     map[string][]string{},
				StatusCode: 200,
				Status:     "200 " + http.StatusText(200),
			}
			res.Body = ioutil.NopCloser(&errorReader{})
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "read failed", err.Error())
}

func TestURLCredentialsProvider_GetCredentials(t *testing.T) {
	originDo := hookDo
	defer func() { hookDo = originDo }()

	// case 0: get previous credentials failed
	p, err := NewURLCredentialsProviderBuilderBuilder().
		WithUrl("http://localhost:8080").
		Build()
	assert.Nil(t, err)

	// case 1: get credentials failed
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.GetCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "mock server error", err.Error())

	// case 2: get invalid expiration
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"AccessKeyId":"akid","AccessKeySecret":"aksecret","Expiration":"invalidexpiration","SecurityToken":"ststoken"}`)
			return
		}
	}
	_, err = p.GetCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "parsing time \"invalidexpiration\" as \"2006-01-02T15:04:05Z\": cannot parse \"invalidexpiration\" as \"2006\"", err.Error())

	// case 3: happy result
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"AccessKeyId":"akid","AccessKeySecret":"aksecret","Expiration":"2021-10-20T04:27:09Z","SecurityToken":"ststoken"}`)
			return
		}
	}
	cc, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", cc.AccessKeyId)
	assert.Equal(t, "aksecret", cc.AccessKeySecret)
	assert.Equal(t, "ststoken", cc.SecurityToken)
	assert.Equal(t, "credential_uri", cc.ProviderName)
	assert.True(t, p.needUpdateCredential())
	// get credentials again
	cc, err = p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", cc.AccessKeyId)
	assert.Equal(t, "aksecret", cc.AccessKeySecret)
	assert.Equal(t, "ststoken", cc.SecurityToken)
	assert.Equal(t, "credential_uri", cc.ProviderName)
	assert.True(t, p.needUpdateCredential())
}

func TestURLCredentialsProviderWithHttpOptions(t *testing.T) {
	p, err := NewURLCredentialsProviderBuilderBuilder().
		WithUrl("http://100.100.100.200").
		WithHttpOptions(&HttpOptions{
			ConnectTimeout: 100 * time.Millisecond,
			ReadTimeout:    100 * time.Millisecond,
		}).
		Build()
	assert.Nil(t, err)
	_, err = p.GetCredentials()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "timeout")
}
