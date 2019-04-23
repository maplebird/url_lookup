# url_lookup

I used this as an excuse to learn GoLang :)

# Requirements

## Docker test environment

 * Linux or OS X
 * Docker
 * MySQL binary (just `mysql` executable), do not need full server

This is the recommended way for local testing, as it only requires 3 shell scripts to be executed.

## Local server
 * Linux or OS X
 * GoLang
 * MySQL server
 
This was only tested using MySQL 5.7 and 8.0, however any version 5.1 and above should work fine.

# Setting up a test environment

## Easy option with Docker

Make sure your docker daemon is started, and local MySQL Server is stopped if running.

Alternatively, set url_lookup to run on a non-standard MySQL port so there are no collisions 
(see [next section](run-mysql-on-non-standard-port)).

You may also need to run `docker login` first.

Run wrapper script that will build app and db containers and then start test server:

```bash
./wrapper.sh
```

There are 3 scripts under wrapper.sh will:
* Build and start test database container, then run schema migrations against it (including populating it with some test data)
* Build url_lookup_server app container
* Start the app container in the current terminal session (all logs will be output to STDOUT)

If you do not wish to automatically create a test database, follow the steps to set up a database from the next section.

Start the test server.

`./start_test_server.sh`

You should now be able to hit the API (see [Using URL Lookup](#using-url_lookup) for instructions)

Cleanup at the end by running `./cleanup.sh`.
This will delete any new docker images or running containers
(with the exception of public `mysql` and `golang` images)

### Run MySQL on non-standard port 

You can set test environment to spawn a database container on a non-standard port 
and have the server container automatically connect to it.

To do this, just set this environment variable before running `wrapper.sh`:

```bash
export URL_LOOKUP_DBPORT=3307 
``` 

Then run the 3 scripts from the previous section.  No further work should be required.

## Local server (not running in docker)

Make sure GoLang is installed (uses local directory for GOPATH to download dependencies and output binary)

For database, either:
* Configure database as per the the [Database setup](#database-setup) section.

Run `./local_build.sh`

This will download any required dependencies to src/, run unit and integration tests, and finally start test server
in the current terminal session.

## Database setup

Follow the steps below to configure database required for url_lookup if not using automatically created test Docker database
or running a remote DB instance.

* Make sure MySQL is started on the db host (tested on OS X with MySQL 5.7).
* Run all `.sql` scripts in `src/migrations` against the database in question
* DB must allow TCP connections other than from 127.0.0.1 (i.e. listen on local IPv4) if using docker app server
* Either execute them in the database shell, or run them as follows:

```bash
for SQL in $(ls migrations/*.sql); do
    mysql -h HOST -u USER -p < ${SQL}
done
```

Where HOST is your database host (`127.0.0.1` for local server) and USER is your admin database user.

You then need to point the server against this DB instance.

Update `src/url_lookup/config.properties` with database configuration.  Default configuration is shown below:

```
DBHost = 127.0.0.1
DBPort = 3306
DBSchema = url_lookup
DBUser = url_lookup_ro
DBPassword = password
```

Alternatively, set any required environment variables before starting url_lookup server:

* `URL_LOOKUP_DBHOST`
* `URL_LOOKUP_DBPORT`
* `URL_LOOKUP_DBSCHEMA`
* `URL_LOOKUP_DBUSER`
* `URL_LOOKUP_DBPASSWORD`

### Important note
`url_analysis` only supports TCP connections to its database.
It will not work with UNIX socket connections, even when running the database on localhost.

# Using url_lookup

URL lookup is a simple REST API accessible via HTTP GET that provides reputation of any known URLs.

Reputation comes in 4 categories:

* Known safe.  Websites that are known good and can be considered free of malware or questionable content.
* Known unsafe.  Websites that are known bad and generally have a large amount of malware or host questionable content.
* Mixed.  Generally, websites that are not necessarily malicious, but can be used to host malicious data
  * When querying mixed reputation websites, full URL path is checked
  * An example of a mixed reputation website is a file hosting service such as www.megaupload.com
    * Some files on it can be safe, some can be unsafe.
    * For example, querying www.megaupload.com/files/not_a_virus would return a safe reputation
    * Querying www.megaupload.com/files/my_virus would return an unsafe reputation
    * If reputation of a specific path on a mixed-reputation website is unknown, url_lookup shows mixed reputation
* Unknown.  Websites that are not in the url_lookup database.

By default, it is configured to listen on port 5000 and can be accessed as such.

## Querying url_lookup

url_lookup can be accessed at the following URL:

`my_host:5000/1/urlinfo/{url_to_query}`

* my_host is hostname or IP of the server url_lookup runs on (or localhost if used in testing)
* {url_to_query} is url of website you'd like to get info about

Example usage:

```
curl localhost:5000/1/urlinfo/www.megaupload.com/files/my_virus
```

It returns a JSON object with the following structure:

```json
{
  "RequestedUrl": "www.megaupload.com/files/my_virus",
  "Reputation": "unsafe"
}
```

### Scheme in lookup request

Including scheme in your request, such as
 
`/1/urlinfo/https://www.megaupload.com/files/my_virus`

Is **NOT** supported.

Please strip out scheme portion of the URL before querying it.

### Ports in lookup request

Ports are supported, but any ports are considered separate websites entirely
and need to be added to the database separately.

For example:

`www.megaupload.com:80/files/my_virus`

and

`www.megaupload.com:443/files/my_virus`

Are two separate websites entirely as far as url_lookup is concerned.

For best results, only store base FQDN and object paths of any website added to the database,
but exclude port.  Then, when querying url_lookup, remove port from your lookup.

This was a conscious design decision, as a website can host the same content on multiple ports,
which can also be accessed in multiple ways (such as `https://foo.com`, `https://foo.com:443`, or 
`http://foo.com`).  This adds significant overhead when managing url_lookup database of website reputations 
but provides little tangible benefit.

### Managing the database

Current schema is very simple, with only two tables: `fqdns`, and `path_lookup`.

`fqdns` table stores a list of any websites using their fully-qualified domain names 
(host.domain-name.TLD), and reputation.

Schema looks like this:

```
+--------------------+------------+
| fqdn               | reputation |
+--------------------+------------+
| get.dogecoin.com   | unsafe     |
| www.github.com     | safe       |
| www.megaupload.com | mixed      |
+--------------------+------------+
```

`path_lookup` stores a list of full object paths for any FQDNs defined in `fqdns` with a mixed reputation.

```
+----+--------------------+--------------------+------------+
| id | fqdn               | path               | reputation |
+----+--------------------+--------------------+------------+
|  1 | www.megaupload.com | /files/not_a_virus | safe       |
|  2 | www.megaupload.com | /files/my_virus    | unsafe     |
|  3 | www.megaupload.com | /files/random_file | unknown    |
+----+--------------------+--------------------+------------+
```

### Adding new websites to the database

If a website has a known safe or unsafe reputation, it only needs to be stored in `fqdns` table.

```sql
INSERT INTO fqdns (fqdn, reputation) VALUES
('my-host.com', 'safe');
```

Note that `fqdn` table can only store unique hostnames.  You cannot store two objects with fqdn `my-host.com`.

If a website has a mixed reputation and you would like to configure multiple objects as safe of unsafe,
you also need to add their reputations to `path_lookup` table.

```sql
INSERT INTO path_lookup (fqdn, path, reputation) VALUES
('my-host.com', '/link/to/safe/object', 'safe'),
('my-host.com', '/link/to/unsafe/object', 'unsafe');
```
