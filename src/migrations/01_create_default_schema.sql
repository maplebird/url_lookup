CREATE SCHEMA url_lookup CHARACTER SET UTF8 COLLATE utf8_bin;

CREATE TABLE url_lookup.fqdns (
    fqdn VARCHAR(200) NOT NULL,
    type CHAR(25) NOT NULL DEFAULT 'unknown',
    PRIMARY KEY (fqdn)
);

CREATE TABLE url_lookup.path_lookup (
    id int NOT NULL AUTO_INCREMENT,
    fqdn VARCHAR(200),
    path VARCHAR(200),
    type CHAR(25) NOT NULL DEFAULT 'unknown',
    PRIMARY KEY (id)
);
