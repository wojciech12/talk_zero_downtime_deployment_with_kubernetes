========
Recreate
========

Demo shows the recreate deployment strategy with Kubernetes.

Prereq
~~~~~~

Get the images to docker hub:

::

  make docker_deploy

Notice: You need to change DOCKER_PROJECT_ID in `Makefile <Makefile>`_.

Demo
~~~~

Notice: it causes downtime.

::

  minikube start
  kubectl config use-context minikube

  kubectl apply -f kube-api-service.yaml

  # deployment
  kubectl apply -f kube-demo-api.yaml

  # notice, it takes longer time due to the initialDelaySeconds
  kubectl get po

  # get the port our service is exposed on by minikube
  export SVC_PORT=$(kubectl get service demo-api --output='jsonpath="{.spec.ports[0].nodePort}"' | tr -d '"')

  # check whether it is correct:
  curl $(minikube ip):${SVC_PORT}

:: 

  # run curl in an endless loop to
  # see the changes from the user perspective:
  watch -n0.3 -x curl -s $(minikube ip):${SVC_PORT}

  # observe 
  watch -n0.3 -x kubectl get po

  # let's get from 1.0.0 to 2.0.0
  kubectl set image  deployment/demo-api app=wojciech11/api-status:2.0.0


  # notice: beacause of the initialDelaySeconds readiness and liveness probes
  # we need to wait longer

Notice:

We would also change the version in `kube-demo-api.yaml <kube-demo-api.yaml>`_ to 2.0.0 and use ``kube apply -f kube-demo-api.yaml``.