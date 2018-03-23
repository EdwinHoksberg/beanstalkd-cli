BUILDPATH=$(CURDIR)
MANPATH="/usr/local/share/man/man1"
GO=$(shell which go)

# Compile time values
PROGRAM=beanstalkd-cli
VERSION=$(shell printf "%s [%s]" `git describe --abbrev=0 --tags` `git rev-parse --verify HEAD`)

# Interpolate the variable values using go link flags
LDFLAGS=-ldflags "-X 'main.Name=${PROGRAM}' -X 'main.Version=${VERSION}'"

build:
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	$(GO) build ${LDFLAGS} -o $(BUILDPATH)/bin/$(PROGRAM)

install:
	@if [ ! -d $(MANPATH) ] ; then mkdir -p $(MANPATH) ; fi

	cp $(BUILDPATH)/bin/$(PROGRAM) /usr/bin/$(PROGRAM)
	gzip -c $(PROGRAM).man | tee $(MANPATH)/$(PROGRAM).1.gz > /dev/null

clean:
	rm -f $(BUILDPATH)/bin/$(PROGRAM)

all: build
