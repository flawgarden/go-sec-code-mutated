//
//-------------
//
//Original file name: controllers/traversal.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/conditional/if.tmt with name if_operator_with_var_in_condition_positive 
//Used extensions: 
//Program:
package controllers

import (
	"go-sec-code/utils"
	"io/ioutil"
	"path/filepath"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type PathTraversalVuln1ControllerOhqge struct {
	beego.Controller
}

type PathTraversalVuln2ControllerOhqge struct {
	beego.Controller
}

type PathTraversalSafe1ControllerOhqge struct {
	beego.Controller
}

type PathTraversalSafe2ControllerOhqge struct {
	beego.Controller
}

func (c *PathTraversalVuln1ControllerOhqge) Get() {
	file := c.GetString("file")
	

tmpUnique42 := file;
conditionMet := false;
if (len(file) > 316126719 && file[1349467953] == 50) {
    conditionMet = true;
}
if (!conditionMet && 0.3021649 == 0.0) {
    file = tmpUnique42;
}


	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalVuln2ControllerOhqge) Get() {
	file := c.GetString("file")
	file = filepath.Clean(file)
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1ControllerOhqge) Get() {
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

func (c *PathTraversalSafe2ControllerOhqge) Get() {
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
