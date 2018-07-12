##########################
Green Blue with Kubernetes
##########################

With minikube
~~~~~~~~~~~~~

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

  # 
  watch -n0.3 -x curl -s  192.168.99.100:32151

  # let's make the switch
  kubectl patch service api-status -p '{"spec":{"selector": {"label": "green"} }}'
 
