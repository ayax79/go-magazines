#!/bin/sh

REDIS_BASE=$(kubectl get pods -l app=redis -o name | grep master | sed -e 's/pod\///g' | sed -e 's/\-redis\-master\-[0-9]*//g')

REDIS_NAME="${REDIS_BASE}-redis"
REDIS_HOST="${REDIS_BASE}-redis-master.default.svc.cluster.local"
REDIS_SECRET="${REDIS_BASE}-redis"

mkdir -p ./build/k8s
cp k8s/* ./build/k8s
find ./build/k8s -name \*yaml \
    -exec sed -i -e "s/REDIS_NAME/$REDIS_NAME/g" {} \; \
    -exec sed -i -e "s/REDIS_HOST/$REDIS_HOST/g" {} \; \
    -exec sed -i -e "s/REDIS_SECRET/$REDIS_SECRET/g" {} \; 