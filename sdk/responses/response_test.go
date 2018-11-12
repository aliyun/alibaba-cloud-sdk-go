package responses

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CommonResponse(t *testing.T) {
	r := NewCommonResponse()
	assert.NotNil(t, r)

	assert.Equal(t, 0, r.GetHttpStatus())
	// assert.Equal(t, nil, r.GetHttpHeaders())
	assert.Equal(t, "", r.GetHttpContentString())
	assert.Equal(t, 0, len(r.GetHttpContentBytes()))
	assert.Nil(t, r.GetOriginHttpResponse())
	assert.False(t, r.IsSuccess())
}

func Test_CommonResponse_parseFromHttpResponse(t *testing.T) {
	r := NewCommonResponse()
	header := make(http.Header)
	status := "200"
	statusCode := 200
	res := &http.Response{
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		Header:     header,
		StatusCode: statusCode,
		Status:     status + " " + http.StatusText(statusCode),
	}
	var noBody io.ReadCloser = ioutil.NopCloser(bytes.NewReader(nil))
	res.Header = make(http.Header)
	res.Header.Add("Server", "GitHub.com")
	res.Body = noBody
	r.parseFromHttpResponse(res)
	expected := `HTTP/1.1 200 OK
Server: GitHub.com


`

	assert.True(t, r.IsSuccess())
	assert.Equal(t, "GitHub.com", r.GetHttpHeaders()["Server"][0])
	assert.Equal(t, expected, r.String())
}
