# Mastodon Stream Hub

This project is supposed to be a highly available and scalable real-time data pipeline that accepts multiple producers and serves many consumers nearly real time.

## Initial Goal

The initial goal was to ensure that microblogging with Mastodon posts from producer to consumer are delivered nearly in real-time. Look at the following video for the final result.

![final project](mastodon-stream-hub.gif)

## Architecture Overview

![architecture](architecture.jpg)

### Real-time Data Processing and Broadcasting Architecture

This architecture is designed to ingest, process, and broadcast real-time data to connected clients efficiently. It leverages Kafka for message streaming, gRPC for inter-service communication, and WebSockets for delivering real-time updates to web clients. The system is built with scalability, fault tolerance, and low latency in mind, making it suitable for applications requiring real-time data dissemination, such as live dashboards, chat applications, or streaming analytics.

### Data Pipeline and Architecture Considerations

This project employs a robust and scalable data processing and broadcasting pipeline, meticulously designed to handle real-time data efficiently across various platforms. Below is an overview of the architecture and the rationale behind the technology choices.


#### 1. **Data Ingestion**

- **Chosen Technology**: Apache Kafka.
- **Role**: Kafka serves as the backbone for real-time data ingestion, reliably collecting messages from various data producers and making them available for processing.
- **Alternative**: Cloud-native solutions like Amazon Kinesis, Google Pub/Sub, or Azure Event Hubs could also serve this purpose, offering managed scalability and integration with cloud ecosystems.
- **Rationale**: Apache Kafka was chosen for its high throughput, durability, and scalability. Its open-source nature and widespread adoption make it a versatile choice for a broad range of use cases, from logging and tracking to event sourcing.

#### 2. **Data Processing**

- **Chosen Technology**: gRPC service with a simple data processor.
- **Role**: This component consumes messages from Kafka, applies necessary transformations or enrichments, and prepares them for broadcasting.
- **Alternative**: Advanced stream processing frameworks like Apache Flink or Kafka Streams were considered for their rich feature sets in handling complex event processing patterns and state management.
- **Rationale**: A gRPC service with a straightforward processing approach was selected to prioritize development simplicity and speed. While Apache Flink or Kafka Streams offer powerful processing capabilities, they introduce additional complexity and operational overhead, which may not be necessary for the initial scope of this project.

#### 3. **Data Storage**

- **Chosen Technology**: MongoDB.
- **Role**: MongoDB acts as the persistent storage layer, offering flexibility, scalability, and real-time data access patterns.
- **Alternative**: Other NoSQL databases like Cassandra or managed cloud solutions such as Amazon DynamoDB or Google Firestore could be used for scalable and flexible data storage.
- **Rationale**: MongoDB was chosen for its rich query capabilities, document-oriented storage model, and the change streams feature, which seamlessly integrates with the WebSocket server for real-time data broadcasting. Its developer-friendly ecosystem and operational simplicity make it an excellent fit for rapid development cycles.

#### 4. **Data Broadcasting**

- **Chosen Technology**: WebSocket server.
- **Role**: The WebSocket server establishes persistent connections with clients, enabling the real-time broadcasting of processed messages to both web (React) and mobile (Flutter) applications.
- **Alternative**: Server-Sent Events (SSE) or long-polling could be alternatives for pushing updates to clients, though they may not offer the same level of efficiency and bidirectional communication as WebSockets.
- **Rationale**: WebSockets were chosen for their ability to provide full-duplex communication channels over a single TCP connection, ensuring low latency and real-time updates. This technology is well-supported across modern browsers and mobile platforms, making it ideal for delivering a seamless user experience.

### Conclusion: Trade-offs and Decision Making

Every technology choice in this project represents a trade-off between functionality, complexity, and development velocity. By selecting a combination of Kafka, gRPC, MongoDB, and WebSockets, the architecture aims to balance these factors, providing a scalable, real-time data processing and broadcasting solution that remains manageable and adaptable.

While alternatives like Apache Flink, Kafka Streams, and cloud-native data ingestion services offer powerful capabilities, they were considered overkill for the initial project scope. Such decisions underline the importance of aligning technology choices with project requirements, scalability needs, and team expertise, ensuring that the architecture can evolve as the project grows.

### Architecture Pattern

In designing the system, I adopted a **microservices architecture** approach, emphasizing the creation of small, independently deployable services that communicate with each other using **Remote Procedure Calls (RPC)**, specifically gRPC. This architecture style was chosen to enhance modularity, allowing for each component of the application to be developed, deployed, and scaled independently. Following microservices best practices, the application was decomposed into smaller, reusable services, each responsible for a distinct piece of functionality. This approach facilitates easier maintenance, faster development cycles, and better scalability. However, due to the project's time constraints, not all aspects of the services were refined to perfection, with some areas left for future improvement.

### Folder Structure

A monorepo approach was chosen for managing the project's codebase, housing all microservices and shared libraries within a single repository. This structure promotes code reuse, simplifies dependency management, and streamlines the development process, especially in the context of microservices. Below is an example folder structure representing this approach:

```
├── README.md
├── backend
    ├── Dockerfile
    ├── Makefile
    ├── cmd
    │   ├── apigateway
    │   ├── dataprocessorservice
    │   ├── kafkaconsumer
    │   ├── mastodonstream
    │   └── producerservice
    ├── go.mod
    ├── go.sum
    ├── pkg
    │   ├── api
    │   ├── config
    │   ├── dataprocessor
    │   ├── gcppublisher
    │   ├── kafkapublisher
    │   ├── kafkasubscriber
    │   ├── mastodonclient
    │   ├── pubsub
    │   ├── pubsubservice
    │   ├── storage
    │   ├── util
    │   └── websocket
    └── protos
        ├── dataprocessor.proto
        ├── mastodonstream.proto
        ├── post.proto
        └── pubsub.proto
├── docker-compose.yml
├── docs
│   └── index.md
├── frontend
│   ├── mobile
│   └── web
├── generate_ts_models.sh
├── infrastructure
│   ├── AWS
│   └── GCP
├── mastodon-stream-hub.gif
├── mastodon-stream-hub.mp4
└── scripts
    └── install_protoc_tools.sh         
```

### Why This Structure?

- **Centralized Management**: A monorepo simplifies managing dependencies and shared libraries, making it easier to coordinate changes across multiple services.
- **Consistent Tooling**: Developers can use the same set of tools and configurations across the entire project, reducing the learning curve and setup time.
- **Improved Collaboration**: Having a single repository encourages more collaboration and code sharing among teams, as all project components are easily accessible.
- **Simplified CI/CD**: Continuous integration and deployment processes can be more easily managed with a monorepo, as a single pipeline can handle building, testing, and deploying all services.

Choosing a monorepo for microservices in this project acknowledges the trade-offs between simplicity and complexity. It offers workflow conducive to the microservices architecture pattern, balancing the benefits of service isolation with the ease of a unified development ecosystem.

### Alternative Architecture: Fully Serverless with Managed Services

A fully serverless architecture using cloud-managed services offers a different approach, emphasizing operational simplicity and scalability without managing infrastructure.

- **Data Ingestion**: Utilizing managed services like Amazon Kinesis, Google Pub/Sub, or Azure Event Hubs.
- **Data Processing**: Leveraging cloud functions (AWS Lambda, Google Cloud Functions, Azure Functions) for processing data.
- **Data Broadcasting**: Using managed WebSocket services or building a serverless API with real-time capabilities through services like Amazon API Gateway WebSocket API.


### Architecture Characteristics Comparison

| Characteristic      | Kafka-Based Solution | Fully Serverless Solution |
|---------------------|----------------------|---------------------------|
| Scalability         | High                 | High                      |
| Fault Tolerance     | High                 | High                      |
| Latency             | Low                  | Low to Medium             |
| Operational Overhead| Medium               | Low                       |
| Cost                | Variable             | Pay-as-you-go             |
| Flexibility         | High                 | Medium                    |
| Maintenance         | Requires Management  | Low (Managed by Cloud Provider) |


## Technologies and how to run the project

This project leverages a diverse stack of technologies, each chosen for its strengths in building scalable, efficient, and cross-platform applications. Below is an overview of each technology, including setup guides and best practices.

### Docker & Docker Compose

**Setup**:

1. **Install Docker**: Follow the [official Docker installation guide](https://docs.docker.com/get-docker/) for your operating system.
2. **Install Docker Compose**: Docker Desktop for Windows and Mac includes Docker Compose. For Linux, follow the [Compose installation instructions](https://docs.docker.com/compose/install/).

**Usage**:

- Define services, networks, and volumes in a `docker-compose.yml` file.
- Use `docker-compose up` to start your services and `docker-compose down` to stop them.

[ ] the microservices are still not implemented or let's say it's half implemented. 

**PLEASE: use Docker compose to run external services but then run microservices in the monorepo to test the application to do that follow the steps: 

1- touch `.env` in the root of the application it contains obviously this is for development and you can create as many env as you like. 

```bash
MASTODON_SERVER="https://mastodon.social"
MASTODON_CLIENT_ID="YOU_MAS_ID"
MASTODON_CLIENT_SECRET="YOUR-MAS-SECRET"
MASTODON_ACCESS_TOKEN="YOUR-MAS-ACCESS"
PUBLISHER_TYPE="KAFKA"
KAFKA_BROKERS="localhost:9092"
PUBSUB_TOPIC="mastodon-posts"
GCP_PROJECT_ID="mastodon-stream-hub"
SERVICE_PORT="50051"
GRPC_SERVICE_ADDR="0.0.0.0:50051"
MONGODB_URI="mongodb://localhost:27017/?replicaSet=rs0"
MONGODB_DATABASE="mastodon"
MONGODB_COLLECTION="posts"
MONGO_INITDB_ROOT_USERNAME="root"
MONGO_INITDB_ROOT_PASSWORD="mysecretpassword"
API_GATEWAY_SERVER_PORT=":8080"
REACT_APP_WS_ENDPOINT=ws://localhost:8080/ws
```

2- `docker-compose up`, make sure all services are running
[ ] essentially all other steps below should be removed and only docker-compose be yet didn't have time to finish. 

3-  go to `backend` and run `go mod download`

3- then for development you can run services by doing like (go 1.21 is expected)

- `go run cmd/producerservice/main.go`
- `go run cmd/mastodonstream/main.go`
- `go run cmd/dataprocessorservice/main.go`
- `go run cmd/apigateway/main.go`

**At the moment for this project only run with this order:**

1. `go run cmd/producerservice/main.go`
2. `go run cmd/mastodonstream/main.go`
3. `go run cmd/apigateway/main.go`

4- finally you can run the web app
- go to `frontend/web` then run `npm run dev` 

or you can use `go build` to build each microservice and be ready to deploy.

### gRPC and Protobuf

**Setup**:

1. **Install Protocol Buffers Compiler (protoc)**: Download the appropriate version for your OS from the [GitHub releases page](https://github.com/protocolbuffers/protobuf/releases).
2. **Install gRPC**: Follow the [gRPC Go Quickstart](https://grpc.io/docs/languages/go/quickstart/) to set up gRPC in your Go environment.


- Define your service interfaces and message types in `.proto` files.
- Use `protoc` with the Go plugins to generate Go code for your services.
- I also had a plan to generate models for TS and Dart for both the web and the Flutter app.


### Makefile

**Setup**: Create a `Makefile` in your project root.

**Usage**:

- Define commands as targets in the `Makefile`.
- Use `make <target>` to execute commands.
- use `make gen-proto` to generate Go lang from protobuf and also Typescript for client 

[ ] typescript generation is half done under `generate_ts_models.sh` file. 

### GoLang

**Setup**:

1. **Install Go**: Follow the [official Go installation instructions](https://golang.org/doc/install).
2. **Set up your Workspace**: Follow Go's workspace directory structure and set the `GOPATH` environment variable.

Use `go build` to compile your programs, `go run` to execute them, and `go test` for testing.

[ ] Testing still left while I think test is part of the code, due to timeframe I wanted to deliver more functions even though it could be buggy and less tested. 

### React for Web

**Setup**:

1. **Install Node.js**: Download from the [official Node.js website](https://nodejs.org/). I myself use `bun` it's much faster. 
2. `npm run dev` in the root of the web project under `frontend/web` and a `vite` server with `react` and `tailwind` setup will run. 

### Flutter for Mobile

**Setup**:

1. **Install Flutter**: Follow the [Flutter installation guide](https://flutter.dev/docs/get-started/install).
2. **Set up an Editor**: Install the Flutter and Dart plugins for your preferred IDE.

[ ]I haven't yet finish the implementation but it's doable to implement gRPC in Flutter too.


### Cloud Infrastructure

1. **Terraform**: Navigate to the `infrastructure` directory and customize the `aws/main.tf` or `gcp/main.tf`

Make sure you have all credentials needed already added to the project.

2. Initialize Terraform.

   ```bash
   terraform init
   ```

3. Apply Terraform configuration (this will provision your cloud resources).

   ```bash
   terraform apply
   ```
[ ] I couldn't finish this but it was on my plan to deploy to both cloud

### Scalability of the Project

The architecture and technology choices in this project lay a solid foundation for scalability, ensuring that the system can adapt to increasing workloads by scaling resources up or out as necessary. Here’s how each component contributes to the project’s scalability and what measures can be taken to further enhance it:

#### Kafka for Scalable Data Ingestion

- **Current Setup**: Apache Kafka, which serves as the backbone for data ingestion, inherently supports high scalability through its distributed nature. Topics can be partitioned and replicated across a cluster of brokers to handle high volumes of data and maintain fault tolerance.
- **Scaling Strategy**: As data volume grows, Kafka clusters can be scaled out by adding more brokers. Partitions can also be increased for topics to distribute the load more effectively across the cluster. This ensures that the ingestion layer can handle larger streams of data without becoming a bottleneck.

#### Microservices with gRPC

- **Current Setup**: The system uses a microservices architecture, where services communicate using gRPC. This setup not only breaks down the application into smaller, manageable pieces but also facilitates independent scaling of services based on demand.
- **Scaling Strategy**: Individual microservices can be scaled horizontally by deploying additional instances. Using container orchestration platforms like Kubernetes simplifies this process by automatically managing service instances based on traffic patterns, ensuring that processing capabilities scale efficiently with demand.

#### MongoDB for Scalable Storage

- **Current Setup**: MongoDB is used for data storage, known for its scalability and flexibility. It supports sharding to distribute data across multiple servers, enabling horizontal scaling.
- **Scaling Strategy**: Implementing sharding in MongoDB allows the database layer to accommodate growing data sizes and query volumes. Proper indexing strategies and shard key selection are crucial for optimizing query performance and ensuring even data distribution.

#### WebSockets for Real-Time Communication

- **Current Setup**: WebSockets provide a full-duplex communication channel between the server and clients, facilitating real-time data broadcasting.
- **Scaling Strategy**: The WebSocket server can be scaled horizontally by deploying multiple instances behind a load balancer that supports WebSocket traffic. This can be efficiently managed in a Kubernetes environment, which also enables SSL termination at the load balancer for secure connections.

#### Kubernetes for Orchestration and Autoscaling

- **Proposed Enhancement**: Deploying the entire stack on Kubernetes offers several advantages for scalability:
  - **Pod Autoscaling**: Automatically adjusts the number of pod replicas in a Deployment or StatefulSet based on CPU usage or custom metrics.
  - **Node Autoscaling**: Automatically adjusts the number of nodes in the cluster based on the demands of the workloads and the available resources.
  - **Service Discovery and Load Balancing**: Efficiently distributes traffic among service instances and integrates with cloud providers' load balancers for external traffic.


## TODOs

## Project Todos

- [ ] **Code Improvements**
  - [ ] Refactor existing code for clarity and efficiency.
  - [ ] Add comments and documentation within the codebase.

- [ ] **Finalize Docker Setup**
  - [ ] Create optimized Dockerfiles for each service.
  - [ ] Ensure Docker Compose setup for local development is complete and functional.

- [ ] **Terraform Infrastructure**
  - [ ] Define cloud resources for the project in Terraform.
  - [ ] Test Terraform scripts for deploying to GCP/AWS.

- [ ] **Integrate Additional Modules**
  - [ ] Evaluate and integrate Apache Flink for stream processing.
  - [ ] Set up GCP Pub/Sub for an alternative to Kafka in cloud environments.

- [ ] **Simplify Build and Deployment**
  - [ ] Combine build and deployment scripts into a Makefile for simplicity.
  - [ ] Ensure Makefile commands cover all necessary actions (build, test, deploy, etc.).

- [ ] **Protobuf for Web and Mobile**
  - [ ] Generate and test Protobuf files for web (React) use.
  - [ ] Generate and test Protobuf files for mobile (Flutter) use.

- [ ] **Scalability Checks**
  - [ ] Review and optimize each service for scalability.
  - [ ] Plan for horizontal scaling in cloud deployment.

- [ ] **Service Builds and Deployment**
  - [ ] Automate the build process for each microservice.
  - [ ] Automate the deployment process, including CI/CD pipelines.

- [ ] **Additional Improvements**
  - [ ] List any other specific tasks or improvements needed for the project.
- [ ] **Error handling**
  - [ ] Ensure all errors are caught and also have a proper response 
  - [ ] I also need to make sure that in case of errors in major services what should happen.

## Contribution 

Feel free to open a PR and make this project even better. 