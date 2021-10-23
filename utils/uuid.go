package utils

import (
	"github.com/twharmon/gouid"
)

func GetUUID() string {
	return string(gouid.MixedCaseAlpha)
}