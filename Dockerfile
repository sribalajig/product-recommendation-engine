# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:1.8.3

# Create a directory inside the container to store all our application and then make it the working directory.
RUN mkdir -p /go/src/home24
WORKDIR /go/src/home24
# Copy the example-app directory (where the Dockerfile lives) into the container.
COPY . /go/src/home24

# Download and install any required third party dependencies into the container.
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get -u golang.org/x/net/context
RUN go get -u gopkg.in/olivere/elastic.v5

RUN export PATH=$PATH:$GOPATH/bin

RUN bee run /go/src/home24/recommendation.api