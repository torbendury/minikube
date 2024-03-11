# Running istio demo with minikube

For a minimal local installation, refer to [LOCAL.md](LOCAL.md). If you want a full-fledged example, continue reading.

```bash
# start a default minikube
minikube start

# wait until everything comes up and is reachable!
# install istio with a demo profile
istioctl install --set profile=demo -y

# wait until everything is settled
watch 'kubectl get pods -A'

# enable istio proxy injection for the default ns
kubectl label namespace default istio-injection=enabled

# install an example application if you want to (or install your own stuff)
kubectl create -f https://raw.githubusercontent.com/istio/istio/release-1.20/samples/bookinfo/platform/kube/bookinfo.yaml

# wait for everything to come up
watch 'kubectl get pods -A'

# if your ingressgateway does not get a LB ip
kubectl get svc -A
# then you need to create a minikube tunnel in a separate terminal
minikube tunnel
# leave it open, leave it alone

# verify application is running and can reach services
kubectl exec "$(kubectl get pod -l app=ratings -o jsonpath='{.items[0].metadata.name}')" -c ratings -- curl -sS productpage:9080/productpage | grep -o "<title>.*</title>"

# create a istio gateway so the bookinfo app is reachable
kubectl create -f https://raw.githubusercontent.com/istio/istio/release-1.20/samples/bookinfo/networking/bookinfo-gateway.yaml

# set some ENV vars for convenience
export INGRESS_NAME=istio-ingressgateway
export INGRESS_NS=istio-system
export INGRESS_HOST=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export SECURE_INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="https")].port}')
export TCP_INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="tcp")].port}')
export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT

# check if your book app is reachable in your browser
echo "http://${GATEWAY_URL}/productpage"
# see that every time you reload, the review stars are different (served by reviews v1,v2,v3) - those are served by a "least request" approach.

# proceed here: https://istio.io/latest/docs/setup/getting-started/#dashboard
```
