#!/bin/sh

trap "exit 0" TERM

# >&2 redirects stdout to stderr( so it wont get written to zathurarc )
cd /root
apk add go >&2
go install github.com/gennaro-tedesco/zathuraconf@latest >&2
./go/bin/zathuraconf -fp ./zathurarc config.json >&2

# this outputs to stdout and can be written to anything
cat zathurarc
