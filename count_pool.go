/*
   Create: 2023/8/11
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

var countPool sync.Map

type poolFile struct {
	Path  string `json:"path"`
	Count int    `json:"count"`
}

func InitPool() {
	countPool = sync.Map{}
	loadPool()
	go startSync()
}

// 从本地文件系统中加载池
func loadPool() {
	_, err := os.Stat(Pool)
	if err != nil {
		log.Printf("load counts from local error: %s\n", err.Error())
		log.Println("create counts pool")
		createPool()
	}
	var p map[string]poolFile
	data, _ := os.ReadFile(Pool)
	_ = json.Unmarshal(data, &p)
	for key, val := range p {
		countPool.Store(key, val)
	}
}

// 更新或加入pool
func updatePool(key string) {
	data, ok := countPool.Load(key)
	if ok {
		p := data.(poolFile)
		countPool.Store(key, poolFile{
			p.Path,
			p.Count + 1,
		})
	} else {
		countPool.Store(key, poolFile{
			key,
			1,
		})
	}
}

func getPool(key string) poolFile {
	data, ok := countPool.Load(key)
	if ok {
		return data.(poolFile)
	}
	return poolFile{}
}

func createPool() {
	os.WriteFile(Pool, []byte{}, 0644)
}

// 存储到本地文件系统
func syncPool() {
	m := make(map[string]poolFile)
	countPool.Range(func(key, value any) bool {
		if key.(string) != "" {
			m[key.(string)] = value.(poolFile)
			return true
		}
		return false
	})
	data, _ := json.MarshalIndent(m, "", "  ")
	_ = os.WriteFile(Pool, data, 0644)
}

func startSync() {
	t := time.Tick(time.Duration(SyncCount) * time.Second)
	for range t {
		log.Printf("start sync pool -> %s\n", Pool)
		syncPool()
	}
}
