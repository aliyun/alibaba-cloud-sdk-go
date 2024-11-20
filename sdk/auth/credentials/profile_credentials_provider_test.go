package credentials

import (
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/internal"
	"github.com/stretchr/testify/assert"
	"gopkg.in/ini.v1"
)

var inistr = `
[default]
enable = true
type = access_key
access_key_id = foo
access_key_secret = bar

[notype]
access_key_id = foo
access_key_secret = bar

[noak]
type = access_key
access_key_secret = bar

[emptyak]
type = access_key
access_key_id =
access_key_secret = bar

[ecs]
type = ecs_ram_role
role_name = EcsRamRoleTest

[noecs]
type = ecs_ram_role

[emptyecs]
type = ecs_ram_role
role_name =

[ram]
type = ram_role_arn
access_key_id = foo
access_key_secret = bar
role_arn = role_arn
role_session_name = session_name

[noram]
type = ram_role_arn
access_key_secret = bar
role_arn = role_arn
role_session_name = session_name

[emptyram]
type = ram_role_arn
access_key_id =
access_key_secret = bar
role_arn = role_arn
role_session_name = session_name

[rsa]
type = rsa_key_pair
public_key_id = publicKeyId
private_key_file = ./pk.pem

[norsa]
type = rsa_key_pair
public_key_id = publicKeyId

[emptyrsa]
type = rsa_key_pair
public_key_id = publicKeyId
private_key_file =

[error_rsa]
type = rsa_key_pair
public_key_id = publicKeyId
private_key_file = ./pk_error.pem

[error_type]
type = error_type
public_key_id = publicKeyId
private_key_file = ./pk_error.pem
`

func TestProfileCredentialsProviderBuilder(t *testing.T) {
	rollback := internal.Memory("ALIBABA_CLOUD_PROFILE")
	defer rollback()

	// profile name from specified
	provider := NewProfileCredentialsProviderBuilder().WithProfileName("custom").Build()
	assert.Equal(t, "custom", provider.profileName)

	// profile name from env
	os.Setenv("ALIBABA_CLOUD_PROFILE", "profile_from_env")
	provider = NewProfileCredentialsProviderBuilder().Build()

	assert.Equal(t, "profile_from_env", provider.profileName)

	// profile name from default
	os.Setenv("ALIBABA_CLOUD_PROFILE", "")
	provider = NewProfileCredentialsProviderBuilder().Build()
	assert.Equal(t, "default", provider.profileName)
}

func TestProfileCredentialsProvider_getCredentialsProvider(t *testing.T) {
	provider := NewProfileCredentialsProviderBuilder().WithProfileName("custom").Build()
	_, err := provider.getCredentialsProvider(ini.Empty())
	assert.NotNil(t, err)
	assert.EqualError(t, err, "ERROR: Can not load sectionsection \"custom\" does not exist")

	file, err := ini.Load([]byte(inistr))
	assert.Nil(t, err)
	assert.NotNil(t, file)

	// no type
	provider = NewProfileCredentialsProviderBuilder().WithProfileName("notype").Build()
	_, err = provider.getCredentialsProvider(file)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "ERROR: Can not find credential typeerror when getting key of section \"notype\": key \"type\" not exists")

	// no ak
	provider = NewProfileCredentialsProviderBuilder().WithProfileName("noak").Build()
	_, err = provider.getCredentialsProvider(file)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "ERROR: Failed to get value")

	// value is empty
	provider = NewProfileCredentialsProviderBuilder().WithProfileName("emptyak").Build()
	_, err = provider.getCredentialsProvider(file)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "ERROR: Value can't be empty")

	// static ak provider
	provider = NewProfileCredentialsProviderBuilder().Build()
	cp, err := provider.getCredentialsProvider(file)
	assert.Nil(t, err)
	akcp, ok := cp.(*StaticAKCredentialsProvider)
	assert.True(t, ok)
	cc, err := akcp.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, &Credentials{
		AccessKeyId:     "foo",
		AccessKeySecret: "bar",
		SecurityToken:   "",
		BearerToken:     "",
		ProviderName:    "static_ak"},
		cc)

	// ecs_ram_role without rolename
	provider = NewProfileCredentialsProviderBuilder().WithProfileName("noecs").Build()
	_, err = provider.getCredentialsProvider(file)
	assert.EqualError(t, err, "ERROR: Failed to get value")

	// ecs_ram_role with rolename
	provider = NewProfileCredentialsProviderBuilder().WithProfileName("ecs").Build()
	cp, err = provider.getCredentialsProvider(file)
	assert.Nil(t, err)
	_, ok = cp.(*ECSRAMRoleCredentialsProvider)
	assert.True(t, ok)

	// ram role arn without keys
	provider = NewProfileCredentialsProviderBuilder().WithProfileName("noram").Build()
	_, err = provider.getCredentialsProvider(file)
	assert.EqualError(t, err, "ERROR: Failed to get value")

	// ram role arn without values
	provider = NewProfileCredentialsProviderBuilder().WithProfileName("emptyram").Build()
	_, err = provider.getCredentialsProvider(file)
	assert.EqualError(t, err, "ERROR: Value can't be empty")

	// normal ram role arn
	provider = NewProfileCredentialsProviderBuilder().WithProfileName("ram").Build()
	cp, err = provider.getCredentialsProvider(file)
	assert.Nil(t, err)
	_, ok = cp.(*RAMRoleARNCredentialsProvider)
	assert.True(t, ok)

	// unsupported type
	provider = NewProfileCredentialsProviderBuilder().WithProfileName("error_type").Build()
	_, err = provider.getCredentialsProvider(file)
	assert.EqualError(t, err, "ERROR: Failed to get credential")
}

func TestProfileCredentialsProviderGetCredentials(t *testing.T) {
	originDo := hookDo
	defer func() { hookDo = originDo }()
	rollback := internal.Memory("ALIBABA_CLOUD_CREDENTIALS_FILE")
	defer func() {
		getHomePath = internal.GetHomePath
		rollback()
	}()

	// testcase: empty home
	getHomePath = func() string {
		return ""
	}
	provider := NewProfileCredentialsProviderBuilder().WithProfileName("custom").Build()
	_, err := provider.GetCredentials()
	assert.EqualError(t, err, "cannot found home dir")

	// testcase: invalid home
	getHomePath = func() string {
		return "/path/invalid/home/dir"
	}

	provider = NewProfileCredentialsProviderBuilder().WithProfileName("custom").Build()
	_, err = provider.GetCredentials()
	assert.EqualError(t, err, "ERROR: Can not open fileopen /path/invalid/home/dir/.alibabacloud/credentials: no such file or directory")

	// testcase: specify credentials file with env
	os.Setenv("ALIBABA_CLOUD_CREDENTIALS_FILE", "/path/to/credentials.invalid")
	provider = NewProfileCredentialsProviderBuilder().WithProfileName("custom").Build()
	_, err = provider.GetCredentials()
	assert.EqualError(t, err, "ERROR: Can not open fileopen /path/to/credentials.invalid: no such file or directory")
	os.Unsetenv("ALIBABA_CLOUD_CREDENTIALS_FILE")

	// get from credentials file
	getHomePath = func() string {
		wd, _ := os.Getwd()
		return path.Join(wd, "fixtures")
	}

	provider = NewProfileCredentialsProviderBuilder().WithProfileName("custom").Build()
	_, err = provider.GetCredentials()
	assert.EqualError(t, err, "ERROR: Can not load sectionsection \"custom\" does not exist")

	provider = NewProfileCredentialsProviderBuilder().Build()
	cc, err := provider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, &Credentials{
		AccessKeyId:     "foo",
		AccessKeySecret: "bar",
		SecurityToken:   "",
		BearerToken:     "",
		ProviderName:    "profile/static_ak",
	}, cc)

	// get credentials again
	cc, err = provider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, &Credentials{
		AccessKeyId:     "foo",
		AccessKeySecret: "bar",
		SecurityToken:   "",
		BearerToken:     "",
		ProviderName:    "profile/static_ak",
	}, cc)

	// the profile ram with invalid ak
	provider = NewProfileCredentialsProviderBuilder().WithProfileName("ram").Build()
	_, err = provider.GetCredentials()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "InvalidAccessKeyId.NotFound")

	hookDo = func(fn do) do {
		return func(req *http.Request) (res *http.Response, err error) {
			res = mockResponse(200, `{"Credentials": {"AccessKeyId":"akid","AccessKeySecret":"aksecret","Expiration":"2021-10-20T04:27:09Z","SecurityToken":"ststoken"}}`)
			return
		}
	}
	provider = NewProfileCredentialsProviderBuilder().WithProfileName("ram").Build()
	cc, err = provider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akid", cc.AccessKeyId)
	assert.Equal(t, "aksecret", cc.AccessKeySecret)
	assert.Equal(t, "ststoken", cc.SecurityToken)
	assert.Equal(t, "profile/ram_role_arn/static_ak", cc.ProviderName)

	provider.innerProvider = new(testProvider)
	cc, err = provider.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "test", cc.AccessKeyId)
	assert.Equal(t, "test", cc.AccessKeySecret)
	assert.Equal(t, "profile/test", cc.ProviderName)

	provider.innerProvider = new(testErrorProvider)
	_, err = provider.GetCredentials()
	assert.Equal(t, "error", err.Error())
}
