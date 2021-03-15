package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/psinthorn/go_smallsite/pkg/configs"
)

var (
	appConfig configs.AppConfig
	session   *scs.SessionManager
)
