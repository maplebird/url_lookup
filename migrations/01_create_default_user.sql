CREATE USER `url_lookup_ro`@`%`
IDENTIFIED  BY 'password';

GRANT SELECT ON url_lookup.* TO `url_lookup_ro`@`%`;

FLUSH PRIVILEGES;