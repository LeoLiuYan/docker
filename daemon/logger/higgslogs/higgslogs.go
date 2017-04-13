package higgslogs

import (
	"github.com/docker/docker/daemon/logger"
	"github.com/Sirupsen/logrus"
)

const (
	name = "higgs"
)

func init() {
	if err := logger.RegisterLogDriver(name, New); err != nil {
		logrus.Fatal(err)
	}
}

func New(info logger.Info) (logger.Logger, error){
	return nil, nil
}
