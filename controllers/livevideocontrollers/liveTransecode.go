package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 直播转码
type LiveTranscodeController struct {
	beego.Controller
}

// @Title 创建转码规则
// @Description CreateLiveTranscodeRule
// @Param AppName body String false "推流路径"
// @Param DomainName body String false "推流域名"
// @Param StreamName body String false "流名称"
// @Param TemplateId body Integer false "指定已有的模板Id"
// @Success 200 {object} models.User
// @router /create [get]
func (l *LiveTranscodeController) CreateLiveTranscodeRule() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewCreateLiveTranscodeRuleRequest()
	DomainName := l.Input().Get("DomainName")
	AppName := l.Input().Get("AppName")
	StreamName := l.Input().Get("StreamName")
	TemplateId := l.Input().Get("TemplateId")
	params := `{\"DomainName\":`+ DomainName +`,\"AppName\":`+ AppName +`,\"StreamName\":`+ StreamName +`,\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateLiveTranscodeRule(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 创建转码模板
// @Description CreateLiveTranscodeTemplate
// @Param TemplateName body String false "模板名称"
// @Param VideoBitrate body Integer false "视频码率"
// @Success 200 {object} models.User
// @router /cltt [get]
func (l *LiveTranscodeController) CreateLiveTranscodeTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewCreateLiveTranscodeTemplateRequest()
	TemplateName := l.Input().Get("TemplateName")
	VideoBitrate := l.Input().Get("VideoBitrate")
	params := `{\"TemplateName\":`+ TemplateName +`,\"VideoBitrate\":`+ VideoBitrate +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateLiveTranscodeTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除转码规则
// @Description DeleteLiveTranscodeRule
// @Param AppName body String false "推流路径"
// @Param DomainName body String false "推流域名"
// @Param StreamName body String false "流名称"
// @Param TemplateId body Integer false "模板Id"
// @Success 200 {object} models.User
// @router /delete [get]
func (l *LiveTranscodeController) DeleteLiveTranscodeRule() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveTranscodeRuleRequest()
	DomainName := l.Input().Get("DomainName")
	AppName := l.Input().Get("AppName")
	StreamName := l.Input().Get("StreamName")
	TemplateId := l.Input().Get("TemplateId")
	params := `{\"DomainName\":`+ DomainName +`,\"AppName\":`+ AppName +`,\"StreamName\":`+ StreamName +`,\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveTranscodeRule(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除转码模板
// @Description DeleteLiveTranscodeTemplate
// @Param TemplateId body Integer false "模板Id"
// @Success 200 {object} models.User
// @router /dltt [get]
func (l *LiveTranscodeController) DeleteLiveTranscodeTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveTranscodeTemplateRequest()
	TemplateId := l.Input().Get("TemplateId")
	params := `{\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveTranscodeTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取转码规则列表
// @Description DescribeLiveTranscodeRules
// @Success 200 {object} models.User
// @router /dltr [get]
func (l *LiveTranscodeController) DescribeLiveTranscodeRules() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveTranscodeRulesRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveTranscodeRules(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取单个转码模板
// @Description DescribeLiveTranscodeTemplate
// @Param TemplateId body Integer false "模板Id"
// @Success 200 {object} models.User
// @router /de [get]
func (l *LiveTranscodeController) DescribeLiveTranscodeTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveTranscodeTemplateRequest()
	TemplateId := l.Input().Get("TemplateId")
	params := `{\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveTranscodeTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取转码模板列表
// @Description DescribeLiveTranscodeTemplates
// @Success 200 {object} models.User
// @router /debtt [get]
func (l *LiveTranscodeController) DescribeLiveTranscodeTemplates() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveTranscodeTemplatesRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveTranscodeTemplates(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改转码模板配置
// @Description ModifyLiveTranscodeTemplate
// @Param TemplateId body Integer false "模板Id"
// @Success 200 {object} models.User
// @router /modify [get]
func (l *LiveTranscodeController) ModifyLiveTranscodeTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewModifyLiveTranscodeTemplateRequest()
	TemplateId := l.Input().Get("TemplateId")
	params := `{\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyLiveTranscodeTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}