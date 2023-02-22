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
