prepare:
	kubectl scale --replicas=3 deploy/api-status-nginx-blue && \
	kubectl scale --replicas=0 deploy/api-status-nginx-green && \
	kubectl patch service api-status --type=json -p='[{"op": "remove", "path": "/spec/selector/label"}]'

watch_curl:
	 export SVC_PORT=$$(kubectl get service api-status --output='jsonpath="{.spec.ports[0].nodePort}"' | tr -d '"') ; \
	 watch -n0.3 -x curl -s $$(minikube ip):$${SVC_PORT}

watch_get_po:
	 watch -n0.3 -x kubectl get po

delete:
	kubectl delete svc api-status
