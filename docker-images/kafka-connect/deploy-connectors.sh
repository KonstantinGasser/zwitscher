
curl -X POST http://localhost:8083/connectors -H "Content-Type: application/json" -d '{
        "name": "tweet_data_source",
        "config": {
                "tasks.max": "1",
                "connector.class": "com.github.jcustenborder.kafka.connect.spooldir.SpoolDirCsvSourceConnector",
                "input.path": "/home/data",
                "error.path": "/home/data/error",
                "finished.path": "/home/data/finished",
                "halt.on.error": "false",
                "errors.tolerance": "all",
                "errors.deadletterqueue.topic.name": "tweet_deadletterqueue",
                "errors.deadletterqueue.topic.replication.factor": "1",
                "empty.poll.wait.ms": "3000",
                "csv.first.row.as.header": "true",
                "schema.generation.enabled": "true",
                "csv.null.field.indicator": "EMPTY_SEPARATORS",
                "csv.separator.char": "44",
                "input.file.pattern": ".*?tweets.*?\\.csv",
                "topic": "tweets"
                }
        }'

curl -X POST http://localhost:8083/connectors -H "Content-Type: application/json" -d '{
        "name": "follower_data_source",
        "config": {
                "tasks.max": "1",
                "connector.class": "com.github.jcustenborder.kafka.connect.spooldir.SpoolDirCsvSourceConnector",
                "input.path": "/home/data",
                "error.path": "/home/data/error",
                "finished.path": "/home/data/finished",
                "halt.on.error": "false",
                "errors.tolerance": "all",
                "errors.deadletterqueue.topic.name": "follower_deadletterqueue",
                "errors.deadletterqueue.topic.replication.factor": "1",
                "empty.poll.wait.ms": "3000",
                "csv.first.row.as.header": "true",
                "schema.generation.enabled": "true",
                "csv.null.field.indicator": "EMPTY_SEPARATORS",
                "csv.separator.char": "32",
                "input.file.pattern": ".*?combine.*?\\.txt",
                "topic": "follows"
                }
        }'
