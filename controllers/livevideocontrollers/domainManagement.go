package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 域名管理
type DomainManagementController struct {
	beego.Controller
}

var credential = common.NewCredential(
	" AKIDDGODyhDajeNnyugSwDH8nDNOApEOuUgt",
	"81f4E9omVzE0XpulQq1xgEZgyugHildW",
)

// @Title 添加域名
// @Description AddLiveDomain
// @Param DomainName body String false "域名名称"
// @Param DomainType body String false "域名类型， 0：推流域名， 1：播放域名"
// @Success 200 {object} models.User
// @router /add [get]
func (u *DomainManagementController) AddLiveDomain() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewAddLiveDomainRequest()
	DomainName := u.Input().Get("DomainName")
	DomainType := u.Input().Get("DomainType")
	params := `{\"DomainName\":`+ DomainName +`,\"DomainType\":`+ DomainType +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.AddLiveDomain(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除域名
// @Description DeleteLiveDomain
// @Param DomainName body String false "要删除的域名"
// @Param DomainType body String false "域名类型， 0：推流， 1：播放"
// @Success 200 {object} models.User
// @router /delete [get]
func (u *DomainManagementController) DeleteLiveDomain() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveDomainRequest()
	DomainName := u.Input().Get("DomainName")
	DomainType := u.Input().Get("DomainType")
	params := `{\"DomainName\":`+ DomainName +`,\"DomainType\":`+ DomainType +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveDomain(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询域名信息
// @Description DescribeLiveDomain
// @Success 200 {object} models.User
// @router /ded [get]
func (u *DomainManagementController) DescribeLiveDomain() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveDomainRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveDomain(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 查询域名列表
// @Description DescribeLiveDomain
// @Success 200 {object} models.User
// @router /deds [get]
func (u *DomainManagementController) DescribeLiveDomains() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveDomainsRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveDomains(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 启动域名
// @Description EnableLiveDomain
// @Param DomainName body String false "待启用的直播域名"
// @Success 200 {object} models.User
// @router /enable [get]
func (u *DomainManagementController) EnableLiveDomain() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewEnableLiveDomainRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.EnableLiveDomain(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 禁用域名
// @Description ForbidLiveDomain
// @Param DomainName body String false "停用的直播域名"
// @Success 200 {object} models.User
// @router /Fobid [get]
func (u *DomainManagementController) ForbidLiveDomain() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewForbidLiveDomainRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ForbidLiveDomain(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改播放域名信息
// @Description ModifyLivePlayDomain
// @Param DomainName body String false "播放域名"
// @Param PlayType body String false "拉流域名类型。1-国内；2-全球；3-境外"
// @Success 200 {object} models.User
// @router /modify [get]
func (u *DomainManagementController) ModifyLivePlayDomain() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewModifyLivePlayDomainRequest()
	DomainName := u.Input().Get("DomainName")
	PlayType := u.Input().Get("DomainType")
	params := `{\"DomainName\":`+ DomainName +`,\"PlayType\":`+ PlayType +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyLivePlayDomain(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}