// 应用程序入口：必须是main包，main方法。路径名无所谓。
package main    // 包

// 引入代码依赖
import (
	"fmt" 
	"os"
)

func main () {
	if len(os.Args) > 1{
		fmt.Println("Hello World", os.Args[1])
	}
	os.Exit(-1)
}
// 二、退出返回值 os.Exit
// 三、通过os.Args获取命令行参数
