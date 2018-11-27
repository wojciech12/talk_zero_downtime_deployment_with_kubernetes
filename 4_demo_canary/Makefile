prepare:
	kubectl scale --replicas=3 deploy/api-status-nginx-blue && \
	kubectl scale --replicas=0 deploy/api-status-nginx-green && \
	kubectl patch service api-status --type=json -p='[{"op": "remove", "path": "/spec/selector/label"}]'

clean:
	kubectl delete svc api-status && \
	kubectl delete deploy api-status-nginx-blue && \
	kubectl delete deploy api-status-nginx-green
