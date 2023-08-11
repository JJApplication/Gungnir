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
)

type ConfigModel struct {
	Root       string            `json:"root"`
	Host       string            `json:"host"`
	Port       int               `json:"port"`
	DenyDirs   []string          `json:"deny_dirs"`   // 禁止访问的目录
	DenyFiles  []string          `json:"deny_files"`  // 禁止访问的文件
	DenyIPs    []string          `json:"deny_ips"`    // 禁止访问的IP
	DenyAgents []string          `json:"deny_agents"` // 禁止访问的user-agent
	Headers    map[string]string `json:"headers"`     // 自定义headers
}

func loadConfig() ConfigModel {
	data, err := os.ReadFile(Config)
	if err != nil {
		log.Printf("load config [%s] error: %s\n", Config, err.Error())
		return ConfigModel{}
	}
	var cf ConfigModel
	if err = json.Unmarshal(data, &cf); err != nil {
		log.Printf("load config [%s] error: %s\n", Config, err.Error())
	}

	return cf
}

func newConfig() {
	data, _ := json.MarshalIndent(ConfigModel{
		Root:       DefaultRoot,
		Host:       DefaultHost,
		Port:       DefaultPort,
		DenyDirs:   nil,
		DenyFiles:  nil,
		DenyIPs:    nil,
		DenyAgents: nil,
		Headers:    nil,
	}, "", "  ")

	if err := os.WriteFile(Config, data, 0644); err != nil {
		log.Printf("generate config file error: %s\n", err.Error())
		return
	}
	log.Printf("generate config file to %s\n", Config)
}

func (c *ConfigModel) Init() {
	Root = c.mustSetString(c.Root, Root)
	Host = c.mustSetString(c.Host, Host)
	Port = c.mustSetInt(c.Port, Port)
	DenyDirs = c.DenyDirs
	DenyFiles = c.DenyFiles
	DenyIPs = c.DenyIPs
	DenyAgents = c.DenyAgents
	Headers = c.Headers
}

func (c *ConfigModel) mustSetString(cf string, def string) string {
	if cf == "" {
		return def
	}

	return cf
}

func (c *ConfigModel) mustSetInt(cf int, def int) int {
	if cf == 0 {
		return def
	}

	return cf
}
