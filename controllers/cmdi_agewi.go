

package controllers

import (
"fmt"
"go-sec-code/utils"
"os/exec"
beego "github.com/beego/beego/v2/server/web"
"sync"
"sync/atomic"
)

type CommandInjectVuln1ControllerAgewi struct {
	beego.Controller
}

type CommandInjectVuln2ControllerAgewi struct {
	beego.Controller
}

type CommandInjectVuln3ControllerAgewi struct {
	beego.Controller
}

type CommandInjectSafe1ControllerAgewi struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerAgewi) Get() {
	dir := c.GetString("dir")
	

counter := int32(0)
stringCopy := dir

if atomic.LoadInt32(&counter) == 0 {
    dir = ""
}

var wg sync.WaitGroup
wg.Add(2)
done := make(chan struct{})

go func() {
    defer wg.Done()
    atomic.AddInt32(&counter, 1)
}()

go func() {
    defer wg.Done()
    atomic.AddInt32(&counter, 1)
}()

if atomic.LoadInt32(&counter) == 2 {
    dir = stringCopy
}
close(done)
<-done
wg.Wait()



	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerAgewi) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerAgewi) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerAgewi) Get() {
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
