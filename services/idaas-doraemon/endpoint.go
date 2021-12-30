package idaas_doraemon

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "central"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-hangzhou": "idaas-doraemon.cn-hangzhou.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
