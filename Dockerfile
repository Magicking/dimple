FROM golang
MAINTAINER Sylvain Laurent

ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR github.com/Magicking/dimple
ENV PROJECT_NAME dimple-server

ADD vendor /usr/local/go/src
ADD cmd /go/src/${PROJECT_DIR}/cmd
ADD models /go/src/${PROJECT_DIR}/models
ADD restapi /go/src/${PROJECT_DIR}/restapi
ADD internal /go/src/${PROJECT_DIR}/internal

WORKDIR /go/src/${PROJECT_DIR}

RUN go build -v -o /go/bin/main /go/src/${PROJECT_DIR}/cmd/${PROJECT_NAME}/main.go

EXPOSE 8090

CMD ["/go/bin/main", "--host=0.0.0.0", "--port=8090" ]
