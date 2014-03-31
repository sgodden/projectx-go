package util

import (
	"os"
)

type Stage int

const (
	STAGE_DEVELOPMENT Stage = iota
	STAGE_PRODUCTION Stage = iota
)

var (
	stage Stage = STAGE_DEVELOPMENT
)

func SetStage(s Stage) {
	stage = s
}

func GetStage() Stage {
	stageEnv := os.Getenv("application.stage")
	ret := STAGE_PRODUCTION
	if len(stageEnv) > 0 && stageEnv == "dev" {
		ret = STAGE_DEVELOPMENT
	}
	return ret
}
