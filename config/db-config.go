package config

import (
	"rocky-springs-86767/x/mongodb"
)

type dbConfig struct {
	Host     string
	Name     string
	Uname    string
	Password string
}

var DBConfig dbConfig

func (d dbConfig) initDB() {
	mongodb.Connect(d.Host, d.Name, d.Uname, d.Password)
}
