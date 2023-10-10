1st design the database, generate codes to talk to the DB in a consistent and reliable way using transactions, understand the DB isolation levels, and how to use it correctly in production. Besides the database,use docker for local development, how to use Git to manage your codes, and how to use GitHub Action to run unit tests automatically.

2nd build a set of RESTful HTTP APIs using Gin - one of the most popular Golang frameworks for building web services. This includes everything from loading app configs, mocking DB for more robust unit tests, handling errors, authenticating users, and securing the APIs with JWT and PASETO access tokens.

3rd build your app with Docker and deploy it to a production Kubernetes cluster on AWS. build a minimal docker image, set up a free-tier AWS account, create a production database, store and retrieve production secrets, create a Kubernetes cluster with EKS, use GitHub Action to automatically build and deploy the image to the EKS cluster, buy a domain name and route the traffics to the service, secure the connection with HTTPS and auto-renew TLS certificate from Let's Encrypt.

4th discuss several advanced backend topics such as managing user sessions, building gRPC APIs, using gRPC gateway to serve both gRPC and HTTP requests at the same time, embedding Swagger documentation as part of the backend service, partially updating a record using optional parameters, and writing structured logger HTTP middlewares and gRPC interceptors.

5th asynchronous processing in Golang using background workers and Redis as its message queue, and how to gracefully shut down the server to protect your processing resources. As this part is still a work in progress, we will keep making and uploading new videos about new topics in the future, such as: sending emails, gracefully shutting down servers, CORS, bulk inserts, etc. So please come back here to check them out from time to time.

How to run this bank?

