---
layout: default
---

# Java Example

Using a generic Docker container

```java {*|1|2|3-6|7|9-11|13-15|*}
GenericContainer container = new GenericContainer("postgres:15")
    .withExposedPorts(5432)
    .waitingFor(new LogMessageWaitStrategy()
        .withRegEx(".*database system is ready to accept connections.*\s")
        .withTimes(2)
        .withStartupTimeout(Duration.of(60, ChronoUnit.SECONDS)));
container.start();

var username = "test";
var password = "test";
var jdbcUrl = "jdbc:postgresql://" + container.getHost() + ":" + container.getMappedPort(5432) + "/test";

// Perform database operations...

container.stop();
```
