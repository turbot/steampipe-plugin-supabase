# Table: supabase_project_network_restrictions

Each Supabase project comes with configurable restrictions on the IP ranges that are allowed to connect to Postgres and PgBouncer ("your database"). These restrictions are enforced before traffic reaches your database. This is a useful feature if you want to limit access to your project to a specific set of users or devices. If a connection is not restricted by IP, it still needs to authenticate successfully with valid database credentials.

## Examples

### Basic info

```sql
select
  project_id,
  entitlement,
  status
from
  supabase_project_network_restrictions;
```
