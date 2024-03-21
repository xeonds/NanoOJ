FROM ubuntu:latest
RUN apt-get update && apt-get install -y \
    g++ gcc openjdk-17-jdk-headless golang python3
WORKDIR /app
EXPOSE 8080

CMD ["./nano-oj-linux-amd64"]
