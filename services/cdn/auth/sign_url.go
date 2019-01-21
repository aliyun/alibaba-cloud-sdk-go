package auth

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"time"
)

// An URLSigner provides URL signing utilities to sign URLs for Aliyun CDN
// resources.
//
// The signer is safe to use concurrently.
type URLSigner struct {
	authType string
	privKey  string
}

// newURLSigner returns a new signer object.
func newURLSigner(authType string, privKey string) *URLSigner {
	return &URLSigner{
		authType: authType,
		privKey:  privKey,
	}
}

// Sign returns a signed aliyuncdn url based on authentication type
func (s URLSigner) Sign(baseURL string, path string, expires time.Time) (string, error) {
	r, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("unable to parse baseurl: %s", r)
	}

	switch s.authType {
	case "a":
		return aTypeSign(r, path, s.privKey, expires), nil
	case "b":
		return bTypeSign(r, path, s.privKey, expires), nil
	case "c":
		return cTypeSign(r, path, s.privKey, expires), nil
	default:
		return "", fmt.Errorf("invalid authentication type")
	}
}

// sign by A type authentication method.
func aTypeSign(r *url.URL, path string, privateKey string, expires time.Time) string {
	//rand is a random uuid without "-"
	rand := GenerateUUID().String()
	// not use, "0" by default
	uid := "0"
	secret := fmt.Sprintf("%s-%d-%s-%s-%s", path, expires.Unix(), rand, uid, privateKey)
	hashValue := md5.Sum([]byte(secret))
	authKey := fmt.Sprintf("%d-%s-%s-%x", expires.Unix(), rand, uid, hashValue)
	u := &url.URL{Scheme: r.Scheme, Host: r.Host, Path: path}
	signURL := fmt.Sprintf("%s?auth_key=%s", u.String(), authKey)
	return signURL
}

// sign by B type authentication method.
func bTypeSign(r *url.URL, path string, privateKey string, expires time.Time) string {
	formatExp := expires.Format("200601021504")
	secret := privateKey + formatExp + path
	hashValue := md5.Sum([]byte(secret))
	signURL := fmt.Sprintf("%s%s/%x%s", r.String(), formatExp, hashValue, path)
	return signURL
}

// sign by C type authentication method.
func cTypeSign(r *url.URL, path string, privateKey string, expires time.Time) string {
	hexExp := fmt.Sprintf("%x", expires.Unix())
	secret := privateKey + path + hexExp
	hashValue := md5.Sum([]byte(secret))
	signURL := fmt.Sprintf("%s%x/%s%s", r.String(), hashValue, hexExp, path)
	return signURL
}
