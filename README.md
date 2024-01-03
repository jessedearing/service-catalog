# Service Catalog

## Getting Started

+ TODO: Add instructions for how to run this <2024-01-02, Jesse Dearing>

## Design

- **Postgres for storage**  
  I've chosen Postgres because I'm most familiar with it as a database and I've scaled application built on top of it in production. Postgres has limitations having a single primary that accepts writes in standard configurations. My approach is to start with a system that's easily understandable and then build in complexity as needed to accommodate scale. For example, Postgres doesn't handle multi-primary writes, but do you need that initially or will a few read replicas across AZs be sufficient?
- **GraphQL for API**  
  I've typically written services that communicate with JSON over HTTP on a TLS transport and I've run gRPC services in production. I've heard a lot of good things about GraphQL so I wanted to give it a shot for this implementation. GraphQL seems to be a more common implementation for public facing APIs and front end services but gRPC seems to be favored for back end service to service communication.
