
package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"go-sec-code/utils"
	"os/exec"
)

type CommandInjectVuln1ControllerWudtu struct {
	beego.Controller
}

type CommandInjectVuln2ControllerWudtu struct {
	beego.Controller
}

type CommandInjectVuln3ControllerWudtu struct {
	beego.Controller
}

type CommandInjectSafe1ControllerWudtu struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerWudtu) Get() {
	dir := c.GetString("dir")
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerWudtu) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerWudtu) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")

stringReturner := func() func() string {
    return func() string {
        return "vWEFo"
    }
}

stringRet := stringReturner()
repoUrl = stringRet()

	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerWudtu) Get() {
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
