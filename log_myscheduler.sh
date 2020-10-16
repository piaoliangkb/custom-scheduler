#!/bin/bash

podname=$(kubectl get po -n kube-system | grep myscheduler | awk '{print $1}')
kubectl logs -f $podname -n kube-system
