/*
  mysql table schema and initialize data management tools.

  example data directory:
    .
    └── table1             # this is the table name
        ├── data           # data scripts
        │   └── v1.sql
        └── schema         # schema scripts
            ├── latest.sql # create table if not exists
            └── v1.sql

  how it works:
    1. create table __tb_schema__ for saving current schema versions
    2. create table __tb_schema_logs__ for logging schema upgrade
    3. load each table version in __tb_schema__
    4. range each table by given directory
    5. build schema
      - if current table is not exists, create table by latest.sql
      - otherwise range v*.sql higher than current version
      - sort by version and upgrade
      - insert or update __tb_schema__.schema_version
      - insert __tb_schema_logs__
    6. build data
      - range v*.sql higher than current version
      - sort by version and upgrade
      - insert or update __tb_schema__.data_version
      - insert __tb_schema_logs__
    7. __tb_schema__ and __tb_schema_logs__ table is the protected table, so you can not use this table
    8. v0.sql is the protected file, so you can not define this version
    9. each sql operation will be the table name in sql file
*/
package mysqlschema
