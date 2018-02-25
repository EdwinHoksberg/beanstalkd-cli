package command

import (
	"fmt"
	"github.com/maxid/beanstalkd"
	"github.com/urfave/cli"
)

func (c *Command) Flush(cli *cli.Context) {
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

	counter := 0
	for {
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

		job, err := queue.Reserve(1)
		if err != nil {
			if err.Error() == "timed out" {
				// If the job deleted counter is still at zero, it means the tube
				// was empty before we even started flushing jobs, so we don't show the error.
				if counter == 0 {
					log.WithError(err).Info("No jobs found in tube, exiting...")
				}

				break
			}

			// The DEADLINE_SOON message just indicates that a job was about to expire.
			// This has no effect on our actions here so we just ignore it.
			if err.Error() != "deadline soon" {
				log.WithError(err).Error("Failed reserving a job")
				break
			}
		}

		if err := queue.Delete(job.Id); err != nil {
			log.WithError(err).WithField("id", job.Id).Error("Failed to delete job")
			break
		}

		counter++
	}

	log.Infof("%d jobs removed", counter)

	queue.Quit()
}
