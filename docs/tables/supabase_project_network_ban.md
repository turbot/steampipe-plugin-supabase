# Table: supabase_project_network_ban

The network bans feature allows users to ban specific IP addresses or IP address ranges from accessing your Supabase project. This is a useful feature if you want to prevent malicious traffic from accessing your project, or if you need to block traffic from specific countries or regions.

## Examples

### Basic info

```sql
select
  project_id,
  address
from
  supabase_project_network_ban;
```
