FROM golang:alpine as builder
WORKDIR /app
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./run .

FROM alpine as runner
WORKDIR /root
# timezone and ca-certificates
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai
#Copy executable from builder
COPY --from=builder /app/run .
#Copy files from builder
COPY --from=builder /app/config .
COPY --from=builder /app/data .
COPY --from=builder /app/resources .
EXPOSE 8000
CMD ["./run"]