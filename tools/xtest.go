package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func xtest() {
	// 要修改的文件路径
	filePath := "config.yaml"

	// 读取文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	// 将文件内容转换为字符串
	fileContent := string(content)

	// 要查找的字段
	targetField := "pwdHashKey: "

	// 查找目标字段的索引
	index := strings.Index(fileContent, targetField)
	if index == -1 {
		fmt.Println("未找到目标字段:", targetField)
		return
	}
	_, salt, _ := RandSalt(8)
	// 在目标字段后面添加内容
	newContent := fileContent[:index+len(targetField)] + salt

	// 将修改后的内容写回文件
	err = ioutil.WriteFile(filePath, []byte(newContent), os.ModePerm)
	if err != nil {
		fmt.Println("写回文件失败:", err)
		return
	}

	fmt.Println("文件内容已成功修改。")

}
