![image](https://hub.steampipe.io/images/plugins/turbot/supabase-social-graphic.png)

# Supabase Plugin for Steampipe

Use SQL to query projects, functions, network restrictions, and more from your Supabase organization.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/supabase)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/supabase/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
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

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-supabase/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Supabase Plugin](https://github.com/turbot/steampipe-plugin-supabase/labels/help%20wanted)
