## Prerequisite

- [Install JDK](https://tecadmin.net/install-openjdk-java-ubuntu/)

- [Install Apache Kafka](https://kafka.apache.org/quickstart)

## Testing

Run following command from your local repository:

```
cd assignment
go run main.go
```

Open another terminal and use CURL to test API - Publish a message to Kafka topic:

```
curl -X POST localhost:8080/send-message \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": "hello"
}'
```

Open another terminal and run following command to start consuming messages from Kafka topic.  
Make sure your current working directory is still in "assignment" folder.

```
go run consumer/consumer.go
```
