---
layout: two-cols-header
---

<style>
    .col-left {
        margin-right: 1rem;
    }
</style>

# User Model

Database schema & fixture data

::left::

```sql
-- data/000-init-schema.sql
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL
);
```

::right::

```sql
--- data/001-add-users.sql
INSERT INTO
    users (id, first_name, last_name, email)
VALUES
    (
        '00000000-0000-0000-0000-000000000001',
        'Harry',
        'Potter',
        'hpotter@hogwarts.edu'
    ),
    (
        '00000000-0000-0000-0000-000000000002',
        'Ron',
        'Weasley',
        'rweasley@hogwarts.edu'
    ),
    (
        '00000000-0000-0000-0000-000000000003',
        'Hermione',
        'Granger',
        'hgranger@hogwarts.edu'
    );
```
