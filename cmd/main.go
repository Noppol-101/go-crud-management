package main

import (
	"fmt"
	"my-crud-management/config"
	"my-crud-management/internal/adapter/logger"
	"my-crud-management/internal/server"
	"os"
	"path/filepath"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/sirupsen/logrus"
)

func init() {
	config.InitTimeZone()
	config.InitConfigEnvironment()
	logger.Set()
}

func main() {

	a := kingpin.New(filepath.Base(os.Args[0]), fmt.Sprintf("%s %s", "ProgramName", "Version"))
	a.HelpFlag.Short('h')

	// Start
	startCmd := a.Command("start", "start server command")

	switch kingpin.MustParse(a.Parse(os.Args[1:])) {
	case startCmd.FullCommand():
		logrus.Info("server starting at ", time.Now().Format(time.RFC3339))
		server.StartServer()
	}
}
