package globals

import (
	"fmt"
	"github.com/mix-go/console"
	"github.com/mix-go/logrus"
	logrus2 "github.com/sirupsen/logrus"
	"io"
	"os"
)

func Init() {
	// logger
	logger := Logger()
	file := logrus.NewFileWriter(fmt.Sprintf("%s/../logs/hydra-wework-auth-server.log", console.App.BasePath), 7, 0)
	writer := io.MultiWriter(os.Stdout, file)
	logger.SetOutput(writer)
	if console.App.Debug {
		logger.SetLevel(logrus2.DebugLevel)
	}

}
