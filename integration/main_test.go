package integration

import (
	"os"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
)

func TestMain(m *testing.M) {
	// 设置测试前的准备工作
	setup()

	// 运行所有测试
	exitCode := m.Run()

	// 清理工作
	teardown()

	// 退出测试程序，exitCode 为0表示成功，非0表示失败
	os.Exit(exitCode)
}

func setup() {
	// 初始化代码
}

func teardown() {
	// 清理代码
	client, err := newRamClient()
	if err != nil {
		return
	}
	// 查询用户
	getUserRequest := ram.CreateGetUserRequest()
	getUserRequest.UserName = username
	getUserRequest.Scheme = "HTTPS"
	_, err = client.GetUser(getUserRequest)

	if err != nil {
		if se, ok := err.(*errors.ServerError); ok {
			// 如果用户不存在，则返回
			if se.ErrorCode() == "EntityNotExist.User" {
				return
			}
		}
	}

	cleanUpUser()
	cleanUpRole()
}

func cleanUpUser() {
	client, err := newRamClient()
	if err != nil {
		return
	}

	listPoliciesForUserRequest := ram.CreateListPoliciesForUserRequest()
	listPoliciesForUserRequest.UserName = username

	response, err := client.ListPoliciesForUser(listPoliciesForUserRequest)
	for _, v := range response.Policies.Policy {
		r := ram.CreateDetachPolicyFromUserRequest()
		r.UserName = username
		r.PolicyName = v.PolicyName
		r.PolicyType = v.PolicyType
		client.DetachPolicyFromUser(r)
	}

	lakr := ram.CreateListAccessKeysRequest()
	lakr.UserName = username
	accessKeysResponse, err := client.ListAccessKeys(lakr)
	for _, v := range accessKeysResponse.AccessKeys.AccessKey {
		r := ram.CreateDeleteAccessKeyRequest()
		r.UserName = username
		r.UserAccessKeyId = v.AccessKeyId
		client.DeleteAccessKey(r)
	}

	dur := ram.CreateDeleteUserRequest()
	dur.UserName = username
	_, err = client.DeleteUser(dur)
}

func cleanUpRole() {
	client, err := newRamClient()
	if err != nil {
		return
	}
	lpfrr := ram.CreateListPoliciesForRoleRequest()
	lpfrr.RoleName = rolename
	response, err := client.ListPoliciesForRole(lpfrr)
	for _, v := range response.Policies.Policy {
		r := ram.CreateDetachPolicyFromRoleRequest()
		r.RoleName = rolename
		r.PolicyName = v.PolicyName
		r.PolicyType = v.PolicyType
		client.DetachPolicyFromRole(r)
	}
	drr := ram.CreateDeleteRoleRequest()
	drr.RoleName = rolename
	_, err = client.DeleteRole(drr)
}
