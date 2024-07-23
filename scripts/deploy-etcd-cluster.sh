#!/bin/bash
set -e

etcd_initial_cluster="etcd-s1=http://172.25.0.101:2380,etcd-s2=http://172.25.0.102:2380,etcd-s3=http://172.25.0.103:2380"
etcd_containers=/data/containers/etcd

./deploy-etcd-conf.sh s1 172.25.0.101 ${etcd_initial_cluster} ${etcd_containers}
./deploy-etcd-conf.sh s2 172.25.0.102 ${etcd_initial_cluster} ${etcd_containers}
./deploy-etcd-conf.sh s3 172.25.0.103 ${etcd_initial_cluster} ${etcd_containers}

# 复制 docker-compose.yml 文件
cp -f ../docker/docker-compose.yml ${etcd_containers}

# 删除 docker 网络
NETWORK_NAME="containers_etcd-net"
if docker network ls | grep -q ${NETWORK_NAME}; then
  docker network rm ${NETWORK_NAME}
fi

# 创建 etcd 服务
cd ${etcd_containers}
docker compose up -d

# 验证 etcd 服务
docker compose ps
