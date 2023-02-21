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

### List functions with verify JWT enabled

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
