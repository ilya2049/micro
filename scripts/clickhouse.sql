CREATE TABLE http_calls (
    time DateTime Codec(DoubleDelta, LZ4),
    date ALIAS toDate(time),
    "/send" Int64 Codec(T64, LZ4),
    "/check" Int64 Codec(T64, LZ4)
) Engine = MergeTree
PARTITION BY toYYYYMM(time)
ORDER BY (time);

CREATE TABLE http_calls_queue (
    time DateTime,
    "/send" Int64 Codec(T64, LZ4),
    "/check" Int64 Codec(T64, LZ4)
)
ENGINE = Kafka
SETTINGS kafka_broker_list = 'kafka:9092',
       kafka_topic_list = 'http_calls',
       kafka_group_name = 'http_calls_consumer_group1',
       kafka_format = 'CSV',
       kafka_max_block_size = 1048576;


CREATE MATERIALIZED VIEW http_calls_queue_mv TO http_calls AS
SELECT time, "/send", "/check"
FROM http_calls_queue;