package ltl

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "ltl.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "ltl.aliyuncs.com",
			"cn-beijing":                  "ltl.aliyuncs.com",
			"cn-shanghai-inner":           "ltl.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "ltl.aliyuncs.com",
			"cn-north-2-gov-1":            "ltl.aliyuncs.com",
			"cn-yushanfang":               "ltl.aliyuncs.com",
			"cn-qingdao-nebula":           "ltl.aliyuncs.com",
			"cn-beijing-finance-pop":      "ltl.aliyuncs.com",
			"cn-wuhan":                    "ltl.aliyuncs.com",
			"cn-zhangjiakou":              "ltl.aliyuncs.com",
			"us-west-1":                   "ltl.aliyuncs.com",
			"cn-zhangbei":                 "ltl.aliyuncs.com",
			"rus-west-1-pop":              "ltl.aliyuncs.com",
			"cn-shanghai-et15-b01":        "ltl.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "ltl.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "ltl.aliyuncs.com",
			"ap-northeast-1":              "ltl.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "ltl.aliyuncs.com",
			"cn-shanghai-et2-b01":         "ltl.aliyuncs.com",
			"ap-southeast-1":              "ltl.aliyuncs.com",
			"ap-southeast-2":              "ltl.aliyuncs.com",
			"ap-southeast-3":              "ltl.aliyuncs.com",
			"ap-southeast-5":              "ltl.aliyuncs.com",
			"us-east-1":                   "ltl.aliyuncs.com",
			"cn-shenzhen-inner":           "ltl.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "ltl.aliyuncs.com",
			"cn-beijing-gov-1":            "ltl.aliyuncs.com",
			"cn-wulanchabu":               "ltl.aliyuncs.com",
			"ap-south-1":                  "ltl.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "ltl.aliyuncs.com",
			"cn-haidian-cm12-c01":         "ltl.aliyuncs.com",
			"cn-qingdao":                  "ltl.aliyuncs.com",
			"cn-hongkong-finance-pop":     "ltl.aliyuncs.com",
			"cn-shanghai-finance-1":       "ltl.aliyuncs.com",
			"cn-hongkong":                 "ltl.aliyuncs.com",
			"eu-central-1":                "ltl.aliyuncs.com",
			"cn-shenzhen":                 "ltl.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "ltl.aliyuncs.com",
			"eu-west-1":                   "ltl.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "ltl.aliyuncs.com",
			"eu-west-1-oxs":               "ltl.aliyuncs.com",
			"cn-beijing-finance-1":        "ltl.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "ltl.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "ltl.aliyuncs.com",
			"cn-shenzhen-finance-1":       "ltl.aliyuncs.com",
			"me-east-1":                   "ltl.aliyuncs.com",
			"cn-chengdu":                  "ltl.aliyuncs.com",
			"cn-hangzhou-test-306":        "ltl.aliyuncs.com",
			"cn-hangzhou-finance":         "ltl.aliyuncs.com",
			"cn-beijing-nu16-b01":         "ltl.aliyuncs.com",
			"cn-edge-1":                   "ltl.aliyuncs.com",
			"cn-huhehaote":                "ltl.aliyuncs.com",
			"cn-fujian":                   "ltl.aliyuncs.com",
			"ap-northeast-2-pop":          "ltl.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
