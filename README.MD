# Hex to RGB Convertor 
A web application written in GO which can be easily deployed on kubernetes cluster with custom metric monitoring using prometheus and grafana

# prerequisite 
Version >=0.18.0 of the Prometheus Operator requires a Kubernetes cluster of version >=1.8.0. If you are just starting out with the Prometheus Operator, it is highly recommended to use the latest version.

# Step1
please ignore if helm is already install on both client and cluster and go to [Step2](#Step2)

To install helm tiller service account 
>kubectl apply -f helm/service-account.yml
To install the rbac role for tiller deployment using helm
>kubectl apply -f helm/role-binding.yml
>helm init --service-account tiller --wait

# Step2
Create a namespace `monitoring` which we will use to deploy prometheus
>kubectl apply -f monitoring/namespace.yml
Make sure helm repo is updated and we deploy latest `stable/prometheus-operator`
>helm repo up

To install promethues 
>helm install stable/prometheus-operator --namespace monitoring --name prometheus

incase it throws error that one of the following mentioned cdr already exists add the `--set  prometheusOperator.createCustomResource=false` this will ignore the CRD installation.

- prometheusrules.monitoring.coreos.com
- servicemonitors.monitoring.coreos.com
- alertmanagers.monitoring.coreos.com
- prometheuses.monitoring.coreos.com
- podmonitors.monitoring.coreos.com

>helm install stable/prometheus-operator --namespace monitoring --name prometheus --set  prometheusOperator.createCustomResource=false

# Step3
Install hle;m charts for Go application

To install go application 
>helm install k8s/go-server/helm-chart



