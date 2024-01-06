# Service Catalog

## Legend

* ❓ — External question. A question on requirements. Typically it would be directed to product or UX, but also may be any external party.
* ⚠️  — Engineering trade off. A decision made that resulted in being able to deliver this take home in a timely fashion but has known long-term implications.

## Getting Started

The app can be launched with docker-compose or podman-compose on Linux.

### Podman

```sh
sudo podman-compose up --build
```

### Docker

```sh
docker-compose up --build
```

Once the Postgres and app container have started, you can go to `http://localhost:8080` and use the GraphQL Playground or use curl from the command line.

## Design

### GraphQL

I've typically written services that communicate with JSON over HTTP on a TLS transport and I've run gRPC services in production. I've heard a lot of good things about GraphQL so I wanted to give it a shot for this implementation. GraphQL seems to be a more common implementation for public facing APIs and front end services but gRPC seems to be favored for back end service to service communication.

This was a new challenge for me and I ended up spending significantly more time than I anticipated on the take home exercise as a result, but I found this to be learning experience and pretty valuable. It helps me to also "get" why it would be favorable to use GraphQL. I don't think I would have had the exposure otherwise working mostly on infrastructure and platform services.

### Resilience

I'm using `github.com/jackc/pgx/v5` over `github.com/lib/pq` because it has a connection pool component. I've found the connection pool to recover in most failure cases. I also make sure that upon application start up it is able to communicate with the database or it will panic because if a new instance of the application comes online then it shouldn't start until it can communicate with the database.

### Health Checks

Health checks are on the `/healthz` endpoint. If this endpoint returns either a 200 or 503. In the event that the DB cannot be reached this could be used to fail a Kubernetes readiness check and cause the incoming requests to fail fast alleviating load on an already broken system.

### Observability

Prometheus metrics are available on the `/metricsz` endpoint. I publish a count of all requests on each resolver and a summary of resolver, the response code (2 for 200 and 5 for 500) their latency.

### Postgres as Database

I've chosen Postgres because I'm most familiar with it as a database and I've scaled application built on top of it in production. Postgres has limitations having a single primary that accepts writes in standard configurations. My approach is to start with a system that's easily understandable and then build in complexity as needed to accommodate scale. For example, Postgres doesn't handle multi-primary writes, but do you need that initially or will a few read replicas across AZs be sufficient?
I also chose Postgres because I can implement the search requirements would be expected in an application. I've used a GIN index with trigrams on the name field. This will let users find services in the service catalog by name but permit some misspellings or other fuzzy matches. I've also used a standard full text GIN index on the description field.

❓ GIN indexes perform slowly on writes. I'm assuming that there wouldn't be a lot of writes, but it is an open product and UX question on how frequently we'd expect users to do updates, creates, or deletes of services. Versions are not susceptible to this since they are in a different table.

⚠️ The pagination is done currently with Postgres' LIMIT and OFFSET. This works fine until the tables grow and retrievals return tens or hundreds of thousands of rows. OFFSET means to read through those rows, but discard them. I've addressed queries written with LIMIT/OFFSET doing full table scans because of large data sets. I've included a sequence column as a mitigation. The idea is you'd use the sequence number in the where clause and then use LIMIT to return the number of rows per page of rows with a sequence greater than the last sequence returned on the previous page.

## Example Queries

If you use the interface on http://localhost:8080 you can copy and paste the queries directly. If you wish to use curl you can do so by wrapping the graphql in a JSON object with a `query` field.

For example, the following are equivalent:

```graphql
query {
  services {
    name
    description
    versions {
      version
    }
  }
}
```

```shell
cat <<'EOF' |tr -d '\n\t' | curl http://localhost:8080/query -X POST -H 'Content-Type: application/json' --data @-
{"query":"{
  services {
    name
    description
    versions {
      version
    }
  }
}"}
EOF
```

Note that you shouldn't include `query` in the `query` itself.

### List all services (paginated)

```graphql
query {
  services(page: 1) {
    name
    description
    versions {
      version
    }
  }
}
```

### Service Fuzzy Name Example

In this example a search query for "boatification" will match "notification"

```graphql
query {
    searchByName(name: "boatification") {
    name
    description
    versions {
      version
    }
  }
}
```

Output

```json
{
  "data": {
    "searchByName": [
      {
        "name": "Notifications",
        "description": "",
        "versions": [
          {
            "version": "0.0.1025"
          },
          {
            "version": "0.0.1026"
          },
          {
            "version": "0.0.1027"
          },
          {
            "version": "0.0.1028"
          }
        ]
      },
      {
        "name": "Notifications",
        "description": "",
        "versions": [
          {
            "version": "6e1663d78057b60e55407b21d9bfd7b5ebaa475e"
          },
          {
            "version": "272fb59d9c759a31da4fe7e5719f2aea5b17f959"
          },
          {
            "version": "a83358ea7a95792c33ab83c54c4a271c223ed030"
          },
          {
            "version": "65b66010f55d72390a88d3af314168f7602bff8b"
          }
        ]
      }
    ]
  }
}
```

### Service Search Name and Description

In this example search for `ipsum` will match multiple service descriptions.

```graphql
query {
    searchAll(query: "ipsum") {
    name
    description
    versions {
      version
    }
  }
}
```

Output

```json
{
  "data": {
    "searchAll": [
      {
        "name": "Contact Us",
        "description": "Lorem ipsum dolor sit amet, consetetur sadipscing",
        "versions": [
          {
            "version": "jezebel"
          },
          {
            "version": "jovial"
          },
          {
            "version": "jersey"
          },
          {
            "version": "jaunty"
          }
        ]
      },
      {
        "name": "Security",
        "description": "Lorem ipsum dolor",
        "versions": [
          {
            "version": "v3"
          },
          {
            "version": "v4"
          },
          {
            "version": "v2"
          },
          {
            "version": "v1"
          }
        ]
      },
      {
        "name": "Locate Us",
        "description": "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
        "versions": [
          {
            "version": "1.1.1"
          },
          {
            "version": "1.2.0"
          },
          {
            "version": "1.0.0"
          },
          {
            "version": "1.1.0"
          }
        ]
      },
      {
        "name": "Contact Us",
        "description": "Lorem ipsum dolor sit amet, consetetur sadipscing",
        "versions": [
          {
            "version": "157.0.0"
          },
          {
            "version": "158.0.0"
          },
          {
            "version": "155.0.0"
          },
          {
            "version": "156.0.0"
          }
        ]
      },
      {
        "name": "Security",
        "description": "Lorem ipsum dolor",
        "versions": [
          {
            "version": "7.7.7"
          },
          {
            "version": "5.5.5"
          },
          {
            "version": "4.4.4"
          },
          {
            "version": "6.6.6"
          }
        ]
      },
      {
        "name": "FX Rates International",
        "description": "Lorem ipsum dolor",
        "versions": [
          {
            "version": "1.0.0-rc1"
          },
          {
            "version": "1.0.0"
          },
          {
            "version": "1.0.0-beta1"
          },
          {
            "version": "1.0.0-alpha1"
          }
        ]
      },
      {
        "name": "FX Rates International",
        "description": "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor",
        "versions": [
          {
            "version": "5.5.51"
          },
          {
            "version": "5.5.50"
          },
          {
            "version": "5.5.53"
          },
          {
            "version": "5.5.52"
          }
        ]
      },
      {
        "name": "Reporting",
        "description": "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
        "versions": [
          {
            "version": "v1.7.0"
          },
          {
            "version": "v1.8.0"
          },
          {
            "version": "v1.6.0"
          },
          {
            "version": "v1.5.0"
          }
        ]
      }
    ]
  }
}
```
