FROM ubuntu:20.04

RUN apt-get update && \
    apt-get install -y gcc g++ openjdk-11-jdk-headless golang python3 && \
    rm -rf /var/lib/apt/lists/*

COPY judge /main/judge
RUN chmod +x /main/judge

WORKDIR /main

ENTRYPOINT ["/bin/bash", "-c", "/main/judge \"$@\"", "--"]
CMD ["$LANG", "$TIME_LIMIT"]
