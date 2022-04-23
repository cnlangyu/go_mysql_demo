package generate_content

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
)


// GetCon4Py 从python脚本获取语句
// locale 语言 如 zh_CN
// nb_word 字数（模糊值）
func GetCon4Py(locale string, nbWord int) (string, error) {
	pyShell := fmt.Sprintf("import faker; print(faker.Faker(locale='%s').sentence(nb_words=%d))", locale, nbWord)
	cmd := exec.Command("python3", "-c", pyShell)
	output, err := cmd.CombinedOutput()
	if err !=nil{
		panic(err)
		return "", err
	}
	return string(output), nil
}

// GetCon4Http 从网络中获取内容
// 默认连接[彩虹屁生成器](https://chp.shadiao.app/api.php)
//Deprecated
func GetCon4Http(url string) (string, error) {
	if len(url) < 1{
		url = "https://chp.shadiao.app/api.php"
	}
	response, err := http.Get(url)
	if err != nil{
		panic(err)
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
		return "", err
	}
	content := string(body)
	fmt.Println(content)
	return content, nil
}
