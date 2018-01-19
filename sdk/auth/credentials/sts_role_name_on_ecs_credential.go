package credentials

type StsRoleNameOnEcsCredential struct {
	RoleName string
}

func NewStsRoleNameOnEcsCredential(RoleName string) *StsRoleNameOnEcsCredential {
	return &StsRoleNameOnEcsCredential{
		RoleName: RoleName,
	}
}
