package command

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func (c *Command) Put(cli *cli.Context) {
	log := c.GetLogger(cli)

	if len(cli.String("data")) < 1 {
		log.Error("No data passed, use the --data parameter to insert data into a tube.")
		return
	}

	client, err := c.GetBeanstalkdClient(cli)
	if err != nil {
		log.WithError(err).Error("Could not connect to beanstalkd server")
		return
	}

	client.Use(cli.String("tube"))
	if err != nil {
		log.WithError(err).WithField("tube", cli.String("tube")).Error("Failed to select tube")
		return
	}

	id, err := client.Put(
		uint32(cli.Int("priority")),
		cli.Duration("delay"),
		cli.Duration("duration"),
		[]byte(cli.String("data")),
	)

	if err != nil {
		log.WithError(err).WithField("tube", cli.String("tube")).Error("Failed to insert job in queue")
		return
	}

	log.WithFields(logrus.Fields{
		"tube": cli.String("tube"),
		"id":   id,
	}).Info("Succesfully inserted job")

	client.Quit()
}
