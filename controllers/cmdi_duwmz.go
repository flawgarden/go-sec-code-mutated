//
//-------------
//
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/atomic.tmt with name atomic_integer_negative 
//Used extensions: 
//Program:
package controllers

import (
"fmt"
"go-sec-code/utils"
"os/exec"
beego "github.com/beego/beego/v2/server/web"
"sync"
"sync/atomic"
)

type CommandInjectVuln1ControllerDuwmz struct {
	beego.Controller
}

type CommandInjectVuln2ControllerDuwmz struct {
	beego.Controller
}

type CommandInjectVuln3ControllerDuwmz struct {
	beego.Controller
}

type CommandInjectSafe1ControllerDuwmz struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerDuwmz) Get() {
	dir := c.GetString("dir")
	

counter := int32(0)

var wg sync.WaitGroup

wg.Add(2)
go func() {
    defer wg.Done()
    atomic.AddInt32(&counter, 1)

}()

go func() {
    defer     wg.Done()
    atomic.AddInt32(&counter, 1)
}()

wg.Wait()

if atomic.LoadInt32(&counter) == 2 {
    dir = ""
}



	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerDuwmz) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerDuwmz) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerDuwmz) Get() {
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
