CREATE SCHEMA url_lookup CHARACTER SET UTF8 COLLATE utf8_bin;

CREATE TABLE fqdns (
    fqdn VARCHAR(200) NOT NULL,
    type CHAR(25) NOT NULL DEFAULT 'safe',
    PRIMARY KEY (fqdn)
);

CREATE TABLE path_lookup (
    id int NOT NULL,
    fqdn VARCHAR(200),
    path VARCHAR(200),
    type CHAR(25) NOT NULL DEFAULT 'safe',
    PRIMARY KEY (id)
);
