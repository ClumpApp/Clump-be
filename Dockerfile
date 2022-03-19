##
## Build
##
FROM golang:1.17 AS build

WORKDIR /go/src/app
COPY . .

RUN go build -ldflags="-s -w" -o /clump 

##
## Containerize
##
FROM busybox:1.34.1

WORKDIR /
COPY --from=build /clump /clump

EXPOSE 8080

ENTRYPOINT ["/clump"]
