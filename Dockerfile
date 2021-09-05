FROM golang:1.17 as building

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o serverapp .

FROM scratch

COPY --from=building /app/serverapp /

EXPOSE 9999

CMD ["/serverapp"]