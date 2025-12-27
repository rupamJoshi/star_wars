# Star Wars GraphQL Service

A Go-based GraphQL service that fetches Star Wars characters from the SWAPI API, allows saving searched character to MongoDB, and exposes queries and mutations to access them.

---

## Features

* Search characters by name from SWAPI.
* Resolve films and vehicles URLs to actual names.
* Save Searched characters to MongoDB.
* Retrieve all saved characters.
* Configurable via YAML file.
* Unit-tested with mock SWAPI and in-memory repository.

---

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/star-wars-graphql.git
cd star-wars-graphql
```

2. Install dependencies:

```bash
go mod tidy
```

3. Ensure MongoDB is running:



```bash
docker run -d -p 27017:27017 --name mongo mongo:latest
```


# Generate GrapQL Schema

```bash 
go run github.com/99designs/gqlgen generate
```

---

## Configuration

The service uses a YAML configuration file:

**config.yaml**

```yaml
host: localhost
port: 8080
schema: http

db_config:
  host: localhost
  port: 27017
  db_name: starwars
  collection_name: characters

swapi_config:
  swapi_url: https://swapi.info/api/
```



## GraphQL Queries & Mutations

### Query: Get Character

```graphql
query {
  searchCharacter(name: "Luke Skywalker") {
    name
    films
    vehicles
  }
}
```

### Mutation: Save Favorite Character

```graphql
mutation save {
  saveSearchResult {
    id
    name
    SavedAt
  }
}
```

### Query: Get All Saved Characters

```graphql
query getALL {
  getSavedResults {
    name
    SavedAt
    films
  }
}
```



