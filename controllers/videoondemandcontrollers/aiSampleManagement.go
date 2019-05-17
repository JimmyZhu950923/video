package videoondemandcontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

// AI样本管理
type AiSampleController struct {
	beego.Controller
}

var credential = common.NewCredential(
" AKIDDGODyhDajeNnyugSwDH8nDNOApEOuUgt",
"81f4E9omVzE0XpulQq1xgEZgyugHildW",
)

// @Title 创建任务样本
// @Description CreatePersonSample
// @Param Name body string false "人物名称"
// @Param FaceContents body string false "人脸图片"
// @Param Usages body string false "人物应用场景"
// @Success 200 {object} models.User
// @router /cps [get]
func (a *AiSampleController) CreatePersonSample() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)
	request := vod.NewCreatePersonSampleRequest()

	Name := a.Input().Get("Name")
	FaceContents := a.Input().Get("FaceContents")
	Usages := a.Input().Get("Usages")

	params := `{\"Name\":`+ Name +`,\"FaceContents\":[`+ FaceContents +`],\"Usages\":[`+ Usages +`]}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreatePersonSample(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 创建关键词样本
// @Description CreateWordSamples
// @Param Usages body string false "关键词应用场景"
// @Param Keyword body string false "关键词"
// @Param Usages body string false "关键词标签"
// @Success 200 {object} models.User
// @router /cws [get]
func (a *AiSampleController) CreateWordSamples() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewCreateWordSamplesRequest()
	Usages := a.Input().Get("Usages")
	Keyword := a.Input().Get("Keyword")
	Tags := a.Input().Get("Tags")
	params := `{\"Usages\":[`+ Usages +`],\"Words\":[{\"Keyword\":`+ Keyword +`,\"Tags\":[`+ Tags +`]}]}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateWordSamples(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除人物样本
// @Description DeletePersonSample
// @Param PersonId body string false "人物ID"
// @Success 200 {object} models.User
// @router /dps [get]
func (a *AiSampleController) DeletePersonSample() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDeletePersonSampleRequest()
	PersonId := a.Input().Get("PersonId")
	params := `{\"PersonId\":`+ PersonId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeletePersonSample(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 删除关键词样本
// @Description DeleteWordSample
// @Param Keywords body string false "关键词"
// @Success 200 {object} models.User
// @router /dws [get]
func (a *AiSampleController) DeleteWordSamples() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDeleteWordSamplesRequest()
	Keywords := a.Input().Get("Keywords")
	params := `{\"Keywords\":[`+ Keywords +`]}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteWordSamples(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取人物样本列表
// @Description DescribePersonSamples
// @Success 200 {object} models.User
// @router /deps [get]
func (a *AiSampleController) DescribePersonSamples() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDescribePersonSamplesRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribePersonSamples(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 获取关键词样本列表
// @Description DescribeWordSamples
// @Success 200 {object} models.User
// @router /dews [get]
func (a *AiSampleController) DescribeWordSamples() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewDescribeWordSamplesRequest()
	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeWordSamples(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改人物样本
// @Description ModifyPersonSample
// @Param PersonId body string false "人物ID"
// @Success 200 {object} models.User
// @router /mps [get]
func (a *AiSampleController) ModifyPersonSample() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewModifyPersonSampleRequest()
	PersonId := a.Input().Get("PersonId")
	params := `{\"PersonId\":`+ PersonId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyPersonSample(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 修改关键词样本
// @Description ModifyWordSample
// @Param Keyword body string false "关键词"
// @Success 200 {object} models.User
// @router /mws [get]
func (a *AiSampleController) ModifyWordSample() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewModifyWordSampleRequest()
	Keyword := a.Input().Get("Keyword")
	params := `{\"Keyword\":`+ Keyword +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ModifyWordSample(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}