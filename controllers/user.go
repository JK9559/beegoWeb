package controllers

import (
	"beegoWeb/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	userlist := models.DataList()
	fmt.Println("用户列表数据：")
	fmt.Println(userlist)
	c.Data["datas"] = userlist
	c.TplName = "home.html"
}

type DeleteHomeController struct {
	beego.Controller
}

func (c *DeleteHomeController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	id, _ := c.GetInt64("Id")
	fmt.Println(id)
	isdelete := models.DeleteById(id)
	if isdelete {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}
	c.Redirect("/", 302)
}

type EditHomeController struct {
	beego.Controller
}

func (c *EditHomeController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	//用户列表
	idval, errId := strconv.ParseInt(c.GetString("Id"), 10, 64);
	if errId != nil {
		fmt.Println("缺少参数id");
	}
	user, err := models.QueryById(idval)
	if err == true {
		fmt.Println("获取模型失败");
		fmt.Println(err);
	} else {
		fmt.Println("获取模型成功");
	}
	/*返回json数据*/
	c.Data["data"] = user
	c.TplName = "edit.html"
}

type UpdateHomeController struct {
	beego.Controller
}

func (c *UpdateHomeController) Post() {
	o := orm.NewOrm()
	o.Using("default")

	fields := make(map[string]interface{})

	//获取参数
	id, _ := c.GetInt("Id")
	name := c.GetString("name")
	nickname := c.GetString("nickname")
	pwd := c.GetString("pwd")
	email := c.GetString("email")
	sex := c.GetString("sex")
	phone := c.GetString("phone")
	roleid := "1";
	status, _ := strconv.ParseInt("1", 10, 64);
	//新增用户
	if id == 0 {
		idval, _ := strconv.ParseInt("0", 10, 64);
		//新增一条数据，并给模型赋值
		user := models.Create(idval, name, nickname, pwd, email, sex, roleid, status, phone)
		fmt.Println(user)
	} else {
		fields["Id"] = id
		name := c.GetString("name")
		fields["Name"] = name
		nickname := c.GetString("nickname")
		fields["Nickname"] = nickname
		pwd := c.GetString("pwd")
		fields["Pwd"] = pwd
		sex := c.GetString("sex")
		fields["Sex"] = sex
		email := c.GetString("email")
		fields["Email"] = email
		//更新
		models.UpdateById(id, "user", fields)
	}
	//路径的跳转
	c.Redirect("/Home/List", 302)
}

//控制器
type UserController struct {
	beego.Controller
}

//get请求用户分页列表 执行的方法
func (c *UserController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	//得到当前分页html的数据
	pa, err := c.GetInt("page")
	if err != nil {
		println(err)
	}
	pre_page := 3
	totals := models.GetDataNum()
	res := models.Paginator(pa, pre_page, totals)
	//得到分页user的数据
	userlist := models.LimitList(3, pa)
	c.Data["datas"] = userlist //用户的数据
	c.Data["paginator"] = res  //分页的数据
	c.Data["totals"] = totals  //分页的数据
	c.TplName = "list.html"
}
