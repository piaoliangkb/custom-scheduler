# custom-scheduler

Custom scheduler based on [Kubernetes Scheduling Framework](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/) with two plugins:

- [httpreq](https://github.com/piaoliangkb/custom-scheduler/tree/v1.18/pkg/httpreq)
- [simplelog](https://github.com/piaoliangkb/custom-scheduler/tree/v1.18/pkg/simplelog)

Run this scheduler in a Kubernetes cluster V1.18. 

See API changes in branch v1.19 for Kubernetes V1.19.

## Build docker image

```
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

```
$ kubectl apply -f deploy/myscheduler.yaml
```

## Using scheduler

```
// Specify Pod.Spec.schedulerName in yaml file defined in myscheduler.yaml
$ kubectl apply -f deploy/nginx-pod-myscheduler.yaml
$ kubectl apply -f deploy/nginx-pod-default-scheduler.yaml
```

## Scheduler log

```
$ ./log_myscheduler.sh
```

## Delete scheduler

```
$ kubectl delete -f deploy/myscheduler.yaml
```
