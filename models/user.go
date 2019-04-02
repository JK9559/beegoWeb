package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math"
)

type User struct {
	Id       int64  `orm:"auto"`
	Name     string `orm:"size(100)"`
	Nickname string `orm:"size(100)"`
	Pwd      string `orm:"size(100)"`
	Email    string `orm:"size(100)"`
	Sex      string `orm:"size(2)"`
	Roleid   string `orm:"size(100)"`
	Status   int64
	Phone    string `orm:"size(16)"`
}

func Create(uid int64, name string, nickname string, pwd string, email string,
	sex string, roleid string, status int64, phone string, ) (user User) {

	user, err := QueryById(uid)
	if err == true {
		return user
	} else {
		o := orm.NewOrm()
		o.Using("default")
		newusr := new(User)
		newusr.Id = uid
		fmt.Println(uid)
		newusr.Name = name
		newusr.Nickname = nickname
		newusr.Pwd = pwd
		newusr.Email = email
		newusr.Sex = sex
		newusr.Roleid = roleid
		newusr.Status = status
		newusr.Phone = phone
		_, err := o.Insert(newusr)
		if err != nil {
			fmt.Println(err)
		}
		return *newusr

	}
}

func QueryById(uid int64) (User, bool) {
	o := orm.NewOrm()
	u := User{Id: uid}

	err := o.Read(&u)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return u, false
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return u, false
	} else {
		fmt.Println(u.Id, u.Name)
		return u, true
	}
}

func DeleteById(id int64) bool {
	o := orm.NewOrm()
	o.Using("default")
	if num, err := o.Delete(&User{Id: id}); err == nil {
		fmt.Println("影响的行数为：")
		fmt.Println(num)
		return true
	} else {
		return false
	}
}

func UpdateById(id int, table string, filed map[string]interface{}) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(table).Filter("Id", id).Update(filed)
	if err == nil {
		return true
	}
	return false
}

func QueryByName(name string) (User, error) {
	var user User

	o := orm.NewOrm()
	qs := o.QueryTable("user")

	err := qs.Filter("Name", name).One(&user)
	fmt.Println(err)
	if err == nil {
		fmt.Println(user.Name)
		return user, nil
	}
	return user, err
}

//根据用户数据列表
func DataList() (users []User) {

	o := orm.NewOrm()
	qs := o.QueryTable("user")

	var us []User
	cnt, err := qs.Filter("id__gt", 0).OrderBy("-id").Limit(10, 0).All(&us)
	if err == nil {
		fmt.Printf("count", cnt)
	}
	return us
}

//查询语句，sql语句的执行
//格式类似于:o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "slene")
//
func QueryBySql(sql string, qarms [] string) bool {

	o := orm.NewOrm()

	//执行sql语句
	o.Raw(sql, qarms)

	return true
}

//根据用户分页数据列表
func LimitList(pagesize int, pageno int) (users []User) {

	o := orm.NewOrm()
	qs := o.QueryTable("user")

	var us []User
	cnt, err := qs.Limit(pagesize, (pageno-1)*pagesize).All(&us)
	if err == nil {
		fmt.Printf("count", cnt)
	}
	return us
}

//根据用户数据总个数
func GetDataNum() int64 {

	o := orm.NewOrm()
	qs := o.QueryTable("user")

	var us []User
	num, err := qs.Filter("id__gt", 0).All(&us)
	if err == nil {
		return num
	} else {
		return 0
	}
}

func Paginator(page, prepage int, nums int64) map[string]interface{} {

	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
		//fmt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	return paginatorMap
}

func init() {
	orm.RegisterModel(new(User))
}
