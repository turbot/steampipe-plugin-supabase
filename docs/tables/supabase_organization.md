---
title: "Steampipe Table: supabase_organization - Query Supabase Organizations using SQL"
description: "Allows users to query Supabase Organizations, providing insights into organization details including name, owner, and creation date."
---

# Table: supabase_organization - Query Supabase Organizations using SQL

Supabase Organizations is a feature within the Supabase platform that allows users to manage multiple projects under one umbrella. It provides an efficient way to manage permissions and access for various Supabase projects. Supabase Organizations help in maintaining a structured approach to project management and access control.

## Table Usage Guide

The `supabase_organization` table provides insights into the organizations within Supabase. As a project manager or team lead, you can explore organization-specific details through this table, including the name of the organization, the owner, and when it was created. Utilize it to manage and monitor the access and permissions of various projects under your organization.

## Examples

### Basic info
Explore the basic details of your Supabase organizations to understand their identity and structure. This can aid in managing your resources effectively and keeping track of your organizational units.

```sql
select
  name,
  id
from
  supabase_organization;
```

### Get the count of projects per organization
Analyze the distribution of projects within your organizations. This can help you understand which organizations have the most projects, allowing you to allocate resources more effectively.

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