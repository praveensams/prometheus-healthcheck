FROM golang

WORKDIR /mnt

COPY *.go .

COPY url.txt /mnt/

RUN go build health.go

EXPOSE 9101

CMD [ "./health" ]

