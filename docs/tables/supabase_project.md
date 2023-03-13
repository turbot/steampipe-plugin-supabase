# Table: supabase_project

Projects in Supabase are designed to organize and manage applications and data in a scalable and secure way.

Within a project, one can create and manage multiple databases, authentication settings, storage buckets, and API keys. Team members can also be invited to collaborate on the project and manage their permissions.

## Examples

### Basic info

```sql
select
  name,
  region,
  created_at,
  organization_id
from
  supabase_project;
```

### Get the count of projects per region

```sql
select
  region,
  count(id) as project_count
from
  supabase_project
group by
  region;
```

### Get the list of banned IPs

```sql
select
  b.address as ip,
  p.name as project
from
  supabase_project_network_bans as b
  join supabase_project as p on b.project_id = p.id;
```
