package make

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var CmdMakeController = &cobra.Command{
	Use:   "controller",
	Short: "Create controller, exmaple: make controller user",
	Run:   runMakeController,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeController(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])

	// 组建目标目录
	dir := "app/controllers/"
	filePath := fmt.Sprintf(dir+"%s_controller.go", args[0])

	dir = path.Dir(filePath)

	os.MkdirAll(dir, os.ModePerm)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "controller", model)
}
