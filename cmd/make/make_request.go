package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file, example make request user user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeRequest(cmd *cobra.Command, args []string) {

	// 格式化模型名称，返回一个 Model 对象
	model := makeModelFromString(args[0])

	dir := "app/requests/"
	filePath := fmt.Sprintf(dir+"%s_request.go", model.PackageName)
	os.MkdirAll(dir, os.ModePerm)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "request", model)
}
