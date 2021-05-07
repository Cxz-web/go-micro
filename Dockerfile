FROM alpine
ADD go-micro /go-micro
ENTRYPOINT [ "/go-micro" ]
