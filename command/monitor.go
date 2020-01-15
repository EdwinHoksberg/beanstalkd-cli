package command

import (
	"bytes"
	"fmt"
	"time"

	"github.com/urfave/cli"
)

var (
	keys = []string{
		"uptime",
		"current-jobs-ready",
		"current-workers",
		"current-producers",
		"current-connections",
	}

	tubeKeys = []string{
		"total-jobs",
		"current-jobs-ready",
		"current-using",
		"current-watching",
	}
)

// Monitor outputs a overview of server and tube statistics, repeated every second.
func (c *Command) Monitor(cli *cli.Context) error {
	log := c.GetLogger(cli)

	// Build and connect to beanstalkd
	client, err := c.GetBeanstalkdClient(cli)
	if err != nil {
		log.WithError(err).Error("Could not connect to beanstalkd server")
		return err
	}

	// Assign a few interesting keys to display.
	// If no keys are passed on the cli we assign a we default keys ourself.
	if len(cli.StringSlice("keys")) > 0 {
		keys = cli.StringSlice("keys")
	}

	if len(cli.StringSlice("tubekeys")) > 0 {
		tubeKeys = cli.StringSlice("tubekeys")
	}

	// Infinite loop, show beanstalkd stats until the users exists the application.
	for {
		// Retrieve stats, break loop if an error occured.
		log.Debug("Retrieving stats...")
		stats, err := client.Stats()

		if err != nil {
			log.WithError(err).Error("Failed to retrieve server stats")
			break
		}

		// Build a buffer for our output.
		var buffer bytes.Buffer

		// Clear the screen.
		buffer.WriteString("\033[H\033[2J")

		// Write global stats.
		for i, value := range keys {
			buffer.WriteString(fmt.Sprintf("%s: %s", value, stats[value]))

			if i < (len(keys) - 1) {
				buffer.WriteString(", ")
			}
		}

		buffer.WriteByte('\n')

		// Get a slice of all existing tubes.
		tubes, err := client.ListTubes()
		if err != nil {
			log.WithError(err).Error("Error listing tubes")
			break
		}

		// Write stats for each tube.
		for _, tube := range tubes {
			log.WithField("tube", tube).Debug("Retrieving tube stats...")

			stats, err := client.StatsTube(tube)
			if err != nil {
				log.WithError(err).WithField("tube", tube).Error("Error reading tube stats")
				break
			}

			buffer.WriteString(fmt.Sprintf("\nName: %s\n", tube))

			for _, value := range tubeKeys {
				buffer.WriteString(fmt.Sprintf("%s: %s\n", value, stats[value]))
			}
		}

		// Output the buffer.
		fmt.Print(buffer.String())

		// Sleep for 1 second.
		time.Sleep(time.Second)
	}

	return nil
}
