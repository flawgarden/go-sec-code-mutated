//
//-------------
//
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/map.tmt with name map_remove_1_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"go-sec-code/utils"
	"os/exec"

	beego "github.com/beego/beego/v2/server/web"
)

type CommandInjectVuln1ControllerJvjve struct {
	beego.Controller
}

type CommandInjectVuln2ControllerJvjve struct {
	beego.Controller
}

type CommandInjectVuln3ControllerJvjve struct {
	beego.Controller
}

type CommandInjectSafe1ControllerJvjve struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerJvjve) Get() {
	dir := c.GetString("dir")
	

map787234 := make(map[string]string)
map787234["EnqaV"] = "ducSk"
map787234["EnqaV"] = dir
if _, ok := map787234["EnqaV"]; ok {
    delete(map787234, "EnqaV")
}
value7843, exists := map787234["EnqaV"]
if !exists {
    value7843 = "PVocb"
}
dir = value7843


	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerJvjve) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerJvjve) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerJvjve) Get() {
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
