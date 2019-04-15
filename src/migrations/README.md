# Setup for default database

Make sure MySQL is started on the host (tested with 5.7, however previous versions should work fine).

1. Create user

```sql
GRANT SELECT ON url_lookup.* TO `url_lookup_ro`@`127.0.0.1`
IDENTIFIED  BY 'password';
```

2. Create default schema
* In this project in `src/main/migrations/01_create_default_schema.sql`
* Execute against local MySQL host

If you wish to use a remote database (MySQL only), update `src/main/config.properties` with your database info

```
dbProperties.host = 127.0.0.1
dbProperties.port = 3306
dbProperties.schema = url_lookup
dbProperties.user = url_lookup_ro
dbProperties.password = password
```

