package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os/exec"
	"sync/atomic"
)

type CommandInjectVuln1Controllermpfrr struct {
	beego.Controller
}

func (c *CommandInjectVuln1Controllermpfrr) Get() {
	dir := c.GetString("dir")
	var ar42 atomic.Value
	ar42.Store(dir)

	go func() {
		ar42.Store(dir)
	}()
	dir = ar42.Load().(string)
	input := fmt.Sprintf("ls %s", dir)
	cmd := exec.Command("bash", "-c", input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(out)
}
