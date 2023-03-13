# Table: supabase_organization

A Supabase organization is a container that allows the user to organize and manage multiple Supabase projects, billing, and team members. It is an entity that represents your company or organization on Supabase.

## Examples

### Basic info

```sql
select
  name,
  id
from
  supabase_organization;
```

### Get the count of projects per organization

```sql
select
  o.name as org_name,
  count(p.id) as project_count
from
  supabase_project as p
  join supabase_organization as o on p.organization_id = o.id
group by
  o.name;
```
