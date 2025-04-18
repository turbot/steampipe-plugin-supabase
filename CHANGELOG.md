## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#48](https://github.com/turbot/steampipe-plugin-supabase/pull/48))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#48](https://github.com/turbot/steampipe-plugin-supabase/pull/48))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#46](https://github.com/turbot/steampipe-plugin-supabase/pull/46))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#46](https://github.com/turbot/steampipe-plugin-supabase/pull/46))

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#40](https://github.com/turbot/steampipe-plugin-supabase/pull/40))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#40](https://github.com/turbot/steampipe-plugin-supabase/pull/40))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-supabase/blob/main/docs/LICENSE). ([#40](https://github.com/turbot/steampipe-plugin-supabase/pull/40))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#39](https://github.com/turbot/steampipe-plugin-supabase/pull/39))

## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#18](https://github.com/turbot/steampipe-plugin-supabase/pull/18))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#14](https://github.com/turbot/steampipe-plugin-supabase/pull/14))
- Recompiled plugin with Go version `1.21`. ([#14](https://github.com/turbot/steampipe-plugin-supabase/pull/14))

## v0.0.1 [2023-03-21]

_What's new?_

- New tables added
  - [supabase_function](https://hub.steampipe.io/plugins/turbot/supabase/tables/supabase_function)
  - [supabase_organization](https://hub.steampipe.io/plugins/turbot/supabase/tables/supabase_organization)
  - [supabase_project](https://hub.steampipe.io/plugins/turbot/supabase/tables/supabase_project)
  - [supabase_project_custom_hostname](https://hub.steampipe.io/plugins/turbot/supabase/tables/supabase_project_custom_hostname)
  - [supabase_project_network_restriction](https://hub.steampipe.io/plugins/turbot/supabase/tables/supabase_project_network_restriction)
  - [supabase_secret](https://hub.steampipe.io/plugins/turbot/supabase/tables/supabase_secret)
