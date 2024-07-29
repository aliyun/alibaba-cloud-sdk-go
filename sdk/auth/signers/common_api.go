package signers

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type commonApiFunc func(request *requests.CommonRequest, signer interface{}) (response *responses.CommonResponse, err error)
