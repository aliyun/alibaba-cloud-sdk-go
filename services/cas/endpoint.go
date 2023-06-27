package cas

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "cas.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "cas.aliyuncs.com",
			"cn-beijing":                  "cas.aliyuncs.com",
			"cn-shanghai-inner":           "cas.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "cas.aliyuncs.com",
			"cn-north-2-gov-1":            "cas.aliyuncs.com",
			"cn-yushanfang":               "cas.aliyuncs.com",
			"cn-qingdao-nebula":           "cas.aliyuncs.com",
			"cn-beijing-finance-pop":      "cas.aliyuncs.com",
			"cn-wuhan":                    "cas.aliyuncs.com",
			"cn-zhangjiakou":              "cas.aliyuncs.com",
			"us-west-1":                   "cas.aliyuncs.com",
			"cn-zhangbei":                 "cas.aliyuncs.com",
			"rus-west-1-pop":              "cas.aliyuncs.com",
			"cn-shanghai-et15-b01":        "cas.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "cas.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "cas.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "cas.aliyuncs.com",
			"cn-shanghai-et2-b01":         "cas.aliyuncs.com",
			"ap-southeast-3":              "cas.aliyuncs.com",
			"ap-southeast-5":              "cas.aliyuncs.com",
			"us-east-1":                   "cas.aliyuncs.com",
			"cn-shenzhen-inner":           "cas.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "cas.aliyuncs.com",
			"cn-beijing-gov-1":            "cas.aliyuncs.com",
			"cn-wulanchabu":               "cas.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "cas.aliyuncs.com",
			"cn-haidian-cm12-c01":         "cas.aliyuncs.com",
			"cn-qingdao":                  "cas.aliyuncs.com",
			"cn-hongkong-finance-pop":     "cas.aliyuncs.com",
			"cn-shanghai":                 "cas.aliyuncs.com",
			"cn-shanghai-finance-1":       "cas.aliyuncs.com",
			"cn-hongkong":                 "cas.aliyuncs.com",
			"cn-shenzhen":                 "cas.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "cas.aliyuncs.com",
			"eu-west-1":                   "cas.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "cas.aliyuncs.com",
			"eu-west-1-oxs":               "cas.aliyuncs.com",
			"cn-beijing-finance-1":        "cas.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "cas.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "cas.aliyuncs.com",
			"cn-shenzhen-finance-1":       "cas.aliyuncs.com",
			"cn-chengdu":                  "cas.aliyuncs.com",
			"cn-hangzhou-test-306":        "cas.aliyuncs.com",
			"cn-hangzhou-finance":         "cas.aliyuncs.com",
			"cn-beijing-nu16-b01":         "cas.aliyuncs.com",
			"cn-edge-1":                   "cas.aliyuncs.com",
			"cn-huhehaote":                "cas.aliyuncs.com",
			"cn-fujian":                   "cas.aliyuncs.com",
			"ap-northeast-2-pop":          "cas.aliyuncs.com",
			"cn-hangzhou":                 "cas.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
