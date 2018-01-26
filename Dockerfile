FROM golang as build
COPY . /go/src/github.com/philoserf/canary/
WORKDIR /go/src/github.com/philoserf/canary/
RUN go build -o /canary

FROM scratch
LABEL maintainer="mark@philoserf.com"
COPY --from=build canary /
CMD ["/canary", "-h"]
