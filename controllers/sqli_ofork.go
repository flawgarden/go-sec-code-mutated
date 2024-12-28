

package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type SqlInjectionVuln1ControllerOfork struct {
	beego.Controller
}

type SqlInjectionVuln2ControllerOfork struct {
	beego.Controller
}

type SqlInjectionVuln3ControllerOfork struct {
	beego.Controller
}

type SqlInjectionVuln4ControllerOfork struct {
	beego.Controller
}

type SqlInjectionSafe1ControllerOfork struct {
	beego.Controller
}

type SqlInjectionSafe2ControllerOfork struct {
	beego.Controller
}

type SqlInjectionSafe3ControllerOfork struct {
	beego.Controller
}

func (c *SqlInjectionVuln1ControllerOfork) Get() {
	id := c.GetString("id")

	set111 := make(map[string]struct{})
	set222 := make(map[string]struct{})
	set111[id] = struct{}{}
	set222[id] = struct{}{}
	set222[id] = struct{}{}
	for k := range set222 {
		delete(set111, k)
	}
	value7849 := "GWiav"
	if len(set111) == 0 {
		value7849 = id
	}
	id = value7849

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *SqlInjectionVuln2ControllerOfork) Get() {
	username := c.GetString("username")

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStr := fmt.Sprintf("select * from user where username=\"%s\"", username)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *SqlInjectionVuln3ControllerOfork) Get() {
	username := c.GetString("username")
	field := c.GetString("field")
	engine, err := xorm.NewEngine("mysql", source)
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)
	user := models.User{}
	session := engine.Prepare().And(fmt.Sprintf("%s like ?", field), username)
	ok, err := session.Get(&user)
	if !ok && err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *SqlInjectionVuln4ControllerOfork) Get() {
	username := c.GetString("username")
	order := c.GetString("order")
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	expression := sq.Select("*").From("user").Where(sq.Eq{"username": username}).OrderBy(order)
	sqlStr, args, err := expression.ToSql()
	fmt.Println(sqlStr)
	if err != nil {
		panic(err)
	}
	user := models.User{}
	err = db.QueryRow(sqlStr, args...).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *SqlInjectionSafe1ControllerOfork) Get() {
	id, err := c.GetInt("id", 1)
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStr := "select * from user where id=?"
	user := models.User{}
	err = db.QueryRow(sqlStr, id).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *SqlInjectionSafe2ControllerOfork) Get() {
	username := c.GetString("username")
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStr := "select * from user where username=?"
	user := models.User{}
	err = db.QueryRow(sqlStr, username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *SqlInjectionSafe3ControllerOfork) Get() {
	username := c.GetString("username")
	field := c.GetString("field")
	o := orm.NewOrm()
	user := models.User{}
	cond := orm.NewCondition().And(field+"__icontains", username)
	qs := o.QueryTable(&models.User{})
	err := qs.SetCond(cond).One(&user)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}
