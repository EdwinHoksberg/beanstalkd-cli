package command

import (
	"fmt"

	"github.com/mpdroog/beanstalkd"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/urfave/cli"
)

// Command holds data for the cli commands.
type Command struct{}

// GetLogger returns a logrus object according to any flags set
// in the application parameters.
func (c Command) GetLogger(cli *cli.Context) *log.Logger {
	if cli.Bool("quiet") {
		logger, _ := test.NewNullLogger()
		return logger
	}

	// Set the default output formatter
	format := &log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
	}

	log.SetFormatter(format)

	// If the verbose flag was enabled, enable debug logging
	if cli.Bool("verbose") {
		log.SetLevel(log.DebugLevel)
	}

	return log.StandardLogger()
}

// GetBeanstalkdClient returns a client object for interacting with a beanstalkd server.
func (c Command) GetBeanstalkdClient(cli *cli.Context) (*beanstalkd.BeanstalkdClient, error) {
	// Build a connection string.
	addr := fmt.Sprintf("%s:%d", cli.GlobalString("server"), cli.GlobalInt("port"))

	// Connect to the beanstalkd server.
	log.Debugf("Connecting to beanstalkd server: %s", addr)
	client, err := beanstalkd.Dial(addr)

	if err != nil {
		return nil, err
	}

	log.Debug("Successfully connected")

	return client, nil
}
