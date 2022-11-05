package app

import (
	"gopkg.in/ini.v1"
	"net/http"
)

var (
	Ini    *ini.File
	Server *http.Server
)
