# Table: supabase_project_custom_hostname

In Supabase, a custom hostname is a user-defined domain name that you can use to access a web application or static website that you have hosted on the Supabase Hosting service.

## Examples

### Basic info

```sql
select
  custom_hostname,
  status,
  data ->> 'created_at' as created_at,
  project_id
from
  supabase_project_custom_hostname;
```

### List custom hostname that are yet to start

```sql
select
  custom_hostname,
  status,
  project_id
from
  supabase_project_custom_hostname
where
  status = '1_not_started';
```

### List custom hostname created in last 3 days

```sql
select
  custom_hostname,
  status,
  data ->> 'created_at' as created_at,
  project_id
from
  supabase_project_custom_hostname
where
  (data ->> 'created_at')::timestamp > (current_date - interval '3 days');
```
