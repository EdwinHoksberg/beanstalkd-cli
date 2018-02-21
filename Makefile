BUILDPATH=$(CURDIR)
GO=$(shell which go)

# Binary name
PROGRAM=beanstalkd-cli

# Compile time values
COMMIT_HASH=`git rev-parse --verify HEAD`

# Interpolate the variable values using go link flags
LDFLAGS=-ldflags "-X main.CommitHash=${COMMIT_HASH} -X main.Name=${PROGRAM}"

build:
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	$(GO) build ${LDFLAGS} -o $(BUILDPATH)/bin/$(PROGRAM)

release:
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi

	GOOS=linux GOARCH=amd64 $(GO) build ${LDFLAGS} -o $(BUILDPATH)/bin/$(PROGRAM)_linux_amd64
	GOOS=linux GOARCH=386 $(GO) build ${LDFLAGS} -o $(BUILDPATH)/bin/$(PROGRAM)_linux_386
	GOOS=windows GOARCH=amd64 $(GO) build ${LDFLAGS} -o $(BUILDPATH)/bin/$(PROGRAM)_windows_amd64
	GOOS=windows GOARCH=386 $(GO) build ${LDFLAGS} -o $(BUILDPATH)/bin/$(PROGRAM)_windows_386

clean:
	rm -f $(BUILDPATH)/bin/$(PROGRAM)

all: build
