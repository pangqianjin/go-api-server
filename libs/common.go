package libs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

// get the dirrectory path of programName
func getExecPath() (path string) {
	file, _ := exec.LookPath(os.Args[0]) // path of programName
	path, _ = filepath.Abs(file)         // convert to absolute path
	path = filepath.Dir(path)            // the path's directory
	return
}

// read file and bind to JSON
func ReadFileToModel(fileName string, model interface{}) (err error) {
	path := fmt.Sprintf("%s/%s", getExecPath(), fileName) // concat the config file path
	file, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("文件：%s不存在，错误：%v+", path, err)
	}

	contentByte, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("读取文件：%s错误：%+v", path, err)
	}

	err = json.Unmarshal(contentByte, &model)
	if err != nil {
		err = fmt.Errorf("文件：%s 的json unmarshal失败: %+v", fileName, err)
	}
	return
}

// read xxx.json or xxx.post.json
func ReadJSON(fileName string) (result string, err error) {
	// path := fmt.Sprintf("%s/%s", getExecPath(), fileName) // concat the config file path
	path := fileName
	file, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("文件：%s不存在，错误：%v+", path, err)
	}

	contentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("读取文件：%s错误：%+v", path, err)
	}

	str := string(contentBytes)
	commentReg := regexp.MustCompile(`(\/\*([\s\S]*?)\*\/|([^:]|^)\/\/(.*)$)`)
	result = commentReg.ReplaceAllString(str, "") // replace commemts in json file

	return
}
