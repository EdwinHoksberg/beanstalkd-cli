package command

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type Command struct{}

func (c Command) GetLogger(cli *cli.Context) *log.Logger {
	format := &log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
	}

	log.SetFormatter(format)

	if cli.GlobalBool("verbose") {
		log.SetLevel(log.DebugLevel)
	}

	return log.StandardLogger()
}
