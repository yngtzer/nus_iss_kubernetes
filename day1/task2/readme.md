# Task 2 : Bring up Northwind Database

## Create network nw-net
`docker network create nw-net`

## Create local volume
`docker volume create --name nwdb-vol`

## Pull northwind db image
`docker pull stackupiss/northwind-db:v1`

## Run northwind db
`docker run --network nw-net --name nwdb -v nwdb-vol:/var/lib/mysql -d stackupiss/northwind-db:v1`

# Run web app in the same network
```
docker run -d --network nw-net -p 8080:3000 -e DB_HOST=nwdb -e DB_USER=root -e DB_PASSWORD=changeit stackupiss/northwind-app:v1
```