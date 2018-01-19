package credentials

type BaseCredential struct {
	AccessKeyId     string
	AccessKeySecret string
}

func NewBaseCredential(AccessKeyId, AccessKeySecret string) *BaseCredential {
	return &BaseCredential{
		AccessKeyId:     AccessKeyId,
		AccessKeySecret: AccessKeySecret,
	}
}
