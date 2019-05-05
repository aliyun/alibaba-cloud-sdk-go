// Package self_cms is used for product cms
package self_cms

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sort"
	"strings"
	"time"

	"github.com/alibabacloud-go/tea/tea"
)

// Client the client type
type Client struct {
	Protocol__      string
	Endpoint__      string
	UserAgent__     string
	AccessKeyId     string
	AccessKeySecret string
}

func (client *Client) GetGMTDateString__() string {
	gmt := time.FixedZone("GMT", 0)
	// Mon, 3 Jan 2010 08:33:47 GMT
	return time.Now().In(gmt).Format("Mon, 02 Jan 2006 15:04:05 GMT")
}

func (client *Client) GetIp__() string {
	ifaces, _ := net.Interfaces()

	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			return ip.String()
		}
	}
	return ""
}

func (client *Client) Stringify__(obj interface{}) string {
	b, _ := json.Marshal(obj)
	return string(b)
}

func (client *Client) GetContentLength__(request *tea.Request) string {
	return fmt.Sprintf("%d", len([]byte(request.Body)))
}

func (client *Client) GetContentMD5__(request *tea.Request) string {
	h := md5.New()
	io.WriteString(h, request.Body)
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}

func shaHmac1(source, secret string) []byte {
	key := []byte(secret)
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(source))
	return hmac.Sum(nil)
}

func base16(src []byte) string {
	cc := []byte("0123456789ABCDEF")
	out := make([]byte, len(src)*2)
	for i, v := range src {
		out[2*i] = cc[(v>>4)&0x0f]
		out[2*i+1] = cc[v&0x0f]
	}

	return string(out)
}

func getCanonicalizedHeaders(headers map[string]string) string {
	keys := []string{}
	for key := range headers {
		if strings.HasPrefix(key, "x-cms-") || strings.HasPrefix(key, "x-acs") {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	out := ""
	for _, value := range keys {
		out += fmt.Sprintf("%s:%s\n", value, headers[value])
	}
	return strings.TrimSpace(out)
}

func getCanonicalizedResource(pathname string, query map[string]string) string {
	out := pathname
	if query == nil || len(query) == 0 {
		return out
	}

	return ""
}

// CanonicalizedResource 的构造方式如下：
// 将 CanonicalizedResource 设置为空字符串（””）；
// 放入要访问的URI,如/event/custom/upload
// 如请求包含查询字符串（QUERY_STRING），则在 CanonicalizedResource 字符串尾部添加 ？ 和查询字符串。
// 其中，QUERY_STRING 是 URL 中请求参数按字典序排序后的字符串，其中参数名和值之间用 = 相隔组成字符串，并对参数名-值对按照字典序升序排序，然后以 & 符号连接构成字符串。其公式化描述如下：

// QUERY_STRING = "KEY1=VALUE1" + "&" + "KEY2=VALUE2"

// GetAuthorization__ get authorization
func (client *Client) GetAuthorization__(request *tea.Request) string {
	// yourAccessKeyId:yourAccessKeySecret
	verb := request.Method
	contentMD5 := request.Headers["content-md5"]
	contentType := request.Headers["content-type"]
	date := request.Headers["date"]
	canonicalizedHeaders := getCanonicalizedHeaders(request.Headers)
	canonicalizedResource := getCanonicalizedResource(request.Pathname, request.Query)
	signString := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", verb, contentMD5, contentType, date, canonicalizedHeaders, canonicalizedResource)
	signature := base16(shaHmac1(signString, client.AccessKeySecret))
	return fmt.Sprintf("%s:%s", client.AccessKeyId, signature)
}

func (client *Client) IsNotOk__(response *tea.Response) bool {
	return response.StatusCode < 200 || response.StatusCode >= 400
}

func (client *Client) Is4xx__(body map[string]interface{}) bool {
	code := body["code"]
	if v := code.(string); v != "200" {
		return true
	}
	return false
}

func (client *Client) ReadJSON__(response *tea.Response) (result map[string]interface{}, err error) {
	body, err := response.ReadBody()
	if err != nil {
		return
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &result)
	return
}

func NewClient(config map[string]string) *Client {
	return &Client{
		Protocol__:      "https",
		AccessKeyId:     config["AccessKeyId"],
		AccessKeySecret: config["AccessKeySecret"],
	}
}
