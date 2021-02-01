.DEFAULT_GOAL := swagger

install_swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

kubectl_up:
	kubectl create -f k8s-postgres-service.yaml -f k8s-postgres-deployment.yaml && kubectl create -f k8s-service.yaml -f k8s-deployment.yaml

kubectl_down:
	kubectl delete -f k8s-service.yaml -f k8s-deployment.yaml && kubectl delete -f k8s-postgres-service.yaml -f k8s-postgres-deployment.yaml

kubectl_display:
	kubectl get deployments,services,pods,endpoints,rs

minikube_env:
	minikube docker-env

replace_activity_collector:
	go mod edit -replace github.com/garcialuis/ActivityCollector=/home/luis/Documents/Projects/GoLang/ActivityCollector