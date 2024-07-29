package credentials

// Deprecated: The RsaKeyPair is no more supported now.
type RsaKeyPairCredential struct {
	PrivateKey        string
	PublicKeyId       string
	SessionExpiration int
}

func NewRsaKeyPairCredential(privateKey, publicKeyId string, sessionExpiration int) *RsaKeyPairCredential {
	return &RsaKeyPairCredential{
		PrivateKey:        privateKey,
		PublicKeyId:       publicKeyId,
		SessionExpiration: sessionExpiration,
	}
}
