# Table: supabase_project_network_restriction

Each Supabase project comes with configurable restrictions on the IP ranges that are allowed to connect to Postgres and PgBouncer ("your database"). These restrictions are enforced before traffic reaches your database. This is a useful feature if you want to limit access to your project to a specific set of users or devices. If a connection is not restricted by IP, it still needs to authenticate successfully with valid database credentials.

## Examples

### Basic info

```sql
select
  project_id,
  entitlement,
  status
from
  supabase_project_network_restriction;
```

### List projects with no access to network restrictions

```sql
select
  project_id,
  entitlement,
  status
from
  supabase_project_network_restriction
where
  entitlement = 'disallowed';
```

### List projects where network restriction configuration is not applied

```sql
select
  p.name as project,
  r.status
from
  supabase_project_network_restriction as r
  join supabase_project as p on r.project_id = p.id
where
  r.status != 'applied';
```

### Get the list of allowed CIDRs

```sql
select
  ip as allowed_cidr,
  project_id,
  status
from
  supabase_project_network_restriction,
  jsonb_array_elements_text(config -> 'dbAllowedCidrs') as ip;
```
