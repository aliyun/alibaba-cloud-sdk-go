package credentials

type StsRoleArnCredential struct {
	AccessKeyId           string
	AccessKeySecret       string
	RoleArn               string
	RoleSessionName       string
	RoleSessionExpiration int
}
