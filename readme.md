## Open-Telemetry Golang Client Library Docs
* https://github.com/open-telemetry/opentelemetry-go
* https://github.com/open-telemetry/opentelemetry-go-contrib
* https://lightstep.com/blog/opentelemetry-101-what-are-metrics/


## Deploy Open-Telemetry-Collector to k8s
### Docs
* https://github.com/open-telemetry/opentelemetry-operator
* https://github.com/open-telemetry/opentelemetry-collector/blob/main/examples/demo/otel-collector-config.yaml
* https://opentelemetry.io/docs/collector/configuration/
* https://cert-manager.io/docs/installation/helm/
### Install
``` sh
# install cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.5.0/cert-manager.crds.yaml
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.5.0
# install open-telemetry operator
kubectl apply -f https://github.com/open-telemetry/opentelemetry-operator/releases/download/v0.31.0/opentelemetry-operator.yaml
# install open-telemetry-collector
kubectl apply -n try -f ./yamls/optl-collector.yaml
```
### Uninstall
``` sh
# uninstall open-telemetry-colector
kubectl delete -n try -f ./yamls/optl-collector.yaml
# uninstall open-telemetry operator
kubectl delete -f https://github.com/open-telemetry/opentelemetry-operator/releases/latest/download/opentelemetry-operator.yaml
# uninstall cert-manager
helm --namespace cert-manager delete cert-manager
kubectl delete -f https://github.com/jetstack/cert-manager/releases/download/v1.5.0/cert-manager.crds.yaml
```


## Test
* Deploy open-telemetry-collector to k8s
* Port-Forward
``` sh
kubectl port-forward $(kubectl -n try get pod | grep "my-optl-collector-collector" | awk '{print $1}') 30080:30080 -n try
```
* Run the test inside trace and metric folder with flag -v to see the tracing output in stdout