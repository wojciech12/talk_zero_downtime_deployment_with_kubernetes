######
Canary
######

Several strategies with Treafik, Istio,... below an example of a manual deployment.

Use the k8s configuration from `the blue-green deployment strategy <../3_demo_bluegreen/>`_.

Assuming the blue is our prod and green is our canary:

::

  kubectl scale --replicas=3 deploy/api-status-nginx-blue
  kubectl scale --replicas=0 deploy/api-status-nginx-green

::

  # let's get the service to blue (our prod)
  # and green (our canary)
  kubectl patch service api-status --type=json -p='[{"op": "remove", "path": "/spec/selector/label"}]'

::

  kubectl scale --replicas=3 deploy/api-status-nginx-blue
  kubectl scale --replicas=1 deploy/api-status-nginx-green

  # no errors, let's continue
  kubectl scale --replicas=2 deploy/api-status-nginx-blue
  kubectl scale --replicas=2 deploy/api-status-nginx-green

  # no errors, let's continue
  kubectl scale --replicas=1 deploy/api-status-nginx-blue
  kubectl scale --replicas=3 deploy/api-status-nginx-green

  # no errors, let's continue
  kubectl scale --replicas=0 deploy/api-status-nginx-blue
  kubectl scale --replicas=3 deploy/api-status-nginx-green

  # now we can delete the deployment blue

::

  export SVC_PORT=$(kubectl get service api-status --output='jsonpath="{.spec.ports[0].nodePort}"' | tr -d '"')
  watch -n0.3 -x curl -s $(minikube ip):${SVC_PORT}
