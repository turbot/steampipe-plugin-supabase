# Table: supabase_bucket

A storage bucket is a scalable and secure object storage service that allows you to store and serve files, images, and other assets for your application. It's similar to services like Amazon S3 or Google Cloud Storage.

Storage buckets in Supabase are designed to be highly scalable and reliable, with built-in redundancy and automatic failover. They also integrate seamlessly with other Supabase services, such as the authentication and API services, to provide a complete backend solution for your application.

## Examples

### Basic info

```sql
select
  name,
  public,
  created_at,
  project_id
from
  supabase_bucket;
```

### List all public buckets

```sql
select
  name,
  public,
  created_at,
  project_id
from
  supabase_bucket
where
  public;
```

### Get the count of buckets per project

```sql
select
  p.name as project,
  count(b.id) as bucket_count
from
  supabase_bucket as b
  join supabase_project as p on b.project_id = p.id
group by
  p.name;
```
