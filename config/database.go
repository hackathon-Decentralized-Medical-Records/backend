package config

import (
	"strconv"
	"strings"
)

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

func (db *Database) Dsn() string {
	var builder strings.Builder
	builder.WriteString(db.User)
	builder.WriteString(":")
	builder.WriteString(db.Password)
	builder.WriteString("@tcp(")
	builder.WriteString(db.Host)
	builder.WriteString(":")
	builder.WriteString(strconv.Itoa(db.Port))
	builder.WriteString(")/")
	builder.WriteString(db.Dbname)
	return builder.String()
}
