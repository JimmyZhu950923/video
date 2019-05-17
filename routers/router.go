package routers

import (
	"tencent/controllers/livevideocontrollers"
	"tencent/controllers/videoondemandcontrollers"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/aiSample",
			beego.NSInclude(
				&videoondemandcontrollers.AiSampleController{},
			),
		),
		beego.NSNamespace("event",
			beego.NSInclude(
				&videoondemandcontrollers.EventNotificationController{},
			),
		),
		beego.NSNamespace("/task",
			beego.NSInclude(
				&videoondemandcontrollers.TaskManagementController{},
			),
		),
		beego.NSNamespace("/others",
			beego.NSInclude(
				&videoondemandcontrollers.OthersController{},
			),
		),
		beego.NSNamespace("/param",
			beego.NSInclude(
				&videoondemandcontrollers.ParamTemplateController{},
			),
		),
		beego.NSNamespace("/media",
			beego.NSInclude(
				&videoondemandcontrollers.MediaController{},
			),
		),
		beego.NSNamespace("/videoUpload",
			beego.NSInclude(
				&videoondemandcontrollers.VideoUploadController{},
			),
		),
		beego.NSNamespace("/videoClassification",
			beego.NSInclude(
				&videoondemandcontrollers.VideoClassificationController{},
			),
		),
		beego.NSNamespace("/videoProcess",
			beego.NSInclude(
				&videoondemandcontrollers.VideoProcessController{},
			),
		),
		beego.NSNamespace("/domain",
			beego.NSInclude(
				&livevideocontrollers.DomainManagementController{},
			),
		),
		beego.NSNamespace("/log",
			beego.NSInclude(
				&livevideocontrollers.RealTimeLogController{},
			),
		),
		beego.NSNamespace("/delay",
			beego.NSInclude(
				&livevideocontrollers.DelayManagementController{},
			),
		),
		beego.NSNamespace("/record",
			beego.NSInclude(
				&livevideocontrollers.RecordManagementController{},
			),
		),
		beego.NSNamespace("/screenshots",
			beego.NSInclude(
				&livevideocontrollers.ScreenshotsController{},
			),
		),
		beego.NSNamespace("/callback",
			beego.NSInclude(
				&livevideocontrollers.LiveCallbackController{},
			),
		),
		beego.NSNamespace("/pull",
			beego.NSInclude(
				&livevideocontrollers.LivePullStreamController{},
			),
		),
		beego.NSNamespace("/stream",
			beego.NSInclude(
				&livevideocontrollers.LiveStreamManagementController{},
			),
		),
		beego.NSNamespace("/stream",
			beego.NSInclude(
				&livevideocontrollers.LiveTranscodeController{},
			),
		),
		beego.NSNamespace("/count",
			beego.NSInclude(
				&livevideocontrollers.CountQueryController{},
			),
		),
		beego.NSNamespace("/cert",
			beego.NSInclude(
				&livevideocontrollers.CertManagementController{},
			),
		),
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&livevideocontrollers.AuthManagementController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
