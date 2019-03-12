package requests

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_AcsRequest(t *testing.T) {
	r := defaultBaseRequest()
	assert.NotNil(t, r)

	// query params
	query := r.GetQueryParams()
	assert.Equal(t, 0, len(query))
	r.addQueryParam("key", "value")
	assert.Equal(t, 1, len(query))
	assert.Equal(t, "value", query["key"])

	// form params
	form := r.GetFormParams()
	assert.Equal(t, 0, len(form))
	r.addFormParam("key", "value")
	assert.Equal(t, 1, len(form))
	assert.Equal(t, "value", form["key"])

	// getter/setter for stringtosign
	assert.Equal(t, "", r.GetStringToSign())
	r.SetStringToSign("s2s")
	assert.Equal(t, "s2s", r.GetStringToSign())

	// content type
	_, contains := r.GetContentType()
	assert.False(t, contains)
	r.SetContentType("application/json")
	ct, contains := r.GetContentType()
	assert.Equal(t, "application/json", ct)
	assert.True(t, contains)

	// default 3 headers & content-type
	headers := r.GetHeaders()
	assert.Equal(t, 4, len(headers))
	r.addHeaderParam("x-key", "x-key-value")
	assert.Equal(t, 5, len(headers))
	assert.Equal(t, "x-key-value", headers["x-key"])

	// GetVersion
	assert.Equal(t, "", r.GetVersion())
	// GetActionName
	assert.Equal(t, "", r.GetActionName())

	// GetMethod
	assert.Equal(t, "GET", r.GetMethod())
	r.Method = "POST"
	assert.Equal(t, "POST", r.GetMethod())

	// Domain
	assert.Equal(t, "", r.GetDomain())
	r.SetDomain("ecs.aliyuncs.com")
	assert.Equal(t, "ecs.aliyuncs.com", r.GetDomain())

	// Region
	assert.Equal(t, "", r.GetRegionId())
	r.RegionId = "cn-hangzhou"
	assert.Equal(t, "cn-hangzhou", r.GetRegionId())

	// AcceptFormat
	assert.Equal(t, "JSON", r.GetAcceptFormat())
	r.AcceptFormat = "XML"
	assert.Equal(t, "XML", r.GetAcceptFormat())

	// GetLocationServiceCode
	assert.Equal(t, "", r.GetLocationServiceCode())

	// GetLocationEndpointType
	assert.Equal(t, "", r.GetLocationEndpointType())

	// GetProduct
	assert.Equal(t, "", r.GetProduct())

	// GetScheme
	assert.Equal(t, "", r.GetScheme())
	r.SetScheme("HTTPS")
	assert.Equal(t, "HTTPS", r.GetScheme())

	// GetReadTimeout
	assert.Equal(t, 0*time.Second, r.GetReadTimeout())
	r.SetReadTimeout(5 * time.Second)
	assert.Equal(t, 5*time.Second, r.GetReadTimeout())

	// GetConnectTimeout
	assert.Equal(t, 0*time.Second, r.GetConnectTimeout())
	r.SetConnectTimeout(5 * time.Second)
	assert.Equal(t, 5*time.Second, r.GetConnectTimeout())

	// GetHTTPSInsecure
	assert.True(t, r.GetHTTPSInsecure() == nil)
	r.SetHTTPSInsecure(true)
	assert.Equal(t, true, *r.GetHTTPSInsecure())

	// GetPort
	assert.Equal(t, "", r.GetPort())

	// GetUserAgent
	r.AppendUserAgent("cli", "1.01")
	assert.Equal(t, "1.01", r.GetUserAgent()["cli"])
	// GetUserAgent
	r.AppendUserAgent("cli", "2.02")
	assert.Equal(t, "2.02", r.GetUserAgent()["cli"])
	// Content
	assert.Equal(t, []byte(nil), r.GetContent())
	r.SetContent([]byte("The Content"))
	assert.True(t, bytes.Equal([]byte("The Content"), r.GetContent()))
}

type AcsRequestTest struct {
	*baseRequest
	Ontology AcsRequest
	Query    string    `position:"Query" name:"Query"`
	Header   string    `position:"Header" name:"Header"`
	Path     string    `position:"Path" name:"Path"`
	Body     string    `position:"Body" name:"Body"`
	TypeAcs  *[]string `position:"type" name:"type" type:"Repeated"`
}

func (r AcsRequestTest) BuildQueries() string {
	return ""
}

func (r AcsRequestTest) BuildUrl() string {
	return ""
}

func (r AcsRequestTest) GetBodyReader() io.Reader {
	return nil
}

func (r AcsRequestTest) GetStyle() string {
	return ""
}

func (r AcsRequestTest) addPathParam(key, value string) {
	return
}

func Test_AcsRequest_InitParams(t *testing.T) {
	r := &AcsRequestTest{
		baseRequest: defaultBaseRequest(),
		Query:       "query value",
		Header:      "header value",
		Path:        "path value",
		Body:        "body value",
	}
	tmp := []string{r.Query, r.Header}
	r.TypeAcs = &tmp
	r.addQueryParam("qkey", "qvalue")
	InitParams(r)

	queries := r.GetQueryParams()
	assert.Equal(t, "query value", queries["Query"])
	headers := r.GetHeaders()
	assert.Equal(t, "header value", headers["Header"])
	// TODO: check the body & path
}

type StartMPUTaskRequest struct {
	*RpcRequest
	OwnerId         Integer   `position:"Query" name:"OwnerId"`
	AppId           string    `position:"Query" name:"AppId"`
	ChannelId       string    `position:"Query" name:"ChannelId"`
	TaskId          string    `position:"Query" name:"TaskId"`
	MediaEncode     Integer   `position:"Query" name:"MediaEncode"`
	BackgroundColor Integer   `position:"Query" name:"BackgroundColor"`
	LayoutIds       []Integer `position:"Query" name:"LayoutIds" type:"Repeated"`
	StreamURL       string    `position:"Query" name:"StreamURL"`
}

func Test_RPCRequest_InitParams(t *testing.T) {
	channelID := "id"
	r := &StartMPUTaskRequest{
		RpcRequest: &RpcRequest{},
	}
	r.init()
	r.Domain = "rtc.aliyuncs.com"
	r.AppId = "app ID"
	r.ChannelId = channelID
	r.TaskId = channelID
	r.MediaEncode = NewInteger(2)
	r.BackgroundColor = NewInteger(0)
	r.StreamURL = fmt.Sprintf("rtmp://video-center.alivecdn.com/galaxy/%s_%s?vhost=fast-live.chinalivestream.top", channelID, channelID)
	var out []Integer
	out = append(out, NewInteger(2))
	r.LayoutIds = out

	InitParams(r)

	queries := r.GetQueryParams()

	assert.Equal(t, "2", queries["LayoutIds.1"])
	assert.Len(t, queries, 7)
}
