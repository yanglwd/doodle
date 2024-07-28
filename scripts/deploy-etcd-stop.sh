#!/bin/bash
set -e

# etcd 容器目录
etcd_containers=/data/containers/etcd

# 停止etcd
cd ${etcd_containers}
docker-compose down -v