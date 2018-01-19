package credentials

type StsRoleArnCredential struct {
	AccessKeyId           string
	AccessKeySecret       string
	RoleArn               string
	RoleSessionName       string
	RoleSessionExpiration int
}

func NewStsRoleArnCredential(AccessKeyId, AccessKeySecret, RoleArn, RoleSessionName string, RoleSessionExpiration int) *StsRoleArnCredential {
	return &StsRoleArnCredential{
		AccessKeyId:           AccessKeyId,
		AccessKeySecret:       AccessKeySecret,
		RoleArn:               RoleArn,
		RoleSessionName:       RoleSessionName,
		RoleSessionExpiration: RoleSessionExpiration,
	}
}
