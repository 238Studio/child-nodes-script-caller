package script

// Script 脚本结构体
type Script struct {
	Interpreter string //解释器(或执行器)
	Path        string //脚本路径
	Options     string //脚本选项
	Flags       string //脚本标志
}
