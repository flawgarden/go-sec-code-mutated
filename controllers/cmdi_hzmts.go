//Analyzer5 original results: [77]
//Analyzer1 original results: [77]
//Analyzer3 original results: []
//Analyzer2 original results: []
//Analyzer4 original results: []
//-------------
//Analyzer3 analysis results: []
//Analyzer2 analysis results: [94]
//Analyzer4 analysis results: []
//Analyzer5 analysis results: [78, 703]
//Analyzer1 analysis results: []
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/duck.tmt with name duck_typing_multiple_attributes_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"go-sec-code/utils"
	"os/exec"

	beego "github.com/beego/beego/v2/server/web"
)

type CommandInjectVuln1ControllerHzmts struct {
	beego.Controller
}

type CommandInjectVuln2ControllerHzmts struct {
	beego.Controller
}

type CommandInjectVuln3ControllerHzmts struct {
	beego.Controller
}

type CommandInjectSafe1ControllerHzmts struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerHzmts) Get() {
	dir := c.GetString("dir")
	input := fmt.Sprintf("ls %s", dir)

d := NewFakeDuckWithAttribute(input)
input = MakeItQuackFieldAttr(d, "tmp_str")

	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerHzmts) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerHzmts) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerHzmts) Get() {
	dir := c.GetString("dir")
	commandInjectFilter := utils.CommandInjectFilter{}
	evil := commandInjectFilter.DoFilter(dir)
	if evil == false {
		c.Ctx.ResponseWriter.Write([]byte("evil input"))
		return
	}
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}