//Analyzer1 original results: [77]
//Analyzer2 original results: [77]
//Analyzer3 original results: []
//Analyzer4 original results: []
//Analyzer5 original results: []
//-------------
//Analyzer3 analysis results: [94]
//Analyzer4 analysis results: []
//Analyzer1 analysis results: [78, 703]
//Analyzer2 analysis results: []
//Analyzer5 analysis results: []
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/default.tmt with name binary_op_interface_default_negative 
//Used extensions: 
//Original file region: 27, 36, null, null
//Mutated file region: 46, 59, null, null

package controllers

import (
	"fmt"
	"go-sec-code/utils"
	"os/exec"

	beego "github.com/beego/beego/v2/server/web"
)

type CommandInjectVuln1ControllerMnarw struct {
	beego.Controller
}

type CommandInjectVuln2ControllerMnarw struct {
	beego.Controller
}

type CommandInjectVuln3ControllerMnarw struct {
	beego.Controller
}

type CommandInjectSafe1ControllerMnarw struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerMnarw) Get() {
	dir := c.GetString("dir")
	input := fmt.Sprintf("ls %s", dir)

var a12341 BinaryOpInterfaceDefault = &BinaryOpInterfaceDefaultImplementation{}
input = a12341.InterfaceCall(input, input)

	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerMnarw) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerMnarw) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerMnarw) Get() {
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