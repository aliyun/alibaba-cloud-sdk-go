package mpaas

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "mpaas.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "mpaas.aliyuncs.com",
			"cn-beijing":                  "mpaas.aliyuncs.com",
			"cn-shanghai-inner":           "mpaas.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "mpaas.aliyuncs.com",
			"cn-north-2-gov-1":            "mpaas.aliyuncs.com",
			"cn-yushanfang":               "mpaas.aliyuncs.com",
			"cn-qingdao-nebula":           "mpaas.aliyuncs.com",
			"cn-beijing-finance-pop":      "mpaas.aliyuncs.com",
			"cn-wuhan":                    "mpaas.aliyuncs.com",
			"cn-zhangjiakou":              "mpaas.aliyuncs.com",
			"us-west-1":                   "mpaas.aliyuncs.com",
			"rus-west-1-pop":              "mpaas.aliyuncs.com",
			"cn-shanghai-et15-b01":        "mpaas.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "mpaas.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "mpaas.aliyuncs.com",
			"ap-northeast-1":              "mpaas.aliyuncs.com",
			"cn-shanghai-et2-b01":         "mpaas.aliyuncs.com",
			"ap-southeast-1":              "mpaas.aliyuncs.com",
			"ap-southeast-2":              "mpaas.aliyuncs.com",
			"ap-southeast-3":              "mpaas.aliyuncs.com",
			"ap-southeast-5":              "mpaas.aliyuncs.com",
			"us-east-1":                   "mpaas.aliyuncs.com",
			"cn-shenzhen-inner":           "mpaas.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "mpaas.aliyuncs.com",
			"cn-beijing-gov-1":            "mpaas.aliyuncs.com",
			"ap-south-1":                  "mpaas.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "mpaas.aliyuncs.com",
			"cn-haidian-cm12-c01":         "mpaas.aliyuncs.com",
			"cn-qingdao":                  "mpaas.aliyuncs.com",
			"cn-hongkong-finance-pop":     "mpaas.aliyuncs.com",
			"cn-shanghai":                 "mpaas.aliyuncs.com",
			"cn-shanghai-finance-1":       "mpaas.aliyuncs.com",
			"cn-hongkong":                 "mpaas.aliyuncs.com",
			"eu-central-1":                "mpaas.aliyuncs.com",
			"cn-shenzhen":                 "mpaas.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "mpaas.aliyuncs.com",
			"eu-west-1":                   "mpaas.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "mpaas.aliyuncs.com",
			"eu-west-1-oxs":               "mpaas.aliyuncs.com",
			"cn-beijing-finance-1":        "mpaas.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "mpaas.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "mpaas.aliyuncs.com",
			"cn-shenzhen-finance-1":       "mpaas.aliyuncs.com",
			"me-east-1":                   "mpaas.aliyuncs.com",
			"cn-chengdu":                  "mpaas.aliyuncs.com",
			"cn-hangzhou-test-306":        "mpaas.aliyuncs.com",
			"cn-hangzhou-finance":         "mpaas.aliyuncs.com",
			"cn-beijing-nu16-b01":         "mpaas.aliyuncs.com",
			"cn-edge-1":                   "mpaas.aliyuncs.com",
			"cn-huhehaote":                "mpaas.aliyuncs.com",
			"cn-fujian":                   "mpaas.aliyuncs.com",
			"ap-northeast-2-pop":          "mpaas.aliyuncs.com",
			"cn-hangzhou":                 "mpaas.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
