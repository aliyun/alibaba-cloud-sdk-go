package credentials

type StsCredential struct {
	AccessKeyId       string
	AccessKeySecret   string
	AccessKeyStsToken string
}

func NewStsCredential(AccessKeyId, AccessKeySecret, AccessKeyStsToken string) *StsCredential {
	return &StsCredential{
		AccessKeyId:       AccessKeyId,
		AccessKeySecret:   AccessKeySecret,
		AccessKeyStsToken: AccessKeyStsToken,
	}
}
