# Workshop 02

- Create the following architecture in kubernetes
  - a database using `deployment` resource in k8s
  - a `service` resource in k8s such that the app(s) can communicate with database using FQDN rather than IP address
  - an app using `deployment` resource in k8s with the following requirements
    - able to connect to database
    - readiness & liveness probes
    - 3 replicas
  - a `service` resource such that the app(s) is expose to the internet
  - perform rolling update


## Steps

### 1. Create namespace
```
apiVersion: v1
kind: Namespace
metadata:
  name: bggns
```

### 2. Create database deployment resource
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: bggdb-deploy
  labels:
    app: bggdb-deploy
  namespace: bggns
spec:
  replicas: 1
  selector:
    matchLabels:
      app: bggdb-deploy-po
  template:
    metadata:
      labels:
        app: bggdb-deploy-po
    spec:
      containers:
      - name: bggdb-deploy-po
        image: stackupiss/bgg-database:v1
        ports:
        - containerPort: 3306
```

### 3. Create database's service resource
```
apiVersion: v1
kind: Service
metadata:
  name: bggdb-svc
  namespace: bggns
spec:
  selector:
    app: bggdb-deploy-po
  ports:
    - protocol: TCP
      port: 3306
      targetPort: 3306
```

### 4. Create database's credential using configmap resource
```
apiVersion: v1
kind: ConfigMap
metadata:
  name: bggapp-cm
  labels:
    name: bggapp-cm
  namespace: bggns
data:
   BGG_DB_USER: root
   BGG_DB_HOST: bggdb-svc
```

### 5. Create database's credential using secret resource
```
apiVersion: v1
kind: Secret
metadata:
  name: bggapp-secret
  labels:
    name: bggapp-secret
  namespace: bggns
type: Opaque
data:
  # echo -n changeit | baset64 -w0 -
  BGG_DB_PASSWORD: Y2hhbmdlaXQ=
```

### 6. Create app deployment resource
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: bggapp-deploy
  labels:
    app: bggapp-deploy
  namespace: bggns
spec:
  replicas: 1
  selector:
    matchLabels:
      app: bggapp-deploy-po
  template:
    metadata:
      labels:
        app: bggapp-deploy-po
    spec:
      containers:
      - name: bggapp-deploy-po
        image: stackupiss/bgg-backend:v1
        ports:
        - containerPort: 3000
        envFrom:
        - configMapRef:
            name: bggapp-cm
        readinessProbe:
          httpGet:
            path: /healthz
            port: 3000
          initialDelaySeconds: 5
          periodSeconds: 10
        livenessProbe:
          httpGet:
            path: /healthz
            port: 3000
          initialDelaySeconds: 15
          periodSeconds: 20
```

### 7. Test app is working fine locally
```
k -n bggns port-forward pod/<POD_NAME> 8080:3000
curl localhost:8080
```

### 8. Create app's service resource
```
apiVersion: v1
kind: Service
metadata:
  name: bggapp-svc
  namespace: bggns
spec:
  selector:
    app: bggapp-deploy-po
  ports:
    - protocol: TCP
      port: 80
      targetPort: 3000
  type: LoadBalancer
```

### 9. Rolling update
```
strategy:
    type: RollingUpdate
    rollingUpdate:
        maxSurge: 1
        maxUnavailable: 0
```
maxSurge, means how many new pod(s) can be created
maxUnavailable, means how many pod(s) can be brought down

### 10. Rollback
```
# check rollout status
k -n bggns rollout status deployment/bggapp-deploy

# check rollout history
k -n bggns rollout history deployment/bggapp-deploy

# execute rollout
k -n bggns rollout undo deployment/bggapp-deploy
```

## Additional

- Notice the value of `BGG_DB_HOST` in the configmap is not written in full form. Because DNS queries may be expanded using the Pod's `/etc/resolv.conf`

```
# content
search bggns.svc.cluster.local svc.cluster.local cluster.local
nameserver x.x.0.10
options ndots:5
```

```
# Final form 
<SVC_NAME>.<NAMESPACE>.svc.cluster-domain.example
bggdb-svc.bggns.svc.cluster.local
```