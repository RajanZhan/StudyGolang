package dao

import (
	"../global"
	"../models"
	"../tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	logindbname = global.Config.DbInfo.LoginDb     //登录数据库
	logindbtye  = global.Config.DbInfo.LoginDbType //数据库类型
	logindb     *gorm.DB                           //数据库连接
)

func init() {
	global.WgDb.Add(1)
	go LoginDbInit()
}

//初始化包
func LoginDbInit() {
	logindbname = global.CurrPath + logindbname
	//fmt.Println("登录数据库地址:",logindbname)
	tdb, err := gorm.Open(logindbtye, logindbname)
	tools.PanicErr(err, "登录数据库初始化")
	logindb = tdb
	if !logindb.HasTable(&models.Login{}) {
		logindb.CreateTable(&models.Login{})
	}
	//fmt.Println("登录数据库初始化完成")
	global.WgDb.Done()
}

func CheckLogin(login *models.Login) (res int) {
	lg := models.Login{}
	logindb.Where(login).First(&lg)
	if lg.Email != "" {
		res = 1
	}
	return
}

func ExistLogin(email string) (res bool) {
	lg := models.Login{}
	logindb.Where(models.Login{Email: email}).First(&lg)
	if lg.Email != "" {
		res = true
	}
	return
}

func AddLogin(login *models.Login) (err error) {
	if login.Email == "" || login.Password == "" {
		err = global.EmptyUserPwd
		return
	}
	logindb.Save(login)
	return
}