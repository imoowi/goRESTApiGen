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
	"time"

	"github.com/imoowi/goRESTApiGen/template"
	"github.com/imoowi/goRESTApiGen/util"
	"github.com/spf13/cobra"
)

type Generator struct {
	ModuleName    string
	AppPath       string
	AppName       string
	ServiceName   string
	ModelName     string
	ModelCollName string
	Err           error
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
	// for k, line := range strings.Split(string(data), "\n") {
	// 	fmt.Println(`line=`, line)
	// 	fmt.Println(k)
	// 	if k {

	// 	}
	// 	// oneRow := strings.Split(line, " ")
	// 	// for _, module := range oneRow {
	// 	// 	if module == `module` {
	// 	// 		g.ModuleName = oneRow[1]

	// 	// 		fmt.Println(`oneRow=`, oneRow)
	// 	// 		break
	// 	// 	}
	// 	// }
	// }
	lines := strings.Split(string(data), "\n")
	g.ModuleName = strings.Replace(lines[0], "module ", "", -1)

	fmt.Println(`module=` + g.ModuleName)
	// return false
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
	g.ModelCollName = g.ModelName
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
	templateModel.ModelCollName = strings.ToLower(g.ModelCollName)
	templateModel.ModelInstanceName = strings.ToLower(g.ModelName) + `Model`
	modelFile := modelpath + `/` + g.ModelName + `.model.go`
	fmt.Println(`modelFile path = `, modelFile)
	_, g.Err = os.Stat(modelFile)
	if os.IsNotExist(g.Err) {
		// fmt.Println(modelFile, ` 文件不存在`)
	} else {
		//重命名model文件
		// fmt.Println(modelFile, ` 文件重命名`)

		// newModelFile := modelFile + `.bak.` + cast.ToString(time.Now().Unix()) + `.go`
		newModelFile := fmt.Sprintf(`%s.bak.%d.go`, modelFile, time.Now().Unix())
		err = os.Rename(modelFile, newModelFile)
		if err != nil {
			fmt.Println(`rename model file err, `, err.Error())
			return
		}
	}
	file, err := os.OpenFile(modelFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	preModel := templateModel.PreModel()
	// fmt.Println(`preModel = `, preModel)
	write.WriteString(preModel)
	write.WriteString(templateModel.PreList())
	write.WriteString(templateModel.PreAdd())
	write.WriteString(templateModel.PreUpdate())
	write.WriteString(templateModel.PreDelete())
	write.WriteString(templateModel.PreGetOne())
	write.Flush()
	fmt.Println(`file[models/` + templateModel.ModelName + `.model.go] generated!`)
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
	fmt.Println(`modelFile path = `, serviceFile)
	_, g.Err = os.Stat(serviceFile)
	if os.IsNotExist(g.Err) {
		// fmt.Println(modelFile, ` 文件不存在`)
	} else {
		//重命名model文件
		// fmt.Println(serviceFile, ` 文件重命名`)

		// newModelFile := modelFile + `.bak.` + cast.ToString(time.Now().Unix()) + `.go`
		newServiceFile := fmt.Sprintf(`%s.bak.%d.go`, serviceFile, time.Now().Unix())
		err = os.Rename(serviceFile, newServiceFile)
		if err != nil {
			fmt.Println(`rename model file err, `, err.Error())
			return
		}
	}
	file, err := os.OpenFile(serviceFile, os.O_WRONLY|os.O_CREATE, 0755)
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
	fmt.Println(`file[services/` + templateService.ServiceName + `.service.go] generated!`)

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

	fmt.Println(`modelFile path = `, appnameFile)
	_, g.Err = os.Stat(appnameFile)
	if os.IsNotExist(g.Err) {
		// fmt.Println(modelFile, ` 文件不存在`)
	} else {
		//重命名model文件
		// fmt.Println(appnameFile, ` 文件重命名`)

		// newModelFile := modelFile + `.bak.` + cast.ToString(time.Now().Unix()) + `.go`
		newappnameFile := fmt.Sprintf(`%s.bak.%d.go`, appnameFile, time.Now().Unix())
		err = os.Rename(appnameFile, newappnameFile)
		if err != nil {
			fmt.Println(`rename model file err, `, err.Error())
			return
		}
	}

	file, err := os.OpenFile(appnameFile, os.O_WRONLY|os.O_CREATE, 0755)
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

	fmt.Println(`file[app/` + templateApp.AppName + `/` + templateApp.AppName + `.handler.go] generated!`)

	routerFile := `./app/` + g.AppPath + `/router.go`

	fmt.Println(`modelFile path = `, routerFile)
	_, g.Err = os.Stat(routerFile)
	if os.IsNotExist(g.Err) {
		// fmt.Println(modelFile, ` 文件不存在`)
	} else {
		//重命名model文件
		// fmt.Println(routerFile, ` 文件重命名`)

		// newModelFile := modelFile + `.bak.` + cast.ToString(time.Now().Unix()) + `.go`
		newrouterFile := fmt.Sprintf(`%s.bak.%d.go`, routerFile, time.Now().Unix())
		err = os.Rename(routerFile, newrouterFile)
		if err != nil {
			fmt.Println(`rename model file err, `, err.Error())
			return
		}
	}

	file, err = os.OpenFile(routerFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write = bufio.NewWriter(file)
	write.WriteString(templateApp.PreRouter())
	write.Flush()
	fmt.Println(`file[app/` + templateApp.AppName + `/router.go] generated!`)
}
