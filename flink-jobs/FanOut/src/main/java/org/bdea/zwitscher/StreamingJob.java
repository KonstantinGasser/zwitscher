/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.bdea.zwitscher;

import org.apache.flink.api.common.functions.MapFunction;
import org.apache.flink.api.common.serialization.SimpleStringSchema;
import org.apache.flink.api.java.utils.ParameterTool;
import org.apache.flink.shaded.jackson2.com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.flink.shaded.jackson2.com.fasterxml.jackson.databind.node.ObjectNode;
import org.apache.flink.streaming.api.datastream.DataStream;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaConsumer;
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaProducer;
import org.apache.flink.streaming.connectors.kafka.internals.KeyedSerializationSchemaWrapper;

import java.util.*;

public class StreamingJob {

	public static void main(String[] args) throws Exception {

		ParameterTool parameterTool = ParameterTool.fromArgs(args);

		final String jobName = parameterTool.get("job-name", "fan_out_job");
		final String inputTopic = parameterTool.get("input-topic", "clean_tweets");
		final String consumerGroup = parameterTool.get("group-id", "zwitscher");
		final String kafkaAddress = parameterTool.get("kafka-address", "kafka:29092");

		final ObjectMapper objectMapper = new ObjectMapper();

		// set up the streaming execution environment
		final StreamExecutionEnvironment env = StreamExecutionEnvironment.getExecutionEnvironment();

		//create a new kafka consumer
		Properties consumerProps = new Properties();
		consumerProps.setProperty("bootstrap.servers", kafkaAddress);
		consumerProps.setProperty("group.id", consumerGroup);
		FlinkKafkaConsumer<String> flinkKafkaConsumer = new FlinkKafkaConsumer<>(inputTopic,
				new SimpleStringSchema(), consumerProps);

		//add the consumer to the environment as a data-source, to get a DataStream
		DataStream<String> dataStream = env.addSource(flinkKafkaConsumer);

		// parse json and assign a flink event-time timestamp to each element
		DataStream<ObjectNode> jsonStream = dataStream
				.map((MapFunction<String, ObjectNode>) value -> (ObjectNode) objectMapper.readTree(value));

		jsonStream.map(value -> {
			String id = value.get("user_id").textValue();
			System.out.println(id);
			List<String> followers = Neo.getFollowers(id);
			for (String f : followers) {
				Memcached.updateTimeline(f, value.toString());
			}
			return value;
		});

		// serialize
//		DataStream<String> outputStream = filteredStream.map(value -> objectMapper.writeValueAsString(value));

//		//create a new kafka producer
//		Properties producerProps = new Properties();
//		producerProps.setProperty("bootstrap.servers", kafkaAddress);
//		FlinkKafkaProducer<String> flinkKafkaProducer = new FlinkKafkaProducer<>(outputTopic,
//				new KeyedSerializationSchemaWrapper<>(new SimpleStringSchema()),
//				producerProps, FlinkKafkaProducer.Semantic.AT_LEAST_ONCE);
//
//		//add the producer to the dataStream as a sink
//		outputStream.addSink(flinkKafkaProducer);

		// execute program
		env.execute(jobName);
	}
}
