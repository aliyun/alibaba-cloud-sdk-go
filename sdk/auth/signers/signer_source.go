package signers

import "github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"

type SourceCredentialSigner struct {
	credential *credentials.SourceCredential
}

func NewSourceSigner(credential *credentials.SourceCredential) *SourceCredentialSigner {
	return &SourceCredentialSigner{
		credential: credential,
	}
}

func (signer *SourceCredentialSigner) GetExtraParam() map[string]string {
	return nil
}
func (*SourceCredentialSigner) GetName() string {
	return "HMAC-SHA1"
}

func (*SourceCredentialSigner) GetType() string {
	return ""
}

func (*SourceCredentialSigner) GetVersion() string {
	return "1.0"
}

func (signer *SourceCredentialSigner) GetAccessKeyId() (accessKeyId string, err error) {
	return signer.credential.AccessKeyId, nil
}

func (signer *SourceCredentialSigner) Sign(stringToSign, secretSuffix string) string {
	secret := signer.credential.AccessKeySecret + secretSuffix
	return ShaHmac1(stringToSign, secret)
}
