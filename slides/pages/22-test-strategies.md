---
layout: two-cols-header
---

# Test Strategies

Initializing and managing _Testcontainers_

::left::

**Test Suite**

- Database state persists between tests.
- Use for read-only tests.
  - `Repo.GetAllUsers`
  - `Repo.GetUser`
- Populate with fixture data.

::right::

**Test Repo**

- Database is created and destroyed each test.
- Use for tests which modify the database.
  - `Repo.CreateUser`
  - `Repo.UpdateUser`
  - `Repo.DeleteUser`
- Can use fixture data or only initialize schema.
- Can be run in parallel.
