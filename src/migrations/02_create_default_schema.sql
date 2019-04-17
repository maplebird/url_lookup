CREATE SCHEMA url_lookup CHARACTER SET UTF8 COLLATE utf8_bin;

CREATE TABLE url_lookup.fqdns (
    fqdn VARCHAR(200) NOT NULL,
    reputation CHAR(25) NOT NULL DEFAULT 'unknown',
    INDEX (fqdn),
    PRIMARY KEY (fqdn)
);

CREATE TABLE url_lookup.path_lookup (
    id int NOT NULL AUTO_INCREMENT,
    fqdn VARCHAR(200),
    path VARCHAR(200),
    reputation CHAR(25) NOT NULL DEFAULT 'unknown',
    INDEX (fqdn, path),
    PRIMARY KEY (id)
);
