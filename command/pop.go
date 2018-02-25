package command

import (
	"fmt"
	"github.com/maxid/beanstalkd"
	"github.com/urfave/cli"
)

func (c *Command) Pop(cli *cli.Context) {
	log := c.GetLogger(cli)

	// Build a connection string.
	addr := fmt.Sprintf("%s:%d", cli.String("server"), cli.Int("port"))

	// Connect to the beanstalkd server .
	log.Debugf("Connecting to beanstalkd server: %s", addr)
	queue, err := beanstalkd.Dial(addr)

	if err != nil {
		log.WithError(err).Error("Could not connect to beanstalkd server")
		return
	}

	if cli.String("tube") != "default" {
		// Watch a specified tube.
		if _, err := queue.Watch(cli.String("tube")); err != nil {
			log.WithError(err).Error("Failed to select tube")
			return
		}

		// By default the default tube is always in the watch list.
		// To prevent flushing any jobs from the default tube we can ignore it
		// after watching a different tube.
		if _, err := queue.Ignore("default"); err != nil {
			log.WithError(err).Error("Failed to ignore default tube")
			return
		}
	}

	// Here we reserve a job to retrieve its contents
	log.Debug("Reserving a job...")
	job, err := queue.Reserve(3)
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
	if err := queue.Delete(job.Id); err != nil {
		log.WithError(err).Error()
		return
	}

	fmt.Println(string(job.Data[:]))

	queue.Quit()
}
