

package controllers

import (
"fmt"
"go-sec-code/utils"
"os/exec"
beego "github.com/beego/beego/v2/server/web"
"sync"
)

type CommandInjectVuln1ControllerQsddx struct {
	beego.Controller
}

type CommandInjectVuln2ControllerQsddx struct {
	beego.Controller
}

type CommandInjectVuln3ControllerQsddx struct {
	beego.Controller
}

type CommandInjectSafe1ControllerQsddx struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerQsddx) Get() {
	dir := c.GetString("dir")
	

w := &Wrapper[string]{Value: dir}
task1 := NewSwitchingTask(w, dir)
task2 := NewSwitchingTask(w, dir)
var wg sync.WaitGroup
wg.Add(2)
go func() {
    defer wg.Done()
    task2.Run()
}()
go func() {
    defer wg.Done()
    task1.Run()
}()
wg.Wait()
dir = w.Value


	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerQsddx) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerQsddx) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerQsddx) Get() {
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
