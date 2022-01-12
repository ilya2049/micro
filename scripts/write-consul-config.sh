curl --request PUT \
"http://127.0.0.1:8500/v1/kv/config/hasherapi" \
--data '{
  "logger": {
    "level": "debug",
    "graylog": {
      "host": "graylog:12201",
      "source": "hasherapi"
    }
  },
  "hasher": {
    "host": "hasher:8090",
    "timeoutSec": 1
  },
  "redis": {
    "host": "redis:6379",
    "password": "123456789"
  },
  "kafka": {
    "host": "kafka:9092",
    "topic": "http_calls"
  }
}'

curl --request PUT \
"http://127.0.0.1:8500/v1/kv/config/hasher" \
--data '{
  "logger": {
    "level": "debug",
    "graylog": {
      "host": "graylog:12201",
      "source": "hasher"
    }
  },
  "grpc": {
    "host": ":8090"
  }
}'