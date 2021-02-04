package main

import (
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {

	producerMSG := make(chan []byte)

	server := APIserver{
		memcache: &memClient{
			client: memcache.New("localhost:11211"),
		},
		neo4j: &neoClient{
			driver: newDriver(),
		},
		kafka: &kafkaClient{
			producer: newProducer(),
			msgChan:  producerMSG,
		},
		mongoC: &mongoClient{
			driver: newMongoClient(),
		},
	}

	go server.kafka.produce(producerMSG)

	server.setUp()

	if err := server.run(); err != nil {
		log.Fatal(err.Error())
	}
}
