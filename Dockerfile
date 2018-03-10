FROM golang:1.10-alpine as builder

ENV DEST $GOPATH/src/github.com/andrepinto/helmsman
WORKDIR $DEST

RUN apk --update add git make upx
RUN go get -u github.com/golang/dep/...

ADD . $DEST
RUN make dep
RUN make build
RUN upx ./bin/helmsman

# ----------------------------------------------------------------------------

FROM alpine:3.7

WORKDIR /app
VOLUME /app/charts
EXPOSE 8000

COPY --from=builder /go/src/github.com/andrepinto/helmsman/bin/helmsman /app/helmsman

CMD /app/helmsman
