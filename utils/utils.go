package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/howeyc/gopass"
)

func GetPassword() string {
	// Get user's password
	fmt.Println("Please enter your password: ")
	pwd, err := gopass.GetPasswd()
	if err != nil {
		panic(err)
	}
	return string(pwd)
}

// 创建目录
func NewDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

// 判断文件夹是否为空
func IsDirEmpty(path string) bool {
	NewDir(path)
	dir, _ := ioutil.ReadDir(path)
	return len(dir) == 0
}

// 复制文件
func CopyFile(src, dst string) {
	srcData, _ := os.Open(src)
	defer srcData.Close()

	dstData, _ := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0644)
	defer dstData.Close()

	io.Copy(dstData, srcData)
}

type ChainConfig struct {
	ChainId string `json:"chain_id"`
}

// 读取 chain_config.json 模版
func GetChainConfigTemp(path string) *ChainConfig {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var chainConfig ChainConfig

	// 创建 json 解码器
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&chainConfig)
	if err != nil {
		log.Fatal(err)
	}

	return &chainConfig
}
