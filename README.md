# Cosmoscan Fetching API

## Prerequisite:

1. Clickhouse (current deployed version: 19.14.3.3)

    #### Deploy on Kubernetes with Docker file:

    - Create a StatefulSet file: `clickhouse.yaml` with the template below:

    ```
        apiVersion: apps/v1
        kind: StatefulSet
        metadata:
        name: clickhouse
        namespace: default
        spec:
        replicas: 1
        selector:
            matchLabels:
            app: clickhouse
        serviceName: clickhouse
        template:
            metadata:
            labels:
                app: clickhouse
            spec:
            containers:
            - image: yandex/clickhouse-server:19.14.3.3
                imagePullPolicy: IfNotPresent
                name: clickhouse
                volumeMounts:
                - mountPath: /var/lib/clickhouse
                name: clickhouse-pvc
                - mountPath: /val/log/clickhouse-server/
                name: clickhouse-logs
                ports:
                - containerPort: 8123
            restartPolicy: Always
            volumes:
            - name: data
                nfs:
                path: __VOLUME_PATH__
                server: __VOLUME_SERVER__
        volumeClaimTemplates:
        - metadata:
            name: clickhouse-pvc
            spec:
            accessModes: [ "ReadWriteOnce" ]
            storageClassName: gp2-encryption
            resources:
                requests:
                storage: 10Gi
        - metadata:
            name: clickhouse-logs
            spec:
            accessModes: [ "ReadWriteOnce" ]
            storageClassName: gp2-encryption
            resources:
                requests:
                storage: 10Gi
        updateStrategy:
            rollingUpdate:
            partition: 0
            type: RollingUpdate

        ---
        apiVersion: v1
        kind: Service
        metadata:
        namespace: default
        name: clickhouse
        spec:
        selector:
        app: clickhouse
        type: NodePort
        ports:
        - name: endpoint
            protocol: TCP
            port: 8123
            targetPort: 8123
    ```

    - Edit the __VOLUME_PATH__ and the __VOLUME_SERVER__ with your information.
    - Run the command to create clickhouse pod & service:
    ```
        kubectl apply -f clickhouse.yaml
    ```

    #### Or install with other methods on: https://clickhouse.com/docs/en/getting-started/install/

2. MySQL

    #### Download: https://dev.mysql.com/downloads/mysql/

3. Aura node

    #### Check the environment here: https://docs.aura.network/environment

4. Golang

    #### Download: https://go.dev/dl/

## How to run

1. Create a `config.json` file beside `main.go` with the template below and edit with your config:
    ```
        {
        "api": {
            "port": "8080",
            "allowed_hosts": [
            "http://localhost:8000"
            ]
        },
        "mysql": {
            "host": "localhost",
            "port": "3306",
            "db": "cosmoscan",
            "user": "root",
            "password": "secret"
        },
        "clickhouse": {
            "protocol": "http",
            "host": "localhost",
            "port": 8123,
            "user": "default",
            "password": "",
            "database": "cosmoshub3"
        },
        "parser": {
            "node": "http://127.0.0.1:1234",
            "batch": 500,
            "fetchers": 5
        },
        "cmc_key": ""
        }
    ```
2. Run

    ### Native way
    
    Run the command: ```go build && ./cosmoscan-api```

    ### Docker way

    Use `Dockerfile`