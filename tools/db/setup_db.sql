BEGIN;
\i /docker-entrypoint-initdb.d/schema.sql;
\i /docker-entrypoint-initdb.d/triggers.sql;
COMMIT;