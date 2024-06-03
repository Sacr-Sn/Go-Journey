package main

// 如果上行有红线，是因为不再path的工作区。解决方法：在“当前模块根目录，如S01”终端执行：go mod init [dirname]
// D:\About_Go\Codes\Go-Journey\S01> go mod init S01
// 然后会在S01中生成一个go.mod文件，这个文件对S01文件夹下所有子目录都有效

import (
	mytest "S01/test"
)

func main() {
	mytest.Test04()
}
