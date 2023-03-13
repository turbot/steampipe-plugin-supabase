# Table: supabase_secret

In Supabase, a secret is a piece of confidential information that is used for authentication, authorization, or other security-related purposes. Secrets can include things like API keys, database passwords, and other sensitive information that should not be publicly exposed.

Supabase allows users to securely store your secrets using environment variables, which can be accessed in your code using the Supabase library or other tools. This helps to protect your application and data from unauthorized access, and is an important part of building secure and reliable applications.

## Examples

### Basic info

```sql
select
  project_id,
  name,
  value
from
  supabase_secret;
```
