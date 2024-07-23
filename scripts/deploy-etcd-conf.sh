#!/bin/bash
set -e

if [ $# -ne 4 ]; then
    echo "Usage: $0 <etcd_name> <etcd_client_ip> <etcd_initial_cluster> <docker_container_dir>"
    exit 1
fi

# 设置 etcd 服务名称
etcd_name=${1}

# 设置 etcd 服务对外 IP
etcd_client_ip=${2}

# 设置 etcd 初始化集群信息
etcd_initial_cluster=${3}

# 设置 containers 基础目录
docker_container_dir=${4}
docker_etcd_dir=${docker_container_dir}/${etcd_name}

# 创建基础目录
mkdir -p ${docker_etcd_dir}/{data,config}

# 创建 etcd 配置文件
function deploy_etcd_config(){
cat > ${docker_etcd_dir}/config/etcd.conf.yml <<-EOF
name: etcd-${etcd_name}
data-dir: /var/etcd
listen-client-urls: http://0.0.0.0:2379
advertise-client-urls: http://${etcd_client_ip}:2379
listen-peer-urls: http://0.0.0.0:2380
initial-advertise-peer-urls: http://${etcd_client_ip}:2380
initial-cluster: ${etcd_initial_cluster}
initial-cluster-token: etcd-cluster
initial-cluster-state: new
logger: zap
log-level: info
#log-outputs: stderr

EOF
}

echo -e "\033[1;32m [1].Deploy etcd config.\n \033[0m"
deploy_etcd_config