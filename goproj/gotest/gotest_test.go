/**
 *
 *	文件名必须是 _test.go 结尾的， 这样再执行 go test 的时候才会执行到相应的代码
 *	你必须inport testing 这个包
 *	所有的测试用例函数必须是 Test 开头
 *	测试用例会按照源代码中写的顺序依次执行
 *	测试函数 TestXxx() 的参数是 testing.T， 我们可以使用该类型来记录错误或者是测试状态
 *	测试格式：func TestXxx(t *testing.T)，Xxx 部分可以为任意的字母数字的组合， 但是首字母不能是小写字幕
 *	函数通过调用 testing.T 的 Error, Errorf, FailNow, Fatal, FatalIf方法， 说明测试不通过
 *
 *
 * 	go test 不会显示测试通过的信息
 * 	go test -v 显示测试完整的信息
 */

package gotest

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除法函数测试没通过")
	} else {
		t.Log("第一个测试通过了")
	}
}

func Test_Division_2(t *testing.T) {
	t.Error("就是不通过")
}
