package credentials

type RsaKeyPairCredential struct {
	PrivateKey        string
	PublicKeyId       string
	SessionExpiration int
}

func NewRsaKeyPairCredential(PrivateKey, PublicKeyId string, SessionExpiration int) *RsaKeyPairCredential {
	return &RsaKeyPairCredential{
		PrivateKey:        PrivateKey,
		PublicKeyId:       PublicKeyId,
		SessionExpiration: SessionExpiration,
	}
}
