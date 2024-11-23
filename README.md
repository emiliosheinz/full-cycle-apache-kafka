# Apache Kafka 

Apache Kafka is an open-source distributed event streaming platform used by thousands of companies to build high-performance data pipelines, stream analytics, and enable crucial data integration.

Each day events are becoming more and more important in the world of software development. The ability to react to events in real-time is crucial for many applications. Events are present in the comunication between simple systems, IOT devices, in the monitoring of complex systems, and many other scenarios.

## Kafka superpowers

- **High throughput**: It can handle a large number of messages per second, which means it scales pretty well;
- **Very low latency (2ms)**: It can process messages very quickly, which is crucial for real-time applications;
- **Storage**: It can store large amounts of data for a long time, which is crucial for data analytics, for example;
- **High availability**: It can be configured to be fault-tolerant, which means it can be used in critical applications;
- **It's available everywhere**: It can be deployed on-premises, in the cloud, and even in Kubernetes. It also connects to many languages and frameworks like Java, Python, Go, and many others.
- **It's open-source**: It's free to use and has a large community behind it.

## The foundation

Kafka is built on top of a few key concepts:

- **Producers**: They are the ones that produce messages to Kafka topics;
- **Consumers**: They are the ones that consume messages from Kafka topics;
- **Topics**: They are the channels where messages are sent to;
- **Partitions**: They are the way Kafka scales topics. Each topic can be divided into partitions, which can be distributed across different brokers;
- **Brokers**: They are the servers that store the messages;
- **ZooKeeper**: It's a service that manages the brokers and the topics;

Below is a diagram that shows how these concepts are connected:

![Kafka architecture](./docs/images/kafka-architecture.png)

### Topics

Topics are the channels where messages are sent to and consumed from. When messages are sent to a topic, they are stored in partitions. Each partition is an ordered, immutable sequence of messages that is continually appended to—a commit log. This allows multiple consumers to read from a topic in parallel, for example.

![topics](./docs/images/topics.png)

#### Anatomy of a record

A record is the basic unit of data in Kafka. It consists of **headers**, a **key**, a **value**, and a **timestamp**. The key and the value are both byte arrays, and the timestamp is a long number. The key is optional, and the value can be anything you want.

### Partitions

Partitions are the way Kafka scales topics. Each topic can be divided into partitions, which can be distributed across different brokers. This allows Kafka to handle a large number of messages per second and store large amounts of data for a long period of time.

#### Ensuring delivery order

Messages are appended to partitions in the order they are sent. This means that messages sent to different partitions can be consumed in a different order than they were sent. If you need to ensure that messages are consumed in the order they were sent, you can use a single partition for a topic which can be done by using the same key for all messages.

#### Distributed partitions

Partitions can be distributed across different brokers. This allows Kafka to scale topics horizontally. Each partition is replicated across multiple brokers to ensure fault tolerance. This means that if a broker goes down, another broker can take over.
 
The number of replicas for a given partition is configurable through the `replication.factor` configuration. The default value is 1, but it's recommended to set it to at least 2 to ensure fault tolerance and 3 or more for mission-critical applications that require high availability.

#### Partition leadership

Each partition has a leader and one or more followers. The leader is responsible for handling read and write requests for the partition and the followers simply replicate the leader's data. If, for any reason, the leader goes down, one of the followers becomes the new leader.

### Delivery guarantees

Kafka provides the following delivery guarantees:

- **At most once**: Messages can be lost but are never redelivered. If the consumer chooses the latest messages, it means it will only see the new messages arriving since it registered. The consumer will miss any messages produced while it is unavailable. This could be desirable when the consumer is interested in the current state and losing old messages is acceptable.

- **At least once**: Messages are never lost but can be redelivered. If losing messages is not acceptable, it can choose to read the earliest messages and in this case, the broker returns the messages since its last committed offset.

- **Exactly once**: Messages are never lost and are never redelivered. This is the most complex guarantee and requires coordination between producers and consumers. It's not always possible to achieve this guarantee, but Kafka provides the tools to make it possible.

### Idempotent producers

Kafka provides a feature called idempotent producers that allows you to produce messages with exactly once semantics. This means that even if a producer sends the same message multiple times, it will only be written once to the partition. This is achieved by assigning a sequence number to each message and deduplicating messages with the same sequence number.

### Consumer groups

Consumers can be grouped together to consume messages from a topic in parallel. Each consumer in the group is assigned a subset of the partitions in the topic. This allows Kafka to scale the number of consumers and handle a large number of messages per second.

Only one consumer in the group can consume messages from a partition at a time. This means that if you have more consumers in the same group than partitions, some consumers will be idle. If you have more partitions than consumers, some consumers will consume messages from multiple partitions. In other words, only consumers in different groups can consume messages from the same partition at the same time.

### Some usefull commands
```
kafka-topics --create --topic=teste --partitions=3 --bootstrap-server=localhost:9092
kafka-topics --list --bootstrap-server=localhost:9092
kafka-topics --topic=teste --describe --bootstrap-server=localhost:9092
kafka-topics --create --bootstrap-server=localhost:9092 --topic=teste --partitions=3

kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning
```
## Kafka Connect

Kafka Connect is a free and open-source component of Apache Kafka that works as a centralized data hub for simple integrations between databases, key-value stores, search indexes and file systems.

### Standalone Workers

Standalone workers in Kafka Connect are used for simpler, single-instance deployments. All configurations, including connectors and offsets, are stored locally on the worker's filesystem. This mode is suitable for development, testing, or small-scale production setups where high availability or scalability is not a concern. As only one instance is running, it cannot share workloads or recover from failures automatically.

### Distributed Workers

Distributed workers operate in a cluster, enabling fault tolerance and scalability. Configurations are stored in Kafka topics, allowing any worker in the cluster to manage tasks. This mode supports rebalancing, where tasks are redistributed automatically if a worker joins or leaves the cluster, making it ideal for production environments with high availability requirements.

### Converters

Tasks utilize converters to translate data between Kafka Connect and the source or sink systems. In other words converters are responsible for serializing and deserializing data, ensuring compatibility between the data formats. Kafka Connect includes built-in converters for common data formats, such as JSON, Avro, and Protobuf, and supports custom converters for other formats.

### DLQ (Dead Letter Queue)

Kafka Connect provides a Dead Letter Queue (DLQ) feature to handle failed records. When a record fails to be processed, it is sent to a separate topic, allowing you to analyze and troubleshoot the issue. This feature is particularly useful for debugging, monitoring, and ensuring data integrity in your data pipelines.

**errors.tolerance config:**

- **none**: The connector will fail immediately if any record fails to be processed.
- **all**: The connector will continue processing records even if some fail, sending failed records to the DLQ.
- **all-or-none**: The connector will fail if all records fail to be processed, but continue processing if only some records fail.

