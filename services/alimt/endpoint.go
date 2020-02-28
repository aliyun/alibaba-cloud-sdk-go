package alimt

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "mt.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "mt.aliyuncs.com",
			"cn-beijing":                  "mt.aliyuncs.com",
			"cn-shanghai-inner":           "mt.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "mt.aliyuncs.com",
			"cn-north-2-gov-1":            "mt.aliyuncs.com",
			"cn-yushanfang":               "mt.aliyuncs.com",
			"cn-qingdao-nebula":           "mt.aliyuncs.com",
			"cn-beijing-finance-pop":      "mt.aliyuncs.com",
			"cn-wuhan":                    "mt.aliyuncs.com",
			"cn-zhangjiakou":              "mt.aliyuncs.com",
			"us-west-1":                   "mt.aliyuncs.com",
			"rus-west-1-pop":              "mt.aliyuncs.com",
			"cn-shanghai-et15-b01":        "mt.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "mt.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "mt.aliyuncs.com",
			"ap-northeast-1":              "mt.aliyuncs.com",
			"cn-shanghai-et2-b01":         "mt.aliyuncs.com",
			"ap-southeast-1":              "mt.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":              "mt.aliyuncs.com",
			"ap-southeast-3":              "mt.aliyuncs.com",
			"ap-southeast-5":              "mt.aliyuncs.com",
			"us-east-1":                   "mt.aliyuncs.com",
			"cn-shenzhen-inner":           "mt.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "mt.aliyuncs.com",
			"cn-beijing-gov-1":            "mt.aliyuncs.com",
			"ap-south-1":                  "mt.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "mt.aliyuncs.com",
			"cn-haidian-cm12-c01":         "mt.aliyuncs.com",
			"cn-qingdao":                  "mt.aliyuncs.com",
			"cn-hongkong-finance-pop":     "mt.aliyuncs.com",
			"cn-shanghai":                 "mt.aliyuncs.com",
			"cn-shanghai-finance-1":       "mt.aliyuncs.com",
			"cn-hongkong":                 "mt.aliyuncs.com",
			"eu-central-1":                "mt.aliyuncs.com",
			"cn-shenzhen":                 "mt.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "mt.aliyuncs.com",
			"eu-west-1":                   "mt.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "mt.aliyuncs.com",
			"eu-west-1-oxs":               "mt.aliyuncs.com",
			"cn-beijing-finance-1":        "mt.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "mt.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "mt.aliyuncs.com",
			"cn-shenzhen-finance-1":       "mt.aliyuncs.com",
			"me-east-1":                   "mt.aliyuncs.com",
			"cn-chengdu":                  "mt.aliyuncs.com",
			"cn-hangzhou-test-306":        "mt.aliyuncs.com",
			"cn-hangzhou-finance":         "mt.aliyuncs.com",
			"cn-beijing-nu16-b01":         "mt.aliyuncs.com",
			"cn-edge-1":                   "mt.aliyuncs.com",
			"cn-huhehaote":                "mt.aliyuncs.com",
			"cn-fujian":                   "mt.aliyuncs.com",
			"ap-northeast-2-pop":          "mt.aliyuncs.com",
			"cn-hangzhou":                 "mt.cn-hangzhou.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
