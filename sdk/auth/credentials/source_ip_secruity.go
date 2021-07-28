package credentials

type SourceCredential struct {
	AccessKeyId       string
	AccessKeySecret   string
	SourceIp          string
	SecurityTransport string
}

func NewSourceCredential(accessKeyId, accessKeySecret, sourceIp, securityTransport string) *SourceCredential {
	return &SourceCredential{
		SourceIp:          sourceIp,
		SecurityTransport: securityTransport,
		AccessKeyId:       accessKeyId,
		AccessKeySecret:   accessKeySecret,
	}
}
