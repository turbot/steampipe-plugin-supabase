---
title: "Steampipe Table: supabase_project - Query Supabase Projects using SQL"
description: "Allows users to query Supabase Projects, providing a comprehensive overview of project details and configurations."
---

# Table: supabase_project - Query Supabase Projects using SQL

Supabase Projects is a feature of Supabase, a real-time and open-source Firebase alternative. It allows developers to manage and control their project settings, configurations, and details in one place. It is an essential tool for developers looking to maintain and organize their projects efficiently.

## Table Usage Guide

The `supabase_project` table offers extensive insights into Supabase projects. As a developer, you can use this table to gain a comprehensive understanding of your project settings, configurations, and details. This can be particularly beneficial for managing multiple projects, ensuring configurations are correct, and identifying any potential inconsistencies or issues.

## Examples

### Basic info
Explore which projects were created in different regions and when, in order to understand the geographical distribution and timeline of your organization's projects.

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
Discover the details of your project's database, such as its host and version, by analyzing the associated information. This can be useful in understanding the overall structure and setup of your projects for better management and optimization.

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
Analyze the distribution of your projects across various regions. This could be beneficial in understanding geographical trends or allocating resources more effectively based on project density.

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
Analyze the settings to understand the configuration of your project's API. This can provide valuable insights into the project's structure, such as the exposed schemas and maximum rows, which can help optimize the project's performance.

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
Analyze the settings to understand the list of allowed network ranges for each Supabase project. This query helps in assessing the network restrictions for database access, providing insights into the security configurations of your projects.

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