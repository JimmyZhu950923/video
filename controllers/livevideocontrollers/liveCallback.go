package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 直播回调
type LiveCallbackController struct {
	beego.Controller
}

// @Title 创建回调规则
// @Description CreateLiveCallbackRule
// @Param DomainName body String false "推流域名"
// @Param AppName body String false "推流路径"
// @Param TemplateId body Integer false "模板ID"
// @Success 200 {object} models.User
// @router /clcr [get]
func (l *LiveCallbackController) CreateLiveCallbackRule() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewCreateLiveCallbackRuleRequest()
	DomainName := l.Input().Get("DomainName")
	AppName := l.Input().Get("AppName")
	TemplateId := l.Input().Get("TemplateId")
	params := `{\"DomainName\":`+ DomainName +`,\"AppName\":`+ AppName +`,\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateLiveCallbackRule(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 创建回调模板
// @Description CreateLiveCallbackTemplate
// @Param TemplateName body String false "模板名称。非空的字符串"
// @Success 200 {object} models.User
// @router /clct [get]
func (l *LiveCallbackController) CreateLiveCallbackTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewCreateLiveCallbackTemplateRequest()
	TemplateName := l.Input().Get("TemplateName")
	params := `{\"TemplateName\":`+ TemplateName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateLiveCallbackTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除回调规则
// @Description DeleteLiveCallbackRule
// @Param DomainName body String false "推流域名"
// @Param AppName body String false "推流路径"
// @Success 200 {object} models.User
// @router /dlcr [get]
func (l *LiveCallbackController) DeleteLiveCallbackRule() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveCallbackRuleRequest()
	DomainName := l.Input().Get("DomainName")
	AppName := l.Input().Get("AppName")
	params := `{\"DomainName\":`+ DomainName +`,\"AppName\":`+ AppName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveCallbackRule(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除回调模板
// @Description DeleteLiveCallbackTemplate
// @Param TemplateId body Integer false "模板Id"
// @Success 200 {object} models.User
// @router /dlct [get]
func (l *LiveCallbackController) DeleteLiveCallbackTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveCallbackTemplateRequest()
	TemplateId := l.Input().Get("TemplateId")
	params := `{\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveCallbackTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取回调规则列表
// @Description DescribeLiveCallbackRules
// @Success 200 {object} models.User
// @router /dlcr [get]
func (l *LiveCallbackController) DescribeLiveCallbackRules() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveCallbackRulesRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveCallbackRules(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 回调单个回调模板
// @Description DescribeLiveCallbackTemplate
// @Param TemplateId body Integer false "模板Id"
// @Success 200 {object} models.User
// @router /dlct [get]
func (l *LiveCallbackController) DescribeLiveCallbackTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveCallbackTemplateRequest()
	TemplateId := l.Input().Get("TemplateId")
	params := `{\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveCallbackTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取回调模板列表
// @Description DescribeLiveCallbackTemplate
// @Success 200 {object} models.User
// @router /dlcts [get]
func (l *LiveCallbackController) DescribeLiveCallbackTemplates() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveCallbackTemplatesRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveCallbackTemplates(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改回调模板
// @Description ModifyLiveCallbackTemplate
// @Param TemplateId body Integer false "模板Id"
// @Success 200 {object} models.User
// @router /mlct [get]
func (l *LiveCallbackController) ModifyLiveCallbackTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewModifyLiveCallbackTemplateRequest()
	TemplateId := l.Input().Get("TemplateId")
	params := `{\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyLiveCallbackTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}