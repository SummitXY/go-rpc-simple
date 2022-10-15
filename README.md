# go-rpc-simple

## build docker image

```
docker build --tag go-run-simple . 
```

如果提示`Cannot connect to the Docker daemon at unix:///var/run/docker.sock`

需要打开docker客户端应用， 会自动运行docker进程

## push docker image

```shell
docker push qxybest/go-rpc-simple:latest
```

## run docker container

```shell
 docker run -p 9001:9001 qxybest/go-rpc-simple
```
> 其中port1:port2, port1是指外部访问的端口， port2是指这个rpc服务listen的端口

## apply k8s service
```shell
// 这个是暴露集群外的端口
kubectl apply -f go-rpc-simple.yaml
// 只暴露对集群内的端口运行这个
kubectl apply -f go-rpc-simple-inter.yaml
```

## 获取该rpc service的对外ip:port
```shell
minikube service go-rpc-simple-service --url 
```
