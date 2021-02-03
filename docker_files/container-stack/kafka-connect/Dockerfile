FROM confluentinc/cp-kafka-connect

COPY ./connectors/jcustenborder-kafka-connect-spooldir-1.0.41 /usr/share/confluent-hub-components

COPY ./connectors/mongodb-kafka-connect-mongodb-1.3.0 /usr/share/confluent-hub-components
COPY ./connectors/neo4j-kafka-connect-neo4j-1.0.9 /usr/share/confluent-hub-components
COPY ./deploy-connectors.sh ./
