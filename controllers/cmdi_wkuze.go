//
//-------------
//
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_remove_if_positive 
//Used extensions: 
//Program:
package controllers

import (
"fmt"
"go-sec-code/utils"
"os/exec"
beego "github.com/beego/beego/v2/server/web"
"container/list"
)

type CommandInjectVuln1ControllerWkuze struct {
	beego.Controller
}

type CommandInjectVuln2ControllerWkuze struct {
	beego.Controller
}

type CommandInjectVuln3ControllerWkuze struct {
	beego.Controller
}

type CommandInjectSafe1ControllerWkuze struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerWkuze) Get() {
	dir := c.GetString("dir")
	

queue787231 := list.New()
queue787231.PushBack("DyZve")
queue787231.PushBack(dir)
for e := queue787231.Front(); e != nil; {
    next := e.Next()
    if e.Value == "DyZve" {
        queue787231.Remove(e)
    }
    e = next
}
dir = queue787231.Front().Value.(string)


	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerWkuze) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerWkuze) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerWkuze) Get() {
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
