package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/psinthorn/go_smallsite/internal/configs"
)

var (
	appConfig configs.AppConfig
	session   *scs.SessionManager
)
