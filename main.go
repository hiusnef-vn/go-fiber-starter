package main

import (
	"github.com/hiusnef-vn/go-fiber-starter/utils/logger"
	"go.uber.org/zap"
)

func main() {
	for v := range 10 {
		logger.GetLogger().Info("Hello", zap.Int("seq", v))
	}
}
