###############################
Rolling-updates with Kubernetes
###############################

Prereq
~~~~~~

Deploy docker images to hub.docker.com, see `Makefile <demo-bluegreen/Makefile>`_

With minikube
~~~~~~~~~~~~~

::

  minikube start
  kubectl config use-context minikube

  kubectl apply -f kube-api-service.yaml

  # deployment
  kubectl apply -f kube-demo-api.yaml

  # notice, it takes longer time due to the initialDelaySeconds
  kubectl get po

  export SVC_PORT=$(kubectl get service demo-api --output='jsonpath="{.spec.ports[0].nodePort}"' | tr -d '"')

  curl $(minikube ip):${SVC_PORT}

:: 

  watch -n0.3 -x curl -s curl $(minikube ip):${SVC_PORT}
  kubectl set image  deployment/demo-api app=wojciech11/api-status:green

  # observe 
  kubectl get po


How to ensure, we wait for all the request to be closed?
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Use:

::

  terminationGracePeriodSeconds: 60

See: `traefik-ds.yaml <https://github.com/containous/traefik/blob/master/examples/k8s/traefik-ds.yaml>`_
