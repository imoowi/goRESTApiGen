package app

import (
	"bufio"
	"fmt"
	"os"

	"github.com/imoowi/goRESTApiGen/template"
	"github.com/spf13/cobra"
)

type Generator struct {
}

func (g *Generator) Init() {
	//在当前目录下创建配置文件
	modelFile := `.gen.conf`
	file, err := os.OpenFile(modelFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(`root = "."`)
	write.WriteString("\n")
	write.WriteString(`module = "go-monitor"`)
	write.Flush()
	fmt.Println(`.gen.conf created,pls alter it`)
}
func (g *Generator) Gen(cmd *cobra.Command, args []string) {
	appname, err := cmd.Flags().GetString(`appname`)
	if err != nil {
		fmt.Println(`pls input appname`)
		return
	}
	if appname == `` {
		fmt.Println(`pls input appname`)
		return
	}
	fmt.Println(`appname=`, appname)

	apppath, err := cmd.Flags().GetString(`path`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if apppath == `` {
		apppath = appname
	}
	fmt.Println(`path=`, apppath)

	service, err := cmd.Flags().GetString(`service`)
	if err != nil {
		fmt.Println(`service input err`)
		return
	}
	if service == `` {
		service = appname
	}
	fmt.Println(`service=`, service)
	model, err := cmd.Flags().GetString(`model`)
	if err != nil {
		fmt.Println(`model input err`)
		return
	}
	if model == `` {
		model = appname
	}
	fmt.Println(`model=`, model)

	//创建model
	modelpath := `./models`
	err = os.MkdirAll(modelpath, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(`modelpath makedir success.`)
	var templateModel template.TemplateModel
	modelFile := modelpath + `/` + model + `.model.go`
	file, err := os.OpenFile(modelFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(templateModel.PreModel(model))
	write.Flush()

	//创建service
	servicepath := `./services/` + service
	err = os.MkdirAll(servicepath, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(`servicepath makedir success.`)

	var templateService template.TemplateService
	serviceFile := servicepath + `/` + service + `.service.go`
	file, err = os.OpenFile(serviceFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write = bufio.NewWriter(file)
	write.WriteString(templateService.PreModel(service))
	write.WriteString(templateService.FuncSearch(service, model))
	write.Flush()

	//创建app
	err = os.MkdirAll(`./app/`+apppath, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(`apppath makedir success.`)

	var templateApp template.TemplateApp
	appnameFile := `./app/` + apppath + `/` + appname + `.handler.go`
	file, err = os.OpenFile(appnameFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write = bufio.NewWriter(file)
	write.WriteString(templateApp.PreModel(``, appname))
	write.WriteString(templateApp.FuncSearch(appname, service))
	write.Flush()

	routerFile := `./app/` + apppath + `/router.go`
	file, err = os.OpenFile(routerFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write = bufio.NewWriter(file)
	write.WriteString(templateApp.PreRouter(``, appname))
	write.Flush()
}
