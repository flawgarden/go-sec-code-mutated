//Analyzer5 original results: [77]
//Analyzer1 original results: [77]
//Analyzer3 original results: []
//Analyzer2 original results: []
//Analyzer4 original results: []
//-------------
//Analyzer3 analysis results: []
//Analyzer2 analysis results: [94]
//Analyzer4 analysis results: []
//Analyzer5 analysis results: [78, 703]
//Analyzer1 analysis results: [563]
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/set.tmt with name set_add_simple_negative 
//Used extensions: MACRO_Create_Set -> ~[MACRO_SetName]~ := make(map[~[TYPE@1]~]struct{}) | MACRO_Add_Fixed_CONST_ToSet -> ~[MACRO_SetName]~[~[CONST_~[TYPE@1]~@1]~] = struct{}{} | MACRO_SetName -> set787231 | MACRO_SetName -> set787231 | MACRO_SetName -> set787231
//Program:
package controllers

import (
	"fmt"
	"go-sec-code/utils"
	"os/exec"

	beego "github.com/beego/beego/v2/server/web"
)

type CommandInjectVuln1ControllerHgymv struct {
	beego.Controller
}

type CommandInjectVuln2ControllerHgymv struct {
	beego.Controller
}

type CommandInjectVuln3ControllerHgymv struct {
	beego.Controller
}

type CommandInjectSafe1ControllerHgymv struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerHgymv) Get() {
	dir := c.GetString("dir")

set787231 := make(map[string]struct{})
set787231["xylRn"] = struct{}{}
dir = func() string {
    for k := range set787231 {
        return k
    }
    return "EYYht"
}()

	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerHgymv) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerHgymv) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerHgymv) Get() {
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
