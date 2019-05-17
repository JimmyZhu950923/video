package videoondemandcontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

// 视频处理
type VideoProcessController struct {
	beego.Controller
}

// @Title 编辑视频
// @Description EditMedia
// @Param InputType body String false "输入视频的类型，可以取的值为 File，Stream 两种"
// @Success 200 {object} models.User
// @router /em [get]
func (v *VideoProcessController) EditMedia() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewEditMediaRequest()
	InputType := v.Input().Get("InputType")
	params := `{\"InputType\":`+ InputType +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.EditMedia(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 视频处理
// @Description ProcessMedia
// @Param FileId body String false "媒体文件 ID"
// @Success 200 {object} models.User
// @router /pm [get]
func (v *VideoProcessController) ProcessMedia() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewProcessMediaRequest()
	FileId := v.Input().Get("FileId")
	params := `{\"FileId\":`+ FileId +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ProcessMedia(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 使用任务流模板进行视频处理
// @Description ProcessMediaByProcedure
// @Param FileId body String false "媒体文件 ID"
// @Param ProcedureName body String false "[任务流模板](/document/product/266/11700#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF)名字"
// @Success 200 {object} models.User
// @router /pmbp [get]
func (v *VideoProcessController) ProcessMediaByProcedure() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewProcessMediaByProcedureRequest()
	FileId := v.Input().Get("FileId")
	ProcedureName := v.Input().Get("ProcedureName")
	params := `{\"FileId\":`+ FileId +`,\"ProcedureName\":`+ ProcedureName +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ProcessMediaByProcedure(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 对指定URL的视频发起视频处理
// @Description ProcessMediaByUrl
// @Success 200 {object} models.User
// @router /pmbu [get]
func (v *VideoProcessController) ProcessMediaByUrl() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewProcessMediaByUrlRequest()

	params := `{}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ProcessMediaByUrl(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}