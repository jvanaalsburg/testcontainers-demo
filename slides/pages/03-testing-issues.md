---
layout: two-cols-header
---

<style>
    .two-cols-header {
        grid-template-rows: auto !important;
    }
</style>

# Testing Issues

"Problems" _Testcontainers_ can help us avoid

::left::

### Anti-Patterns

- Using development database for testing
- Failing to revert changes after modifying data
- Schema changes and migrations

::right::

<v-click>

### _Testcontainers_ Approach

- Dedicated databases for testing
- Database are created before each test
- Configuration is done in test code

</v-click>
