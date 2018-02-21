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

//current-jobs-urgent: 0
//current-jobs-ready: 0
//current-jobs-reserved: 0
//current-jobs-delayed: 0
//current-jobs-buried: 0
//cmd-put: 0
//cmd-peek: 0
//cmd-peek-ready: 0
//cmd-peek-delayed: 0
//cmd-peek-buried: 0
//cmd-reserve: 3
//cmd-reserve-with-timeout: 21
//cmd-delete: 0
//cmd-release: 0
//cmd-use: 0
//cmd-watch: 0
//cmd-ignore: 0
//cmd-bury: 0
//cmd-kick: 0
//cmd-touch: 0
//cmd-stats: 248
//cmd-stats-job: 0
//cmd-stats-tube: 14
//cmd-list-tubes: 0
//cmd-list-tube-used: 0
//cmd-list-tubes-watched: 0
//cmd-pause-tube: 0
//job-timeouts: 0
//total-jobs: 0
//max-job-size: 65535
//current-tubes: 1
//current-connections: 1
//current-producers: 0
//current-workers: 0
//current-waiting: 0
//total-connections: 54
//pid: 15157
//version: 1.10
//rusage-utime: 0.009677
//rusage-stime: 0.032257
//uptime: 3799
//binlog-oldest-index: 0
//binlog-current-index: 0
//binlog-records-migrated: 0
//binlog-records-written: 0
//binlog-max-size: 10485760
//id: 708a57be1d24db1d
//hostname: edwin-desktop

//name: default
//current-jobs-urgent: 0
//current-jobs-ready: 0
//current-jobs-reserved: 0
//current-jobs-delayed: 0
//current-jobs-buried: 0
//total-jobs: 0
//current-using: 1
//current-watching: 1
//current-waiting: 0
//cmd-delete: 0
//cmd-pause-tube: 0
//pause: 0
//pause-time-left: 0
