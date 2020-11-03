# custom-scheduler

This repo is a sample for custom scheduler using Kubernetes Scheduling Framework.

## Build docker image

```shell
// 编译和构建镜像
$ make image

// 对镜像打标签
$ docker tag myscheduler piaoliangkb/myscheduler:v1.18

// 打包镜像并拷贝到 worker 节点
$ docker image save piaoliangkb/myscheduler:v1.18 -o myscheduler_v1.18.tar.gz

// worker 节点导出镜像
$ docker image load -i myscheduler_v1.18.tar.gz
```

## Deploy scheduler

```shell
$ kubectl apply -f deploy/myscheduler.yaml
```

## Using scheduler

```shell
// pod 指定 myscheduler.yaml 文件中定义的 schedulerName
$ kubectl apply -f deploy/nginx-pod-myscheduler.yaml
$ kubectl apply -f deploy/nginx-pod-default-scheduler.yaml
```

## Scheduler log

查看 scheduler log

```shell
$ ./log_myscheduler.sh
```

## Delete scheduler

```shell
$ kubectl delete -f deploy/myscheduler.yaml
```
