package common

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"
)

type ReverseConfig struct {
	Routes []struct {
		Path string `json:"path"`
		Url  string `json:"url"`
	} `json:"routes"`
}

var ReverseConfigHandle *ReverseConfig

// 读取路由配置文件
func init() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("读取应用目录失败:%s \n", err))
	}

	// 读取配置文件
	configData, err := os.ReadFile(workDir + "/reverse.json")
	if err != nil {
		fmt.Println("Failed to read config file:", err)
		return
	}
	// 解析配置文件
	err = jsoniter.Unmarshal(configData, &ReverseConfigHandle)
	if err != nil {
		fmt.Println("Failed to parse config file:", err)
		return
	}
}
