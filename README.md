# Evaluating Performance of gRPC-Web vs grpc-gateway

This repository contains 2 equal APIs: using gRPC-Web with Envoy and using grpc-gateway. The goal is to run benchmarks for 2 approaches and compare them. APIs have 1 endpoint to ping. Benchmarks also include response parsing.

## Requirements

- Docker
- docker-compose

## Description

Using nodejs client and `benchmarkify` lib to benchmark.

Deployment model:

Flow definition:

- Direct request: call direct HTTP request to backend. Review spring-grpc-http to know more about this back. Result in this test will be seen like a baseline.
- Using gRPC-Web: 

## Results


