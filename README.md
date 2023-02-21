![image](https://hub.steampipe.io/images/plugins/turbot/supabase-social-graphic.png)

# Supabase Plugin for Steampipe

Use SQL to query projects, functions, network restrictions, and more from your Supabase organization.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/supabase)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/supabase/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-supabase/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install supabase
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/supabase#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/supabase#configuration).

Run a query:

```sql
select
  name,
  region,
  created_at,
  organization_id
from
  supabase_project;
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
