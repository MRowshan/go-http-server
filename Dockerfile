FROM golang
COPY go/src/app/app /opt/go-http-server/app
COPY go/src/app/static /opt/go-http-server/static
WORKDIR "/opt/go-http-server"
ENTRYPOINT ["/opt/go-http-server/app"]
