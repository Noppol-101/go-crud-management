package config

import (
	"fmt"
	"time"
)

func InitTimeZone() {

	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(fmt.Errorf("fatal error init time zone: %w", err))
	}
	time.Local = ict
}
