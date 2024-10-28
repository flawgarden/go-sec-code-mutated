//
//-------------
//
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/atomic.tmt with name atomic_integer_restore_positive 
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

type CommandInjectVuln1ControllerUlevk struct {
	beego.Controller
}

type CommandInjectVuln2ControllerUlevk struct {
	beego.Controller
}

type CommandInjectVuln3ControllerUlevk struct {
	beego.Controller
}

type CommandInjectSafe1ControllerUlevk struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerUlevk) Get() {
	dir := c.GetString("dir")
	

counter := int32(0)
stringCopy := dir

if atomic.LoadInt32(&counter) == 0 {
    dir = ""
}

var wg sync.WaitGroup
wg.Add(2)

go func() {
    defer wg.Done()
    atomic.AddInt32(&counter, 1)
}()

go func() {
    defer wg.Done()
    atomic.AddInt32(&counter, 1)
}()

wg.Wait()

if atomic.LoadInt32(&counter) == 2 {
    dir = stringCopy
}



	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerUlevk) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerUlevk) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerUlevk) Get() {
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
