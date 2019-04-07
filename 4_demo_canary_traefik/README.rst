===================
Canary with Traefik
===================

Run after `4_demo_canary <../4_demo_canary>`_, see `Makefile <Makefile>`_:

::

	kubectl apply -f *r.yaml
	kubectl apply -f *svc.yaml
	kubectl apply -f *_rbac.yaml

	kubectl get svc --namespace=kube-system

::

  kubectl scale --replicas=1 deploy/api-status-nginx-blue
  kubectl scale --replicas=1 deploy/api-status-nginx-green

::

	export SVC_PORT=$(kubectl get service traefik-ingress-service -n kube-system --output='jsonpath="{.spec.ports[0].nodePort}"' | tr -d '"')
	watch -n0.3 -x curl -s $(minikube ip):${SVC_PORT}/status

::

  # watch
  watch -n0.3 -x kubectl get po

::

  # change ingress.yaml
  kubectl apply -f ingress.yaml
