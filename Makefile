BUILDPATH=$(CURDIR)
GO=$(shell which go)

# Binary name
PROGRAM=beanstalkd-cli

# Compile time values
COMMIT_HASH=`git rev-parse --verify HEAD`
VERSION=`git describe --abbrev=0 --tags`

# Interpolate the variable values using go link flags
LDFLAGS=-ldflags "-X main.CommitHash=${COMMIT_HASH} -X main.Name=${PROGRAM} -X main.Version=${VERSION}"

build:
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	$(GO) build ${LDFLAGS} -o $(BUILDPATH)/bin/$(PROGRAM)

clean:
	rm -f $(BUILDPATH)/bin/$(PROGRAM)

all: build
