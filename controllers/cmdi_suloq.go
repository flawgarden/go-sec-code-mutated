//Analyzer1 original results: [77]
//Analyzer2 original results: [77]
//Analyzer3 original results: []
//Analyzer4 original results: []
//Analyzer5 original results: []
//-------------
//Analyzer3 analysis results: [94]
//Analyzer4 analysis results: []
//Analyzer1 analysis results: [78, 703]
//Analyzer2 analysis results: [78]
//Analyzer5 analysis results: []
//Original file name: controllers/cmdi.go
//Original file CWE's: [77]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/lambdas/mutation.tmt with name nested_unary_lambda_mutation_positive 
//Used extensions: 
//Original file region: 27, 36, null, null
//Mutated file region: 46, 70, null, null


package controllers

import (
	"fmt"
	"go-sec-code/utils"
	"os/exec"

	beego "github.com/beego/beego/v2/server/web"
)

type CommandInjectVuln1ControllerSuloq struct {
	beego.Controller
}

type CommandInjectVuln2ControllerSuloq struct {
	beego.Controller
}

type CommandInjectVuln3ControllerSuloq struct {
	beego.Controller
}

type CommandInjectSafe1ControllerSuloq struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerSuloq) Get() {
	dir := c.GetString("dir")
	

s23423 := dir
a12341 := &StringHolder{}
lambda1231 := func(s *StringHolder) {
    innerLambda1231 := func(p *StringHolder) {
        p.value = ""
    }
    innerLambda1231(s)
    s.value = s23423
}
lambda1231(a12341)
dir = a12341.value


	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln2ControllerSuloq) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerSuloq) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerSuloq) Get() {
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