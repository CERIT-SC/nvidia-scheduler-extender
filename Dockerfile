FROM golang:1.18-alpine as build

WORKDIR /go/src/github.com/CERIT-SC/nvidia-scheduler-extender
COPY . .

RUN go mod download && \
    go build -o /go/bin/nvidia-scheduler-extender cmd/*.go

FROM alpine

COPY --from=build /go/bin/nvidia-scheduler-extender /usr/bin/nvidia-scheduler-extender

CMD ["nvidia-scheduler-extender"]
