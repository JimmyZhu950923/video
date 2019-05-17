package livevideocontrollers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
)

// 实时日志
type RealTimeLogController struct {
	beego.Controller
}

// @Title 批量获取日志
// @Description DescribeLogDownloadList
// @Param StartTime body String false "开始时间，北京时间。 格式：yyyy-mm-dd HH:MM:SS"
// @Param EndTime body String false "结束时间，北京时间。 格式：yyyy-mm-dd HH:MM:SS。 注意：结束时间 - 开始时间 <=7天"
// @Success 200 {object} models.User
// @router /de [get]
func (r *RealTimeLogController) DescribeLogDownloadList() {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "live.tencentcloudapi.com"
	client, _ := live.NewClient(credential, "ap-shanghai", cpf)

	request := live.NewDescribeLogDownloadListRequest()
	StartTime := r.Input().Get("StartTime")
	EndTime := r.Input().Get("EndTime")
	PlayDomains := r.Input().Get("PlayDomains")
	params := `{\"StartTime\":`+ StartTime +`,\"EndTime\":`+ EndTime +`,\"PlayDomains\":[`+ PlayDomains +`]}`
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeLogDownloadList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
