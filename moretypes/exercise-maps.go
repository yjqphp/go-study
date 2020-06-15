package main
import (
	"golang.org/x/tour/wc"
	"strings"
)
/**
	实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。
	函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。
 */
func WordCount(s string) map[string]int  {
	strMap := make(map[string]int)
	for _, str := range strings.Fields(s) {
		_, ok := strMap[str]
		if ok {
			strMap[str] = strMap[str] + 1
		} else {
			strMap[str] = 1
		}
	}
	return strMap
	//return map[string]int{"x": 1}
}

func  main()  {
	wc.Test(WordCount)
}