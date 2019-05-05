package tea

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// CastError is used for cast type fails
type CastError struct {
	Message string
}

// NewCastError is used for cast type fails
func NewCastError(message string) (err error) {
	return &CastError{
		Message: message,
	}
}

func (err *CastError) Error() string {
	return err.Message
}

func firstDownCase(name string) string {
	return strings.ToLower(string(name[0])) + name[1:]
}

// Convert is use convert map[string]interface object to struct
func Convert(in map[string]interface{}, out interface{}) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr {
		return NewCastError("The out parameter must be pointer")
	}

	for i := 0; i < v.Elem().NumField(); i++ {
		fieldInfo := v.Elem().Type().Field(i)
		name := firstDownCase(fieldInfo.Name)
		if value, ok := in[name]; ok {
			if reflect.ValueOf(value).Kind() == v.Elem().FieldByName(fieldInfo.Name).Kind() {
				v.Elem().FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			} else {
				currentType := reflect.ValueOf(value).Type()
				expectType := v.Elem().FieldByName(fieldInfo.Name).Type()
				return NewCastError(fmt.Sprintf("Convert type fails for field: %s, expect type: %s, current type: %s", name, expectType, currentType))
			}
		}
	}

	out = v.Interface()
	return nil
}

// Request is used wrap http request
type Request struct {
	Protocol string
	Port     int
	Method   string
	Pathname string
	Headers  map[string]string
	Query    map[string]string
	Body     string
}

// NewRequest is used shortly create Request
func NewRequest() (req *Request) {
	return &Request{
		Headers: map[string]string{},
		Query:   map[string]string{},
	}
}

// Response is use d wrap http response
type Response struct {
	httpResponse  *http.Response
	StatusCode    int
	StatusMessage string
}

// NewResponse is create response with http response
func NewResponse(httpResponse *http.Response) (res *Response) {
	res = &Response{
		httpResponse: httpResponse,
	}
	res.StatusCode = httpResponse.StatusCode
	res.StatusMessage = httpResponse.Status
	return
}

// ReadBody is used read response body
func (response *Response) ReadBody() (body []byte, err error) {
	res := response.httpResponse
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	return
}

// DoRequest is used send request to server
func DoRequest(request *Request) (response *Response, err error) {
	requestMethod := request.Method
	if requestMethod == "" {
		requestMethod = "GET"
	}

	protocol := "http"
	if request.Protocol != "" {
		protocol = strings.ToLower(request.Protocol)
	}

	port := 0
	if protocol == "http" {
		port = 80
	} else if protocol == "https" {
		port = 443
	}

	if request.Port != 0 {
		port = request.Port
	}

	domain := request.Headers["host"]
	requestURL := fmt.Sprintf("%s://%s:%d%s", protocol, domain, port, request.Pathname)
	queryParams := request.Query
	// sort QueryParams by key
	q := url.Values{}
	for key, value := range queryParams {
		q.Add(key, value)
	}
	querystring := q.Encode()
	if len(querystring) > 0 {
		requestURL = fmt.Sprintf("%s?%s", requestURL, querystring)
	}

	fmt.Println(requestMethod)
	fmt.Println(requestURL)
	httpRequest, err := http.NewRequest(requestMethod, requestURL, strings.NewReader(request.Body))
	if err != nil {
		return
	}
	for key, value := range request.Headers {
		httpRequest.Header[key] = []string{value}
		fmt.Printf("> %s: %s\n", key, value)
	}
	httpRequest.Host = domain

	httpClient := &http.Client{}
	res, err := httpClient.Do(httpRequest)
	if res != nil {
		fmt.Printf("< HTTP/1.1 %s\n", res.Status)
		for key, value := range res.Header {
			fmt.Printf("< %s: %s\n", key, strings.Join(value, ""))
		}
	}

	if err != nil {
		return
	}

	response = NewResponse(res)
	return
}

// SDKError struct is used save error code and message
type SDKError struct {
	Code    string
	Message string
}

func (err *SDKError) Error() string {
	return fmt.Sprintf("SDKError: %s %s", err.Code, err.Message)
}

// NewSDKError is used for shortly create SDKError object
func NewSDKError(obj map[string]interface{}) *SDKError {
	err := &SDKError{}
	if val, ok := obj["code"].(int); ok {
		err.Code = strconv.Itoa(val)
	} else if val, ok := obj["code"].(string); ok {
		err.Code = val
	}

	if obj["message"] != nil {
		err.Message = obj["message"].(string)
	}
	return err
}
