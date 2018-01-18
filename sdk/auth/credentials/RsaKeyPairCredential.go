package credentials

type RsaKeyPairCredential struct {
	PrivateKey        string
	PublicKeyId       string
	SessionExpiration int
}
