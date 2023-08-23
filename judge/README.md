# Judge Server

## Introduction

This is the judge server for NanoOJ. It uses docker to compile and run the code. It uses distributed workers to run the code.

## Structure

```text
- Nginx Server(Reverse Proxy & Load Balancer)
- Database Server
- Core Web Server 1
- Core Web Server 2
- ...
- Docker Judge Server 1
    - Worker 1
    - Worker 2
    - ...
    - Worker N
- Docker Judge Server 2
    - Worker 1
    - Worker 2
    - ...
    - Worker N
- ...
```

The system is designed to be **stateless**. All the data is stored in the database server. The judge server is designed to be scalable. You can add more judge servers to the system to increase the performance, or add more web servers to the system to increase the throughput.

If you want, you can also deploy the whole system on a single machine. Or, you can deploy the whole system on a cluster using Kubernetes.
