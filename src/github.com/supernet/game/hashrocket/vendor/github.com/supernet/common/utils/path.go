package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v2"
)

// GetCurrentDir 获取当前应用程序所在目录
func GetCurrentDir() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}

//CreateFileIfNotExist check file exist or create it if not exist
func CreateFileIfNotExist(filePath string) error {
	_, err := os.Stat(filePath) //os.Stat获取文件信息
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		CreateIfNotExist(filepath.Dir(filePath))
		fd, err := os.Create(filePath)
		defer func() {
			if fd != nil {
				fd.Close()
			}
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateIfNotExist check exist or create it if not exist
func CreateIfNotExist(path string) bool {
	success := true
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				success = false
			}
		} else {
			success = false
		}
	}
	return success
}

//FileExist 判断文件是否存在
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// LoadYamlFile 加载解析yaml配置文件(多个文件采用拼接)，res需传入已初始化空间的指针
func LoadYamlFile(res interface{}, paths ...string) error {
	if res == nil || reflect.TypeOf(res).Kind() != reflect.Ptr {
		return errors.New("invalid result receiver")
	}

	var datas []byte
	for i := 0; i < len(paths); i++ {
		data, err := ioutil.ReadFile(paths[i])
		if err != nil {
			return err
		}
		data = append(data, '\n') //防止拼接时前一个文件末尾无换行 或后一个文件开头无换行
		datas = append(datas, data...)
	}
	return yaml.Unmarshal(datas, res)
}
