# Table: supabase_function

A Supabase function allows users to write and deploy serverless functions that execute on the edge, which means the functions run closer to the client or user, rather than on a central server.

Edge Functions are lightweight and can be used for a variety of purposes, such as serving cached content, transforming or modifying responses, adding security headers, and more. They can be written in JavaScript, TypeScript, or any language that can be transpiled to JavaScript, and can be deployed using Supabase's CLI or API.

## Examples

### Basic info

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
