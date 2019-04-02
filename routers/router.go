package routers

import (
	"beegoWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/Home/Edit", &controllers.EditHomeController{})
	beego.Router("/Home/Update", &controllers.UpdateHomeController{})
	beego.Router("/Home/Delete", &controllers.DeleteHomeController{})
	beego.Router("/Home/List", &controllers.UserController{})
}
