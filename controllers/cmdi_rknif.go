//Analyzer1 original results: [77]
//Analyzer2 original results: [77]
//Analyzer3 original results: []
//Analyzer4 original results: []
//Analyzer5 original results: []
//-------------
//Analyzer3 analysis results: []
//Analyzer4 analysis results: []
//Analyzer1 analysis results: [78, 703]
//Analyzer2 analysis results: [563]
//Analyzer5 analysis results: []
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/map.tmt with name map_simple_get_1_negative 
//Used extensions: 
//Original file region: 49, 56, null, null
//Mutated file region: 68, 80, null, null

package controllers

import (
	"fmt"
	"go-sec-code/utils"
	"os/exec"

	beego "github.com/beego/beego/v2/server/web"
)

type CommandInjectVuln1ControllerRknif struct {
	beego.Controller
}

type CommandInjectVuln2ControllerRknif struct {
	beego.Controller
}

type CommandInjectVuln3ControllerRknif struct {
	beego.Controller
}

type CommandInjectSafe1ControllerRknif struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerRknif) Get() {
	dir := c.GetString("dir")
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerRknif) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerRknif) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")

map787234 := make(map[string]string)
map787234["bYSJc"] = "vTLgm"
repoUrl = map787234["bYSJc"]

	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerRknif) Get() {
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