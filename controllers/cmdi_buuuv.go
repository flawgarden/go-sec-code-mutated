//
//-------------
//
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/channels.tmt with name channel_struct_data_int_positive 
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

type CommandInjectVuln1ControllerBuuuv struct {
	beego.Controller
}

type CommandInjectVuln2ControllerBuuuv struct {
	beego.Controller
}

type CommandInjectVuln3ControllerBuuuv struct {
	beego.Controller
}

type CommandInjectSafe1ControllerBuuuv struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerBuuuv) Get() {
	dir := c.GetString("dir")
	

dataChannel := make(chan DataInt, 1)
dataChannel <- DataInt{Value: 1}

var wg sync.WaitGroup
wg.Add(1)

go func() {
    defer wg.Done()
    data := <-dataChannel
    data.Value += 1
    dataChannel <- data
}()

wg.Wait()

result := <-dataChannel
if (result.Value == 1) {
    dir = "constant_string"
}


	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerBuuuv) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerBuuuv) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerBuuuv) Get() {
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
