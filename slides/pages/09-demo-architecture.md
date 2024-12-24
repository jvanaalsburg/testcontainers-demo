---
layout: default
---

<style>
    .slidev-code-wrapper {
        max-height: 400px;
        overflow-y: auto;
    }
</style>

# Demo Architecture

Local development environment

````md magic-move
```yaml {*|5-6|7-8} {maxHeight:'400px'}
# docker-compose.yml
services:
  api:
    build: # ...
    volumes:
      - ./src:/app
    ports:
      - 1323:1323
```

```yaml {10-11|12-15|16-17,20-21|16,18|7-8|*} {maxHeight:'400px'}
# docker-compose.yml
services:
  api:
    build: # ...
    volumes: # ...
    ports: # ...
    depends_on:
      - db

  db:
    image: postgres:16-alpine
    environment:
      POSTGRES_DB: demo
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
    volumes:
      - pg-data:/var/lib/postgresql/data
      - ./data:/docker-entrypoint-initdb.d

volumes:
  pg-data:
```
````
