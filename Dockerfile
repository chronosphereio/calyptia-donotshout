FROM golang:alpine AS build

WORKDIR /src
COPY . ./
RUN go mod vendor && \
	go build

FROM alpine:latest
COPY --from=build /src/donotshout /donotshout
ENTRYPOINT ["/donotshout"]
