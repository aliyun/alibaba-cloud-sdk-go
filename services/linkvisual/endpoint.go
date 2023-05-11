package linkvisual

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "linkvisual.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "linkvisual.aliyuncs.com",
			"cn-beijing":                  "linkvisual.aliyuncs.com",
			"cn-shanghai-inner":           "linkvisual.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "linkvisual.aliyuncs.com",
			"cn-north-2-gov-1":            "linkvisual.aliyuncs.com",
			"cn-yushanfang":               "linkvisual.aliyuncs.com",
			"cn-qingdao-nebula":           "linkvisual.aliyuncs.com",
			"cn-beijing-finance-pop":      "linkvisual.aliyuncs.com",
			"cn-wuhan":                    "linkvisual.aliyuncs.com",
			"cn-zhangjiakou":              "linkvisual.aliyuncs.com",
			"us-west-1":                   "linkvisual.aliyuncs.com",
			"rus-west-1-pop":              "linkvisual.aliyuncs.com",
			"cn-shanghai-et15-b01":        "linkvisual.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "linkvisual.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "linkvisual.aliyuncs.com",
			"ap-northeast-1":              "linkvisual.aliyuncs.com",
			"cn-shanghai-et2-b01":         "linkvisual.aliyuncs.com",
			"ap-southeast-1":              "linkvisual.aliyuncs.com",
			"ap-southeast-2":              "linkvisual.aliyuncs.com",
			"ap-southeast-3":              "linkvisual.aliyuncs.com",
			"ap-southeast-5":              "linkvisual.aliyuncs.com",
			"us-east-1":                   "linkvisual.aliyuncs.com",
			"cn-shenzhen-inner":           "linkvisual.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "linkvisual.aliyuncs.com",
			"cn-beijing-gov-1":            "linkvisual.aliyuncs.com",
			"ap-south-1":                  "linkvisual.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "linkvisual.aliyuncs.com",
			"cn-haidian-cm12-c01":         "linkvisual.aliyuncs.com",
			"cn-qingdao":                  "linkvisual.aliyuncs.com",
			"cn-hongkong-finance-pop":     "linkvisual.aliyuncs.com",
			"cn-shanghai-finance-1":       "linkvisual.aliyuncs.com",
			"cn-hongkong":                 "linkvisual.aliyuncs.com",
			"eu-central-1":                "linkvisual.aliyuncs.com",
			"cn-shenzhen":                 "linkvisual.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "linkvisual.aliyuncs.com",
			"eu-west-1":                   "linkvisual.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "linkvisual.aliyuncs.com",
			"eu-west-1-oxs":               "linkvisual.aliyuncs.com",
			"cn-beijing-finance-1":        "linkvisual.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "linkvisual.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "linkvisual.aliyuncs.com",
			"cn-shenzhen-finance-1":       "linkvisual.aliyuncs.com",
			"me-east-1":                   "linkvisual.aliyuncs.com",
			"cn-chengdu":                  "linkvisual.aliyuncs.com",
			"cn-hangzhou-test-306":        "linkvisual.aliyuncs.com",
			"cn-hangzhou-finance":         "linkvisual.aliyuncs.com",
			"cn-beijing-nu16-b01":         "linkvisual.aliyuncs.com",
			"cn-edge-1":                   "linkvisual.aliyuncs.com",
			"cn-huhehaote":                "linkvisual.aliyuncs.com",
			"cn-fujian":                   "linkvisual.aliyuncs.com",
			"ap-northeast-2-pop":          "linkvisual.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
