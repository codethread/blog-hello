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

FROM base as build
# dropped these in to cache mod download
COPY go.mod $APP_HOME
COPY go.sum $APP_HOME
RUN go mod download
COPY ./ $APP_HOME
RUN go build -o bin

FROM base as prod

COPY --from=build $APP_HOME/assets ./assets
COPY --from=build $APP_HOME/templates ./templates
COPY --from=build $APP_HOME/certs ./certs
COPY --from=build $APP_HOME/bin .

WORKDIR $APP_HOME
CMD ./bin

