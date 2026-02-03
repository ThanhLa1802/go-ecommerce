package main

import (
	"go.uber.org/zap"
)

func main() {
	sugar := zap.NewExample().Sugar()
	sugar.Infof("Hello name")

	logger := zap.NewExample()
	logger.Info("Hello")
}
