# url_lookup

I used this as an excuse to learn GoLang :)

# Requirements
 * Linux or OS X
 * Docker
 * GoLang
 * MySQL binary (do not need MySQL server, just `mysql` CLI executable)
 
If not using test docker database or a local MySQL server, the only requirement is Linux/OS X with installed GoLang.

Make sure to configure remote database using the steps in its corresponding section.

This was tested using MySQL 5.7 and 8.0, however 5.6 should work fine.
 
# Setting up a test environment

## Easy option

Make sure your Docker daemon is started.

In the main folder, run build_test_database.sh
* You may need to run `docker login` and start your docker daemon before running this script

`./build_test_database.sh`

If you do not wish to automatically create a test database, follow the steps to set up a database from the next section.

Start the test server.

`./start_test_server.sh`

You should now be able to hit the API (see [Using URL Lookup](#using-url_lookup) for instructions)

Make sure to cleanup at the end by running `./cleanup.sh`.  This will delete any build artifacts and test database container.

## Database setup if not using Docker script.

Follow the steps below to configure database required for url_lookup if not using test Docker database

* Make sure MySQL is started on the db host (tested with 5.7 and 8.0, however previous versions should work fine).
* Run all `.sql` scripts in `src/migrations` against the database in question
* Either execute them directly in the database shell, or run them as follows:

```bash
for SQL in $(ls src/migrations/*.sql); do
    mysql -h HOST -u USER -p < ${SQL}
done
```

Where HOST is your database host and USER is your database user.

If you wish to use a database on a remote host or non-standard port, update `src/main/config.properties` with your configuration

```
dbProperties.host = 127.0.0.1
dbProperties.port = 3306
dbProperties.schema = url_lookup
dbProperties.user = url_lookup_ro
dbProperties.password = password
```

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

#### Adding new websites to the database

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

