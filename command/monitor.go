package command

import (
	"bytes"
	"fmt"
	"github.com/maxid/beanstalkd"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"time"
)

var logger *logrus.Logger

func (c *Command) Monitor(cli *cli.Context) {
	logger = c.GetLogger(cli)

	// Build a connection string.
	addr := fmt.Sprintf("%s:%d", cli.String("server"), cli.Int("port"))

	// Connect to the beanstalkd server .
	logger.Debugf("Connecting to beanstalkd server: %s", addr)
	queue, err := beanstalkd.Dial(addr)

	if err != nil {
		logger.WithError(err).Error("Could not connect to beanstalkd server")
		return
	}

	if len(cli.String("tube")) < 1 {
		monitor(cli, queue)
	} else {
		monitorTube(cli, queue)
	}
}

func monitor(cli *cli.Context, queue *beanstalkd.BeanstalkdClient) {
	// Assign a few interesting keys to display.
	// If no keys are passed on the cli we assign a we default keys ourself.
	keys := cli.StringSlice("keys")
	if len(keys) == 0 {
		keys = []string{
			"current-jobs-ready",
			"current-workers",
			"current-producers",
			"current-connections",
		}
	}

	// Infinite loop, show beanstalkd stats until the users exists the application.
	for {
		// Retrieve stats, break loop if an error occured.
		logger.Debug("Retrieving stats...")
		stats, err := queue.Stats()

		if err != nil {
			logger.WithError(err).Error()
			break
		}

		// Build a buffer for our output.
		var buffer bytes.Buffer
		for _, value := range keys {
			buffer.WriteString(fmt.Sprintf("%s: %s, ", value, stats[value]))
		}

		buffer.WriteByte('\r')

		// Output the buffer.
		fmt.Print(buffer.String())

		// Sleep for 1 second.
		time.Sleep(time.Second)
	}
}

func monitorTube(cli *cli.Context, queue *beanstalkd.BeanstalkdClient) {
	// Assign a few interesting keys to display.
	// If no keys are passed on the cli we assign a we default keys ourself.
	keys := cli.StringSlice("keys")
	if len(keys) == 0 {
		keys = []string{
			"total-jobs",
			"current-jobs-ready",
			"current-using",
			"current-waiting",
		}
	}

	// Infinite loop, show beanstalkd stats until the users exists the application.
	for {
		// Retrieve stats, break loop if an error occured.
		stats, err := queue.StatsTube(cli.String("tube"))
		if err != nil {
			logger.WithError(err).Error()
			break
		}

		// Build a buffer for our output.
		var buffer bytes.Buffer
		buffer.WriteString(fmt.Sprintf("%s: ", cli.String("tube")))

		for _, value := range keys {
			buffer.WriteString(fmt.Sprintf("%s: %s, ", value, stats[value]))
		}

		buffer.WriteByte('\r')

		// Output the buffer.
		fmt.Print(buffer.String())

		// Sleep for 1 second.
		time.Sleep(time.Second)
	}
}
