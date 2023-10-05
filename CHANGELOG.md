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
