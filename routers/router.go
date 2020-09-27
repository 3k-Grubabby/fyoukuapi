package routers

import (
	"github.com/astaxie/beego"
	"youku/controllers"
)

func init() {
	beego.Include(&controllers.UserControler{})
	beego.Include(&controllers.VideoController{})
	beego.Include(&controllers.BaseController{})
	beego.Include(&controllers.CommentController{})
}
