package videoondemandcontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

// 视屏上传
type VideoUploadController struct {
	beego.Controller
}

// @Title 申请上传
// @Description ApplyUpload
// @Param MediaType body String false "媒体类型"
// @Success 200 {object} models.User
// @router /au [get]
func (v *VideoUploadController) ApplyUpload() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewApplyUploadRequest()
	MediaType := v.Input().Get("MediaType")
	params := `{\"MediaType\":`+ MediaType +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.ApplyUpload(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}

// @Title 确认上传
// @Description CommitUpload
// @Param VodSessionKey body String false "点播会话，取申请上传接口的返回值 VodSessionKey"
// @Success 200 {object} models.User
// @router /cu [get]
func (v *VideoUploadController) CommitUpload() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vod.NewClient(credential, "ap-shanghai", cpf)

	request := vod.NewCommitUploadRequest()
	VodSessionKey := v.Input().Get("VodSessionKey")
	params := `{\"VodSessionKey\":`+ VodSessionKey +`}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CommitUpload(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}