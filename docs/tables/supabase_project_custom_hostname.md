---
title: "Steampipe Table: supabase_project_custom_hostname - Query Supabase Project Custom Hostnames using SQL"
description: "Allows users to query Supabase Project Custom Hostnames, specifically to retrieve details regarding the custom domain names associated with a Supabase project."
---

# Table: supabase_project_custom_hostname - Query Supabase Project Custom Hostnames using SQL

A Supabase Project Custom Hostname is an attribute of a Supabase project, which allows users to associate custom domain names with their projects. It enables the customization of the project's URL, thereby enhancing its accessibility and user-friendliness. This feature contributes to the flexibility and scalability of Supabase projects.

## Table Usage Guide

The `supabase_project_custom_hostname` table provides insights into the custom domain names associated with a Supabase project. As a project manager or developer, you can explore details about these custom hostnames through this table, including their verification status, creation time, and associated project ID. Utilize it to manage and monitor the custom domain names of your Supabase projects.

## Examples

### Basic info
Explore which custom hostnames are associated with your Supabase projects, along with their status and creation date. This can be useful in managing and tracking your project's custom domain configurations.

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
Discover the segments that include custom hostnames in your projects that are yet to begin. This is beneficial in managing project timelines and identifying potential delays.

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
Uncover the details of recent custom hostnames to monitor the development of your projects. This is particularly useful for keeping track of any new custom hostnames that have been created within the last three days.

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