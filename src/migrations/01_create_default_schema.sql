# noinspection SqlNoDataSourceInspectionForFile

CREATE SCHEMA url_lookup CHARACTER SET UTF8 COLLATE utf8_bin;

CREATE TABLE url_lookup.fqdns (
    fqdn VARCHAR(200) NOT NULL UNIQUE,
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

# Populate some default values for testing
INSERT INTO url_lookup.fqdns (fqdn, reputation) VALUES
('www.github.com', 'safe'),
('get.dogecoin.com', 'unsafe'),
('www.megaupload.com', 'mixed');

INSERT INTO url_lookup.path_lookup (fqdn, path, reputation) VALUES
('www.megaupload.com', '/files/not_a_virus', 'safe'),
('www.megaupload.com', '/files/my_virus', 'unsafe'),
('www.megaupload.com', '/files/random_file', 'unknown');

