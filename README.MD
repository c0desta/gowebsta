# Hex to RGB Convertor 
A web application written in GO which can be easily deployed on kubernetes cluster with custom metric monitoring using prometheus and grafana

![Preview of Grafana](grafana2.jpg?raw=true "Title")

# Please read notes at the end of the file incase you face any issues

# prerequisite 
Version >=0.18.0 of the Prometheus Operator requires a Kubernetes cluster of version >=1.8.0. If you are just starting out with the Prometheus Operator, it is highly recommended to use the latest version.

# Step1
please ignore if helm is already install on both client and cluster and go to [Step2](#Step2)

To install helm tiller service account 
>kubectl apply -f helm/service-account.yml
To install the rbac role for tiller deployment using helm
```
kubectl apply -f helm/role-binding.yml
```
```
helm init --service-account tiller --wait
```
# Step2
Create a namespace `monitoring` which we will use to deploy prometheus
```
kubectl apply -f monitoring/namespace.yml
```
Make sure helm repo is updated and we deploy latest `stable/prometheus-operator`
```
helm repo up
```
To install promethues 
```
helm install stable/prometheus-operator --namespace monitoring --name prometheus
```
incase it throws error that one of the following mentioned cdr already exists add the `--set  prometheusOperator.createCustomResource=false` this will ignore the CRD installation.

- prometheusrules.monitoring.coreos.com
- servicemonitors.monitoring.coreos.com
- alertmanagers.monitoring.coreos.com
- prometheuses.monitoring.coreos.com
- podmonitors.monitoring.coreos.com

>helm install stable/prometheus-operator --namespace monitoring --name prometheus --set  prometheusOperator.createCustomResource=false

# Step3

Build Docker image replace username with your dockerhub username
```
docker build k8s/ -t username/hex2rgb:latest
```
push image to docker
```
docker push username/hex2rgb:latest
```

# Step4
Install hle;m charts for Go application

To install go application 
```
helm install k8s/go-server/helm-chart
```
# Step5

Open service/deployment on local There are two ways to do it. Using kubectl port-forward pod port:port or make a Service which expose the deployment 

example 1
get list of pods
```
kubectl get pods
```
`knobby-kiwi-go-app-k8s-67df99c5cf-8sg82       1/1     Running   0          94m`
```
kubectl port-forward knobby-kiwi-go-app-k8s-67df99c5cf-8sg82  3333:3000
```

and browsing it `http://localhost:3333/convert/hex-rgb/ffcc77`  this will return RGB equivalent `RGB(255, 204, 119)`

To get the deployment list (specify namespace if deployment in specific namespace)
>kubectl.exe get deployment
```
kubectl expose deployment knobby-kiwi-go-app-k8s --type=LoadBalancer --name=app-service
```

# Step6
access grafana `http://localhost:3000/login`
login id `admin`
password `prom-operator`

you can add a custom dashboard and import/edit panel json and copy paste the file ./grafana-panel.json


Notes: There are certain areas of improvement like building go application and later pushing it instead of including the go library in docker but for development this can be useful if we not sure what OS developer might be using.
