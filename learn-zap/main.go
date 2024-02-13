package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("Failed to initialize zap logger")
	}

	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("Unable to read file",
		"path", "/var/log/awesome.log",
	)

	sugar.Infof("Unable to read file: %s", "/var/log/awesome.log")

	logger.Info("Unable to read file",
		zap.String("path", "/var/log/awesome.log"),
	)
}
