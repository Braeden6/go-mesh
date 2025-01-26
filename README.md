# Mesh Demo

```bash
docker build -t mesh-demo/frontend:v1  frontend/
docker build -t mesh-demo/backend:v1  backend/
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.0.0/standard-install.yaml
kubectl apply -f deploy.yaml


# install and setup istio
curl -L https://istio.io/downloadIstio | sh -
kubectl apply -f samples/addons/kiali.yaml
kubectl apply -f samples/addons/prometheus.yaml  
kubectl apply -f samples/addons/jaeger.yaml      
kubectl apply -f samples/addons/grafana.yaml 
cd istio-*
sudo mv bin/istioctl /usr/local/bin/
echo 'export PATH=$PATH:/usr/local/bin' >> ~/.zshrc
istioctl dashboard kiali

 
kubectl scale deployment frontend --replicas=1
kubectl scale deployment backend --replicas=1




docker build -t mesh-demo/backend:dev -f backend/Dockerfile.dev backend/
docker build -t mesh-demo/frontend:dev -f frontend/Dockerfile.dev frontend/
kubectl create namespace dev
kubectl apply -f backend/deploy.dev.yaml
kubectl apply -f frontend/deploy.dev.yaml
kubectl apply -f expose.dev.yaml

