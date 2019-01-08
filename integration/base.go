package integration

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"strings"

	"os"
)
var role_doc = `{
		"Statement": [{
		    "Action": "sts:AssumeRole",
		    "Effect": "Allow",
		    "Principal": {
		     	"RAM": [
				      "acs:ram::%s:root"
		        ]
            }
	    }],
	   "Version": "1"
	}`

func createRole(userid string) error{
	listRequest :=ram.CreateListRolesRequest()
	listRequest.Scheme = "HTTPS"
	client, err := ram.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		return err
	}
	listResponse, err :=client.ListRoles(listRequest)
	if err != nil {
		return err
	}
	for _, role := range listResponse.Roles.Role {
		if strings.ToLower(role.RoleName) == "testrole" {
			return nil
		}
	}
	createRequest := ram.CreateCreateRoleRequest()
	createRequest.Scheme = "HTTPS"
	createRequest.RoleName = "testrole"
	createRequest.AssumeRolePolicyDocument = fmt.Sprintf(role_doc, userid)
	_, err = client.CreateRole(createRequest)
	if err != nil {
		return err
	}
	return nil
}

func createUser()(error){
	listRequest := ram.CreateListUsersRequest()
	listRequest.Scheme = "HTTPS"
	client, err := ram.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		return err
	}
	listResponse, err := client.ListUsers(listRequest)
	if err != nil {
		return err
	}
	for _, user := range listResponse.Users.User {
		if user.UserName == "alice" {
			return nil
		}
	}
	createRequest := ram.CreateCreateUserRequest()
	createRequest.Scheme = "HTTPS"
	createRequest.UserName = "alice"
	_, err = client.CreateUser(createRequest)
	if err != nil {
		return err
	}
	return nil
}

func createAttachPolicyToUser()error{
	listRequest := ram.CreateListPoliciesForUserRequest()
	listRequest.UserName ="alice"
	listRequest.Scheme = "HTTPS"
	client, err := ram.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		return err
	}
	listResponse, err := client.ListPoliciesForUser(listRequest)
	if err != nil {
		return err
	}
	for _, policy := range listResponse.Policies.Policy {
		if policy.PolicyName == "AliyunSTSAssumeRoleAccess" {
			return nil
		}
	}
	createRequest := ram.CreateAttachPolicyToUserRequest()
	createRequest.Scheme = "HTTPS"
	createRequest.PolicyName = "AliyunSTSAssumeRoleAccess"
	createRequest.UserName = "alice"
	createRequest.PolicyType = "System"
	_, err = client.AttachPolicyToUser(createRequest)
	if err != nil {
		return err
	}
	return nil
}

//func createAccessKey()(string, string, error){
//	client, err := ram.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
//	if err != nil {
//		return "", "", err
//	}
//	listrequest := ram.CreateListAccessKeysRequest()
//	listrequest.UserName = "alice"
//	listrequest.Scheme = "HTTPS"
//	listresponse, err := client.ListAccessKeys(listrequest)
//	if err != nil {
//		return "", "", err
//	}
//	if listresponse.AccessKeys.AccessKey != nil {
//		accesskey := listresponse.AccessKeys.AccessKey[0]
//		if accesskey.AccessKeySecret != "" && accesskey.AccessKeySecret != "---" {
//			return accesskey.AccessKeyId, accesskey.AccessKeySecret, err
//		} else {
//			deleterequest := ram.CreateDeleteAccessKeyRequest()
//			deleterequest.UserAccessKeyId = accesskey.AccessKeyId
//			deleterequest.UserName = "alice"
//			_, err := client.DeleteAccessKey(deleterequest)
//			if err != nil {
//				return "", "", err
//			}
//		}
//	}
//	request := ram.CreateCreateAccessKeyRequest()
//	request.Scheme = "HTTPS"
//	request.UserName = "alice"
//	response, err := client.CreateAccessKey(request)
//	if err != nil {
//		return "", "",err
//	}
//
//	return response.AccessKey.AccessKeyId, response.AccessKey.AccessKeySecret, nil
//}

func createAssumeRole()(*sts.AssumeRoleResponse, error){
	err := createUser()
	if err != nil {
		return nil,err
	}
	err = createRole(os.Getenv("USER_ID"))
	if err != nil {
		return nil,err
	}
	err = createAttachPolicyToUser()
	if err != nil {
		return nil,err
	}
	//subaccesskeyid, subaccesskeysecret, err := createAccessKey()
	//if err != nil {
	//	return "","","",err
	//}
	request := sts.CreateAssumeRoleRequest()
	request.RoleArn = fmt.Sprintf("acs:ram::%s:role/testrole", os.Getenv("USER_ID"))
	request.RoleSessionName = "alice_test"
	request.Scheme = "HTTPS"
	client, err := sts.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	response, err := client.AssumeRole(request)
	if err != nil {
		return nil,err
	}
	return response, nil
}

