package admin

import (
	"github.com/gin-gonic/gin"
	"ws/app/http/requests"
	"ws/app/repositories"
	"ws/app/resource"
	"ws/app/util"
)

type SettingHandler struct {
}

func (handler *SettingHandler) Update(c *gin.Context) {
	var form = struct {
		Value string `json:"value" form:"value" binding:"required"`
	}{}
	err := c.ShouldBind(&form)
	if err != nil {
		util.RespValidateFail(c, err.Error())
		return
	}
	admin := requests.GetAdmin(c)
	id := c.Param("id")
	setting := repositories.ChatSettingRepo.First([]*repositories.Where{
		{
		Filed: "group_id = ?",
		Value: admin.GetGroupId(),
		},
		{
			Filed: "id = ?",
			Value: id,
		},
	}, []string{})
	if setting == nil {
		util.RespNotFound(c)
		return
	}
	setting.Value = form.Value
	repositories.ChatSettingRepo.Save(setting)
	util.RespSuccess(c, gin.H{})
}

func (handler *SettingHandler) Index(c *gin.Context) {
	admin := requests.GetAdmin(c)
	settings := repositories.ChatSettingRepo.Get([]*repositories.Where{
		{
			Filed: "group_id = ?",
			Value: admin.GetGroupId(),
		},
	}, -1, []string{}, []string{})
	resp := make([]*resource.ChatSetting, len(settings), len(settings))
	for index, setting := range settings {
		resp[index] = setting.ToJson()
	}
	util.RespSuccess(c, resp)
}