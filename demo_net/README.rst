###########
Demo DotNet
###########

Kubernetes
----------

::

  make docker_build
  make docker_deploy

  minikube start
  kubectl config use-context minikube

  kubectl apply -f kube-deployment.yaml
  kubectl apply -f kube-service.yaml

::

  # in the first terminal
  watch kubectl get po

::

  # in the second terminal
  export SVC_PORT=$(kubectl get service zero-demo-net --output='jsonpath="{.spec.ports[0].nodePort}"' | tr -d '"')
  watch -n0.3 -x curl -s $(minikube ip):${SVC_PORT}/ready

::

  kubectl exec -it zero-demo-net-745cdd7db8-7xfhp  /bin/bash

  # inform k8s we are not ready
  curl 127.0.0.1/doNotReady

  # after few seconds you should see one of the node
  # stopping handling requests

  # inform k8s we are ready:
  curl 127.0.0.1/doReady

::

  kubectl exec -it zero-demo-net-745cdd7db8-7xfhp  /bin/bash

  # let's inform k8s we are unhealthy
  curl 127.0.0.1/doUnHealthz

  # k8s will kill our pod

::

  # while observing logs
  kubectl logs zero-demo-net-745cdd7db8-7xfhp -f
  
  # kill the pod
  kubectl delete po zero-demo-net-745cdd7db8-7xfhp

Docker
------

::

  $(make docker_run_echo)
  docker logs zero-kube-demo-net -f
  docker stop zero-kube-demo-net
  docker stop --time=10 zero-kube-demo-net

::

  watch -n0.3 -x curl -s -I curl 127.0.0.1:8090/ready
