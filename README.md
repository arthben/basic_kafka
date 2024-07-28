# Basic Kafka Publisher & Consumer

## Setting network for kafka 
```console
# create network on docker
docker network create --subnet=192.168.44.0/24 my_kafka_net

# make sure the network created 
docker network ls | grep my_kafka_net 
```

## Run Kafka multiple node with Kafka-UI 
```console
docker-compose up
```

## Add port binding to kafka 
```console

# Add to /etc/host in local laptop
echo "127.0.0.1 kafka-1 kafka-2 kafka-3" >> /etc/hosts
```

## Access Kafka-UI 
Open http://127.0.0.1:8080
