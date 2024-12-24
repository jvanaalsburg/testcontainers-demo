# Testcontainers Demo

### Overview

This a simple JSON API that demonstrates using [Testcontainers][testcontainers] in integration tests. The server is written in Go, using the [Echo][echo] framework.

### Installation and Usage

You must have [Docker][docker] installed on your system in order to run the server and the tests.

#### Running the Server Locally

Use the following command to run the development server.

```bash
docker compose up -d api
```

This will create containers for the API and a PostreSQL database. The API should be running on port `1323`. You can verify that the service is running my making a `GET` request to the `/users` endpoint.

```bash
curl localhost:1323/users 
```

You should get the following response.

```json
[
  {
    "id": "00000000-0000-0000-0000-000000000001",
    "first_name": "Harry",
    "last_name": "Potter",
    "email": "hpotter@hogwarts.edu"
  },
  {
    "id": "00000000-0000-0000-0000-000000000002",
    "first_name": "Ron",
    "last_name": "Weasley",
    "email": "rweasley@hogwarts.edu"
  },
  {
    "id": "00000000-0000-0000-0000-000000000003",
    "first_name": "Hermione",
    "last_name": "Granger",
    "email": "hgranger@hogwarts.edu"
  }
]
```

#### Running the Tests

The integration tests can be run with the following command.

```bash
docker run --rm test
```

If you run `watch docker ps` before running the command above, you should be able to see the Testcontainers get created before the tests are run, and then get automatically destroyed once they have finished.

**NOTE:** Since the integration tests all use Testcontainers, you can modify the data stored in the development database without affecting the test results.

### Slides

The slides from the Technical Seminar are located in the `slides/` directory. You can run the following to launch the dev server and view the presentation locally.

```bash
cd slides/
yarn install 
yarn dev
```

Then open http://localhost:3030/ in your browser to view the presentation.

[docker]: https://www.docker.com/
[echo]: https://echo.labstack.com/
[testcontainers]: https://testcontainers.com/

