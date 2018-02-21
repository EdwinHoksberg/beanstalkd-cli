package main

import (
	"fmt"
	"github.com/edwinhoksberg/beanstalkd-cli/command"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

var (
	Name       string
	CommitHash string
)

func main() {
	command := new(command.Command)

	app := cli.NewApp()

	app.Name = Name
	app.HelpName = Name
	app.Version = "1.0.0"
	app.Usage = fmt.Sprintf(`A simple cli interface for managing beanstalkd queues.

   Homepage: https://github.com/edwinhoksberg/beanstalkd-cli
   Commit:   %s`, CommitHash)

	app.Commands = []cli.Command{
		{
			Name:   "monitor",
			Usage:  "Monitor the beanstalkd queues",
			Action: command.Monitor,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "server",
					Usage: "The server name where beanstalkd is running",
					Value: "127.0.0.1",
				},
				cli.IntFlag{
					Name:  "port",
					Usage: "The port on which beanstalkd is listening",
					Value: 11300,
				},
				cli.StringFlag{
					Name:  "tube",
					Usage: "The name of the tube to monitor",
				},
				cli.StringSliceFlag{
					Name:  "keys",
					Usage: "Different keys to monitor",
				},
			},
		},
		{
			Name:   "flush",
			Usage:  "Completely remove all jobs from a tube",
			Action: command.Flush,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "server",
					Usage: "The server name where beanstalkd is running",
					Value: "127.0.0.1",
				},
				cli.IntFlag{
					Name:  "port",
					Usage: "The port on which beanstalkd is listening",
					Value: 11300,
				},
				cli.StringFlag{
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
				cli.StringFlag{
					Name:  "server",
					Usage: "The server name where beanstalkd is running",
					Value: "127.0.0.1",
				},
				cli.IntFlag{
					Name:  "port",
					Usage: "The port on which beanstalkd is listening",
					Value: 11300,
				},
				cli.StringFlag{
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
				cli.StringFlag{
					Name:  "server",
					Usage: "The server name where beanstalkd is running",
					Value: "127.0.0.1",
				},
				cli.IntFlag{
					Name:  "port",
					Usage: "The port on which beanstalkd is listening",
					Value: 11300,
				},
				cli.StringFlag{
					Name:  "tube",
					Usage: "The name of the tube to pop an item from",
					Value: "default",
				},
				cli.StringFlag{
					Name:  "data",
					Usage: "The message to write to the queue",
				},
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "set this to enable debug logging",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Error()
	}
}
