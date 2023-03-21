# Table: supabase_project

Projects in Supabase are designed to organize and manage applications and data in a scalable and secure way.

Within a project, one can create and manage multiple databases, authentication settings, storage buckets, and API keys. Team members can also be invited to collaborate on the project and manage their permissions.

## Examples

### Basic info

```sql
select
  name,
  id,
  region,
  created_at,
  organization_id
from
  supabase_project;
```

### Get project database information

```sql
select
  name,
  id,
  region,
  created_at,
  database ->> 'host' as database_host,
  database ->> 'version' as database_version,
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

### Get project API settings

```sql
select
  name,
  id,
  region,
  api_settings ->> 'db_schema' as exposed_schemas,
  api_settings ->> 'db_extra_search_path' as extra_search_path,
  api_settings ->> 'max_rows' as max_rows,
  organization_id
from
  supabase_project;
```

### Get the list of allowed CIDRs

```sql
select
  p.name,
  p.id,
  p.database ->> 'host' as database_host,
  p.database ->> 'version' as database_version,
  r.config -> 'dbAllowedCidrs' as allowed_cidrs
from
  supabase_project as p
  left join supabase_project_network_restriction as r on p.id = r.project_id;
```
