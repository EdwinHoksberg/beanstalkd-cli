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
		_, err := queue.Reserve(3)
		if err != nil {
			if err.Error() == "timed out" {
				log.Info("No jobs found in tube, exiting...")
				break
			}

			log.WithError(err).Error()
			break
		}

		counter++
	}

	log.Debugf("%d jobs removed", counter)
}
