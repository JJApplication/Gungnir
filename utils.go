/*
   Create: 2023/8/9
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

func getFilePath(root, dir, f string) string {
	return filepath.Join(root, dir, f)
}

// 获取文件大小
func getFileInfo(f string) (string, string) {
	stat, err := os.Stat(f)
	if os.IsNotExist(err) {
		return "", ""
	}

	return calcFileSize(stat.Size()), getModTime(stat.ModTime())
}

// 计算文件大小
func calcFileSize(size int64) string {
	if size/TB > 0 {
		return fmt.Sprintf("%sTB", strconv.Itoa(int(size/TB)))
	} else if size/GB > 0 {
		return fmt.Sprintf("%sGB", strconv.Itoa(int(size/GB)))
	} else if size/MB > 0 {
		return fmt.Sprintf("%sMB", strconv.Itoa(int(size/MB)))
	} else if size/KB > 0 {
		return fmt.Sprintf("%sKB", strconv.Itoa(int(size/KB)))
	}

	return fmt.Sprintf("%dB", size)
}

func getExt(name string) string {
	ext := filepath.Ext(name)
	if ext == "" {
		return "void"
	}
	return ext
}

// 获取修改时间
func getModTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func calcMD5(f []byte) string {
	h := md5.New()
	h.Write(f)
	return string(h.Sum(nil))
}
