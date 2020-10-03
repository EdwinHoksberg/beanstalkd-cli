#!/usr/bin/env

if [[ "$GIMME_OS" = "linux" ]]; then
    shelltest --with ./bin/beanstalkd-cli_${GIMME_OS}_${GIMME_ARCH}${EXT} --diff --color --all tests
fi
