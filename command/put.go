package command

import (
	"fmt"
	"github.com/maxid/beanstalkd"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func (c *Command) Put(cli *cli.Context) {
	log := c.GetLogger(cli)

	if len(cli.String("data")) < 1 {
		log.Error("No data passed, use the --data parameter to insert data into a tube.")
		return
	}

	// Build a connection string.
	addr := fmt.Sprintf("%s:%d", cli.String("server"), cli.Int("port"))

	// Connect to the beanstalkd server .
	log.Debugf("Connecting to beanstalkd server: %s", addr)
	queue, err := beanstalkd.Dial(addr)

	if err != nil {
		log.WithError(err).Error("Could not connect to beanstalkd server")
		return
	}

	queue.Use(cli.String("tube"))
	if err != nil {
		log.WithError(err).Error("Failed to select tube")
		return
	}

	id, err := queue.Put(0, 0, 10000, []byte(cli.String("data")))
	if err != nil {
		log.WithError(err).Error()
		return
	}

	log.WithFields(logrus.Fields{
		"tube": cli.String("tube"),
		"id":   id,
	}).Info("Succesfully inserted job")

	queue.Quit()
}
