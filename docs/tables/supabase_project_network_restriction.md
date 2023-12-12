---
title: "Steampipe Table: supabase_project_network_restriction - Query Supabase Project Network Restrictions using SQL"
description: "Allows users to query Project Network Restrictions in Supabase, providing insights into network restriction settings for each project."
---

# Table: supabase_project_network_restriction - Query Supabase Project Network Restrictions using SQL

Supabase Project Network Restrictions is a feature within Supabase that allows you to manage and restrict network access to your projects. It provides a way to set up and control network restrictions for different Supabase resources, including databases, web applications, and more. Supabase Project Network Restrictions helps you maintain the security and performance of your Supabase resources by ensuring only authorized networks can access your projects.

## Table Usage Guide

The `supabase_project_network_restriction` table provides insights into network restriction settings within Supabase. As a DevOps engineer, explore project-specific network restriction details through this table, including allowed networks, restricted networks, and associated metadata. Utilize it to uncover information about network restrictions, such as those with specific IP ranges, the allowed networks for each project, and the verification of network restriction policies.

## Examples

### Basic info
Explore which projects have network restrictions by assessing their entitlement status. This can help in understanding the level of access control applied to the projects.

```sql+postgres
select
  project_id,
  entitlement,
  status
from
  supabase_project_network_restriction;
```

```sql+sqlite
select
  project_id,
  entitlement,
  status
from
  supabase_project_network_restriction;
```

### List projects with no access to network restrictions
Explore which projects have been denied network access, providing valuable insights into potential security measures or restrictions in place. This could be particularly useful for assessing compliance with internal policies or identifying areas for improvement in network security.

```sql+postgres
select
  project_id,
  entitlement,
  status
from
  supabase_project_network_restriction
where
  entitlement = 'disallowed';
```

```sql+sqlite
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
Analyze the settings to understand which projects have not applied network restriction configurations, helping to identify potential security vulnerabilities.

```sql+postgres
select
  p.name as project,
  r.status
from
  supabase_project_network_restriction as r
  join supabase_project as p on r.project_id = p.id
where
  r.status != 'applied';
```

```sql+sqlite
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
Uncover the details of permitted network addresses within your project, helping you maintain security by understanding which IP ranges have access. This can be particularly useful in identifying any unusual or unexpected network permissions that could potentially compromise your project's security.

```sql+postgres
select
  ip as allowed_cidr,
  project_id,
  status
from
  supabase_project_network_restriction,
  jsonb_array_elements_text(config -> 'dbAllowedCidrs') as ip;
```

```sql+sqlite
select
  ip.value as allowed_cidr,
  project_id,
  status
from
  supabase_project_network_restriction,
  json_each(supabase_project_network_restriction.config, '$.dbAllowedCidrs') as ip;
```