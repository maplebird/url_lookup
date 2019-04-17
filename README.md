# Requirements
 * Linux or OS X
 * Docker
 * GoLang
 * MySQL binary (do not need MySQL server, just `mysql` CLI executable)
 
# Set up a test environment

Make sure your Docker daemon is started.

In the main folder, run build_test_database.sh
* You may need to run `docker login` and start your docker daemon before running this script

`./build_test_database.sh`

Then start test server

`./start_test_server.sh`

You should now be able to hit the API

# Setup for default database

Follow the steps below to configure database required for url_lookup if not using test Docker database

Make sure MySQL is started on the host (tested with 5.7 and 8.0, however previous versions should work fine).

1. Create user

```sql
GRANT SELECT ON url_lookup.* TO `url_lookup_ro`@`127.0.0.1`
IDENTIFIED  BY 'password';
```

2. Create default schema and populate with test data
* Two 
* In this project in `src/main/migrations/02_create_default_schema.sql`
* Execute against local MySQL host

3. Grant user permissions to read data

```sql
GRANT SELECT ON url_lookup.* TO `url_lookup_ro`@`127.0.0.1`;
``` 

If you wish to use a remote database (MySQL only), update `src/main/config.properties` with your database info

```
dbProperties.host = 127.0.0.1
dbProperties.port = 3306
dbProperties.schema = url_lookup
dbProperties.user = url_lookup_ro
dbProperties.password = password
```

