

package controllers

import (
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type PathTraversalVuln1ControllerLidej struct {
	beego.Controller
}

type PathTraversalVuln2ControllerLidej struct {
	beego.Controller
}

type PathTraversalSafe1ControllerLidej struct {
	beego.Controller
}

type PathTraversalSafe2ControllerLidej struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerLidej) Get() {
	file := c.GetString("file")
	

var a12341 BaseBinaryOpClass
if strings.EqualFold("OLOpQ", "YXMJl") {
    a12341 = &DerivedBinaryOpClass1{}
} else {
    a12341 = &DerivedBinaryOpClass2{}
}
file = a12341.VirtualCall(file, file)


	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerLidej) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerLidej) Get() {
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

func (c *PathTraversalSafe2ControllerLidej) Get() {
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
