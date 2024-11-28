//Analyzer1 original results: [77]
//Analyzer2 original results: [77]
//Analyzer3 original results: []
//Analyzer4 original results: []
//Analyzer5 original results: []
//-------------
//Analyzer3 analysis results: []
//Analyzer4 analysis results: []
//Analyzer1 analysis results: [78, 703]
//Analyzer2 analysis results: []
//Analyzer5 analysis results: []
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/embedding.tmt with name embedding_anonymous_fields_negative 
//Used extensions: 
//Original file region: 49, 56, null, null
//Mutated file region: 68, 84, null, null

package controllers

import (
	"fmt"
	"go-sec-code/utils"
	"os/exec"

	beego "github.com/beego/beego/v2/server/web"
)

type CommandInjectVuln1ControllerKdkkv struct {
	beego.Controller
}

type CommandInjectVuln2ControllerKdkkv struct {
	beego.Controller
}

type CommandInjectVuln3ControllerKdkkv struct {
	beego.Controller
}

type CommandInjectSafe1ControllerKdkkv struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerKdkkv) Get() {
	dir := c.GetString("dir")
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerKdkkv) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerKdkkv) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")

a := AnonFieldHolder{
    Value2: "const_str",
    Anon: Anon{
        Value1: repoUrl,
    },
}
repoUrl = a.Value2

	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerKdkkv) Get() {
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