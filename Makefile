.PHONY: all cluster apps deploy
all: cluster apps deploy

cluster:
	minikube start
	kubectl ctx minikube
	istioctl install -y
	kubectl create ns a || true
	kubectl create ns b || true

apps:
	docker build -t torbendury/service-a:latest -f applications/Dockerfile --build-arg SERVICE_DIRECTORY=service-a applications/service-a
	docker image save torbendury/service-a:latest -o service-a.tar
	docker build -t torbendury/service-b:latest -f applications/Dockerfile --build-arg SERVICE_DIRECTORY=service-b applications/service-b
	docker image save torbendury/service-b:latest -o service-b.tar
	minikube image load service-a.tar
	minikube image load service-b.tar

deploy:
	kubectl apply -k k8s/kustomize
	istioctl install -f k8s/istio/istiooperator.yaml -y

destroy:
	minikube delete
