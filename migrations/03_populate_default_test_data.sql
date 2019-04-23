# Populate some default values for testing
INSERT INTO url_lookup.fqdns (fqdn, reputation) VALUES
('www.github.com', 'safe'),
('get.dogecoin.com', 'unsafe'),
('www.megaupload.com', 'mixed');

INSERT INTO url_lookup.lookup_paths (fqdn, path, reputation) VALUES
('www.megaupload.com', '/files/not_a_virus', 'safe'),
('www.megaupload.com', '/files/my_virus', 'unsafe'),
('www.megaupload.com', '/files/random_file', 'unknown');

