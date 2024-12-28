
package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"go-sec-code/utils"
	"os/exec"
)

type CommandInjectVuln1ControllerFyyku struct {
	beego.Controller
}

type CommandInjectVuln2ControllerFyyku struct {
	beego.Controller
}

type CommandInjectVuln3ControllerFyyku struct {
	beego.Controller
}

type CommandInjectSafe1ControllerFyyku struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerFyyku) Get() {
	dir := c.GetString("dir")

counter := func() func(str string) string {
    count := 0
    return func(str string) string {
        count++
        if count == 1 {
            return str
        } else {
            return "fixed_string"
        }

    }
}()
counter(dir)
dir = counter(dir)

	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerFyyku) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerFyyku) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerFyyku) Get() {
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

