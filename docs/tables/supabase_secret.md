---
title: "Steampipe Table: supabase_secret - Query Supabase Secrets using SQL"
description: "Allows users to query Supabase Secrets, particularly the secret key and the status of the secret. This provides insights into the security and access management within a Supabase project."
---

# Table: supabase_secret - Query Supabase Secrets using SQL

Supabase Secrets are essentially the keys used to access and manage various aspects of a Supabase project. These keys are used to authenticate users, manage roles, and perform various operations within the project. Understanding and managing these secrets is crucial for maintaining security and access control within the Supabase environment.

## Table Usage Guide

The `supabase_secret` table provides insights into the secrets within a Supabase project. As a security engineer or a project manager, you can explore secret-specific details through this table, including the secret key and its status. Use it to manage and monitor the security and access control of your Supabase project, ensuring the right access is provided to the right users and roles.

## Examples

### Basic info
Explore which project secrets are being used in your Supabase setup. This can help you manage and track the use of sensitive information across your projects.

```sql+postgres
select
  project_id,
  name,
  value
from
  supabase_secret;
```

```sql+sqlite
select
  project_id,
  name,
  value
from
  supabase_secret;
```

### List all secrets of a specific project
Discover the segments that contain all hidden or confidential information for a specific project. This can be useful in assessing data security and ensuring sensitive data is adequately protected.

```sql+postgres
select
  project_id,
  name,
  value
from
  supabase_secret
where
  project_id = 'pljlooizchwsteampipe';
```

```sql+sqlite
select
  project_id,
  name,
  value
from
  supabase_secret
where
  project_id = 'pljlooizchwsteampipe';
```