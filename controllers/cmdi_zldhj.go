
package controllers

import (
	"fmt"
	"go-sec-code/utils"
	"os/exec"

	beego "github.com/beego/beego/v2/server/web"
)

type CommandInjectVuln1ControllerZldhj struct {
	beego.Controller
}

type CommandInjectVuln2ControllerZldhj struct {
	beego.Controller
}

type CommandInjectVuln3ControllerZldhj struct {
	beego.Controller
}

type CommandInjectSafe1ControllerZldhj struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerZldhj) Get() {
	dir := c.GetString("dir")
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerZldhj) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerZldhj) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")

arr4124 := []string{repoUrl}
nested7231 := NewNestedFields3FromArray(arr4124)
repoUrl = nested7231.nested1.nested1.nested1.values[0]

	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerZldhj) Get() {
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
