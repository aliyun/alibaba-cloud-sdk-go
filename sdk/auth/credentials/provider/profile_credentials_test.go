package provider_test

import (
	"os"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"

	"github.com/stretchr/testify/assert"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials/provider"
)

var inistr = `
[default]              
enable = true                    
type = access_key                
access_key_id = foo               
access_key_secret = bar            
                   
[client1]                         
type = ecs_ram_role                
role_name = EcsRamRoleTest       

[client2]                                         
type = ram_role_arn                
access_key_id = foo
access_key_secret = bar
role_arn = role_arn
role_session_name = session_name

[client3]                          
type = bearer_token                
bearer_token = bearer_token        

[client4]                          
type = rsa_key_pair               
public_key_id = publicKeyId       
private_key_file = ./pk.pem
`
var privatekey = `this is privatekey`

func TestProfileProvider(t *testing.T) {
	p := provider.NewProfileProvider()
	value, ok := p.(*provider.ProfileProvider)
	assert.True(t, ok)
	assert.Equal(t, value.Profile, "default")

	p = provider.NewProfileProvider("first")
	value, ok = p.(*provider.ProfileProvider)
	assert.True(t, ok)
	assert.Equal(t, value.Profile, "first")

	c, err := p.Resolve()
	assert.Nil(t, c)
	assert.Nil(t, err)

	os.Setenv(provider.ENVCredentialFile, "./credentials")

	file, err := os.Create("./credentials")
	assert.Nil(t, err)
	file.WriteString(inistr)
	file.Close()
	defer os.Remove("./credentials")

	p = provider.NewProfileProvider()
	c, err = p.Resolve()
	assert.Equal(t, credentials.NewAccessKeyCredential("foo", "bar"), c)
	assert.Nil(t, err)

	p = provider.NewProfileProvider("client1")
	c, err = p.Resolve()
	assert.Equal(t, credentials.NewEcsRamRoleCredential("EcsRamRoleTest"), c)
	assert.Nil(t, err)

	p = provider.NewProfileProvider("client2")
	c, err = p.Resolve()
	assert.Equal(t, credentials.NewRamRoleArnCredential("foo", "bar", "role_arn", "session_name", 3600), c)
	assert.Nil(t, err)

	file, err = os.Create("./pk.pem")
	assert.Nil(t, err)
	file.WriteString(privatekey)
	file.Close()

	p = provider.NewProfileProvider("client4")
	c, err = p.Resolve()
	assert.Equal(t, credentials.NewRsaKeyPairCredential("", "publicKeyId", 3600), c)
	assert.Nil(t, err)
	defer os.Remove(`./pk.pem`)

	// p = provider.NewProfileProvider("client3")
	// c, err = p.Resolve()
	// assert.Equal(t, credentials.NewBearerCredential("bearer_token"), c)
	// assert.Nil(t, err)
}
