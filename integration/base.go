package integration

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"

	"fmt"
	"os"
	"strings"
)

var role_doc = `{
		"Statement": [
			{
				"Action": "sts:AssumeRole",
				"Effect": "Allow",
				"Principal": {
					"RAM": [
						"acs:ram::%s:root"
					]
				}
			}
		],
		"Version": "1"
	}`

var (
	username = "test-user-" + os.Getenv("CONCURRENT_ID")
	rolename = "test-role-" + os.Getenv("CONCURRENT_ID")
	rolearn  = fmt.Sprintf("acs:ram::%s:role/%s", os.Getenv("USER_ID"), rolename)
)

var ecsEndpoint = "ecs." + os.Getenv("REGION_ID") + ".aliyuncs.com"

func createRole(userid string) (name string, arn string, err error) {
	listRequest := ram.CreateListRolesRequest()
	listRequest.Scheme = "HTTPS"
	client, err := ram.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		return
	}
	listResponse, err := client.ListRoles(listRequest)
	if err != nil {
		return
	}
	for _, role := range listResponse.Roles.Role {
		if strings.ToLower(role.RoleName) == rolename {
			name = role.RoleName
			arn = role.Arn
			return
		}
	}
	createRequest := ram.CreateCreateRoleRequest()
	createRequest.Scheme = "HTTPS"
	createRequest.RoleName = rolename
	createRequest.AssumeRolePolicyDocument = fmt.Sprintf(role_doc, userid)
	res, err := client.CreateRole(createRequest)
	if err != nil {
		return
	}
	name = res.Role.RoleName
	arn = res.Role.Arn
	return
}

func createUser() (err error) {
	listRequest := ram.CreateListUsersRequest()
	listRequest.Scheme = "HTTPS"
	client, err := ram.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		return
	}
	listResponse, err := client.ListUsers(listRequest)
	if err != nil {
		return
	}
	for _, user := range listResponse.Users.User {
		if user.UserName == username {
			return
		}
	}
	createRequest := ram.CreateCreateUserRequest()
	createRequest.Scheme = "HTTPS"
	createRequest.UserName = username
	_, err = client.CreateUser(createRequest)
	return
}

func createAttachPolicyToUser() error {
	listRequest := ram.CreateListPoliciesForUserRequest()
	listRequest.UserName = username
	listRequest.Scheme = "HTTPS"
	client, err := ram.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
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
	createRequest.UserName = username
	createRequest.PolicyType = "System"
	_, err = client.AttachPolicyToUser(createRequest)
	if err != nil {
		return err
	}
	return nil
}

func createAttachPolicyToRole() error {
	listRequest := ram.CreateListPoliciesForRoleRequest()
	listRequest.RoleName = rolename
	listRequest.Scheme = "HTTPS"
	client, err := ram.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		return err
	}
	listResponse, err := client.ListPoliciesForRole(listRequest)
	if err != nil {
		return err
	}
	for _, policy := range listResponse.Policies.Policy {
		if policy.PolicyName == "AdministratorAccess" {
			return nil
		}
	}
	createRequest := ram.CreateAttachPolicyToRoleRequest()
	createRequest.Scheme = "HTTPS"
	createRequest.PolicyName = "AdministratorAccess"
	createRequest.RoleName = rolename
	createRequest.PolicyType = "System"
	_, err = client.AttachPolicyToRole(createRequest)
	if err != nil {
		return err
	}
	return nil
}

func createAccessKey() (id string, secret string, err error) {
	client, err := ram.NewClientWithAccessKey(os.Getenv("REGION_ID"), os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		return
	}
	listrequest := ram.CreateListAccessKeysRequest()
	listrequest.UserName = username
	listrequest.Scheme = "HTTPS"
	listresponse, err := client.ListAccessKeys(listrequest)
	if err != nil {
		return
	}
	if listresponse.AccessKeys.AccessKey != nil {
		if len(listresponse.AccessKeys.AccessKey) >= 2 {
			accesskey := listresponse.AccessKeys.AccessKey[0]
			deleterequest := ram.CreateDeleteAccessKeyRequest()
			deleterequest.UserAccessKeyId = accesskey.AccessKeyId
			deleterequest.UserName = username
			deleterequest.Scheme = "HTTPS"
			_, err = client.DeleteAccessKey(deleterequest)
			if err != nil {
				return
			}
		}
	}
	request := ram.CreateCreateAccessKeyRequest()
	request.Scheme = "HTTPS"
	request.UserName = username
	response, err := client.CreateAccessKey(request)
	if err != nil {
		return
	}

	id = response.AccessKey.AccessKeyId
	secret = response.AccessKey.AccessKeySecret
	return
}

func createAssumeRole() (response *sts.AssumeRoleResponse, err error) {
	err = createUser()
	if err != nil {
		return
	}

	subaccesskeyid, subaccesskeysecret, err := createAccessKey()
	if err != nil {
		return
	}

	_, _, err = createRole(os.Getenv("USER_ID"))
	if err != nil {
		return
	}
	err = createAttachPolicyToUser()
	if err != nil {
		return
	}
	request := sts.CreateAssumeRoleRequest()
	request.RoleArn = rolearn
	request.RoleSessionName = "alice_test"
	request.Scheme = "HTTPS"
	client, err := sts.NewClientWithAccessKey(os.Getenv("REGION_ID"), subaccesskeyid, subaccesskeysecret)
	if err != nil {
		return
	}
	response, err = client.AssumeRole(request)
	return
}
