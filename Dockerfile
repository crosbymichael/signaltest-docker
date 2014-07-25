FROM crosbymichael/golang

COPY main.go /

RUN go build  main.go

ENTRYPOINT ["./main"]
