
package controllers

import (
"fmt"
"go-sec-code/utils"
"os/exec"
beego "github.com/beego/beego/v2/server/web"
"sync"
)

type CommandInjectVuln1ControllerAixjc struct {
	beego.Controller
}

type CommandInjectVuln2ControllerAixjc struct {
	beego.Controller
}

type CommandInjectVuln3ControllerAixjc struct {
	beego.Controller
}

type CommandInjectSafe1ControllerAixjc struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerAixjc) Get() {
	dir := c.GetString("dir")

dataChannel := make(chan Data, 1)
dataChannel <- Data{Value: dir}

var wg sync.WaitGroup
wg.Add(1)

go func() {
    data := <-dataChannel
    data.Value = "constant_string"
    dataChannel <- data
}()

wg.Wait()

readData := <-dataChannel
dir = readData.Value

	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerAixjc) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerAixjc) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerAixjc) Get() {
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
