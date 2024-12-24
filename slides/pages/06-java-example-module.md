---
layout: default
---

# Java Example

Using a _Testcontainers_ module

```java {*|1-2|4-6|*}
PostgreSQLContainer postgres = new PostgreSQLContainer("postgres:15");
postgres.start();

var username = postgres.getUsername();
var password = postgres.getPassword();
var jdbcUrl = postgres.getJdbcUrl();

// Perform database operations...

postgres.stop();
```
