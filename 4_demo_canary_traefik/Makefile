apply_traefik:
	kubectl apply -f traefik-ingress-controller_rbac.yaml
	kubectl apply -f traefik-ingress-controller_svc.yaml
	kubectl apply -f traefik-ingress-controller.yaml

# for better demo
prepare:
	kubectl scale --replicas=1 deploy/api-status-nginx-blue
	kubectl scale --replicas=1 deploy/api-status-nginx-green

apply_ingress:
	kubectl apply -f ingress.yaml

apply: prepare apply_traefik
	kubectl apply -f prod-svc.yaml
	kubectl apply -f canary-svc.yaml
	$(MAKE) apply_ingress

watch_curl:
	export SVC_PORT=$$(kubectl get service traefik-ingress-service -n kube-system --output='jsonpath="{.spec.ports[0].nodePort}"' | tr -d '"') ; \
	watch -n0.3 -x curl -s $$(minikube ip):$${SVC_PORT}/status

watch_get_po:
	 watch -n0.3 -x kubectl get po

delete:
	kubectl delete -f prod-svc.yaml || true
	kubectl delete -f canary-svc.yaml
