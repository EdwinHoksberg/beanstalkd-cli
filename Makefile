BUILDPATH=$(CURDIR)
GO=$(shell which go)

# Binary name
PROGRAM=beanstalkd-cli

# Compile time values
COMMIT_HASH=446f49e78f79a164079a99a88446182703331c75

# Interpolate the variable values using go link flags
LDFLAGS=-ldflags "-X main.CommitHash=${COMMIT_HASH} -X main.Name=${PROGRAM}"

build:
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	$(GO) build ${LDFLAGS} -o $(BUILDPATH)/bin/$(PROGRAM)

clean:
	rm -f $(BUILDPATH)/bin/$(PROGRAM)

all: build
