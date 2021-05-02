# custom-scheduler

Custom scheduler based on [Kubernetes Scheduling Framework](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/)

## Build docker image

```shell
// Compile and build docker image
$ make image

// Tag docker image
$ docker tag myscheduler piaoliangkb/myscheduler:v1.18

// Save docker image and copy to k8s worker node
$ docker image save piaoliangkb/myscheduler:v1.18 -o myscheduler_v1.18.tar.gz

// Load image on worker node
$ docker image load -i myscheduler_v1.18.tar.gz
```

## Deploy scheduler

```shell
$ kubectl apply -f deploy/myscheduler.yaml
```

## Using scheduler

```shell
// Specify Pod.Spec.schedulerName in yaml file defined in myscheduler.yaml
$ kubectl apply -f deploy/nginx-pod-myscheduler.yaml
$ kubectl apply -f deploy/nginx-pod-default-scheduler.yaml
```

## Scheduler log

```shell
$ ./log_myscheduler.sh
```

## Delete scheduler

```shell
$ kubectl delete -f deploy/myscheduler.yaml
```
