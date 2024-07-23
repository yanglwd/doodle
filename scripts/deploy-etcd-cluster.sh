#!/bin/bash
set -e

etcd_initial_cluster="etcd-s1=http://172.25.0.101:2380,etcd-s2=http://172.25.0.102:2380,etcd-s3=http://172.25.0.103:2380"
etcd_containers=/data/containers/etcd

ip_s1=172.25.0.101
ip_s2=172.25.0.102
ip_s3=172.25.0.103

./deploy-etcd-conf.sh s1 ${ip_s1} ${etcd_initial_cluster} ${etcd_containers}
./deploy-etcd-conf.sh s2 ${ip_s2} ${etcd_initial_cluster} ${etcd_containers}
./deploy-etcd-conf.sh s3 ${ip_s3} ${etcd_initial_cluster} ${etcd_containers}

# 复制 docker-compose.yml 文件
cp -f ../docker/docker-compose.yml ${etcd_containers}

# 修改 docker-compose.yml 文件
sed -i 's/ETCD_IPV4_S1/'"${ip_s1}"'/g' ${etcd_containers}/docker-compose.yml
sed -i 's/ETCD_IPV4_S2/'"${ip_s2}"'/g' ${etcd_containers}/docker-compose.yml
sed -i 's/ETCD_IPV4_S3/'"${ip_s3}"'/g' ${etcd_containers}/docker-compose.yml
sed -i 's/ETCD_IPV4_SUBNET/172.25.0.0\/16/g' ${etcd_containers}/docker-compose.yml
sed -i 's/ETCD_IPV4_GATEWAY/172.25.0.1/g' ${etcd_containers}/docker-compose.yml

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
