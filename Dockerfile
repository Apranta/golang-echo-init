FROM golang:alpine

ARG APPNAME="echostarter_example"

ADD . $GOPATH/src/"${APPNAME}"
WORKDIR $GOPATH/src/"${APPNAME}"

RUN apk add --update git gcc libc-dev;
#  tzdata wget gcc libc-dev make openssl py-pip;

RUN go get -u github.com/golang/dep/cmd/dep

CMD dep ensure -v \
    && go build -v -o $GOPATH/bin/"${APPNAME}" \
    # run app mode
    && "${APPNAME}" run ;

EXPOSE 8050