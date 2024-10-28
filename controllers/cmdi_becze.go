//
//-------------
//
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/concurrency.tmt with name two_set_threads_in_sequence_negative 
//Used extensions: 
//Program:
package controllers

import (
"fmt"
"go-sec-code/utils"
"os/exec"
beego "github.com/beego/beego/v2/server/web"
"sync"
)

type CommandInjectVuln1ControllerBecze struct {
	beego.Controller
}

type CommandInjectVuln2ControllerBecze struct {
	beego.Controller
}

type CommandInjectVuln3ControllerBecze struct {
	beego.Controller
}

type CommandInjectSafe1ControllerBecze struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerBecze) Get() {
	dir := c.GetString("dir")
	

w := &Wrapper[string]{Value: dir}
task1 := NewSettingTask(w, "")
task2 := NewSettingTask(w, dir)
var wg sync.WaitGroup
wg.Add(2)
done := make(chan struct{})
go func() {
    defer wg.Done()
    task2.Run()
    close(done)
}()
go func() {
    defer wg.Done()
    <-done
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

func (c *CommandInjectVuln2ControllerBecze) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerBecze) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerBecze) Get() {
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
