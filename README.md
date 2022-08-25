# echo
Simple tcp echo server that outputs connection information via log:

```
2022/08/25 07:13:17 localAddr: 10.88.5.25:8007 <-> RemoteAddr: 10.132.0.9:55289 - CONNECTED
2022/08/25 07:13:19 localAddr: 10.88.5.25:8007 <-> RemoteAddr: 10.132.0.9:55289 - DISCONNECTED (bytes read/written: 15)
2022/08/25 07:13:21 localAddr: 10.88.5.25:8007 <-> RemoteAddr: 10.132.0.9:47769 - CONNECTED
2022/08/25 07:13:22 localAddr: 10.88.5.25:8007 <-> RemoteAddr: 10.132.0.9:47769 - DISCONNECTED (bytes read/written: 19)
2022/08/25 07:13:29 localAddr: 10.88.5.25:8007 <-> RemoteAddr: 10.132.0.9:49645 - CONNECTED
2022/08/25 07:13:30 localAddr: 10.88.5.25:8007 <-> RemoteAddr: 10.132.0.9:49645 - DISCONNECTED (bytes read/written: 18)
2022/08/25 07:38:26 localAddr: 10.88.5.25:8007 <-> RemoteAddr: 10.132.0.9:35059 - CONNECTED
2022/08/25 07:40:24 localAddr: 10.88.5.25:8007 <-> RemoteAddr: 10.132.0.9:35059 - DISCONNECTED (bytes read/written: 70)
2022/08/25 07:40:31 localAddr: 10.88.5.25:8007 <-> RemoteAddr: 10.132.0.9:42767 - CONNECTED
```

## usage

Default port to listen on is `8007`. To overwrite use environment variable `PORT` 

Run from source via `go`
```
go run main.go
```

Install binary via `go`
```
go install github.com/mrbuk/echo
```

Docker
```
docker run --rm -p 8007:8007 mrbuk/echo:0.2
```

Kubernetes exposed via a `LoadBalancer`
```
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: echoserver
spec:
  replicas: 3
  selector:
    matchLabels:
      app: echoserver
  template:
    metadata:
      labels:
        app: echoserver
    spec:
      containers:
      - image: mrbuk/echo:0.2
        imagePullPolicy: IfNotPresent
        name: echoserver
        ports:
        - containerPort: 8007
        env:
        - name: PORT
          value: "8007"
---
apiVersion: v1
kind: Service
metadata:
  name: echoserver
spec:
  ports:
    - port: 8007
      targetPort: 8007
      protocol: TCP
  type: LoadBalancer
  selector:
    app: echoserver
```
