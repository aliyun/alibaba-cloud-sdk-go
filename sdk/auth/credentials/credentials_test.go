package credentials

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/internal"
	"github.com/stretchr/testify/assert"
)

func mockResponse(statusCode int, content string) (res *http.Response) {
	status := strconv.Itoa(statusCode)
	res = &http.Response{
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		Header:     map[string][]string{"sdk": {"test"}},
		StatusCode: statusCode,
		Status:     status + " " + http.StatusText(statusCode),
	}
	res.Body = ioutil.NopCloser(bytes.NewReader([]byte(content)))
	return
}

func TestStaticAKCredentialsProvider(t *testing.T) {
	p := NewStaticAKCredentialsProvider("akid", "aksecret")
	c, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", c.AccessKeyId)
	assert.Equal(t, "aksecret", c.AccessKeySecret)
	assert.Equal(t, "", c.SecurityToken)
	assert.Equal(t, "", c.BearerToken)

	_, err = NewStaticAKCredentialsProviderBuilder().
		Build()
	assert.EqualError(t, err, "[SDK.InvalidParam] The access key id is empty")

	_, err = NewStaticAKCredentialsProviderBuilder().
		WithAccessKeyId("akid").
		Build()
	assert.EqualError(t, err, "[SDK.InvalidParam] The access key secret is empty")

	provider, err := NewStaticAKCredentialsProviderBuilder().
		WithAccessKeyId("accessKeyId").
		WithAccessKeySecret("accessKeySecret").
		Build()
	assert.Nil(t, err)
	assert.Equal(t, "static_ak", provider.GetProviderName())

	cred, err := provider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "accessKeyId", cred.AccessKeyId)
	assert.Equal(t, "accessKeySecret", cred.AccessKeySecret)
	assert.Equal(t, "", cred.SecurityToken)
	assert.Equal(t, "static_ak", cred.ProviderName)
}

func TestStaticAKCredentialsProviderWithEnv(t *testing.T) {
	originAKID := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	originAKSecret := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
	defer func() {
		os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_ID", originAKID)
		os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET", originAKSecret)
	}()

	os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_ID", "akid_from_env")
	os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET", "aksecret_from_env")
	provider, err := NewStaticAKCredentialsProviderBuilder().
		Build()
	assert.Nil(t, err)
	assert.Equal(t, "static_ak", provider.GetProviderName())

	cred, err := provider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid_from_env", cred.AccessKeyId)
	assert.Equal(t, "aksecret_from_env", cred.AccessKeySecret)
	assert.Equal(t, "", cred.SecurityToken)
	assert.Equal(t, "static_ak", cred.ProviderName)
}

func TestStaticSTSCredentialsProvider(t *testing.T) {
	p := NewStaticSTSCredentialsProvider("akid", "aksecret", "token")
	c, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", c.AccessKeyId)
	assert.Equal(t, "aksecret", c.AccessKeySecret)
	assert.Equal(t, "token", c.SecurityToken)
	assert.Equal(t, "", c.BearerToken)

	_, err = NewStaticSTSCredentialsProviderBuilder().
		Build()
	assert.EqualError(t, err, "[SDK.InvalidParam] The access key id is empty")

	_, err = NewStaticSTSCredentialsProviderBuilder().
		WithAccessKeyId("akid").
		Build()
	assert.EqualError(t, err, "[SDK.InvalidParam] The access key secret is empty")

	_, err = NewStaticSTSCredentialsProviderBuilder().
		WithAccessKeyId("akid").
		WithAccessKeySecret("aksecret").
		Build()
	assert.EqualError(t, err, "[SDK.InvalidParam] The security token is empty")

	provider, err := NewStaticSTSCredentialsProviderBuilder().
		WithAccessKeyId("accessKeyId").
		WithAccessKeySecret("accessKeySecret").
		WithSecurityToken("securityToken").
		Build()
	assert.Nil(t, err)
	assert.Equal(t, "static_sts", provider.GetProviderName())

	cred, err := provider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "accessKeyId", cred.AccessKeyId)
	assert.Equal(t, "accessKeySecret", cred.AccessKeySecret)
	assert.Equal(t, "securityToken", cred.SecurityToken)
	assert.Equal(t, "static_sts", cred.ProviderName)
}

func TestStaticSTSCredentialsProviderWithEnv(t *testing.T) {
	originAKID := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	originAKSecret := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
	originToken := os.Getenv("ALIBABA_CLOUD_SECURITY_TOKEN")
	defer func() {
		os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_ID", originAKID)
		os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET", originAKSecret)
		os.Setenv("ALIBABA_CLOUD_SECURITY_TOKEN", originToken)
	}()

	os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_ID", "akid_from_env")
	os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET", "aksecret_from_env")
	os.Setenv("ALIBABA_CLOUD_SECURITY_TOKEN", "token_from_env")
	provider, err := NewStaticSTSCredentialsProviderBuilder().
		Build()
	assert.Nil(t, err)
	assert.Equal(t, "static_sts", provider.GetProviderName())

	cred, err := provider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid_from_env", cred.AccessKeyId)
	assert.Equal(t, "aksecret_from_env", cred.AccessKeySecret)
	assert.Equal(t, "token_from_env", cred.SecurityToken)
	assert.Equal(t, "static_sts", cred.ProviderName)
}

func TestBearerTokenCredentialsProvider(t *testing.T) {
	p := NewBearerTokenCredentialsProvider("bearertoken")
	c, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "", c.AccessKeyId)
	assert.Equal(t, "", c.AccessKeySecret)
	assert.Equal(t, "", c.SecurityToken)
	assert.Equal(t, "bearertoken", c.BearerToken)
}

type errorReader struct {
}

func (r *errorReader) Read(p []byte) (n int, err error) {
	err = errors.New("read failed")
	return
}

func TestRSAKeyPairCredentialsProvider(t *testing.T) {
	_, err := NewRSAKeyPairCredentialsProvider("publicKey", "privateKey", 3600)
	assert.Nil(t, err)
	// invalid durationSeconds
	_, err = NewRSAKeyPairCredentialsProvider("publicKey", "privateKey", 100)
	assert.NotNil(t, err)
	assert.Equal(t, "[SDK.InvalidParam] Key Pair session duration should be in the range of 15min - 1hr", err.Error())
	// default durationSeconds
	p, err := NewRSAKeyPairCredentialsProvider("publicKey", "privateKey", 0)
	assert.Nil(t, err)
	assert.Equal(t, 3600, p.durationSeconds)

	privateKey := `
MIICeQIBADANBgkqhkiG9w0BAQEFAASCAmMwggJfAgEAAoGBAOJC+2WXtkXZ+6sa
3+qJp4mDOsiZb3BghHT9nVbjTeaw4hsZWHYxQ6l6XDmTg4twPB59LOGAlAjYrT31
3pdwEawnmdf6zyF93Zvxxpy7lO2HoxYKSjbtXO4I0pcq3WTnw2xlbhqHvrcuWwt+
FqH9akzcnwHjc03siZBzt/dwDL3vAgMBAAECgYEAzwgZPqFuUEYgaTVDFDl2ynYA
kNMMzBgUu3Pgx0Nf4amSitdLQYLcdbQXtTtMT4eYCxHgwkpDqkCRbLOQRKNwFo0I
oaCuhjZlxWcKil4z4Zb/zB7gkeuXPOVUjFSS3FogsRWMtnNAMgR/yJRlbcg/Puqk
Magt/yDk+7cJCe6H96ECQQDxMT4S+tVP9nOw//QT39Dk+kWe/YVEhnWnCMZmGlEq
1gnN6qpUi68ts6b3BVgrDPrPN6wm/Z9vpcKNeWpIvxXRAkEA8CcT2UEUwDGRKAUu
WVPJqdAJjpjc072eRF5g792NyO+TAF6thBlDKNslRvFQDB6ymLsjfy8JYCnGbbSb
WqbHvwJBAIs7KeI6+jiWxGJA3t06LpSABQCqyOut0u0Bm8YFGyXnOPGtrXXwzMdN
Fe0zIJp5e69zK+W2Mvt4bL7OgBROeoECQQDsE+4uLw0gFln0tosmovhmp60NcfX7
bLbtzL2MbwbXlbOztF7ssgzUWAHgKI6hK3g0LhsqBuo3jzmSVO43giZvAkEA08Nm
2TI9EvX6DfCVfPOiKZM+Pijh0xLN4Dn8qUgt3Tcew/vfj4WA2ZV6qiJqL01vMsHc
vftlY0Hs1vNXcaBgEA==`
	p, err = NewRSAKeyPairCredentialsProvider("publicKey", privateKey, 3600)
	assert.Nil(t, err)

	originNewRequest := hookNewRequest
	defer func() { hookNewRequest = originNewRequest }()

	// case 1: mock new http request failed
	hookNewRequest = func(fn newReuqest) newReuqest {
		return func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("new http request failed")
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "new http request failed", err.Error())
	// reset new request
	hookNewRequest = originNewRequest

	originDo := hookDo
	defer func() { hookDo = originDo }()

	// case 2: mock read response error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			status := strconv.Itoa(200)
			res = &http.Response{
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				Header:     map[string][]string{},
				StatusCode: 200,
				Status:     status + " " + http.StatusText(200),
			}
			res.Body = ioutil.NopCloser(&errorReader{})
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "read failed", err.Error())

	// case 3: server error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "mock server error", err.Error())

	// case 4: 4xx error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(400, "4xx error")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "SDK.ServerError\nErrorCode: \nRecommend: refresh temp ak failed\nRequestId: \nMessage: 4xx error\nRespHeaders: map[]", err.Error())

	// case 5: invalid json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, "invalid json")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh temp ak err, json.Unmarshal fail: invalid character 'i' looking for beginning of value", err.Error())

	// case 6: empty response json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, "null")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh temp ak token err, fail to get credentials", err.Error())

	// case 7: empty session ak response json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"SessionAccessKey": {}}`)
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh temp ak token err, fail to get credentials", err.Error())

	// case 8: mock ok value
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"SessionAccessKey": {"SessionAccessKeyId":"saki","SessionAccessKeySecret":"saks","Expiration":"2021-10-20T04:27:09Z"}}`)
			return
		}
	}
	creds, err := p.getCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "saki", *creds.SessionAccessKeyId)
	assert.Equal(t, "saks", *creds.SessionAccessKeySecret)
	assert.Equal(t, "2021-10-20T04:27:09Z", *creds.Expiration)

	// needUpdateCredential
	assert.True(t, p.needUpdateCredential())
	p.expirationTimestamp = time.Now().Unix()
	assert.True(t, p.needUpdateCredential())

	p.expirationTimestamp = time.Now().Unix() + 300
	assert.False(t, p.needUpdateCredential())
}

func TestRSAKeyPairCredentialsProviderGetCredentials(t *testing.T) {
	privateKey := `
MIICeQIBADANBgkqhkiG9w0BAQEFAASCAmMwggJfAgEAAoGBAOJC+2WXtkXZ+6sa
3+qJp4mDOsiZb3BghHT9nVbjTeaw4hsZWHYxQ6l6XDmTg4twPB59LOGAlAjYrT31
3pdwEawnmdf6zyF93Zvxxpy7lO2HoxYKSjbtXO4I0pcq3WTnw2xlbhqHvrcuWwt+
FqH9akzcnwHjc03siZBzt/dwDL3vAgMBAAECgYEAzwgZPqFuUEYgaTVDFDl2ynYA
kNMMzBgUu3Pgx0Nf4amSitdLQYLcdbQXtTtMT4eYCxHgwkpDqkCRbLOQRKNwFo0I
oaCuhjZlxWcKil4z4Zb/zB7gkeuXPOVUjFSS3FogsRWMtnNAMgR/yJRlbcg/Puqk
Magt/yDk+7cJCe6H96ECQQDxMT4S+tVP9nOw//QT39Dk+kWe/YVEhnWnCMZmGlEq
1gnN6qpUi68ts6b3BVgrDPrPN6wm/Z9vpcKNeWpIvxXRAkEA8CcT2UEUwDGRKAUu
WVPJqdAJjpjc072eRF5g792NyO+TAF6thBlDKNslRvFQDB6ymLsjfy8JYCnGbbSb
WqbHvwJBAIs7KeI6+jiWxGJA3t06LpSABQCqyOut0u0Bm8YFGyXnOPGtrXXwzMdN
Fe0zIJp5e69zK+W2Mvt4bL7OgBROeoECQQDsE+4uLw0gFln0tosmovhmp60NcfX7
bLbtzL2MbwbXlbOztF7ssgzUWAHgKI6hK3g0LhsqBuo3jzmSVO43giZvAkEA08Nm
2TI9EvX6DfCVfPOiKZM+Pijh0xLN4Dn8qUgt3Tcew/vfj4WA2ZV6qiJqL01vMsHc
vftlY0Hs1vNXcaBgEA==`
	p, err := NewRSAKeyPairCredentialsProvider("publicKey", privateKey, 3600)
	assert.Nil(t, err)

	originDo := hookDo
	defer func() { hookDo = originDo }()

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
			res = mockResponse(200, `{"SessionAccessKey": {"SessionAccessKeyId":"saki","SessionAccessKeySecret":"saks","Expiration":"invalidexpiration"}}`)
			return
		}
	}
	_, err = p.GetCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "parsing time \"invalidexpiration\" as \"2006-01-02T15:04:05Z\": cannot parse \"invalidexpiration\" as \"2006\"", err.Error())

	// case 3: happy result
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"SessionAccessKey": {"SessionAccessKeyId":"saki","SessionAccessKeySecret":"saks","Expiration":"2021-10-20T04:27:09Z"}}`)
			return
		}
	}
	cc, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "saki", cc.AccessKeyId)
	assert.Equal(t, "saks", cc.AccessKeySecret)
	assert.True(t, p.needUpdateCredential())
}

func TestRSAKeyPairCredentialsProviderGetCredentialsWithError(t *testing.T) {
	privateKey := `
MIICeQIBADANBgkqhkiG9w0BAQEFAASCAmMwggJfAgEAAoGBAOJC+2WXtkXZ+6sa
3+qJp4mDOsiZb3BghHT9nVbjTeaw4hsZWHYxQ6l6XDmTg4twPB59LOGAlAjYrT31
3pdwEawnmdf6zyF93Zvxxpy7lO2HoxYKSjbtXO4I0pcq3WTnw2xlbhqHvrcuWwt+
FqH9akzcnwHjc03siZBzt/dwDL3vAgMBAAECgYEAzwgZPqFuUEYgaTVDFDl2ynYA
kNMMzBgUu3Pgx0Nf4amSitdLQYLcdbQXtTtMT4eYCxHgwkpDqkCRbLOQRKNwFo0I
oaCuhjZlxWcKil4z4Zb/zB7gkeuXPOVUjFSS3FogsRWMtnNAMgR/yJRlbcg/Puqk
Magt/yDk+7cJCe6H96ECQQDxMT4S+tVP9nOw//QT39Dk+kWe/YVEhnWnCMZmGlEq
1gnN6qpUi68ts6b3BVgrDPrPN6wm/Z9vpcKNeWpIvxXRAkEA8CcT2UEUwDGRKAUu
WVPJqdAJjpjc072eRF5g792NyO+TAF6thBlDKNslRvFQDB6ymLsjfy8JYCnGbbSb
WqbHvwJBAIs7KeI6+jiWxGJA3t06LpSABQCqyOut0u0Bm8YFGyXnOPGtrXXwzMdN
Fe0zIJp5e69zK+W2Mvt4bL7OgBROeoECQQDsE+4uLw0gFln0tosmovhmp60NcfX7
bLbtzL2MbwbXlbOztF7ssgzUWAHgKI6hK3g0LhsqBuo3jzmSVO43giZvAkEA08Nm
2TI9EvX6DfCVfPOiKZM+Pijh0xLN4Dn8qUgt3Tcew/vfj4WA2ZV6qiJqL01vMsHc
vftlY0Hs1vNXcaBgEA==`
	p, err := NewRSAKeyPairCredentialsProvider("publicKey", privateKey, 3600)
	assert.Nil(t, err)
	_, err = p.GetCredentials()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "SDK.ServerError")
}

func TestNewRAMRoleARNCredentialsProvider(t *testing.T) {
	akProvider := NewStaticAKCredentialsProvider("akid", "aksecret")
	p, err := NewRAMRoleARNCredentialsProvider(akProvider, "roleArn", "rsn", 1000, "", "", "")
	assert.Nil(t, err)
	assert.Equal(t, "rsn", p.roleSessionName)
	assert.Equal(t, 1000, p.durationSeconds)

	p, err = NewRAMRoleARNCredentialsProvider(akProvider, "roleArn", "", 1000, "", "", "")
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(p.roleSessionName, "aliyun-go-sdk-"))

	p, err = NewRAMRoleARNCredentialsProvider(akProvider, "roleArn", "", 0, "", "", "")
	assert.Nil(t, err)
	assert.Equal(t, 3600, p.durationSeconds)

	_, err = NewRAMRoleARNCredentialsProvider(akProvider, "roleArn", "", 100, "", "", "")
	assert.NotNil(t, err)
	assert.Equal(t, "[SDK.InvalidParam] Assume Role session duration should be in the range of 15min - 1hr", err.Error())

	rollback := internal.Memory("ALIBABA_CLOUD_STS_REGION", "ALIBABA_CLOUD_VPC_ENDPOINT_ENABLED")
	defer func() {
		rollback()
	}()
	// case 1: no credentials provider
	_, err = NewRAMRoleARNCredentialsProviderBuilder().
		Build()
	assert.EqualError(t, err, "[SDK.InvalidParam] Must specify a previous credentials provider to assume role")

	// case 2: no role arn
	akProvider, err = NewStaticAKCredentialsProviderBuilder().
		WithAccessKeyId("akid").
		WithAccessKeySecret("aksecret").
		Build()
	assert.Nil(t, err)
	_, err = NewRAMRoleARNCredentialsProviderBuilder().
		WithCredentialsProvider(akProvider).
		Build()
	assert.EqualError(t, err, "[SDK.InvalidParam] The RoleArn is empty")

	// case 3: check default role session name
	p, err = NewRAMRoleARNCredentialsProviderBuilder().
		WithCredentialsProvider(akProvider).
		WithRoleArn("roleArn").
		Build()
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(p.roleSessionName, "aliyun-go-sdk-"))

	// case 4: check default duration seconds
	p, err = NewRAMRoleARNCredentialsProviderBuilder().
		WithCredentialsProvider(akProvider).
		WithRoleArn("roleArn").Build()
	assert.Nil(t, err)
	assert.Equal(t, 3600, p.durationSeconds)

	// case 5: check invalid duration seconds
	_, err = NewRAMRoleARNCredentialsProviderBuilder().
		WithCredentialsProvider(akProvider).
		WithRoleArn("roleArn").
		WithDurationSeconds(100).
		Build()
	assert.EqualError(t, err, "[SDK.InvalidParam] Session duration should be in the range of 900s - max session duration")

	// case 6: check all duration seconds
	p, err = NewRAMRoleARNCredentialsProviderBuilder().
		WithCredentialsProvider(akProvider).
		WithRoleArn("roleArn").
		WithStsRegion("cn-hangzhou").
		WithEnableVpc(true).
		WithPolicy("policy").
		WithExternalId("externalId").
		WithRoleSessionName("rsn").
		WithDurationSeconds(1000).
		Build()
	assert.Nil(t, err)
	assert.Equal(t, "rsn", p.roleSessionName)
	assert.Equal(t, "roleArn", p.roleArn)
	assert.Equal(t, "policy", p.policy)
	assert.Equal(t, "externalId", p.externalId)
	assert.Equal(t, "cn-hangzhou", p.stsRegion)
	assert.Equal(t, 1000, p.durationSeconds)
	// sts endpoint with sts region
	assert.Equal(t, "sts-vpc.cn-hangzhou.aliyuncs.com", p.stsEndpoint)

	// case 7: check default sts endpoint
	os.Setenv("ALIBABA_CLOUD_VPC_ENDPOINT_ENABLED", "1")
	p, err = NewRAMRoleARNCredentialsProviderBuilder().
		WithCredentialsProvider(akProvider).
		WithRoleArn("roleArn").
		WithPolicy("policy").
		WithExternalId("externalId").
		WithRoleSessionName("rsn").
		WithDurationSeconds(1000).
		Build()
	assert.Nil(t, err)
	assert.Equal(t, "rsn", p.roleSessionName)
	assert.Equal(t, "roleArn", p.roleArn)
	assert.Equal(t, "policy", p.policy)
	assert.Equal(t, "externalId", p.externalId)
	assert.Equal(t, "", p.stsRegion)
	assert.Equal(t, 1000, p.durationSeconds)
	assert.Equal(t, "sts.aliyuncs.com", p.stsEndpoint)

	// case 8: check sts endpoint with env
	os.Setenv("ALIBABA_CLOUD_STS_REGION", "cn-hangzhou")
	os.Setenv("ALIBABA_CLOUD_VPC_ENDPOINT_ENABLED", "True")
	p, err = NewRAMRoleARNCredentialsProviderBuilder().
		WithCredentialsProvider(akProvider).
		WithRoleArn("roleArn").
		WithPolicy("policy").
		WithExternalId("externalId").
		WithRoleSessionName("rsn").
		WithDurationSeconds(1000).
		Build()
	assert.Nil(t, err)
	assert.Equal(t, "sts-vpc.cn-hangzhou.aliyuncs.com", p.stsEndpoint)

	// case 9: check sts endpoint with sts endpoint
	p, err = NewRAMRoleARNCredentialsProviderBuilder().
		WithCredentialsProvider(akProvider).
		WithRoleArn("roleArn").
		WithStsEndpoint("sts.cn-shanghai.aliyuncs.com").
		WithPolicy("policy").
		WithExternalId("externalId").
		WithRoleSessionName("rsn").
		WithDurationSeconds(1000).
		Build()
	assert.Nil(t, err)
	assert.Equal(t, "rsn", p.roleSessionName)
	assert.Equal(t, "roleArn", p.roleArn)
	assert.Equal(t, "policy", p.policy)
	assert.Equal(t, "externalId", p.externalId)
	assert.Equal(t, "", p.stsRegion)
	assert.Equal(t, 1000, p.durationSeconds)
	assert.Equal(t, "sts.cn-shanghai.aliyuncs.com", p.stsEndpoint)

	// case 10: check ak&sk
	p, err = NewRAMRoleARNCredentialsProviderBuilder().
		WithAccessKeyId("ak").
		WithAccessKeySecret("sk").
		WithRoleArn("roleArn").
		WithStsEndpoint("sts.cn-shanghai.aliyuncs.com").
		WithPolicy("policy").
		WithExternalId("externalId").
		WithRoleSessionName("rsn").
		WithDurationSeconds(1000).
		Build()
	assert.Nil(t, err)
	cre, err := p.credentialsProvider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "ak", cre.AccessKeyId)
	assert.Equal(t, "sk", cre.AccessKeySecret)
	assert.Equal(t, "static_ak", cre.ProviderName)

	// case 11: check ak&sk&token
	p, err = NewRAMRoleARNCredentialsProviderBuilder().
		WithAccessKeyId("ak").
		WithAccessKeySecret("sk").
		WithSecurityToken("token").
		WithRoleArn("roleArn").
		WithStsEndpoint("sts.cn-shanghai.aliyuncs.com").
		WithPolicy("policy").
		WithExternalId("externalId").
		WithRoleSessionName("rsn").
		WithDurationSeconds(1000).
		Build()
	assert.Nil(t, err)
	cre, err = p.credentialsProvider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "ak", cre.AccessKeyId)
	assert.Equal(t, "sk", cre.AccessKeySecret)
	assert.Equal(t, "token", cre.SecurityToken)
	assert.Equal(t, "static_sts", cre.ProviderName)
}

func TestRAMRoleARNCredentialsProvider_getCredentials(t *testing.T) {
	akProvider := NewStaticAKCredentialsProvider("akid", "aksecret")
	p, err := NewRAMRoleARNCredentialsProvider(akProvider, "roleArn", "rsn", 1000, "", "", "")
	assert.Nil(t, err)

	cc, err := akProvider.GetCredentials()
	assert.Nil(t, err)

	originNewRequest := hookNewRequest
	defer func() { hookNewRequest = originNewRequest }()

	// case 1: mock new http request failed
	hookNewRequest = func(fn newReuqest) newReuqest {
		return func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("new http request failed")
		}
	}
	_, err = p.getCredentials(cc)
	assert.NotNil(t, err)
	assert.Equal(t, "new http request failed", err.Error())
	// reset new request
	hookNewRequest = originNewRequest

	originDo := hookDo
	defer func() { hookDo = originDo }()

	// case 2: server error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.getCredentials(cc)
	assert.NotNil(t, err)
	assert.Equal(t, "mock server error", err.Error())

	// case 3: mock read response error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			status := strconv.Itoa(200)
			res = &http.Response{
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				Header:     map[string][]string{},
				StatusCode: 200,
				Status:     status + " " + http.StatusText(200),
			}
			res.Body = ioutil.NopCloser(&errorReader{})
			return
		}
	}
	_, err = p.getCredentials(cc)
	assert.NotNil(t, err)
	assert.Equal(t, "read failed", err.Error())

	// case 4: 4xx error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(400, "4xx error")
			return
		}
	}
	_, err = p.getCredentials(cc)
	assert.NotNil(t, err)
	assert.Equal(t, "SDK.ServerError\nErrorCode: \nRecommend: refresh session token failed\nRequestId: \nMessage: 4xx error\nRespHeaders: map[]", err.Error())

	// case 5: invalid json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, "invalid json")
			return
		}
	}
	_, err = p.getCredentials(cc)
	assert.NotNil(t, err)
	assert.Equal(t, "refresh RoleArn sts token err, json.Unmarshal fail: invalid character 'i' looking for beginning of value", err.Error())

	// case 6: empty response json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, "null")
			return
		}
	}
	_, err = p.getCredentials(cc)
	assert.NotNil(t, err)
	assert.Equal(t, "refresh RoleArn sts token err, fail to get credentials", err.Error())

	// case 7: empty session ak response json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"Credentials": {}}`)
			return
		}
	}
	_, err = p.getCredentials(cc)
	assert.NotNil(t, err)
	assert.Equal(t, "refresh RoleArn sts token err, fail to get credentials", err.Error())

	// case 8: mock ok value
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"Credentials": {"AccessKeyId":"saki","AccessKeySecret":"saks","Expiration":"2021-10-20T04:27:09Z","SecurityToken":"token"}}`)
			return
		}
	}
	creds, err := p.getCredentials(cc)
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
}

func TestRAMRoleARNCredentialsProvider_getCredentialsWithRequestCheck(t *testing.T) {
	originDo := hookDo
	defer func() { hookDo = originDo }()

	akProvider := NewStaticSTSCredentialsProvider("akid", "aksecret", "ststoken")
	p, err := NewRAMRoleARNCredentialsProvider(akProvider, "roleArn", "rsn", 1000, "policy", "cn-beijing", "externalId")
	assert.Nil(t, err)

	// case 1: server error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			assert.Equal(t, "sts.cn-beijing.aliyuncs.com", req.Host)
			assert.Contains(t, req.URL.String(), "SecurityToken=ststoken")
			body, err := ioutil.ReadAll(req.Body)
			assert.Nil(t, err)
			bodyString := string(body)
			assert.Contains(t, bodyString, "Policy=policy")
			assert.Contains(t, bodyString, "RoleArn=roleArn")
			assert.Contains(t, bodyString, "RoleSessionName=rsn")
			assert.Contains(t, bodyString, "DurationSeconds=1000")

			err = errors.New("mock server error")
			return
		}
	}

	cc, err := akProvider.GetCredentials()
	assert.Nil(t, err)
	_, err = p.getCredentials(cc)
	assert.NotNil(t, err)
	assert.Equal(t, "mock server error", err.Error())
}

type errorCredentialsProvider struct {
}

func (p *errorCredentialsProvider) GetCredentials() (cc *Credentials, err error) {
	err = errors.New("get credentials failed")
	return
}

func (p *errorCredentialsProvider) GetProviderName() string {
	return "error_mock"
}
func TestRAMRoleARNCredentialsProviderGetCredentials(t *testing.T) {
	originDo := hookDo
	defer func() { hookDo = originDo }()

	// case 0: get previous credentials failed
	p, err := NewRAMRoleARNCredentialsProvider(&errorCredentialsProvider{}, "roleArn", "rsn", 1000, "", "", "")
	assert.Nil(t, err)
	_, err = p.GetCredentials()
	assert.Equal(t, "get credentials failed", err.Error())

	akProvider := NewStaticAKCredentialsProvider("akid", "aksecret")
	p, err = NewRAMRoleARNCredentialsProvider(akProvider, "roleArn", "rsn", 1000, "", "", "")
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
			res = mockResponse(200, `{"Credentials": {"AccessKeyId":"akid","AccessKeySecret":"aksecret","Expiration":"invalidexpiration","SecurityToken":"ststoken"}}`)
			return
		}
	}
	_, err = p.GetCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "parsing time \"invalidexpiration\" as \"2006-01-02T15:04:05Z\": cannot parse \"invalidexpiration\" as \"2006\"", err.Error())

	// case 3: happy result
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"Credentials": {"AccessKeyId":"akid","AccessKeySecret":"aksecret","Expiration":"2021-10-20T04:27:09Z","SecurityToken":"ststoken"}}`)
			return
		}
	}
	cc, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", cc.AccessKeyId)
	assert.Equal(t, "aksecret", cc.AccessKeySecret)
	assert.Equal(t, "ststoken", cc.SecurityToken)
	assert.True(t, p.needUpdateCredential())
	assert.Equal(t, "ram_role_arn/static_ak", cc.ProviderName)

	pp, err := NewRAMRoleARNCredentialsProviderBuilder().
		WithCredentialsProvider(p).
		WithRoleArn("roleArn").
		WithRoleSessionName("rsn").
		WithDurationSeconds(1000).
		Build()
	assert.Nil(t, err)
	cc, err = pp.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", cc.AccessKeyId)
	assert.Equal(t, "aksecret", cc.AccessKeySecret)
	assert.Equal(t, "ststoken", cc.SecurityToken)
	assert.True(t, pp.needUpdateCredential())
	assert.Equal(t, "ram_role_arn/ram_role_arn/static_ak", cc.ProviderName)
}

func TestRAMRoleARNCredentialsProviderGetCredentialsWithError(t *testing.T) {
	akProvider := NewStaticAKCredentialsProvider("akid", "aksecret")
	p, err := NewRAMRoleARNCredentialsProvider(akProvider, "roleArn", "rsn", 1000, "", "", "")
	assert.Nil(t, err)
	_, err = p.GetCredentials()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "SDK.ServerError")
}

func TestRAMRoleARNCredentialsProviderWithHttpOptions(t *testing.T) {
	akProvider, err := NewStaticAKCredentialsProviderBuilder().
		WithAccessKeyId("akid").
		WithAccessKeySecret("aksecret").
		Build()
	assert.Nil(t, err)
	p, err := NewRAMRoleARNCredentialsProviderBuilder().
		WithCredentialsProvider(akProvider).
		WithRoleArn("roleArn").
		WithRoleSessionName("rsn").
		WithDurationSeconds(1000).
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

func TestNewECSRAMRoleCredentialsProvider(t *testing.T) {
	p := NewECSRAMRoleCredentialsProvider("")
	assert.Equal(t, "", p.roleName)

	p = NewECSRAMRoleCredentialsProvider("role")
	assert.Equal(t, "role", p.roleName)

	assert.True(t, p.needUpdateCredential())

	rollback := internal.Memory("ALIBABA_CLOUD_ECS_METADATA_DISABLED", "ALIBABA_CLOUD_ECS_METADATA", "ALIBABA_CLOUD_IMDSV1_DISABLED")
	defer func() {
		rollback()
	}()
	p, err := NewECSRAMRoleCredentialsProviderBuilder().Build()
	assert.Nil(t, err)
	assert.Equal(t, "", p.roleName)

	os.Setenv("ALIBABA_CLOUD_ECS_METADATA", "rolename")
	p, err = NewECSRAMRoleCredentialsProviderBuilder().Build()
	assert.Nil(t, err)
	assert.Equal(t, "rolename", p.roleName)

	p, err = NewECSRAMRoleCredentialsProviderBuilder().WithRoleName("role").Build()
	assert.Nil(t, err)
	assert.Equal(t, "role", p.roleName)
	assert.False(t, p.disableIMDSv1)

	os.Setenv("ALIBABA_CLOUD_IMDSV1_DISABLED", "True")
	p, err = NewECSRAMRoleCredentialsProviderBuilder().Build()
	assert.Nil(t, err)
	assert.True(t, p.disableIMDSv1)

	os.Setenv("ALIBABA_CLOUD_IMDSV1_DISABLED", "1")
	p, err = NewECSRAMRoleCredentialsProviderBuilder().WithDisableIMDSv1(true).Build()
	assert.Nil(t, err)
	assert.True(t, p.disableIMDSv1)

	os.Setenv("ALIBABA_CLOUD_ECS_METADATA_DISABLED", "True")
	_, err = NewECSRAMRoleCredentialsProviderBuilder().Build()
	assert.Equal(t, "IMDS credentials is disabled", err.Error())

	assert.True(t, p.needUpdateCredential())
}

func TestECSRAMRoleCredentialsProvider_getRoleName(t *testing.T) {
	p := NewECSRAMRoleCredentialsProvider("")

	originNewRequest := hookNewRequest
	defer func() { hookNewRequest = originNewRequest }()

	// case 1: mock new http request failed
	hookNewRequest = func(fn newReuqest) newReuqest {
		return func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("new http request failed")
		}
	}
	_, err := p.getRoleName()
	assert.NotNil(t, err)
	assert.Equal(t, "get role name failed: new http request failed", err.Error())
	// reset new request
	hookNewRequest = originNewRequest

	originDo := hookDo
	defer func() { hookDo = originDo }()

	// case 2: server error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.getRoleName()
	assert.NotNil(t, err)
	assert.Equal(t, "get role name failed: mock server error", err.Error())

	// case 3: 4xx error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(400, "4xx error")
			return
		}
	}

	_, err = p.getRoleName()
	assert.NotNil(t, err)
	assert.Equal(t, "get role name failed: request http://100.100.100.200/latest/meta-data/ram/security-credentials/ 400", err.Error())

	// case 4: mock read response error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			status := strconv.Itoa(200)
			res = &http.Response{
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				Header:     map[string][]string{},
				StatusCode: 200,
				Status:     status + " " + http.StatusText(200),
			}
			res.Body = ioutil.NopCloser(&errorReader{})
			return
		}
	}
	_, err = p.getRoleName()
	assert.NotNil(t, err)
	assert.Equal(t, "read failed", err.Error())

	// case 5: value json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, "rolename")
			return
		}
	}
	roleName, err := p.getRoleName()
	assert.Nil(t, err)
	assert.Equal(t, "rolename", roleName)
}

func TestECSRAMRoleCredentialsProvider_getRoleNameWithMetadataV2(t *testing.T) {
	originDo := hookDo
	defer func() { hookDo = originDo }()

	p, err := NewECSRAMRoleCredentialsProviderBuilder().WithDisableIMDSv1(true).Build()
	assert.Nil(t, err)

	// case 1: get metadata token failed
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			err = errors.New("mock server error")
			return
		}
	}

	_, err = p.getRoleName()
	assert.NotNil(t, err)
	assert.Equal(t, "get metadata token failed: mock server error", err.Error())

	// case 2: return token
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			if req.URL.Path == "/latest/api/token" {
				res = mockResponse(200, "tokenxxxxx")
			} else {
				assert.Equal(t, "tokenxxxxx", req.Header.Get("x-aliyun-ecs-metadata-token"))
				res = mockResponse(200, "rolename")
			}
			return
		}
	}

	roleName, err := p.getRoleName()
	assert.Nil(t, err)
	assert.Equal(t, "rolename", roleName)
}

func TestECSRAMRoleCredentialsProvider_getCredentials(t *testing.T) {
	originDo := hookDo
	defer func() { hookDo = originDo }()

	p := NewECSRAMRoleCredentialsProvider("")

	// case 1: server error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			err = errors.New("mock server error")
			return
		}
	}
	_, err := p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "get role name failed: mock server error", err.Error())

	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			if req.URL.Path == "/latest/meta-data/ram/security-credentials/" {
				res = mockResponse(200, "rolename")
				return
			}

			err = errors.New("mock server error")
			return
		}
	}

	originNewRequest := hookNewRequest
	defer func() { hookNewRequest = originNewRequest }()

	// case 2: mock new http request failed
	hookNewRequest = func(fn newReuqest) newReuqest {
		return func(method, url string, body io.Reader) (*http.Request, error) {
			if url == "http://100.100.100.200/latest/meta-data/ram/security-credentials/rolename" {
				return nil, errors.New("new http request failed")
			}
			return http.NewRequest(method, url, body)
		}
	}

	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh Ecs sts token err: new http request failed", err.Error())

	hookNewRequest = originNewRequest

	// case 3
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			if req.URL.Path == "/latest/meta-data/ram/security-credentials/" {
				res = mockResponse(200, "rolename")
				return
			}

			if req.URL.Path == "/latest/meta-data/ram/security-credentials/rolename" {
				err = errors.New("mock server error")
				return
			}
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh Ecs sts token err: mock server error", err.Error())

	// case 4: mock read response error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			if req.URL.Path == "/latest/meta-data/ram/security-credentials/" {
				res = mockResponse(200, "rolename")
				return
			}

			if req.URL.Path == "/latest/meta-data/ram/security-credentials/rolename" {
				status := strconv.Itoa(200)
				res = &http.Response{
					Proto:      "HTTP/1.1",
					ProtoMajor: 1,
					Header:     map[string][]string{},
					StatusCode: 200,
					Status:     status + " " + http.StatusText(200),
				}
				res.Body = ioutil.NopCloser(&errorReader{})
				return
			}
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "read failed", err.Error())

	// case 4: 4xx error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			if req.URL.Path == "/latest/meta-data/ram/security-credentials/" {
				res = mockResponse(200, "rolename")
				return
			}

			if req.URL.Path == "/latest/meta-data/ram/security-credentials/rolename" {
				res = mockResponse(400, "4xx error")
				return
			}
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh Ecs sts token err, httpStatus: 400, message = 4xx error", err.Error())

	// case 5: invalid json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			if req.URL.Path == "/latest/meta-data/ram/security-credentials/" {
				res = mockResponse(200, "rolename")
				return
			}

			if req.URL.Path == "/latest/meta-data/ram/security-credentials/rolename" {
				res = mockResponse(200, "invalid json")
				return
			}
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh Ecs sts token err, json.Unmarshal fail: invalid character 'i' looking for beginning of value", err.Error())

	// case 6: empty response json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			if req.URL.Path == "/latest/meta-data/ram/security-credentials/" {
				res = mockResponse(200, "rolename")
				return
			}

			if req.URL.Path == "/latest/meta-data/ram/security-credentials/rolename" {
				res = mockResponse(200, "null")
				return
			}
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh Ecs sts token err, fail to get credentials", err.Error())

	// case 7: empty session ak response json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			if req.URL.Path == "/latest/meta-data/ram/security-credentials/" {
				res = mockResponse(200, "rolename")
				return
			}

			if req.URL.Path == "/latest/meta-data/ram/security-credentials/rolename" {
				res = mockResponse(200, `{}`)
				return
			}
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh Ecs sts token err, fail to get credentials", err.Error())

	// case 8: non-success response
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			if req.URL.Path == "/latest/meta-data/ram/security-credentials/" {
				res = mockResponse(200, "rolename")
				return
			}

			if req.URL.Path == "/latest/meta-data/ram/security-credentials/rolename" {
				res = mockResponse(200, `{"AccessKeyId":"saki","AccessKeySecret":"saks","Expiration":"2021-10-20T04:27:09Z","SecurityToken":"token","Code":"Failed"}`)
				return
			}
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh Ecs sts token err, Code is not Success", err.Error())

	// case 8: mock ok value
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			if req.URL.Path == "/latest/meta-data/ram/security-credentials/" {
				res = mockResponse(200, "rolename")
				return
			}

			if req.URL.Path == "/latest/meta-data/ram/security-credentials/rolename" {
				res = mockResponse(200, `{"AccessKeyId":"saki","AccessKeySecret":"saks","Expiration":"2021-10-20T04:27:09Z","SecurityToken":"token","Code":"Success"}`)
				return
			}
			err = errors.New("mock server error")
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
}

func TestECSRAMRoleCredentialsProvider_getCredentialsWithMetadataV2(t *testing.T) {
	originDo := hookDo
	defer func() { hookDo = originDo }()

	p, err := NewECSRAMRoleCredentialsProviderBuilder().WithDisableIMDSv1(true).WithRoleName("rolename").Build()
	assert.Nil(t, err)

	// case 1: get metadata token failed
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			err = errors.New("mock server error")
			return
		}
	}

	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "get metadata token failed: mock server error", err.Error())

	// case 2: return token
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			if req.URL.Path == "/latest/api/token" {
				res = mockResponse(200, "tokenxxxxx")
			} else if req.URL.Path == "/latest/meta-data/ram/security-credentials/rolename" {
				assert.Equal(t, "tokenxxxxx", req.Header.Get("x-aliyun-ecs-metadata-token"))
				res = mockResponse(200, `{"AccessKeyId":"saki","AccessKeySecret":"saks","Expiration":"2021-10-20T04:27:09Z","SecurityToken":"token","Code":"Success"}`)
			}
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
}

func TestECSRAMRoleCredentialsProviderGetCredentials(t *testing.T) {
	originDo := hookDo
	defer func() { hookDo = originDo }()

	p := NewECSRAMRoleCredentialsProvider("rolename")
	// case 1: get credentials failed
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			err = errors.New("mock server error")
			return
		}
	}
	_, err := p.GetCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh Ecs sts token err: mock server error", err.Error())

	// case 2: get invalid expiration
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"AccessKeyId":"saki","AccessKeySecret":"saks","Expiration":"invalidexpiration","SecurityToken":"token","Code":"Success"}`)
			return
		}
	}
	_, err = p.GetCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "parsing time \"invalidexpiration\" as \"2006-01-02T15:04:05Z\": cannot parse \"invalidexpiration\" as \"2006\"", err.Error())

	// case 3: happy result
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"AccessKeyId":"akid","AccessKeySecret":"aksecret","Expiration":"2021-10-20T04:27:09Z","SecurityToken":"token","Code":"Success"}`)
			return
		}
	}
	cc, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", cc.AccessKeyId)
	assert.Equal(t, "aksecret", cc.AccessKeySecret)
	assert.Equal(t, "token", cc.SecurityToken)
	assert.True(t, p.needUpdateCredential())
}

func TestECSRAMRoleCredentialsProvider_getMetadataToken(t *testing.T) {
	rollback := internal.Memory("ALIBABA_CLOUD_IMDSV1_DISABLED")
	defer func() {
		rollback()
	}()

	originDo := hookDo
	defer func() { hookDo = originDo }()

	p, err := NewECSRAMRoleCredentialsProviderBuilder().Build()
	assert.Nil(t, err)

	// case 1: server error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			err = errors.New("mock server error")
			return
		}
	}

	_, err = p.getMetadataToken()
	assert.Nil(t, err)

	p, err = NewECSRAMRoleCredentialsProviderBuilder().WithDisableIMDSv1(false).Build()
	assert.Nil(t, err)

	_, err = p.getMetadataToken()
	assert.Nil(t, err)

	os.Setenv("ALIBABA_CLOUD_IMDSV1_DISABLED", "true")
	p, err = NewECSRAMRoleCredentialsProviderBuilder().Build()
	assert.Nil(t, err)

	_, err = p.getMetadataToken()
	assert.NotNil(t, err)

	os.Setenv("ALIBABA_CLOUD_IMDSV1_DISABLED", "")
	p, err = NewECSRAMRoleCredentialsProviderBuilder().Build()
	assert.Nil(t, err)

	_, err = p.getMetadataToken()
	assert.Nil(t, err)

	p, err = NewECSRAMRoleCredentialsProviderBuilder().WithDisableIMDSv1(true).Build()
	assert.Nil(t, err)

	_, err = p.getMetadataToken()
	assert.NotNil(t, err)

	assert.Equal(t, "get metadata token failed: mock server error", err.Error())

	// case 2: return token
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, "tokenxxxxx")
			return
		}
	}
	metadataToken, err := p.getMetadataToken()
	assert.Nil(t, err)
	assert.Equal(t, "tokenxxxxx", metadataToken)

	// case 3: return 404
	p, err = NewECSRAMRoleCredentialsProviderBuilder().WithDisableIMDSv1(false).Build()
	assert.Nil(t, err)

	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(404, "not found")
			return
		}
	}
	metadataToken, err = p.getMetadataToken()
	assert.Nil(t, err)
	assert.Equal(t, "", metadataToken)

	p, err = NewECSRAMRoleCredentialsProviderBuilder().WithDisableIMDSv1(true).Build()
	assert.Nil(t, err)

	metadataToken, err = p.getMetadataToken()
	assert.NotNil(t, err)
	assert.Equal(t, "", metadataToken)
}

func TestNewECSRAMRoleCredentialsProviderWithHttpOptions(t *testing.T) {
	p, err := NewECSRAMRoleCredentialsProviderBuilder().
		WithRoleName("test").
		WithHttpOptions(&HttpOptions{
			ConnectTimeout: 100 * time.Millisecond,
			ReadTimeout:    100 * time.Millisecond,
		}).
		Build()
	assert.Nil(t, err)

	_, err = p.getRoleName()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "timeout")

	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "timeout")

	_, err = p.GetCredentials()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "timeout")
}

func TestOIDCCredentialsProviderGetCredentialsWithError(t *testing.T) {
	wd, _ := os.Getwd()
	p, err := NewOIDCCredentialsProviderBuilder().
		// read a normal token
		WithOIDCTokenFilePath(path.Join(wd, "fixtures/mock_oidctoken")).
		WithOIDCProviderARN("provider-arn").
		WithRoleArn("roleArn").
		WithRoleSessionName("rsn").
		WithPolicy("policy").
		WithDurationSeconds(1000).
		Build()

	assert.Nil(t, err)
	_, err = p.GetCredentials()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "SDK.ServerError")
}

func TestNewOIDCCredentialsProvider(t *testing.T) {
	rollback := internal.Memory("ALIBABA_CLOUD_OIDC_TOKEN_FILE", "ALIBABA_CLOUD_OIDC_PROVIDER_ARN", "ALIBABA_CLOUD_ROLE_ARN", "ALIBABA_CLOUD_STS_REGION", "ALIBABA_CLOUD_VPC_ENDPOINT_ENABLED")
	defer func() {
		rollback()
	}()

	_, err := NewOIDCCredentialsProviderBuilder().Build()
	assert.NotNil(t, err)
	assert.Equal(t, "[SDK.InvalidParam] OIDCTokenFilePath can not be empty", err.Error())

	_, err = NewOIDCCredentialsProviderBuilder().WithOIDCTokenFilePath("/path/to/invalid/oidc.token").Build()
	assert.NotNil(t, err)
	assert.Equal(t, "[SDK.InvalidParam] OIDCProviderARN can not be empty", err.Error())

	_, err = NewOIDCCredentialsProviderBuilder().
		WithOIDCTokenFilePath("/path/to/invalid/oidc.token").
		WithOIDCProviderARN("provider-arn").
		Build()
	assert.NotNil(t, err)
	assert.Equal(t, "[SDK.InvalidParam] RoleArn can not be empty", err.Error())

	p, err := NewOIDCCredentialsProviderBuilder().
		WithOIDCTokenFilePath("/path/to/invalid/oidc.token").
		WithOIDCProviderARN("provider-arn").
		WithRoleArn("roleArn").
		Build()
	assert.Nil(t, err)

	assert.Equal(t, "/path/to/invalid/oidc.token", p.oidcTokenFilePath)
	assert.True(t, strings.HasPrefix(p.roleSessionName, "aliyun-go-sdk-"))
	assert.Equal(t, 3600, p.durationSeconds)

	_, err = NewOIDCCredentialsProviderBuilder().
		WithOIDCTokenFilePath("/path/to/invalid/oidc.token").
		WithOIDCProviderARN("provider-arn").
		WithRoleArn("roleArn").
		WithDurationSeconds(100).
		Build()
	assert.NotNil(t, err)
	assert.Equal(t, "[SDK.InvalidParam] Assume Role session duration should be in the range of 15min - 1hr", err.Error())

	os.Setenv("ALIBABA_CLOUD_OIDC_TOKEN_FILE", "/path/from/env")
	os.Setenv("ALIBABA_CLOUD_OIDC_PROVIDER_ARN", "provider_arn_from_env")
	os.Setenv("ALIBABA_CLOUD_ROLE_ARN", "role_arn_from_env")

	p, err = NewOIDCCredentialsProviderBuilder().
		Build()
	assert.Nil(t, err)

	assert.Equal(t, "/path/from/env", p.oidcTokenFilePath)
	assert.Equal(t, "provider_arn_from_env", p.oidcProviderARN)
	assert.Equal(t, "role_arn_from_env", p.roleArn)
	// sts endpoint: default
	assert.Equal(t, "sts.aliyuncs.com", p.stsEndpoint)

	// sts endpoint: with sts endpoint env
	os.Setenv("ALIBABA_CLOUD_STS_REGION", "cn-hangzhou")
	os.Setenv("ALIBABA_CLOUD_VPC_ENDPOINT_ENABLED", "true")
	p, err = NewOIDCCredentialsProviderBuilder().
		Build()
	assert.Nil(t, err)
	assert.Equal(t, "sts-vpc.cn-hangzhou.aliyuncs.com", p.stsEndpoint)

	// sts endpoint: with sts endpoint
	p, err = NewOIDCCredentialsProviderBuilder().
		WithSTSEndpoint("sts.cn-shanghai.aliyuncs.com").
		Build()
	assert.Nil(t, err)
	assert.Equal(t, "sts.cn-shanghai.aliyuncs.com", p.stsEndpoint)

	// sts endpoint: with sts regionId
	p, err = NewOIDCCredentialsProviderBuilder().
		WithStsRegion("cn-beijing").
		WithEnableVpc(true).
		Build()
	assert.Nil(t, err)
	assert.Equal(t, "sts-vpc.cn-beijing.aliyuncs.com", p.stsEndpoint)

	os.Setenv("ALIBABA_CLOUD_VPC_ENDPOINT_ENABLED", "false")
	p, err = NewOIDCCredentialsProviderBuilder().
		WithOIDCTokenFilePath("/path/to/invalid/oidc.token").
		WithOIDCProviderARN("provider-arn").
		WithRoleArn("roleArn").
		WithRoleSessionName("rsn").
		WithStsRegion("cn-hangzhou").
		WithPolicy("policy").
		Build()
	assert.Nil(t, err)

	assert.Equal(t, "/path/to/invalid/oidc.token", p.oidcTokenFilePath)
	assert.Equal(t, "provider-arn", p.oidcProviderARN)
	assert.Equal(t, "roleArn", p.roleArn)
	assert.Equal(t, "rsn", p.roleSessionName)
	assert.Equal(t, "cn-hangzhou", p.stsRegion)
	assert.Equal(t, "sts.cn-hangzhou.aliyuncs.com", p.stsEndpoint)
	assert.Equal(t, "policy", p.policy)
	assert.Equal(t, 3600, p.durationSeconds)
}

func TestOIDCCredentialsProvider_getCredentials(t *testing.T) {
	// case 0: invalid oidc token file path
	p, err := NewOIDCCredentialsProviderBuilder().
		WithOIDCTokenFilePath("/path/to/invalid/oidc.token").
		WithOIDCProviderARN("provider-arn").
		WithRoleArn("roleArn").
		WithRoleSessionName("rsn").
		WithStsRegion("cn-hangzhou").
		WithPolicy("policy").
		Build()
	assert.Nil(t, err)

	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "open /path/to/invalid/oidc.token: no such file or directory", err.Error())

	// case 1: mock new http request failed
	wd, _ := os.Getwd()
	p, err = NewOIDCCredentialsProviderBuilder().
		// read a normal token
		WithOIDCTokenFilePath(path.Join(wd, "fixtures/mock_oidctoken")).
		WithOIDCProviderARN("provider-arn").
		WithRoleArn("roleArn").
		WithRoleSessionName("rsn").
		WithStsRegion("cn-hangzhou").
		WithPolicy("policy").
		Build()
	assert.Nil(t, err)

	originNewRequest := hookNewRequest
	defer func() { hookNewRequest = originNewRequest }()

	hookNewRequest = func(fn newReuqest) newReuqest {
		return func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("new http request failed")
		}
	}

	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "new http request failed", err.Error())

	// reset new request
	hookNewRequest = originNewRequest

	originDo := hookDo
	defer func() { hookDo = originDo }()

	// case 2: server error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			err = errors.New("mock server error")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "mock server error", err.Error())

	// case 3: mock read response error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			status := strconv.Itoa(200)
			res = &http.Response{
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				Header:     map[string][]string{},
				StatusCode: 200,
				Status:     status + " " + http.StatusText(200),
			}
			res.Body = ioutil.NopCloser(&errorReader{})
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "read failed", err.Error())

	// case 4: 4xx error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(400, "4xx error")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "SDK.ServerError\nErrorCode: \nRecommend: get session token failed\nRequestId: \nMessage: 4xx error\nRespHeaders: map[]", err.Error())

	// case 5: invalid json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, "invalid json")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "get oidc sts token err, json.Unmarshal fail: invalid character 'i' looking for beginning of value", err.Error())

	// case 6: empty response json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, "null")
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "get oidc sts token err, fail to get credentials", err.Error())

	// case 7: empty session ak response json
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"Credentials": {}}`)
			return
		}
	}
	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "refresh RoleArn sts token err, fail to get credentials", err.Error())

	// case 8: mock ok value
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"Credentials": {"AccessKeyId":"saki","AccessKeySecret":"saks","Expiration":"2021-10-20T04:27:09Z","SecurityToken":"token"}}`)
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
}

func TestOIDCCredentialsProvider_getCredentialsWithRequestCheck(t *testing.T) {
	originDo := hookDo
	defer func() { hookDo = originDo }()

	// case 1: mock new http request failed
	wd, _ := os.Getwd()
	p, err := NewOIDCCredentialsProviderBuilder().
		// read a normal token
		WithOIDCTokenFilePath(path.Join(wd, "fixtures/mock_oidctoken")).
		WithOIDCProviderARN("provider-arn").
		WithRoleArn("roleArn").
		WithRoleSessionName("rsn").
		WithPolicy("policy").
		WithDurationSeconds(1000).
		Build()

	assert.Nil(t, err)

	// case 1: server error
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			assert.Equal(t, "sts.aliyuncs.com", req.Host)
			assert.Contains(t, req.URL.String(), "Action=AssumeRoleWithOIDC")
			body, err := ioutil.ReadAll(req.Body)
			assert.Nil(t, err)
			bodyString := string(body)
			assert.Contains(t, bodyString, "Policy=policy")
			assert.Contains(t, bodyString, "RoleArn=roleArn")
			assert.Contains(t, bodyString, "RoleSessionName=rsn")
			assert.Contains(t, bodyString, "DurationSeconds=1000")

			err = errors.New("mock server error")
			return
		}
	}

	_, err = p.getCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "mock server error", err.Error())
}

func TestOIDCCredentialsProviderGetCredentials(t *testing.T) {
	originDo := hookDo
	defer func() { hookDo = originDo }()

	// case 1: mock new http request failed
	wd, _ := os.Getwd()
	p, err := NewOIDCCredentialsProviderBuilder().
		// read a normal token
		WithOIDCTokenFilePath(path.Join(wd, "fixtures/mock_oidctoken")).
		WithOIDCProviderARN("provider-arn").
		WithRoleArn("roleArn").
		WithRoleSessionName("rsn").
		WithPolicy("policy").
		WithDurationSeconds(1000).
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
			res = mockResponse(200, `{"Credentials": {"AccessKeyId":"akid","AccessKeySecret":"aksecret","Expiration":"invalidexpiration","SecurityToken":"ststoken"}}`)
			return
		}
	}
	_, err = p.GetCredentials()
	assert.NotNil(t, err)
	assert.Equal(t, "parsing time \"invalidexpiration\" as \"2006-01-02T15:04:05Z\": cannot parse \"invalidexpiration\" as \"2006\"", err.Error())

	// case 3: happy result
	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"Credentials": {"AccessKeyId":"akid","AccessKeySecret":"aksecret","Expiration":"2021-10-20T04:27:09Z","SecurityToken":"ststoken"}}`)
			return
		}
	}
	cc, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", cc.AccessKeyId)
	assert.Equal(t, "aksecret", cc.AccessKeySecret)
	assert.Equal(t, "ststoken", cc.SecurityToken)
	assert.True(t, p.needUpdateCredential())
}

func TestOIDCCredentialsProviderGetCredentialsWithHttpOptions(t *testing.T) {
	wd, _ := os.Getwd()
	p, err := NewOIDCCredentialsProviderBuilder().
		// read a normal token
		WithOIDCTokenFilePath(path.Join(wd, "fixtures/mock_oidctoken")).
		WithOIDCProviderARN("provider-arn").
		WithRoleArn("roleArn").
		WithRoleSessionName("rsn").
		WithPolicy("policy").
		WithDurationSeconds(1000).
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
