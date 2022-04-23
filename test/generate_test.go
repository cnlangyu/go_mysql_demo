package test

import (
	"fmt"
	"github.com/cnlangyu/go_mysql_demo/generate_content"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestPyGenerate(t *testing.T) {
	for i := 0; i < 10; i++ {
		content, err := generate_content.GetCon4Py("zh_CN", 2)
		if err != nil{
			panic(err)
			return
		}
		index := strings.Index(content, "\n")
		if index != -1 && index != 0{
			content = content[:index-1]
		}
		fmt.Println("len = ",len(content), ", string len = ", utf8.RuneCountInString(content), ", con = ", content)
	}
}

func TestNetGenerate(t *testing.T) {
	for i := 0; i< 50; i++{
		content, err := generate_content.GetCon4Http("")
		if err != nil{
			fmt.Println("获取彩虹屁错误:", err)
			return
		}
		fmt.Println("len = ", len(content), ", con = ", content)
	}
}
