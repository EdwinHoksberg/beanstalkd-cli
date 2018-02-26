package command

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func (c *Command) Pop(cli *cli.Context) {
	log := c.GetLogger(cli)

	client, err := c.GetBeanstalkdClient(cli)
	if err != nil {
		log.WithError(err).Error("Could not connect to beanstalkd server")
		return
	}

	if cli.String("tube") != "default" {
		// Watch a specified tube.
		if _, err := client.Watch(cli.String("tube")); err != nil {
			log.WithError(err).WithField("tube", cli.String("tube")).Error("Failed to select tube")
			return
		}

		// By default the default tube is always in the watch list.
		// To prevent flushing any jobs from the default tube we can ignore it
		// after watching a different tube.
		if _, err := client.Ignore("default"); err != nil {
			log.WithError(err).Error("Failed to ignore default tube")
			return
		}
	}

	// Here we reserve a job to retrieve its contents
	log.Debug("Reserving a job...")
	job, err := client.Reserve(3)
	if err != nil {
		if err.Error() == "timed out" {
			log.Info("No jobs found in tube, exiting...")
			return
		}

		log.WithError(err).Error()
		return
	}

	// After reserving a job, we can delete it
	log.WithField("id", job.Id).Debug("Deleting the reserved job")
	if err := client.Delete(job.Id); err != nil {
		log.WithError(err).WithFields(logrus.Fields{
			"tube": cli.String("tube"),
			"job":  job.Id,
		}).Error("Failed to delete job")
		return
	}

	fmt.Println(string(job.Data[:]))

	client.Quit()
}
