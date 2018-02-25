# beanstalkd-cli
A simple cli interface for managing beanstalkd queues.

## Download
You can find the latest release including binaries [here](https://github.com/EdwinHoksberg/beanstalkd-cli/releases/latest).

## Usage
```
NAME:
   beanstalkd-cli - A simple cli interface for managing beanstalkd queues.

   Homepage: https://github.com/edwinhoksberg/beanstalkd-cli

USAGE:
   beanstalkd-cli [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     monitor  Monitor the beanstalkd queues
     flush    Completely remove all jobs from a tube
     pop      Remove a job from a queue and display it
     put      Write a job to a queue
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbose      set this to enable debug logging
   --help, -h     show help
   --version, -v  print the version
```

## Development
This program is written in [Go](https://golang.org/), using these dependencies:
- [sirupsen/logrus](https://github.com/sirupsen/logrus) - Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger.
- [urfave/cli](https://github.com/urfave/cli) - cli is a simple, fast, and fun package for building command line apps in Go.
- [maxid/beanstalkd](https://github.com/maxid/beanstalkd) - A golang beanstalkd client.

## License
[MIT](LICENSE.md)
