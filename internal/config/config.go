package config

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"github.com/viking311/books/pkg/database"
)

type ServerConf struct {
	Host string
	Port int
}

func (sc *ServerConf) GetAddr() string {
	return fmt.Sprintf("%s:%d", sc.Host, sc.Port)
}

type Config struct {
	Database database.PostgresConfig
	Server   ServerConf
}

var Cfg Config

func init() {

	configPath := flag.String("c", "configs/main.yml", "path to config file")
	flag.Parse()

	configFile := filenameWithoutExtension(filepath.Base(*configPath))
	configDir := filepath.Dir(*configPath)

	viper.AddConfigPath(configDir)
	viper.SetConfigName(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	Cfg.Database.Host = viper.GetString("database.host")
	Cfg.Database.Port = viper.GetInt("database.port")
	Cfg.Database.DBName = viper.GetString("database.dbname")
	Cfg.Database.SSLMode = viper.GetString("database.sslmode")
	Cfg.Database.Username = viper.GetString("database.username")
	Cfg.Database.Password = viper.GetString("database.password")

	Cfg.Server.Host = viper.GetString("server.host")
	Cfg.Server.Port = viper.GetInt("server.port")

	if err := envconfig.Process("db", &Cfg.Database); err != nil {
		panic(err)
	}

	if err := envconfig.Process("server", &Cfg.Server); err != nil {
		panic(err)
	}

}

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, filepath.Ext(fn))
}
