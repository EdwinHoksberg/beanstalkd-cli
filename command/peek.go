package command

import (
	"fmt"

	"github.com/urfave/cli"
)

// Peek displays the next job in the tube without removing it.
func (c *Command) Peek(cli *cli.Context) error {
	log := c.GetLogger(cli)

	// Build and connect to beanstalkd
	client, err := c.GetBeanstalkdClient(cli)
	if err != nil {
		log.WithError(err).Error("Could not connect to beanstalkd server")
		return err
	}

	// Select a tube to peek from
	log.Debug("Selecting correct tube...")
	client.Use(cli.String("tube"))
	if err != nil {
		log.WithError(err).WithField("tube", cli.String("tube")).Error("Failed to select tube")
		return err
	}

	// Here we peek a job from the watched tube
	log.Debug("Peeking job...")
	job, err := client.PeekReady()
	if err != nil {
		log.WithError(err).WithField("tube", cli.String("tube")).Error("Failed to peek job")
		return err
	}

	fmt.Println(string(job.Data[:]))

	client.Quit()

	return nil
}
