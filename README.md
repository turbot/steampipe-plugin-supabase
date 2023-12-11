![image](https://hub.steampipe.io/images/plugins/turbot/supabase-social-graphic.png)

# Supabase Plugin for Steampipe

Use SQL to query projects, functions, network restrictions, and more from your Supabase organization.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/supabase)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/supabase/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-supabase/issues)

## Quick start

### Install

Download and install the latest Supabase plugin:

```bash
steampipe plugin install supabase
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/supabase#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/supabase#configuration).

Configure your account details in `~/.steampipe/config/supabase.spc`:

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

Or through environment variables:

```sh
export SUPABASE_ACCESS_TOKEN=sbp_123a45b6c78d901e2345f6steampipe45i432101
```

Run steampipe:

```shell
steampipe query
```

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

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/index) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/index) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/index) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-supabase.git
cd steampipe-plugin-supabase
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/supabase.spc
```

Try it!

```shell
steampipe query
> .inspect supabase
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-supabase/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-supabase/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Supabase Plugin](https://github.com/turbot/steampipe-plugin-supabase/labels/help%20wanted)
