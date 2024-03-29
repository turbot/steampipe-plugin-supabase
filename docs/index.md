---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/supabase.svg"
brand_color: "#3ECF8E"
display_name: "Supabase"
short_name: "supabase"
description: "Steampipe plugin to query projects, functions, network restrictions, and more from your Supabase organization."
og_description: "Query Supabase with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/supabase-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Supabase + Steampipe

[Supabase](https://supabase.com) is an open-source alternative to Firebase, provides a suite of tools and services that help developers build applications with backend functionalities.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

List projects in your Supabase organization:

```sql
select
  name,
  region,
  created_at,
  organization_id
from
  supabase_project;
```

```
+-----------+------------+---------------------------+----------------------+
| name      | region     | created_at                | organization_id      |
+-----------+------------+---------------------------+----------------------+
| Steampunk | us-east-1  | 2023-02-15T20:19:50+05:30 | zuluktedwinzftfztsub |
| Steampipe | ap-south-1 | 2023-02-13T21:29:46+05:30 | zuluktedwinzftfztsub |
+-----------+------------+---------------------------+----------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/supabase/tables)**

## Quick start

### Install

Download and install the latest Supabase plugin:

```bash
steampipe plugin install supabase
```

### Credentials

| Item        | Description                                                                                                                                                                          |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Credentials | All API requests require a Supabase [Access Token](https://supabase.com/docs/reference/api/introduction#authentication) to be included in the `Authorization` header.                |
| Permissions | API tokens carry the same privileges as your user account, and if the user permissions change, the API token permissions also change.                                                |
| Radius      | Each connection represents a single Supabase Installation.                                                                                                                           |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/supabase.spc`)<br />2. Credentials specified in environment variables, e.g., `SUPABASE_ACCESS_TOKEN`. |

### Configuration

Installing the latest supabase plugin will create a config file (`~/.steampipe/config/supabase.spc`) with a single connection named `supabase`:

```hcl
connection "supabase" {
  plugin = "supabase"

  # The Supabase personal token.
  # All API requests require a Supabase Personal token to be included in the Authorization header.
  # To generate or manage your API token, visit your account page: https://app.supabase.com/account/tokens
  # This can also be set via the `SUPABASE_ACCESS_TOKEN` environment variable.
  # access_token = "sbp_123a45b6c78d901e2345f6steampipe45i432101"
}
```

Alternatively, you can also use the standard Supabase environment variables to obtain credentials only if other argument (`access_token`) is not specified in the connection:

```sh
export SUPABASE_ACCESS_TOKEN=sbp_123a45b6c78d901e2345f6steampipe45i432101
```


