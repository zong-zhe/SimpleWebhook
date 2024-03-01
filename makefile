mkdir simple-webhook
cd simple-webhook
go mod init simple-webhook
go build -o webhook-server
openssl genrsa -out tls.key 2048
openssl req -new -key tls.key -subj "/CN=simple-webhook.default.svc" -out tls.csr
openssl x509 -req -in tls.csr -signkey tls.key -out tls.crt
cat tls.crt | base64 | tr -d '\n'
cat tls.key | base64 | tr -d '\n'
docker build -t simple-webhook:latest .
docker push simple-webhook:latest
kubectl apply -f webhook-server.yaml
kubectl apply -f validating-webhook.yaml
kubectl run test-pod --image=nginx
k3d cluster create --config ./mycluster.yaml