package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils"
)

type assumedRoleUser struct {
}

type credentials struct {
	SecurityToken   *string `json:"SecurityToken"`
	Expiration      *string `json:"Expiration"`
	AccessKeySecret *string `json:"AccessKeySecret"`
	AccessKeyId     *string `json:"AccessKeyId"`
}

type ecsRAMRoleCredentials struct {
	SecurityToken   *string `json:"SecurityToken"`
	Expiration      *string `json:"Expiration"`
	AccessKeySecret *string `json:"AccessKeySecret"`
	AccessKeyId     *string `json:"AccessKeyId"`
	LastUpdated     *string `json:"LastUpdated"`
	Code            *string `json:"Code"`
}

type assumeRoleResponse struct {
	RequestID       *string          `json:"RequestId"`
	AssumedRoleUser *assumedRoleUser `json:"AssumedRoleUser"`
	Credentials     *credentials     `json:"Credentials"`
}

type SessionCredentials struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
	Expiration      string
}

type Credentials struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
	BearerToken     string
}

type CredentialsProvider interface {
	GetCredentials() (cc *Credentials, err error)
}

type StaticAKCredentialsProvider struct {
	AccessKeyId     string
	AccessKeySecret string
}

func NewStaticAKCredentialsProvider(accessKeyId, accessKeySecret string) *StaticAKCredentialsProvider {
	return &StaticAKCredentialsProvider{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
}

func (provider *StaticAKCredentialsProvider) GetCredentials() (cc *Credentials, err error) {
	cc = &Credentials{
		AccessKeyId:     provider.AccessKeyId,
		AccessKeySecret: provider.AccessKeySecret,
	}
	return
}

type StaticSTSCredentialsProvider struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}

func NewStaticSTSCredentialsProvider(accessKeyId, accessKeySecret, securityToken string) *StaticSTSCredentialsProvider {
	return &StaticSTSCredentialsProvider{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
	}
}

func (provider *StaticSTSCredentialsProvider) GetCredentials() (cc *Credentials, err error) {
	cc = &Credentials{
		AccessKeyId:     provider.AccessKeyId,
		AccessKeySecret: provider.AccessKeySecret,
		SecurityToken:   provider.SecurityToken,
	}
	return
}

type BearerTokenCredentialsProvider struct {
	BearerToken string
}

func NewBearerTokenCredentialsProvider(bearerToken string) *BearerTokenCredentialsProvider {
	return &BearerTokenCredentialsProvider{
		BearerToken: bearerToken,
	}
}

func (provider *BearerTokenCredentialsProvider) GetCredentials() (cc *Credentials, err error) {
	cc = &Credentials{
		BearerToken: provider.BearerToken,
	}
	return
}

type RSAKeyPairCredentialsProvider struct {
	PublicKeyId       string
	PrivateKeyId      string
	sessionExpiration int
}

func NewRSAKeyPairCredentialsProvider(publicKeyId, privateKeyId string, sessionExpiration int) *RSAKeyPairCredentialsProvider {
	return &RSAKeyPairCredentialsProvider{
		PublicKeyId:       publicKeyId,
		PrivateKeyId:      privateKeyId,
		sessionExpiration: sessionExpiration,
	}
}

func (provider *RSAKeyPairCredentialsProvider) GetCredentials() (cc *Credentials, err error) {
	cc = &Credentials{
		// TODO:
	}
	return
}

type RAMRoleARNCredentialsProvider struct {
	credentialsProvider   CredentialsProvider
	RoleArn               string
	RoleSessionName       string
	RoleSessionExpiration int
	Policy                string
	StsRegion             string
	ExternalId            string
	lastUpdateTimestamp   int64
	sessionCredentials    *SessionCredentials
}

func NewRAMRoleARNCredentialsProvider(credentialsProvider CredentialsProvider, roleArn, roleSessionName string, roleSessionExpiration int, policy, stsRegion, externalId string) (provider *RAMRoleARNCredentialsProvider, err error) {
	provider = &RAMRoleARNCredentialsProvider{
		credentialsProvider:   credentialsProvider,
		RoleArn:               roleArn,
		RoleSessionExpiration: roleSessionExpiration,
		Policy:                policy,
		StsRegion:             stsRegion,
		ExternalId:            externalId,
	}

	if len(roleSessionName) > 0 {
		provider.RoleSessionName = roleSessionName
	} else {
		provider.RoleSessionName = "aliyun-go-sdk-" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)
	}

	if roleSessionExpiration > 0 {
		if roleSessionExpiration >= 900 && roleSessionExpiration <= 3600 {
			provider.RoleSessionExpiration = roleSessionExpiration
		} else {
			err = errors.NewClientError(errors.InvalidParamErrorCode, "Assume Role session duration should be in the range of 15min - 1Hr", nil)
		}
	} else {
		// default to 3600
		provider.RoleSessionExpiration = 3600
	}

	return
}

func (provider *RAMRoleARNCredentialsProvider) getCredentials(cc *Credentials) (sessionCredentials *SessionCredentials, err error) {
	method := "POST"
	var host string
	if provider.StsRegion != "" {
		host = fmt.Sprintf("sts.%s.aliyuncs.com", provider.StsRegion)
	} else {
		host = "sts.aliyuncs.com"
	}

	queries := make(map[string]string)
	queries["Version"] = "2015-04-01"
	queries["Action"] = "AssumeRole"
	queries["Format"] = "JSON"
	queries["Timestamp"] = utils.GetTimeInFormatISO8601()
	queries["SignatureMethod"] = "HMAC-SHA1"
	queries["SignatureVersion"] = "1.0"
	queries["SignatureNonce"] = utils.GetNonce()
	queries["AccessKeyId"] = cc.AccessKeyId
	if cc.SecurityToken != "" {
		queries["SecurityToken"] = cc.SecurityToken
	}

	bodyForm := make(map[string]string)
	bodyForm["RoleArn"] = provider.RoleArn
	if provider.Policy != "" {
		bodyForm["Policy"] = provider.Policy
	}
	if provider.ExternalId != "" {
		bodyForm["ExternalId"] = provider.ExternalId
	}
	bodyForm["RoleSessionName"] = provider.RoleSessionName
	bodyForm["DurationSeconds"] = strconv.Itoa(provider.RoleSessionExpiration)

	// caculate signature
	signParams := make(map[string]string)
	for key, value := range queries {
		signParams[key] = value
	}
	for key, value := range bodyForm {
		signParams[key] = value
	}

	stringToSign := utils.GetUrlFormedMap(signParams)
	stringToSign = strings.Replace(stringToSign, "+", "%20", -1)
	stringToSign = strings.Replace(stringToSign, "*", "%2A", -1)
	stringToSign = strings.Replace(stringToSign, "%7E", "~", -1)
	stringToSign = url.QueryEscape(stringToSign)
	stringToSign = method + "&%2F&" + stringToSign
	secret := cc.AccessKeySecret + "&"
	queries["Signature"] = utils.ShaHmac1(stringToSign, secret)

	querystring := utils.GetUrlFormedMap(queries)
	// do request
	httpUrl := fmt.Sprintf("https://%s/?%s", host, querystring)

	body := utils.GetUrlFormedMap(bodyForm)

	httpRequest, err := http.NewRequest(method, httpUrl, strings.NewReader(body))
	if err != nil {
		return
	}

	// set headers
	httpRequest.Header["Accept-Encoding"] = []string{"identity"}
	httpRequest.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	httpClient := &http.Client{}

	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return
	}

	defer httpResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}

	if httpResponse.StatusCode != http.StatusOK {
		message := "refresh session token failed"
		err = errors.NewServerError(httpResponse.StatusCode, string(responseBody), message)
		return
	}
	var data assumeRoleResponse
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		err = fmt.Errorf("refresh RoleArn sts token err, json.Unmarshal fail: %s", err.Error())
		return
	}
	if data.Credentials == nil {
		err = fmt.Errorf("refresh RoleArn sts token err, fail to get credentials")
		return
	}

	if data.Credentials.AccessKeyId == nil || data.Credentials.AccessKeySecret == nil || data.Credentials.SecurityToken == nil {
		err = fmt.Errorf("refresh RoleArn sts token err, fail to get credentials")
		return
	}

	sessionCredentials = &SessionCredentials{
		AccessKeyId:     *data.Credentials.AccessKeyId,
		AccessKeySecret: *data.Credentials.AccessKeySecret,
		SecurityToken:   *data.Credentials.SecurityToken,
		Expiration:      *data.Credentials.Expiration,
	}
	return
}

func (provider *RAMRoleARNCredentialsProvider) needUpdateCredential() (result bool) {
	return time.Now().Unix()-provider.lastUpdateTimestamp >= int64(provider.RoleSessionExpiration)-180
}

func (provider *RAMRoleARNCredentialsProvider) GetCredentials() (cc *Credentials, err error) {
	if provider.sessionCredentials == nil || provider.needUpdateCredential() {
		// 获取前置凭证
		previousCredentials, err1 := provider.credentialsProvider.GetCredentials()
		if err1 != nil {
			return nil, err1
		}
		sessionCredentials, err2 := provider.getCredentials(previousCredentials)
		if err2 != nil {
			return nil, err2
		}

		provider.sessionCredentials = sessionCredentials
		provider.lastUpdateTimestamp = time.Now().Unix()
	}

	cc = &Credentials{
		AccessKeyId:     provider.sessionCredentials.AccessKeyId,
		AccessKeySecret: provider.sessionCredentials.AccessKeySecret,
		SecurityToken:   provider.sessionCredentials.SecurityToken,
	}
	return
}

type ECSRAMRoleCredentialsProvider struct {
	RoleName           string
	sessionCredentials *SessionCredentials
	expirationTime     int64
}

func NewECSRAMRoleCredentialsProvider(roleName string) *ECSRAMRoleCredentialsProvider {
	return &ECSRAMRoleCredentialsProvider{
		RoleName: roleName,
	}
}

func (provider *ECSRAMRoleCredentialsProvider) needUpdateCredential() bool {
	if provider.expirationTime == 0 {
		return true
	}

	return provider.expirationTime-time.Now().Unix() < 180
}

func (provider *ECSRAMRoleCredentialsProvider) getRoleName() (roleName string, err error) {
	var securityCredURL = "http://100.100.100.200/latest/meta-data/ram/security-credentials/"
	httpRequest, err := http.NewRequest(requests.GET, securityCredURL, strings.NewReader(""))
	if err != nil {
		err = fmt.Errorf("get role name failed: %s", err.Error())
		return
	}
	httpClient := &http.Client{}
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		err = fmt.Errorf("get role name failed: %s", err.Error())
		return
	}

	if httpResponse.StatusCode != http.StatusOK {
		err = fmt.Errorf("get role name failed: request %s %d", securityCredURL, httpResponse.StatusCode)
	}

	defer httpResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}

	roleName = strings.TrimSpace(string(responseBody))
	return
}

func (provider *ECSRAMRoleCredentialsProvider) getCredentials() (sessionCredentials *SessionCredentials, err error) {
	roleName := provider.RoleName
	if roleName == "" {
		roleName, err = provider.getRoleName()
		if err != nil {
			return
		}
	}

	var requestUrl = "http://100.100.100.200/latest/meta-data/ram/security-credentials/" + roleName
	httpRequest, err := http.NewRequest(requests.GET, requestUrl, strings.NewReader(""))
	if err != nil {
		err = fmt.Errorf("refresh Ecs sts token err: %s", err.Error())
		return
	}
	httpClient := &http.Client{}
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		err = fmt.Errorf("refresh Ecs sts token err: %s", err.Error())
		return
	}

	defer httpResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}

	if httpResponse.StatusCode != http.StatusOK {
		err = fmt.Errorf("refresh Ecs sts token err, httpStatus: %d, message = %s", httpResponse.StatusCode, string(responseBody))
		return
	}

	var data ecsRAMRoleCredentials
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		err = fmt.Errorf("refresh Ecs sts token err, json.Unmarshal fail: %s", err.Error())
		return
	}

	if data.AccessKeyId == nil || data.AccessKeySecret == nil || data.SecurityToken == nil {
		err = fmt.Errorf("refresh Ecs sts token err, fail to get credentials")
		return
	}

	if *data.Code != "Success" {
		err = fmt.Errorf("refresh Ecs sts token err, Code is not Success")
		return
	}

	sessionCredentials = &SessionCredentials{
		AccessKeyId:     *data.AccessKeyId,
		AccessKeySecret: *data.AccessKeySecret,
		SecurityToken:   *data.SecurityToken,
		Expiration:      *data.Expiration,
	}
	return
}

func (provider *ECSRAMRoleCredentialsProvider) GetCredentials() (cc *Credentials, err error) {
	if provider.sessionCredentials == nil || provider.needUpdateCredential() {
		sessionCredentials, err1 := provider.getCredentials()
		if err1 != nil {
			return nil, err1
		}

		provider.sessionCredentials = sessionCredentials
		expirationTime, err2 := time.Parse("2006-01-02T15:04:05Z", sessionCredentials.Expiration)
		if err2 != nil {
			return nil, err2
		}
		provider.expirationTime = expirationTime.Unix()
	}

	cc = &Credentials{
		AccessKeyId:     provider.sessionCredentials.AccessKeyId,
		AccessKeySecret: provider.sessionCredentials.AccessKeySecret,
		SecurityToken:   provider.sessionCredentials.SecurityToken,
	}
	return
}

type OIDCCredentialsProvider struct {
	OIDCProviderARN       string
	OIDCTokenFilePath     string
	RoleArn               string
	RoleSessionName       string
	RoleSessionExpiration int
	Policy                string
	StsRegion             string
	lastUpdateTimestamp   int64
	sessionCredentials    *SessionCredentials
}

type OIDCCredentialsProviderBuilder struct {
	provider *OIDCCredentialsProvider
}

func NewOIDCCredentialsProviderBuilder() *OIDCCredentialsProviderBuilder {
	return &OIDCCredentialsProviderBuilder{
		provider: &OIDCCredentialsProvider{},
	}
}

func (b *OIDCCredentialsProviderBuilder) WithOIDCProviderARN(oidcProviderArn string) *OIDCCredentialsProviderBuilder {
	b.provider.OIDCProviderARN = oidcProviderArn
	return b
}

func (b *OIDCCredentialsProviderBuilder) WithOIDCTokenFilePath(oidcTokenFilePath string) *OIDCCredentialsProviderBuilder {
	b.provider.OIDCTokenFilePath = oidcTokenFilePath
	return b
}

func (b *OIDCCredentialsProviderBuilder) WithRoleArn(roleArn string) *OIDCCredentialsProviderBuilder {
	b.provider.RoleArn = roleArn
	return b
}

func (b *OIDCCredentialsProviderBuilder) WithRoleSessionName(roleSessionName string) *OIDCCredentialsProviderBuilder {
	b.provider.RoleSessionName = roleSessionName
	return b
}

func (b *OIDCCredentialsProviderBuilder) WithRoleSessionExpiration(roleSessionExpiration int) *OIDCCredentialsProviderBuilder {
	b.provider.RoleSessionExpiration = roleSessionExpiration
	return b
}

func (b *OIDCCredentialsProviderBuilder) WithStsRegion(region string) *OIDCCredentialsProviderBuilder {
	b.provider.StsRegion = region
	return b
}

func (b *OIDCCredentialsProviderBuilder) WithPolicy(policy string) *OIDCCredentialsProviderBuilder {
	b.provider.Policy = policy
	return b
}

func (b *OIDCCredentialsProviderBuilder) Build() (provider *OIDCCredentialsProvider, err error) {
	provider = b.provider

	if provider.RoleSessionName == "" {
		provider.RoleSessionName = "aliyun-go-sdk-" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)
	}

	if provider.OIDCTokenFilePath == "" {
		provider.OIDCTokenFilePath = os.Getenv("ALIBABA_CLOUD_OIDC_TOKEN_FILE")
	}

	if provider.OIDCTokenFilePath == "" {
		err = errors.NewClientError(errors.InvalidParamErrorCode, "OIDCTokenFilePath can not be empty", nil)
		return
	}

	if provider.OIDCProviderARN == "" {
		provider.OIDCProviderARN = os.Getenv("ALIBABA_CLOUD_OIDC_PROVIDER_ARN")
	}

	if provider.OIDCProviderARN == "" {
		err = errors.NewClientError(errors.InvalidParamErrorCode, "OIDCProviderARN can not be empty", nil)
		return
	}

	if provider.RoleArn == "" {
		provider.RoleArn = os.Getenv("ALIBABA_CLOUD_ROLE_ARN")
	}

	if provider.RoleArn == "" {
		err = errors.NewClientError(errors.InvalidParamErrorCode, "RoleArn can not be empty", nil)
		return
	}

	if provider.RoleSessionExpiration == 0 {
		provider.RoleSessionExpiration = 3600
	} else {
		err = errors.NewClientError(errors.InvalidParamErrorCode, "Assume Role session duration should be in the range of 15min - 1Hr", nil)
		return
	}

	return
}

func (provider *OIDCCredentialsProvider) getCredentials() (sessionCredentials *SessionCredentials, err error) {
	method := "POST"
	var host string
	if provider.StsRegion != "" {
		host = fmt.Sprintf("sts.%s.aliyuncs.com", provider.StsRegion)
	} else {
		host = "sts.aliyuncs.com"
	}

	queries := make(map[string]string)
	queries["Version"] = "2015-04-01"
	queries["Action"] = "AssumeRoleWithOIDC"
	queries["Format"] = "JSON"
	queries["Timestamp"] = utils.GetTimeInFormatISO8601()

	bodyForm := make(map[string]string)
	bodyForm["RoleArn"] = provider.RoleArn
	bodyForm["OIDCProviderArn"] = provider.OIDCProviderARN
	token, err := ioutil.ReadFile(provider.OIDCTokenFilePath)
	if err != nil {
		return
	}

	bodyForm["OIDCToken"] = string(token)
	if provider.Policy != "" {
		bodyForm["Policy"] = provider.Policy
	}

	bodyForm["RoleSessionName"] = provider.RoleSessionName
	bodyForm["DurationSeconds"] = strconv.Itoa(provider.RoleSessionExpiration)

	// caculate signature
	signParams := make(map[string]string)
	for key, value := range queries {
		signParams[key] = value
	}
	for key, value := range bodyForm {
		signParams[key] = value
	}

	querystring := utils.GetUrlFormedMap(queries)
	// do request
	httpUrl := fmt.Sprintf("https://%s/?%s", host, querystring)

	body := utils.GetUrlFormedMap(bodyForm)

	httpRequest, err := http.NewRequest(method, httpUrl, strings.NewReader(body))
	if err != nil {
		return
	}

	// set headers
	httpRequest.Header["Accept-Encoding"] = []string{"identity"}
	httpRequest.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	httpClient := &http.Client{}

	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return
	}

	defer httpResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}

	if httpResponse.StatusCode != http.StatusOK {
		message := "get session token failed"
		err = errors.NewServerError(httpResponse.StatusCode, string(responseBody), message)
		return
	}
	var data assumeRoleResponse
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		err = fmt.Errorf("get oidc sts token err, json.Unmarshal fail: %s", err.Error())
		return
	}
	if data.Credentials == nil {
		err = fmt.Errorf("get oidc sts token err, fail to get credentials")
		return
	}

	if data.Credentials.AccessKeyId == nil || data.Credentials.AccessKeySecret == nil || data.Credentials.SecurityToken == nil {
		err = fmt.Errorf("refresh RoleArn sts token err, fail to get credentials")
		return
	}

	sessionCredentials = &SessionCredentials{
		AccessKeyId:     *data.Credentials.AccessKeyId,
		AccessKeySecret: *data.Credentials.AccessKeySecret,
		SecurityToken:   *data.Credentials.SecurityToken,
		Expiration:      *data.Credentials.Expiration,
	}
	return
}

func (provider *OIDCCredentialsProvider) needUpdateCredential() (result bool) {
	return time.Now().Unix()-provider.lastUpdateTimestamp >= int64(provider.RoleSessionExpiration)-180
}

func (provider *OIDCCredentialsProvider) GetCredentials() (cc *Credentials, err error) {
	if provider.sessionCredentials == nil || provider.needUpdateCredential() {
		sessionCredentials, err2 := provider.getCredentials()
		if err2 != nil {
			return nil, err2
		}

		provider.sessionCredentials = sessionCredentials
	}

	cc = &Credentials{
		AccessKeyId:     provider.sessionCredentials.AccessKeyId,
		AccessKeySecret: provider.sessionCredentials.AccessKeySecret,
		SecurityToken:   provider.sessionCredentials.SecurityToken,
	}
	return
}
