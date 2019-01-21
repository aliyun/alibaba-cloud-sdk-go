package auth

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/docker/distribution/uuid"
)

// sign by A type authentication method.
func aTypeSign(r *url.URL, path string, privateKey string, expires time.Time) string {
	//random uuid
	rand := strings.Replace(uuid.Generate().String(), "-", "", -1)
	// "0" by default, other value is ok
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
