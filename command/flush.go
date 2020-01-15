package command

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Flush empties an entire tube of jobs.
func (c *Command) Flush(cli *cli.Context) error {
	log := c.GetLogger(cli)

	// Build and connect to beanstalkd
	client, err := c.GetBeanstalkdClient(cli)
	if err != nil {
		log.WithError(err).Error("Could not connect to beanstalkd server")
		return err
	}

	log.Debug("Selecting correct tube...")

	if cli.String("tube") != "default" {
		// Watch a specified tube.
		if _, err := client.Watch(cli.String("tube")); err != nil {
			log.WithError(err).WithField("tube", cli.String("tube")).Error("Failed to select tube")
			return err
		}

		// By default the default tube is always in the watch list.
		// To prevent flushing any jobs from the default tube we can ignore it
		// after watching a different tube.
		if _, err := client.Ignore("default"); err != nil {
			log.WithError(err).Error("Failed to ignore default tube")
			return err
		}
	}

	log.Debug("Starting flush loop...")

	counter := 0
	for {
		log.Debug("[Loop] Reserving job...")

		job, err := client.Reserve(1)
		if err != nil {
			if err.Error() == "timed out" {
				// If the job deleted counter is still at zero, it means the tube
				// was empty before we even started flushing jobs, so we don't show the error.
				if counter == 0 {
					log.WithError(err).WithField("tube", cli.String("tube")).Info("No jobs found in tube, exiting...")
				}

				break
			}

			// The DEADLINE_SOON message just indicates that a job was about to expire.
			// This has no effect on our actions here so we just ignore it.
			if err.Error() != "deadline soon" {
				log.WithError(err).WithField("tube", cli.String("tube")).Error("Failed reserving a job")
				break
			}
		}

		log.Debug("[Loop] Deleting job...")

		// Delete the job from the tube
		if err := client.Delete(job.Id); err != nil {
			log.WithError(err).WithFields(logrus.Fields{
				"tube": cli.String("tube"),
				"job":  job.Id,
			}).Error("Failed to delete job")
			break
		}

		log.Debug("[Loop] Done")

		counter++
	}

	log.Infof("%d jobs removed", counter)

	client.Quit()

	return nil
}
