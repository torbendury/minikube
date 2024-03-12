.PHONY: all cluster apps deploy
all: cluster apps deploy

cluster:
	minikube start
	kubectl ctx minikube
	istioctl install -y
	kubectl create ns a || true
	kubectl create ns b || true

apps:
	docker build -t torbendury/service-a:latest -f applications/service-a/Dockerfile applications/service-a
	docker image save torbendury/service-a:latest -o service-a.tar
	docker build -t torbendury/service-b:latest -f applications/service-b/Dockerfile applications/service-b
	docker image save torbendury/service-b:latest -o service-b.tar
	docker build -t torbendury/service-b-canary:latest -f applications/service-b/Dockerfile applications/service-b
	docker image save torbendury/service-b-canary:latest -o service-b-canary.tar
	minikube image load service-a.tar
	minikube image load service-b.tar
	minikube image load service-b-canary.tar
	rm service-a.tar service-b.tar service-b-canary.tar

deploy:
	kubectl apply -k k8s/kustomize
	istioctl install -f k8s/istio/istiooperator.yaml -y

destroy:
	minikube delete
