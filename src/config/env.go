package config

import (
	"os"
	"strings"
)

type envType int

//go:generate stringer -type=envType
const (
	DEV envType = iota
	STAGE
	PRO
)

var env = (func() envType {
	e := os.Getenv("ENV")
	if e == "" {
		e = Settings().Environment.Env
	}
	if res, ok := map[string]envType{
		DEV.String():   DEV,
		STAGE.String(): STAGE,
		PRO.String():   PRO,
	}[strings.ToUpper(e)]; ok {
		return res
	}
	return DEV
})()

// Env is current App's Env
func Env() envType {
	return env
}
