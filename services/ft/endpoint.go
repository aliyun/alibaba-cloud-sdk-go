package ft

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "ft.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "ft.aliyuncs.com",
			"cn-beijing":                  "ft.aliyuncs.com",
			"cn-shanghai-inner":           "ft.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "ft.aliyuncs.com",
			"cn-yushanfang":               "ft.aliyuncs.com",
			"cn-qingdao-nebula":           "ft.aliyuncs.com",
			"cn-beijing-finance-pop":      "ft.aliyuncs.com",
			"cn-wuhan":                    "ft.aliyuncs.com",
			"us-west-1":                   "ft.aliyuncs.com",
			"cn-zhangbei":                 "ft.aliyuncs.com",
			"rus-west-1-pop":              "ft.aliyuncs.com",
			"cn-shanghai-et15-b01":        "ft.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "ft.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "ft.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "ft.aliyuncs.com",
			"cn-shanghai-et2-b01":         "ft.aliyuncs.com",
			"ap-southeast-1":              "ft.aliyuncs.com",
			"ap-southeast-2":              "ft.aliyuncs.com",
			"ap-southeast-3":              "ft.aliyuncs.com",
			"ap-southeast-5":              "ft.aliyuncs.com",
			"cn-shenzhen-inner":           "ft.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "ft.aliyuncs.com",
			"cn-beijing-gov-1":            "ft.aliyuncs.com",
			"cn-wulanchabu":               "ft.aliyuncs.com",
			"ap-south-1":                  "ft.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "ft.aliyuncs.com",
			"cn-haidian-cm12-c01":         "ft.aliyuncs.com",
			"cn-qingdao":                  "ft.aliyuncs.com",
			"cn-hongkong-finance-pop":     "ft.aliyuncs.com",
			"cn-shanghai-finance-1":       "ft.aliyuncs.com",
			"eu-central-1":                "ft.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "ft.aliyuncs.com",
			"eu-west-1":                   "ft.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "ft.aliyuncs.com",
			"eu-west-1-oxs":               "ft.aliyuncs.com",
			"cn-beijing-finance-1":        "ft.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "ft.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "ft.aliyuncs.com",
			"cn-shenzhen-finance-1":       "ft.aliyuncs.com",
			"me-east-1":                   "ft.aliyuncs.com",
			"cn-chengdu":                  "ft.aliyuncs.com",
			"cn-hangzhou-test-306":        "ft.aliyuncs.com",
			"cn-hangzhou-finance":         "ft.aliyuncs.com",
			"cn-beijing-nu16-b01":         "ft.aliyuncs.com",
			"cn-edge-1":                   "ft.aliyuncs.com",
			"cn-huhehaote":                "ft.aliyuncs.com",
			"cn-fujian":                   "ft.aliyuncs.com",
			"ap-northeast-2-pop":          "ft.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
