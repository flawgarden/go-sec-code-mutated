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
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/channels.tmt with name channel_with_cycle_negative 
//Used extensions: 
//Original file region: 27, 36, null, null
//Mutated file region: 46, 81, null, null


package controllers

import (
"fmt"
"go-sec-code/utils"
"os/exec"
beego "github.com/beego/beego/v2/server/web"
"sync"
)

type CommandInjectVuln1ControllerNkyfr struct {
	beego.Controller
}

type CommandInjectVuln2ControllerNkyfr struct {
	beego.Controller
}

type CommandInjectVuln3ControllerNkyfr struct {
	beego.Controller
}

type CommandInjectSafe1ControllerNkyfr struct {
	beego.Controller
}

func (c *CommandInjectVuln1ControllerNkyfr) Get() {
	dir := c.GetString("dir")
	

dataChannel := make(chan int, 5)
dataChannel <- 3
var wg sync.WaitGroup
wg.Add(1)

go func() {
    defer wg.Done()
    ind123 := <- dataChannel
    for i:=0; i < ind123; i++ {
        dataChannel <- i
    }
    close(dataChannel)
}()

wg.Wait()
readData := 0
for d := range dataChannel {
    readData = d
}
if readData != 2 {
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

func (c *CommandInjectVuln2ControllerNkyfr) Get() {
	host := c.Ctx.Request.Host
	input := fmt.Sprintf("curl %s", host)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectVuln3ControllerNkyfr) Get() {
	repoUrl := c.GetString("repoUrl", "--upload-pack=${touch /tmp/pwnned}")
	out, err := exec.Command("git", "ls-remote", repoUrl, "refs/heads/main").CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}

func (c *CommandInjectSafe1ControllerNkyfr) Get() {
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