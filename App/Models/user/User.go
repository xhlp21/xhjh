package user

import (
	"gin/Config"
	"log"
	"strconv"
)

type User struct {
	UserId     int    `json:"user_id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
}

// Login 账号注册
func (u User) Login() string {
	var count int
	err := Config.SqlDB.QueryRow("SELECT COUNT(*) FROM xhj_user WHERE account = ?", u.Account).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		return "账号已存在！"
	}

	res, err := Config.SqlDB.Exec("INSERT INTO xhj_user (name, image, account, password, create_time, update_time) VALUES (?, ?, ?, ?, ?, ?)", u.Name, u.Image, u.Account, u.Password, u.CreateTime, u.UpdateTime)
	if err != nil {
		log.Fatal(err)
	}

	ids, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return "账号注册成功！自增ID为：" + strconv.FormatInt(ids, 10)
}

// Log 账号登录
func (u User) Log() string {
	var count int
	err := Config.SqlDB.QueryRow("SELECT COUNT(*) FROM xhj_user WHERE account = ?", u.Account).Scan(&count)
	if err != nil {
		return "查询出错：" + err.Error()
	}
	if count == 0 {
		return "账号不存在！"
	}
	err = Config.SqlDB.QueryRow("SELECT COUNT(*) FROM xhj_user WHERE account = ? AND password = ?", u.Account, u.Password).Scan(&count)
	if err != nil {
		return "查询出错：" + err.Error()
	}
	if count == 0 {
		return "账号或密码错误！"
	}
	return "登录成功！"
}
