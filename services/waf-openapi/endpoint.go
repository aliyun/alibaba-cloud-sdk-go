package waf_openapi

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "wafopenapi.cn-hangzhou.aliyuncs.com",
			"cn-beijing":            "wafopenapi.cn-hangzhou.aliyuncs.com",
			"cn-wulanchabu":         "wafopenapi.cn-hangzhou.aliyuncs.com",
			"ap-south-1":            "wafopenapi.ap-southeast-1.aliyuncs.com",
			"eu-west-1":             "wafopenapi.ap-southeast-1.aliyuncs.com",
			"cn-shenzhen-finance-1": "wafopenapi.cn-hangzhou.aliyuncs.com",
			"me-east-1":             "wafopenapi.ap-southeast-1.aliyuncs.com",
			"cn-chengdu":            "wafopenapi.cn-hangzhou.aliyuncs.com",
			"cn-north-2-gov-1":      "wafopenapi.cn-hangzhou.aliyuncs.com",
			"cn-qingdao":            "wafopenapi.cn-hangzhou.aliyuncs.com",
			"cn-shanghai":           "wafopenapi.cn-hangzhou.aliyuncs.com",
			"cn-shanghai-finance-1": "wafopenapi.cn-hangzhou.aliyuncs.com",
			"cn-hongkong":           "wafopenapi.ap-southeast-1.aliyuncs.com",
			"cn-heyuan":             "wafopenapi.cn-hangzhou.aliyuncs.com",
			"ap-southeast-1":        "wafopenapi.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":        "wafopenapi.ap-southeast-1.aliyuncs.com",
			"ap-southeast-3":        "wafopenapi.ap-southeast-1.aliyuncs.com",
			"eu-central-1":          "wafopenapi.ap-southeast-1.aliyuncs.com",
			"cn-huhehaote":          "wafopenapi.cn-hangzhou.aliyuncs.com",
			"ap-southeast-5":        "wafopenapi.ap-southeast-1.aliyuncs.com",
			"us-east-1":             "wafopenapi.ap-southeast-1.aliyuncs.com",
			"cn-zhangjiakou":        "wafopenapi.cn-hangzhou.aliyuncs.com",
			"us-west-1":             "wafopenapi.ap-southeast-1.aliyuncs.com",
			"cn-hangzhou":           "wafopenapi.cn-hangzhou.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
