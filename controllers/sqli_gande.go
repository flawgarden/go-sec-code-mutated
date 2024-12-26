

package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"go-sec-code/models"
	"xorm.io/xorm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type SqlInjectionVuln1ControllerGande struct {
	beego.Controller
}

type SqlInjectionVuln2ControllerGande struct {
	beego.Controller
}

type SqlInjectionVuln3ControllerGande struct {
	beego.Controller
}

type SqlInjectionVuln4ControllerGande struct {
	beego.Controller
}

type SqlInjectionSafe1ControllerGande struct {
	beego.Controller
}

type SqlInjectionSafe2ControllerGande struct {
	beego.Controller
}

type SqlInjectionSafe3ControllerGande struct {
	beego.Controller
}

func (c *SqlInjectionVuln1ControllerGande) Get() {
	id := c.GetString("id")
	id = foo2(id)
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

func (c *SqlInjectionVuln2ControllerGande) Get() {
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

func (c *SqlInjectionVuln3ControllerGande) Get() {
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

func (c *SqlInjectionVuln4ControllerGande) Get() {
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

func (c *SqlInjectionSafe1ControllerGande) Get() {
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

func (c *SqlInjectionSafe2ControllerGande) Get() {
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

func (c *SqlInjectionSafe3ControllerGande) Get() {
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
