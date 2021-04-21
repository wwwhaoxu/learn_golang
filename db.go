package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"strings"
)

var ErrorNotFound = errors.New("Record not found")

type User struct {
	id    int64
	name  string
	age   int8
	sex   int8
	phone string
}

//数据库配置
const (
	userName = "henry"
	password = "xxx@1234"
	ip       = "xxx"
	port     = "3306"
	dbName   = "wang"
)

//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB() error {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//fmt.Println(path)
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)

	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		return errors.Wrap(err, "数据库连接失败")
	}
	fmt.Println("connnect success")
	return nil
}

func main() {
	err := InitDB()
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack strace: \n%+v\n", err)
	}

	user, err := Query()
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack strace: \n%+v\n", err)
	}
	fmt.Println(user)

}

func CustomErr(s string) error {
	return errors.Wrap(ErrorNotFound, s)
}

func Query() (User, error) {
	var user User
	err := DB.QueryRow("select * from user where  id = ?", 2).Scan(&user.id, &user.name, &user.age, &user.sex, &user.phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, CustomErr("查询id=1的记录没有找到")
		}
		return User{}, errors.Wrap(err, "不是记录查不到的错误")
	}
	return user, nil
	//rows, err := DB.Query("select * from user where  id = ?", 3)
	//defer rows.Close()
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		return []User{}, CustomErr("记录没有找到")
	//	}
	//	fmt.Println("aaaaa")
	//}
	//var userList []User
	//var user User
	//for rows.Next() {
	//	rows.Scan(&user.id, &user.name, &user.age, &user.sex, &user.phone)
	//	userList = append(userList, user)
	//}
	//return userList, rows.Err()
}
