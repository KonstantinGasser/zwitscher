FROM confluentinc/cp-kafka-connect

COPY ./connectors/jcustenborder-kafka-connect-spooldir-1.0.41 /usr/share/confluent-hub-components

COPY ./connectors/mongodb-kafka-connect-mongodb-1.3.0 /usr/share/confluent-hub-components
COPY ./deploy-connectors.sh ./
