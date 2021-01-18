package dg

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "dg.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "dg.aliyuncs.com",
			"cn-beijing":                  "dg.aliyuncs.com",
			"cn-shanghai-inner":           "dg.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "dg.aliyuncs.com",
			"cn-north-2-gov-1":            "dg.aliyuncs.com",
			"cn-yushanfang":               "dg.aliyuncs.com",
			"cn-qingdao-nebula":           "dg.aliyuncs.com",
			"cn-beijing-finance-pop":      "dg.aliyuncs.com",
			"cn-wuhan":                    "dg.aliyuncs.com",
			"cn-zhangjiakou":              "dg.aliyuncs.com",
			"us-west-1":                   "dg.aliyuncs.com",
			"rus-west-1-pop":              "dg.aliyuncs.com",
			"cn-shanghai-et15-b01":        "dg.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "dg.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "dg.aliyuncs.com",
			"ap-northeast-1":              "dg.aliyuncs.com",
			"cn-shanghai-et2-b01":         "dg.aliyuncs.com",
			"ap-southeast-1":              "dg.aliyuncs.com",
			"ap-southeast-2":              "dg.aliyuncs.com",
			"ap-southeast-3":              "dg.aliyuncs.com",
			"ap-southeast-5":              "dg.aliyuncs.com",
			"us-east-1":                   "dg.aliyuncs.com",
			"cn-shenzhen-inner":           "dg.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "dg.aliyuncs.com",
			"cn-beijing-gov-1":            "dg.aliyuncs.com",
			"ap-south-1":                  "dg.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "dg.aliyuncs.com",
			"cn-haidian-cm12-c01":         "dg.aliyuncs.com",
			"cn-qingdao":                  "dg.aliyuncs.com",
			"cn-hongkong-finance-pop":     "dg.aliyuncs.com",
			"cn-shanghai":                 "dg.aliyuncs.com",
			"cn-shanghai-finance-1":       "dg.aliyuncs.com",
			"cn-hongkong":                 "dg.aliyuncs.com",
			"eu-central-1":                "dg.aliyuncs.com",
			"cn-shenzhen":                 "dg.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "dg.aliyuncs.com",
			"eu-west-1":                   "dg.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "dg.aliyuncs.com",
			"eu-west-1-oxs":               "dg.aliyuncs.com",
			"cn-beijing-finance-1":        "dg.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "dg.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "dg.aliyuncs.com",
			"cn-shenzhen-finance-1":       "dg.aliyuncs.com",
			"me-east-1":                   "dg.aliyuncs.com",
			"cn-chengdu":                  "dg.aliyuncs.com",
			"cn-hangzhou-test-306":        "dg.aliyuncs.com",
			"cn-hangzhou-finance":         "dg.aliyuncs.com",
			"cn-beijing-nu16-b01":         "dg.aliyuncs.com",
			"cn-edge-1":                   "dg.aliyuncs.com",
			"cn-huhehaote":                "dg.aliyuncs.com",
			"cn-fujian":                   "dg.aliyuncs.com",
			"ap-northeast-2-pop":          "dg.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
