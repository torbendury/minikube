# Minimal install

```bash
make
```

As minimal as it gets! :-)

This needs `minikube`, `kubectl`, `istioctl`, `docker` and `make` installed.

It will create a Minikube cluster, install a default istio installation, build two example applications and deploy them to the cluster using kustomize.
