/*
iniはconfigファイルを読み込むためのパッケージ
config.iniとか。サーバー関係でよく使うファイル。
*/
package main

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

// init()はmain()よりも前に呼び出される初期化用の関数。
func init() {
	// .iniファイルをLoadし、iniファイルで設定したセクションの値を指定して呼び出す。
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
