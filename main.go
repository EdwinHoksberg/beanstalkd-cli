package main

import (
	"fmt"
	"os"
	"time"

	"github.com/edwinhoksberg/beanstalkd-cli/command"
	"github.com/urfave/cli"
)

var (
	Name    string
	Version string
)

func main() {
	command := new(command.Command)

	app := cli.NewApp()

	app.Name = Name
	app.Usage = ""
	app.HelpName = Name
	app.Version = Version
	app.Authors = []cli.Author{
		{
			Name:  "Edwin Hoksberg",
			Email: "mail@edwinhoksberg.nl",
		},
	}
	app.UsageText = fmt.Sprintf(`A simple cli interface for managing beanstalkd queues.

   Homepage: https://github.com/edwinhoksberg/beanstalkd-cli`)

	app.Commands = []cli.Command{
		{
			Name:   "monitor",
			Usage:  "Monitor the beanstalkd queues",
			Action: command.Monitor,
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:  "keys",
					Usage: "Different keys to monitor",
				},
				&cli.StringSliceFlag{
					Name:  "tubekeys",
					Usage: "Different tube keys to monitor",
				},
			},
		},
		{
			Name:   "flush",
			Usage:  "Completely remove all jobs from a tube",
			Action: command.Flush,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "tube",
					Usage: "The name of the tube to flush",
					Value: "default",
				},
			},
		},
		{
			Name:   "pop",
			Usage:  "Remove a job from a queue and display it",
			Action: command.Pop,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "tube",
					Usage: "The name of the tube to pop an item from",
					Value: "default",
				},
			},
		},
		{
			Name:   "put",
			Usage:  "Write a job to a queue",
			Action: command.Put,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "tube",
					Usage: "The name of the tube where a new job should be inserted",
					Value: "default",
				},
				&cli.StringFlag{
					Name:  "data",
					Usage: "The message to write to the queue (default reads from stdin)",
					Value: "-",
				},
				&cli.IntFlag{
					Name:  "priority",
					Usage: "Job priority, a lower value will be scheduled before jobs with a larger priority",
					Value: 1024,
				},
				&cli.DurationFlag{
					Name:  "delay",
					Usage: "How many seconds to wait before putting the job in the queue, e.g. 300s or 2h15m",
					Value: 0,
				},
				&cli.DurationFlag{
					Name:  "ttr",
					Usage: "The number of seconds to allow a worker to run this job, e.g. 300s or 2h15m",
					Value: time.Hour,
				},
			},
		},
		{
			Name:   "peek",
			Usage:  "Display a job from the queue without removing it",
			Action: command.Peek,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "tube",
					Usage: "The name of the tube to peek from",
					Value: "default",
				},
			},
		},
	}

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "verbose",
			Usage: "set this to enable debug logging",
		},
		&cli.BoolFlag{
			Name:  "quiet",
			Usage: "set this to disable logging",
		},
		&cli.StringFlag{
			Name:  "server",
			Usage: "The server name where beanstalkd is running",
			Value: "127.0.0.1",
		},
		&cli.IntFlag{
			Name:  "port",
			Usage: "The port on which beanstalkd is listening",
			Value: 11300,
		},
	}

	app.Run(os.Args)
}
