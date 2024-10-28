//Analyzer5 original results: [77]
//Analyzer1 original results: [77]
//Analyzer3 original results: []
//Analyzer2 original results: []
//Analyzer4 original results: []
//-------------
//Analyzer3 analysis results: []
//Analyzer2 analysis results: []
//Analyzer4 analysis results: []
//Analyzer5 analysis results: [78, 703]
//Analyzer1 analysis results: [563]
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/varargs/varargs.tmt with name varargs_with_generics_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"go-sec-code/utils"
	"os/exec"
)

type CommandInjectVuln1ControllerJqyfl struct {
	beego.Controller
}

type CommandInjectVuln2ControllerJqyfl struct {
	beego.Controller
}

type CommandInjectVuln3ControllerJqyfl struct {
	beego.Controller
}

type CommandInjectSafe1ControllerJqyfl struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerJqyfl) Get() {
	dir := c.GetString("dir")
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerJqyfl) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerJqyfl) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")

repoUrl = varargsWithGenerics("FejAj", "beGvw")

	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerJqyfl) Get() {
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