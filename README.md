# Mesh Demo


## setup k3s

```bash
curl -sfL https://get.k3s.io | sh -

# setup k3s with docker (or optional use containerd cli)
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="--docker" sh -

# check status
sudo systemctl status k3s

# allow kubectl to use k3s
mkdir -p ~/.kube
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
sudo chown $USER:$USER ~/.kube/config
chmod 600 ~/.kube/config
export KUBECONFIG=~/.kube/config
```

```bash

# Install Gateway API CRDs
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.0.0/standard-install.yaml

# Install Istio
curl -L https://istio.io/downloadIstio | sh -
cd istio-*
sudo mv bin/istioctl /usr/local/bin/
echo 'export PATH=$PATH:/usr/local/bin' >> ~/.zshrc
source ~/.zshrc

# install istio on k8s
# might need this if you are getting certificate errors
# export KUBECONFIG=/etc/rancher/k3s/k3s.yaml 
istioctl install --set profile=demo -y

# install helm
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# install metrics-server (optional,k3s has metrics-server built in)
helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server/
helm repo update
helm install metrics-server metrics-server/metrics-server \
  --namespace kube-system \
  --set args="{--kubelet-insecure-tls}" \
  --set apiService.insecureSkipTLSVerify=true

# # install keda
# helm repo add kedacore https://kedacore.github.io/charts
# helm repo update
# helm install keda kedacore/keda --namespace keda --create-namespace

# dev setup
kubectl create namespace dev
kubectl label namespace dev istio-injection=enabled
source dev.sh
kubectl apply -f expose.dev.yaml
# if you get value mount issue check dev.sh and update code_root
```

## Setup istio dashboard

```bash
# install istio addons
kubectl apply -f samples/addons/kiali.yaml
kubectl apply -f samples/addons/prometheus.yaml  
kubectl apply -f samples/addons/jaeger.yaml      
kubectl apply -f samples/addons/grafana.yaml 

# Access dashboards (run in separate terminals as needed)
istioctl dashboard kiali        # Service mesh visualization
istioctl dashboard jaeger      # Distributed tracing
istioctl dashboard grafana     # Metrics and monitoring
```


## Stress test
```bash
apt-get install apache2-utils
# On Mac
# brew install apache2-utils
ulimit -n 4096 
ab -n 100000 -c 2000 http://localhost:9090/
```