package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 录制管理
type RecordManagementController struct {
	beego.Controller
}

// @Title 创建录制任务
// @Description CreateLiveRecord
// @Param StreamName body String false "流名称"
// @Success 200 {object} models.User
// @router /create [get]
func (r *RecordManagementController) CreateLiveRecord() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewCreateLiveRecordRequest()
	StreamName := r.Input().Get("StreamName")
	params := `{\"StreamName\":`+ StreamName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateLiveRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 创建录制规则
// @Description CreateLiveRecordRule
// @Param DomainName body String false "推流名称"
// @Param TemplateId body Integer false "模板ID"
// @Success 200 {object} models.User
// @router /clrr [get]
func (r *RecordManagementController) CreateLiveRecordRule() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewCreateLiveRecordRuleRequest()
	DomainName := r.Input().Get("DomainName")
	TemplateId := r.Input().Get("TemplateId")
	params := `{\"DomainName\":`+ DomainName +`,\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateLiveRecordRule(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 创建录制模板
// @Description CreateLiveRecordTemplate
// @Param TemplateName body String false "模板名"
// @Success 200 {object} models.User
// @router /clrt [get]
func (r *RecordManagementController) CreateLiveRecordTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewCreateLiveRecordTemplateRequest()
	TemplateName := r.Input().Get("TemplateName")
	params := `{\"TemplateName\":`+ TemplateName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateLiveRecordTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除录制任务
// @Description DeleteLiveRecord
// @Param StreamName body String false "流名称"
// @Param TaskId body Integer false "任务ID"
// @Success 200 {object} models.User
// @router /delete [get]
func (r *RecordManagementController) DeleteLiveRecord() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveRecordRequest()
	StreamName := r.Input().Get("StreamName")
	TaskId := r.Input().Get("TaskId")
	params := `{\"StreamName\":`+ StreamName +`,\"TaskId\":`+ TaskId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除录制规则
// @Description DeleteLiveRecordRule
// @Param DomainName body String false "推流名称"
// @Success 200 {object} models.User
// @router /dlrr [get]
func (r *RecordManagementController) DeleteLiveRecordRule() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveRecordRuleRequest()
	DomainName := r.Input().Get("DomainName")
	params := `{\"DomainName\":`+ DomainName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveRecordRule(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除录制模板
// @Description DeleteLiveRecordTemplate
// @Param TemplateId body Integer false "模板ID"
// @Success 200 {object} models.User
// @router /dlrt [get]
func (r *RecordManagementController) DeleteLiveRecordTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDeleteLiveRecordTemplateRequest()
	TemplateId := r.Input().Get("TemplateId")
	params := `{\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteLiveRecordTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取录制规则列表
// @Description DescribeLiveRecordRules
// @Success 200 {object} models.User
// @router /delrr [get]
func (r *RecordManagementController) DescribeLiveRecordRules() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveRecordRulesRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveRecordRules(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取单个录制模板
// @Description DeleteLiveRecordTemplate
// @Param TemplateId body Integer false "模板ID"
// @Success 200 {object} models.User
// @router /delrt [get]
func (r *RecordManagementController) DescribeLiveRecordTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveRecordTemplateRequest()
	TemplateId := r.Input().Get("TemplateId")
	params := `{\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveRecordTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取录制模板列表
// @Description DescribeLiveRecordTemplates
// @Success 200 {object} models.User
// @router /delrt [get]
func (r *RecordManagementController) DescribeLiveRecordTemplates() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLiveRecordTemplatesRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLiveRecordTemplates(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改录制模板配置
// @Description ModifyLiveRecordTemplate
// @Param TemplateId body Integer false "模板ID"
// @Success 200 {object} models.User
// @router /modify [get]
func (r *RecordManagementController) ModifyLiveRecordTemplate() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewModifyLiveRecordTemplateRequest()
	TemplateId := r.Input().Get("TemplateId")
	params := `{\"TemplateId\":`+ TemplateId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyLiveRecordTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 终止录制任务
// @Description StopLiveRecord
// @Param StreamName body String false "流名称"
// @Param TaskId body Integer false "任务ID"
// @Success 200 {object} models.User
// @router /stop [get]
func (r *RecordManagementController) StopLiveRecord() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewStopLiveRecordRequest()
	StreamName := r.Input().Get("StreamName")
	TaskId := r.Input().Get("TaskId")
	params := `{\"StreamName\":`+ StreamName +`,\"TaskId\":`+ TaskId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.StopLiveRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}