##########################
Green Blue with Kubernetes
##########################

Prereq
~~~~~~

Get the images to docker hub:

::

  make docker_deploy

Notice: You need to change DOCKER_PROJECT_ID in `Makefile <Makefile>`_.

Demo
~~~~

::

  minikube start
  kubectl config use-context minikube

  # let's get our service running
  kubectl apply -f kube-service-status.yaml

  # our first version: blue
  kubectl apply -f kube-nginx-blue.yaml

  # let's connect to our blue 

  # get the IP
  minikube ip

  # get the port
  export SVC_PORT=$(kubectl get service api-status --output='jsonpath="{.spec.ports[0].nodePort}"' | tr -d '"')

  curl $(minikube ip):${SVC_PORT}

  # let's get the green running
  kubectl apply -f kube-nginx-green.yaml

  # in another terminal window
  watch -n0.3 -x kubectl get po

  # 
  watch -n0.3 -x curl -s $(minikube ip):${SVC_PORT}

  # let's make the switch
  kubectl patch service api-status -p '{"spec":{"selector": {"label": "green"} }}'

Notice: we could also accomplish this strategy using Ingress instead of Service.