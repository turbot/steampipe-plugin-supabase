---
title: "Steampipe Table: supabase_function - Query Supabase Functions using SQL"
description: "Allows users to query Supabase Functions, specifically to retrieve information about user-defined functions in a Supabase database."
---

# Table: supabase_function - Query Supabase Functions using SQL

Supabase Functions are user-defined operations that extend the capabilities of a Supabase database. These functions can be used to perform operations that would normally require multiple queries and round trips in a single call. They are stored in the database and can be invoked using the SQL language.

## Table Usage Guide

The `supabase_function` table provides insights into user-defined functions within a Supabase database. As a database administrator or developer, you can explore function-specific details through this table, including the function's name, schema, data type, and source code. Use it to uncover information about functions, such as their implementation, dependencies, and usage, to help optimize database operations and performance.

## Examples

### Basic info
Explore the status and version of various functions within your Supabase environment. This can be useful to keep track of function updates and monitor their operational status.

```sql
select
  name,
  slug,
  status,
  version
from
  supabase_function;
```

### List throttled functions
Identify instances where certain functions have been throttled or limited in their operations. This can help in diagnosing performance issues and maintaining optimal functionality.

```sql
select
  name,
  slug,
  status,
  version
from
  supabase_function
where
  status = 'THROTTLED';
```

### List functions with JWT verification enabled
Determine the areas in which JWT verification is enabled for various functions. This is useful for ensuring security measures are in place across your functions.

```sql
select
  name,
  slug,
  status,
  version
from
  supabase_function
where
  verify_jwt;
```

### List functions not updated in last 30 days
Discover the functions that have not been updated in the past 30 days. This is useful to identify areas that may require attention or updates to ensure optimal performance and security.

```sql
select
  name,
  slug,
  status,
  version,
  project_id
from
  supabase_function
where
  updated_at < (current_date - interval '30 days');
```

### Get a specific function
Explore the status and version of a specific function within a given project. This can help in tracking the function's progress and managing updates effectively.

```sql
select
  name,
  slug,
  status,
  version
from
  supabase_function
where
  project_id = 'pljlooizchwsteampipe'
  and slug = 'test-function';
```