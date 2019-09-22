package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	FilePath  string
	FileTypes map[string]string

	UserName string
	PassWord string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	JwtSecret    string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	loadBase()
	loadServer()
	loadApp()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadApp() {
	FilePath = Cfg.Section("file").Key("PATH").MustString("src/")
	JwtSecret = Cfg.Section("").Key("JWT_SECRET").MustString("xoxoxoxo")
	FileTypes = map[string]string{
		"key":  "ppt",
		"ppt":  "ppt",
		"pptx": "ppt",
		"pdf":  "article",
		"doc":  "article",
		"docx": "article",
	}

	UserName = Cfg.Section("admin").Key("UserName").MustString("admin")
	PassWord = Cfg.Section("admin").Key("PassWord").MustString("admin")
}
