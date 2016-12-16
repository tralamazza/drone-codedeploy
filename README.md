# drone-codedeploy

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-codedeploy/status.svg)](http://beta.drone.io/drone-plugins/drone-codedeploy)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-codedeploy?status.svg)](http://godoc.org/github.com/drone-plugins/drone-codedeploy)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-codedeploy)](https://goreportcard.com/report/github.com/drone-plugins/drone-codedeploy)
[![Join the chat at https://gitter.im/drone/drone](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/drone/drone)

Drone plugin to deploy or update a project on AWS CodeDeploy. For the
usage information and a listing of the available options please take a look at
[the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build --rm=true -t plugins/codedeploy .
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-codedeploy' not found or does not exist..
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -e PLUGIN_REGION=eu-west-1 \
  -e PLUGIN_DEPLOYMENT_GROUP=mywebservers \
  -e AWS_ACCESS_KEY_ID=<token> \
  -e AWS_SECRET_ACCESS_KEY=<secret> \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/codedeploy --dry-run
```
