apply:
	kubectl apply -f kube-api-service.yaml
	kubectl apply -f kube-demo-api.yaml

watch_curl:
	 export SVC_PORT=$$(kubectl get service demo-api --output='jsonpath="{.spec.ports[0].nodePort}"' | tr -d '"') ; \
	 watch -n0.3 -x curl -s $$(minikube ip):$${SVC_PORT}

watch_get_po:
	 watch -n0.3 -x kubectl get po

delete:
	kubectl delete -f kube-api-service.yaml || true
	kubectl delete -f kube-demo-api.yaml