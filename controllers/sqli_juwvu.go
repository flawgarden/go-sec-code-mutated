// -------------
//
// Original file name: controllers/sqli.go
// Original file CWE's: [89]
// Original file kind: fail
// Mutation info: Insert template from templates-db/languages/go/sensitivity/conditional/while.tmt with name while_operator_positive
// Used extensions:
// Program:
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

type SqlInjectionVuln1ControllerJuwvu struct {
	beego.Controller
}

type SqlInjectionVuln2ControllerJuwvu struct {
	beego.Controller
}

type SqlInjectionVuln3ControllerJuwvu struct {
	beego.Controller
}

type SqlInjectionVuln4ControllerJuwvu struct {
	beego.Controller
}

type SqlInjectionSafe1ControllerJuwvu struct {
	beego.Controller
}

type SqlInjectionSafe2ControllerJuwvu struct {
	beego.Controller
}

type SqlInjectionSafe3ControllerJuwvu struct {
	beego.Controller
}

func (c *SqlInjectionVuln1ControllerJuwvu) Get() {
	id := c.GetString("id")

	tmpUnique42 := id
	for -1912206201 > 0 {
		id = tmpUnique42
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

func (c *SqlInjectionVuln2ControllerJuwvu) Get() {
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

func (c *SqlInjectionVuln3ControllerJuwvu) Get() {
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

func (c *SqlInjectionVuln4ControllerJuwvu) Get() {
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

func (c *SqlInjectionSafe1ControllerJuwvu) Get() {
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

func (c *SqlInjectionSafe2ControllerJuwvu) Get() {
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

func (c *SqlInjectionSafe3ControllerJuwvu) Get() {
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
