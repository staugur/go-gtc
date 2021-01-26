package ufc

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const version = "0.1.0"

// PathExist 检测路径是否存在
func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// PathNotExist 检测路径是否不存在
func PathNotExist(path string) bool {
	return !PathExist(path)
}

// IsDir 是否为目录
func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// IsFile 是否为文件（含普通、设备）
func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

// IsCommonFile 是否为普通文件
func IsCommonFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.Mode().IsRegular()
}

// CreateDir 创建文件夹（无递归），如果已存在则直接返回，否则按照0755权限创建
func CreateDir(path string) error {
	if PathNotExist(path) {
		return os.Mkdir(path, 0755)
	}
	return nil
}

// FileReadByte 读取指定的文件，并返回[]byte
func FileReadByte(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	return ioutil.ReadAll(fi)
}

// FileReadStr 读取指定的文件，并返回字符串
func FileReadStr(path string) (string, error) {
	raw, err := FileReadByte(path)
	return string(raw), err
}

// FileCopy 复制文件，会自动创建并覆盖目标文件
func FileCopy(dstName, srcName string) (written int64, err error) {
	if !IsFile(srcName) {
		return 0, errors.New("src file does not exist")
	}

	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// FileCopyN 按字节复制文件，会自动创建并覆盖目标文件
func FileCopyN(dstName, srcName string, n int64) (written int64, err error) {
	if !IsFile(srcName) {
		return 0, errors.New("src file does not exist")
	}

	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.CopyN(dst, src, n)
}

// IsTrue 仅当值为 1、t、、T、true、True、TRUE、on 时返回布尔值true，其他（错误）返回false
func IsTrue(v string) bool {
	if strings.ToLower(v) == "on" {
		return true
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return b
}

// NotTrue 非IsTrue则是false
// 如0、f、false、False、FALSE、其他字符串或发生错误时都返回布尔值true
func NotTrue(v string) bool {
	return !IsTrue(v)
}

// IsFalse 仅当值为 0、f、F、false、False、FALSE、off 时返回布尔值true，其他（错误）返回false
func IsFalse(v string) bool {
	if strings.ToLower(v) == "off" {
		return true
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	if b == true {
		return false
	}
	// b is false
	return true
}

// StrInSlice 判断字符串是否在切片中
func StrInSlice(val string, slice []string) bool {
	for _, b := range slice {
		if b == val {
			return true
		}
	}
	return false
}

// InArraySlice 判断值是否在数组或切片中，嵌套类型采用reflect.DeepEqual深度对比结果
func InArraySlice(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	kind := reflect.TypeOf(array).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}
