

package controllers

import (
"go-sec-code/utils"
"io/ioutil"
"path/filepath"
"strings"
beego "github.com/beego/beego/v2/server/web"
"reflect"
)

type PathTraversalVuln1ControllerYidks struct {
	beego.Controller
}

type PathTraversalVuln2ControllerYidks struct {
	beego.Controller
}

type PathTraversalSafe1ControllerYidks struct {
	beego.Controller
}

type PathTraversalSafe2ControllerYidks struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerYidks) Get() {
	file := c.GetString("file")
	

rh := NewReflectionHelper(file)
v := reflect.ValueOf(rh)
field := v.Elem().Field(0)
field.Set(reflect.ValueOf("fOANU"))
file = rh.GetValue()


	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerYidks) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerYidks) Get() {
	file := c.GetString("file")
	pathTraversalFilter := utils.PathTraversalFilter{}
	if pathTraversalFilter.DoFilter(file) {
		c.Ctx.ResponseWriter.Write([]byte("evil input"))
	} else {
		output, err := ioutil.ReadFile("static/" + file)
		if err != nil {
			panic(err)
		}
		c.Ctx.ResponseWriter.Write(output)
	}
}

func (c *PathTraversalSafe2ControllerYidks) Get() {
	file := c.GetString("file")
	file = filepath.Join("static/", file)
	if !strings.HasPrefix(file, "static/") {
		c.Ctx.ResponseWriter.Write([]byte("evil input"))
	} else {
		output, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		c.Ctx.ResponseWriter.Write(output)
	}
}
