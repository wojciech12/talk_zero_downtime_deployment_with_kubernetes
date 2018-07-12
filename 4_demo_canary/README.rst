######
Canary
######

Several strategies with Treafik, Istio,... below an example of a manual deployment.

We can use our example for `blue green deployments <../3_demo_bluegreen/>`_:

::

  kubectl scale --replicas=3 deploy/api-status-nginx-blue
  kubectl scale --replicas=0 deploy/api-status-nginx-green

  kubectl patch service api-status --type=json -p='[{"op": "remove", "path": "/spec/selector/label"}]'

  kubectl scale --replicas=3 deploy/api-status-nginx-blue
  kubectl scale --replicas=1 deploy/api-status-nginx-green

  # no errors, let's continoue
  kubectl scale --replicas=2 deploy/api-status-nginx-blue
  kubectl scale --replicas=2 deploy/api-status-nginx-green

  # no errors, let's continoue
  kubectl scale --replicas=1 deploy/api-status-nginx-blue
  kubectl scale --replicas=3 deploy/api-status-nginx-green

  # no errors, let's continoue
  kubectl scale --replicas=0 deploy/api-status-nginx-blue
  kubectl scale --replicas=3 deploy/api-status-nginx-green

  # now we can delete the deployment blue
