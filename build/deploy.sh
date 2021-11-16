#!/bin/bash

currentPath=$(dirname $(readlink -f "$0"))
cd $currentPath
if [ $1 = "develop" ]; then
  deploy_file='./deploy-dev.yaml'
  container_name="imagedesign-dev"
elif [ ${1:0:1} = "v" ]; then
  deploy_file='./deploy.yaml'
  container_name="imagedesign"
else
  echo "THIS CI_COMMIT_REF_NAME $1 DOES NOT REQUIRE DEPLOYMENT"
  exit 0
fi

if [ -z "$(kubectl get deployment ${container_name} 2>/dev/null)" ]; then
  echo "deployment \"${container_name}\"  not exists. prepare for create"
  kubectl apply -f ${deploy_file} --validate=false
fi
kubectl patch deployment ${container_name} -p '{"spec":{"template":{"spec":{"containers":[{"name":"'${container_name}'","image": "registry.cn-hangzhou.aliyuncs.com/branch/imagedesign:'${1}'","env":[{"name":"RESTART_TIME","value":"'$(date +%s)'"}]}]}}}}'
