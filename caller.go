package script

import (
	"os/exec"

	"github.com/238Studio/child-nodes-error-manager/errpack"
)

// RegisterScriptCaller 注册脚本调用器
func RegisterScriptCaller(interpreter, options, flags, path string) (script Script) {
	return Script{
		Interpreter: interpreter,
		Options:     options,
		Flags:       flags,
		Path:        path,
	}
}

func (s *Script) Caller(args ...string) (response []byte, err error) {
	if s.Interpreter == "" && s.Path != "" {
		//检验options和flags是否为空
		if s.Options != "" {
			args = append([]string{s.Options}, args...)
		}
		if s.Flags != "" {
			args = append(args, s.Flags)
		}
		return Caller(s.Path, args...)
	} else {
		//检验options和flags是否为空
		if s.Options != "" {
			args = append([]string{s.Options}, args...)
		}
		if s.Flags != "" {
			args = append(args, s.Flags)
		}
		return Caller(s.Interpreter, args...)
	}
}

// Caller 脚本调用器
// 传入: 解释器or执行器,op,flag,脚本路径, 参数
// 返回: 脚本执行返回值,错误信息
func Caller(interpreter string, args ...string) (response []byte, err error) {
	//创建命令对象
	cmd := exec.Command(interpreter)

	//设置参数
	cmd.Args = append(cmd.Args, args...)

	//执行命令并捕获输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errpack.NewError(errpack.TrivialException, errpack.Script, err)
	}

	return output, nil
}
