FROM ubuntu:latest
RUN apt-get update && apt-get install -y \
    g++ \
    python3
WORKDIR /app
EXPOSE 8080

CMD ["./nano-oj-linux-amd64-1.1.0"]