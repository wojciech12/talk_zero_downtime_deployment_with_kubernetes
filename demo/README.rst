###########
Demo Golang
###########

Basics
------

::

  make build
  ./demo

::

  curl 127.0.0.1:8090/healthz
  curl 127.0.0.1:8090/doUnHealthz
  curl 127.0.0.1:8090/healthz
  curl 127.0.0.1:8090/doHealthz

::

  curl 127.0.0.1:8090/ready
  curl 127.0.0.1:8090/doNotReady
  curl 127.0.0.1:8090/ready
  curl 127.0.0.1:8090/doReady

Docker
------

::

  docker run --name zero-kube-demo -p 8090:8090 -d wojciech11/zero-kube-demo:1.0.0
  docker logs zero-kube-demo -f
  docker stop zero-kube-demo
  docker stop --time=10 zero-kube-demo
  docker stop zero-kube-demo ; docker rm zero-kube-demo

Kubernetes
----------

::

  kubectl apply -f kube-service.yaml
  kubectl apply -f kube-deployment.yaml

  export SVC_PORT=$(kubectl get service zero-demo --output='jsonpath="{.spec.ports[0].nodePort}"' | tr -d '"')
  curl $(minikube ip):${SVC_PORT}/ready

  watch -n0.3 -x curl -s $(minikube ip):${SVC_PORT}/ready
  watch -n0.3 -x kubectl get po

  kubectl exec -it zero-demo-ff967d59f-wjldc   /bin/bash

  # no request getting to our pod
  $ curl 127.0.0.1:8090/doNotReady

  # now, let's inform kubernetes that we are ill:
  # kubernetes will restart the pod
  $ curl 127.0.0.1:8090/doUnHealthz

::

  # cleanup before the next demo
  kubectl delete -f kube-deployment.yaml
  kubectl apply -f kube-service.yaml
