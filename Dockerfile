FROM golang:1.4
MAINTAINER Eric Holmes <eric@remind101.com>

ADD inspector.go /tmp/inspector.go

ENV PORT 80
EXPOSE 80

CMD ["go", "run", "/tmp/inspector.go"]
