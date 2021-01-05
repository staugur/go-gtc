package ufc

import (
	"io"
	"io/ioutil"
	"os"
)

// PathExist 检测路径是否存在
func PathExist(path string) bool {
	if _, err := os.Stat(path); os.IsExist(err) {
		return true
	}
	return false
}

// PathNotExist 检测路径是否不存在
func PathNotExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
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

// IsDir 是否为目录
func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()

}

// CreateDir 创建文件夹，如果已存在则直接返回，否则按照0755权限创建
func CreateDir(path string) error {
	if PathNotExist(path) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// FileReadByte 读取指定的文件，并返回[]byte
func FileReadByte(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		//		panic(err)
		return nil, err
	} else {
		defer fi.Close()
		return ioutil.ReadAll(fi)
	}
}

// FileReadStr 读取指定的文件，并返回字符串
func FileReadStr(path string) (string, error) {
	raw, err := FileReadByte(path)
	return string(raw), err
}

// FileCopy 复制文件
func FileCopy(dstName, srcName string) (written int64, err error) {
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

// FileCopyN 按字节复制文件
func FileCopyN(dstName, srcName string, n int64) (written int64, err error) {
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
