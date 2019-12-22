package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
)

const configFile = "server-tool.ini"

func normalizePathNoCheck(path string) string {
	cmd := exec.Command("sh", "-c", "realpath "+path)
	output, _ := cmd.Output()
	return strings.Trim(string(output), "\r\n")
}

func getOpt(options map[string]string, key, def string) string {
	if val, ok := options[key]; ok {
		return val
	}
	return def
}

func main() {
	var execType = filepath.Base(os.Args[0])
	if len(os.Args) > 1 {
		execType = os.Args[1]
	}

	log.Printf("starting: %v", execType)

	var configPath = filepath.Dir(os.Args[0]) + "/" + configFile
	if _, err := os.Stat(configPath); err != nil {
		configPath = normalizePathNoCheck("~/" + configFile)
	}

	var err error
	var cfg *ini.File
	if cfg, err = ini.Load(configPath); err != nil {
		log.Fatalf("config file not loaded: %v", err)
	}

	switch execType {
	case "ssh-server":
		sshServer(cfg.Section(execType).KeysHash())
	case "ftp-server":
		ftpServer(cfg.Section(execType).KeysHash())
	default:
		log.Fatalf("unknown target: %v", execType)
	}
}
