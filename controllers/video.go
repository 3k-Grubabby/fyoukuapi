package controllers

import (
	"github.com/astaxie/beego"
	"youku/models"
)

type VideoController struct {
	beego.Controller
}

//频道页 - 获取顶部广告
// @router /channel/advert [*]
func (this *VideoController) ChannelAdvert() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	num, videos, err := models.GetChannelAdvert(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "请求数据失败,请稍后重试")
		this.ServeJSON()
	}

}

// 频道页--获取正在热播
// @router /channel/hot [*]
func (this *VideoController) ChannelHotList() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	num, videos, err := models.GetChannelHotList(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}

}

// 频道页-根据频道地区获取推荐的视频
// @router /channel/recommend/region [*]
func (this *VideoController) ChannelRecommendRegionList() {
	channelId, _ := this.GetInt("channelId")
	regionId, _ := this.GetInt("regionId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	if regionId == 0 {
		this.Data["json"] = ReturnError(4002, "必须指定频道地区")
		this.ServeJSON()
	}
	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

//频道页-根据频道类型获取推荐视频
// @router /channel/recommend/type [*]
func (this *VideoController) GetChannelRecomendTypeList() {
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	if typeId == 0 {
		this.Data["json"] = ReturnError(4002, "必须指定频道类型")
		this.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}

}

//根据传入参数获取视频列表
// @router /channel/video [*]
func (this *VideoController) ChannelVideo() {
	//获取频道ID
	channelId, _ := this.GetInt("channelId")
	//获取频道地区ID
	regionId, _ := this.GetInt("regionId")
	//获取频道类型ID
	typeId, _ := this.GetInt("typeId")
	//获取状态
	end := this.GetString("end")
	//获取排序
	sort := this.GetString("sort")
	//获取页码信息
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}

	if limit == 0 {
		limit = 12
	}

	num, videos, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}

}

// 获取视频详情
// @router /video/info [*]
func (this *VideoController) VideoInfo() {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频id")
		this.ServeJSON()
	}
	video, err := models.GetVideoInfo(videoId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", video, 1)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试")
		this.ServeJSON()
	}
}

//获取视频剧集列表
// @router /video/episodes/list [*]
func (this *VideoController) VideoEpisodesList() {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频ID")
		this.ServeJSON()
	}
	num, episodes, err := models.GetVideoEpisodesList(videoId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", episodes, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试~")
		this.ServeJSON()
	}
}
