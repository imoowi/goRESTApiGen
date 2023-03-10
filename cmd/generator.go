/*
Copyright © 2023 yuanjun <imoowi@qq.com>

*/
package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/imoowi/goRESTApiGen/template"
	"github.com/imoowi/goRESTApiGen/util"
	"github.com/spf13/cobra"
)

type Generator struct {
	ModuleName  string
	AppPath     string
	AppName     string
	ServiceName string
	ModelName   string
	Err         error
}

func (g *Generator) Init(cmd *cobra.Command, args []string) bool {
	moduleFile := `go.mod`
	_, g.Err = os.Stat(moduleFile)
	if os.IsNotExist(g.Err) {
		fmt.Println(`项目根目录下没有 go.mod 文件`)
		return false
	}
	data, err := ioutil.ReadFile(moduleFile)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	for _, line := range strings.Split(string(data), "\n") {
		// fmt.Println(line)
		row := strings.Split(line, " ")
		for k, module := range row {
			if module == `module` {
				g.ModuleName = row[1]
				break
			}
			fmt.Println(k, `->`, module)
		}
	}

	if g.ModuleName == `` {
		fmt.Println(`没有找到go.mod里的module配置`)
		return false
	}

	g.AppName, err = cmd.Flags().GetString(`appname`)
	if err != nil {
		fmt.Println(`pls input appname`)
		return false
	}
	if g.AppName == `` {
		fmt.Println(`pls input appname`)
		return false
	}
	fmt.Println(`appname=`, g.AppName)

	g.AppPath, err = cmd.Flags().GetString(`path`)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if g.AppPath == `` {
		g.AppPath = g.AppName
	}
	fmt.Println(`path=`, g.AppPath)

	g.ServiceName, err = cmd.Flags().GetString(`service`)
	if err != nil {
		fmt.Println(`service input  err`)
		return false
	}
	if g.ServiceName == `` {
		g.ServiceName = g.AppName
	}
	fmt.Println(`service=`, g.ServiceName)

	g.ModelName, err = cmd.Flags().GetString(`model`)
	if err != nil {
		fmt.Println(`model input  err`)
		return false
	}
	if g.ModelName == `` {
		g.ModelName = g.AppName
	}
	fmt.Println(`model=`, g.ModelName)
	return true
}
func (g *Generator) Gen(cmd *cobra.Command, args []string) {
	if !g.Init(cmd, args) {
		return
	}

	g.GenModel()

	g.GenService()

	g.GenApp()
}

//创建model
func (g *Generator) GenModel() {
	modelpath := `./models`
	err := os.MkdirAll(modelpath, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(`modelpath makedir success.`)
	var templateModel template.TemplateModel
	templateModel.ModuleName = g.ModuleName
	templateModel.ServiceName = util.FirstUpper(g.ServiceName) + `Service`
	templateModel.ModelName = util.FirstUpper(g.ModelName) + `Model`
	templateModel.ModelInstanceName = strings.ToLower(g.ModelName) + `Model`
	modelFile := modelpath + `/` + g.ModelName + `.model.go`
	file, err := os.OpenFile(modelFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(templateModel.PreModel())
	write.WriteString(templateModel.PreList())
	write.WriteString(templateModel.PreAdd())
	write.WriteString(templateModel.PreUpdate())
	write.WriteString(templateModel.PreDelete())
	write.WriteString(templateModel.PreGetOne())
	write.Flush()
}

//创建service
func (g *Generator) GenService() {
	servicepath := `./services`
	err := os.MkdirAll(servicepath, os.ModePerm)
	if err != nil {
		fmt.Println(g.Err.Error())
		return
	}
	fmt.Println(`servicepath makedir success.`)

	var templateService template.TemplateService
	templateService.ModuleName = g.ModuleName
	templateService.ServiceName = util.FirstUpper(g.ServiceName) + `Service`
	templateService.ModelName = util.FirstUpper(g.ModelName) + `Model`
	templateService.ModelInstanceName = strings.ToLower(g.ModelName) + `Model`
	serviceFile := servicepath + `/` + g.ServiceName + `.service.go`
	file, err := os.OpenFile(serviceFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(templateService.PreModel())
	write.WriteString(templateService.PreList())
	write.WriteString(templateService.PreAdd())
	write.WriteString(templateService.PreUpdate())
	write.WriteString(templateService.PreDelete())
	write.WriteString(templateService.PreGetOne())
	write.Flush()

}

//创建app
func (g *Generator) GenApp() {
	err := os.MkdirAll(`./app/`+g.AppPath, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(`apppath makedir success.`)

	var templateApp template.TemplateApp

	templateApp.ModuleName = g.ModuleName
	templateApp.AppName = g.AppName
	templateApp.ServiceName = util.FirstUpper(g.ServiceName) + `Service`
	templateApp.ModelName = util.FirstUpper(g.ModelName) + `Model`
	templateApp.ModelInstanceName = strings.ToLower(g.ModelName) + `Model`
	appnameFile := `./app/` + g.AppPath + `/` + g.AppName + `.handler.go`
	file, err := os.OpenFile(appnameFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(templateApp.PreModel())
	write.WriteString(templateApp.PreList())
	write.WriteString(templateApp.PreAdd())
	write.WriteString(templateApp.PreUpdate())
	write.WriteString(templateApp.PreDelete())
	write.WriteString(templateApp.PreGetOne())
	write.Flush()

	routerFile := `./app/` + g.AppPath + `/router.go`
	file, err = os.OpenFile(routerFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write = bufio.NewWriter(file)
	write.WriteString(templateApp.PreRouter())
	write.Flush()
}
