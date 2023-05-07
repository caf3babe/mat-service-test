# Quickstart

```bash
# run local docker registry
docker run -d -p 5001:5000 --restart=always --name registry registry:2
# build app
docker build -t localhost:5001/counter .
# push image
docker push localhost:5001/counter
# start cluster
minikube start --insecure-registry "host.docker.internal:5001"
# deploy app
kubectl apply -f deployment.yaml
# deploy service
kubectl apply -f service.yaml
# make minikube network available on host
minikube tunnel
```

Run tests
```bash

for i in {1..1000}
do
  curl http://localhost/increment
done

for i in {1..10}
do
  curl http://localhost/show 2>/dev/null | jq
done

```

Example output
```json
{
  "counter": 1208
}
{
  "counter": 1208
}
{
  "counter": 1195
}
{
  "counter": 1194
}
{
  "counter": 1195
}
{
  "counter": 1195
}
{
  "counter": 1162
}
{
  "counter": 1208
}
{
  "counter": 1229
}
{
  "counter": 1208
}
```