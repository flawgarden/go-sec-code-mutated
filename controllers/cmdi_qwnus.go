//-------------
//
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/exceptions/panic.tmt with name arithmetic_panic_positive
//Used extensions:
//Program:
package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"go-sec-code/utils"
	"os/exec"
)

type CommandInjectVuln1ControllerQwnus struct {
	beego.Controller
}

type CommandInjectVuln2ControllerQwnus struct {
	beego.Controller
}

type CommandInjectVuln3ControllerQwnus struct {
	beego.Controller
}

type CommandInjectSafe1ControllerQwnus struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerQwnus) Get() {
	dir := c.GetString("dir")

	var err error
	func() {
		defer handlePanic(&err) //Отложенный вызов для обработки паники
		divide_simple(0)
	}()
	if err != nil {
		dir = dir + "suffix"
	} else {
		dir = "777"
	}

	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerQwnus) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerQwnus) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerQwnus) Get() {
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
