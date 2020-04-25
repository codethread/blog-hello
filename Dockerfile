FROM golang:1.14 as base

ENV APP_HOME=/app
RUN mkdir $APP_HOME
WORKDIR $APP_HOME

FROM base as dev
COPY go.mod $APP_HOME
COPY go.sum $APP_HOME

COPY tools.go $APP_HOME
RUN go mod download
RUN cat tools.go | grep _ | awk -F'"' '{print $2}' | xargs -tI % go install %

FROM dev as test
# copy in src so we can run tests and lint etc
COPY ./ $APP_HOME
