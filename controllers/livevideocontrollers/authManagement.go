package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 鉴权管理
type AuthManagementController struct {
	beego.Controller
}

// @Title 查询播放鉴权key
// @Description DescribeLivePlayAuthKey
// @Param DomainName body String false "域名"
// @Success 200 {object} models.User
// @router /deplay [get]
func (a *AuthManagementController) DescribeLivePlayAuthKey() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLivePlayAuthKeyRequest()
	DomainName := a.Input().Get("DomainName")
	params := `{\"DomainName\":`+ DomainName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLivePlayAuthKey(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询推流鉴权key
// @Description DescribeLivePlayAuthKey
// @Param DomainName body String false "域名"
// @Success 200 {object} models.User
// @router /depush [get]
func (a *AuthManagementController) DescribeLivePushAuthKey() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLivePushAuthKeyRequest()
	DomainName := a.Input().Get("DomainName")
	params := `{\"DomainName\":`+ DomainName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLivePushAuthKey(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改播放鉴权key
// @Description ModifyLivePlayAuthKey
// @Param DomainName body String false "域名"
// @Success 200 {object} models.User
// @router /mplay [get]
func (a *AuthManagementController) ModifyLivePlayAuthKey() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewModifyLivePlayAuthKeyRequest()
	DomainName := a.Input().Get("DomainName")
	params := `{\"DomainName\":`+ DomainName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyLivePlayAuthKey(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改推流鉴权key
// @Description ModifyLivePushAuthKey
// @Param DomainName body String false "域名"
// @Success 200 {object} models.User
// @router /mpush [get]
func (a *AuthManagementController) ModifyLivePushAuthKey() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewModifyLivePushAuthKeyRequest()
	DomainName := a.Input().Get("DomainName")
	params := `{\"DomainName\":`+ DomainName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyLivePushAuthKey(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}