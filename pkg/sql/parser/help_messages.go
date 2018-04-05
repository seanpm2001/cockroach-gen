// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1061
	`ALTER`: {
		//line sql.y: 1062
		Category: hGroup,
		//line sql.y: 1063
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1077
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1078
		Category: hDDL,
		//line sql.y: 1079
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause>
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

`,
		//line sql.y: 1101
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1113
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1114
		Category: hDDL,
		//line sql.y: 1115
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1117
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1124
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1125
		Category: hDDL,
		//line sql.y: 1126
		Text: `
ALTER SEQUENCE [IF EXISTS] <name>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START <start>]
  [[NO] CYCLE]
ALTER SEQUENCE [IF EXISTS] <name> RENAME TO <newname>
`,
	},
	//line sql.y: 1149
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1150
		Category: hPriv,
		//line sql.y: 1151
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1153
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1158
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1159
		Category: hDDL,
		//line sql.y: 1160
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1162
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1173
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1174
		Category: hDDL,
		//line sql.y: 1175
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1183
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1501
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1502
		Category: hCCL,
		//line sql.y: 1503
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1520
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1528
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1529
		Category: hCCL,
		//line sql.y: 1530
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1546
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1564
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1565
		Category: hCCL,
		//line sql.y: 1566
		Text: `
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   comma = '...'          [CSV-specific]
   comment = '...'        [CSV-specific]
   nullif = '...'         [CSV-specific]

`,
		//line sql.y: 1584
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1675
	`CANCEL`: {
		//line sql.y: 1676
		Category: hGroup,
		//line sql.y: 1677
		Text: `CANCEL JOB, CANCEL QUERY, CANCEL SESSION
`,
	},
	//line sql.y: 1684
	`CANCEL JOB`: {
		ShortDescription: `cancel a background job`,
		//line sql.y: 1685
		Category: hMisc,
		//line sql.y: 1686
		Text: `CANCEL JOB <jobid>
`,
		//line sql.y: 1687
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOB
`,
	},
	//line sql.y: 1695
	`CANCEL QUERY`: {
		ShortDescription: `cancel a running query`,
		//line sql.y: 1696
		Category: hMisc,
		//line sql.y: 1697
		Text: `CANCEL QUERY [IF EXISTS] <queryid>
`,
		//line sql.y: 1698
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1710
	`CANCEL SESSION`: {
		ShortDescription: `cancel an open session`,
		//line sql.y: 1711
		Category: hMisc,
		//line sql.y: 1712
		Text: `CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 1713
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 1741
	`CREATE`: {
		//line sql.y: 1742
		Category: hGroup,
		//line sql.y: 1743
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 1764
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 1765
		Category: hMisc,
		//line sql.y: 1766
		Text: `
CREATE STATISTICS <statisticname>
  ON <colname> [, ...]
  FROM <tablename>
`,
	},
	//line sql.y: 1781
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 1782
		Category: hDML,
		//line sql.y: 1783
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 1787
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 1802
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 1803
		Category: hCfg,
		//line sql.y: 1804
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 1816
	`DROP`: {
		//line sql.y: 1817
		Category: hGroup,
		//line sql.y: 1818
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 1834
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 1835
		Category: hDDL,
		//line sql.y: 1836
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1837
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 1849
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 1850
		Category: hDDL,
		//line sql.y: 1851
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1852
		SeeAlso: `DROP
`,
	},
	//line sql.y: 1864
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 1865
		Category: hDDL,
		//line sql.y: 1866
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1867
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 1879
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 1880
		Category: hDDL,
		//line sql.y: 1881
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1882
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 1902
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 1903
		Category: hDDL,
		//line sql.y: 1904
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 1905
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 1925
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 1926
		Category: hPriv,
		//line sql.y: 1927
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 1928
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 1940
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 1941
		Category: hPriv,
		//line sql.y: 1942
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 1943
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 1965
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 1966
		Category: hMisc,
		//line sql.y: 1967
		Text: `
EXPLAIN <statement>
EXPLAIN [( [PLAN ,] <planoptions...> )] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN, EXECUTE

Plan options:
    TYPES, EXPRS, METADATA, QUALIFY, INDENT, VERBOSE, DIST_SQL

`,
		//line sql.y: 1978
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2039
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2040
		Category: hMisc,
		//line sql.y: 2041
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2042
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2064
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2065
		Category: hMisc,
		//line sql.y: 2066
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2067
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2090
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2091
		Category: hMisc,
		//line sql.y: 2092
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2093
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2113
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2114
		Category: hPriv,
		//line sql.y: 2115
		Text: `
Grant privileges:
  GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>
Grant role membership (CCL only):
  GRANT <roles...> TO <grantees...> [WITH ADMIN OPTION]

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2128
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2144
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2145
		Category: hPriv,
		//line sql.y: 2146
		Text: `
Revoke privileges:
  REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>
Revoke role membership (CCL only):
  REVOKE [ADMIN OPTION FOR] <roles...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2159
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2228
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2229
		Category: hCfg,
		//line sql.y: 2230
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2231
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2243
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2244
		Category: hCfg,
		//line sql.y: 2245
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2246
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2255
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2256
		Category: hCfg,
		//line sql.y: 2257
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2260
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2277
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2278
		Category: hExperimental,
		//line sql.y: 2279
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2287
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2293
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2294
		Category: hExperimental,
		//line sql.y: 2295
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2303
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2311
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2312
		Category: hExperimental,
		//line sql.y: 2313
		Text: `
SCRUB TABLE <tablename>
            [AS OF SYSTEM TIME <expr>]
            [WITH OPTIONS <option> [, ...]]

Options:
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT (<constraint>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS PHYSICAL
`,
		//line sql.y: 2324
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2379
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2380
		Category: hCfg,
		//line sql.y: 2381
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2382
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2403
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2404
		Category: hCfg,
		//line sql.y: 2405
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }

`,
		//line sql.y: 2410
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2427
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2428
		Category: hTxn,
		//line sql.y: 2429
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2436
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 2601
	`SHOW`: {
		//line sql.y: 2602
		Category: hGroup,
		//line sql.y: 2603
		Text: `
SHOW SESSION, SHOW CLUSTER SETTING, SHOW DATABASES, SHOW TABLES, SHOW COLUMNS, SHOW INDEXES,
SHOW CONSTRAINTS, SHOW CREATE TABLE, SHOW CREATE VIEW, SHOW CREATE SEQUENCE, SHOW USERS,
SHOW TRANSACTION, SHOW BACKUP, SHOW JOBS, SHOW QUERIES, SHOW ROLES, SHOW SESSIONS, SHOW SYNTAX,
SHOW TRACE
`,
	},
	//line sql.y: 2636
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 2637
		Category: hCfg,
		//line sql.y: 2638
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 2639
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 2660
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics`,
		//line sql.y: 2661
		Category: hMisc,
		//line sql.y: 2662
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 2669
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 2681
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram`,
		//line sql.y: 2682
		Category: hMisc,
		//line sql.y: 2683
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 2687
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 2700
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 2701
		Category: hCCL,
		//line sql.y: 2702
		Text: `SHOW BACKUP <location>
`,
		//line sql.y: 2703
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 2711
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 2712
		Category: hCfg,
		//line sql.y: 2713
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 2716
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2733
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 2734
		Category: hDDL,
		//line sql.y: 2735
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 2736
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 2744
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 2745
		Category: hDDL,
		//line sql.y: 2746
		Text: `SHOW DATABASES
`,
		//line sql.y: 2747
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 2755
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 2756
		Category: hPriv,
		//line sql.y: 2757
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 2763
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 2776
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 2777
		Category: hDDL,
		//line sql.y: 2778
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 2779
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 2797
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 2798
		Category: hDDL,
		//line sql.y: 2799
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 2800
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 2813
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 2814
		Category: hMisc,
		//line sql.y: 2815
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 2816
		SeeAlso: `CANCEL QUERY
`,
	},
	//line sql.y: 2832
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 2833
		Category: hMisc,
		//line sql.y: 2834
		Text: `SHOW JOBS
`,
		//line sql.y: 2835
		SeeAlso: `CANCEL JOB, PAUSE JOB, RESUME JOB
`,
	},
	//line sql.y: 2843
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 2844
		Category: hMisc,
		//line sql.y: 2845
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
SHOW [COMPACT] [KV] TRACE FOR <statement>
`,
		//line sql.y: 2848
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 2878
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 2879
		Category: hMisc,
		//line sql.y: 2880
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 2881
		SeeAlso: `CANCEL SESSION
`,
	},
	//line sql.y: 2897
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 2898
		Category: hDDL,
		//line sql.y: 2899
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ]
`,
		//line sql.y: 2900
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 2926
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 2927
		Category: hDDL,
		//line sql.y: 2928
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 2940
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 2941
		Category: hMisc,
		//line sql.y: 2942
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 2951
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 2952
		Category: hCfg,
		//line sql.y: 2953
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 2954
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 2973
	`SHOW CREATE TABLE`: {
		ShortDescription: `display the CREATE TABLE statement for a table`,
		//line sql.y: 2974
		Category: hDDL,
		//line sql.y: 2975
		Text: `SHOW CREATE TABLE <tablename>
`,
		//line sql.y: 2976
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 2984
	`SHOW CREATE VIEW`: {
		ShortDescription: `display the CREATE VIEW statement for a view`,
		//line sql.y: 2985
		Category: hDDL,
		//line sql.y: 2986
		Text: `SHOW CREATE VIEW <viewname>
`,
		//line sql.y: 2987
		SeeAlso: `WEBDOCS/show-create-view.html
`,
	},
	//line sql.y: 2995
	`SHOW CREATE SEQUENCE`: {
		ShortDescription: `display the CREATE SEQUENCE statement for a sequence`,
		//line sql.y: 2996
		Category: hDDL,
		//line sql.y: 2997
		Text: `SHOW CREATE SEQUENCE <seqname>
`,
	},
	//line sql.y: 3005
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3006
		Category: hPriv,
		//line sql.y: 3007
		Text: `SHOW USERS
`,
		//line sql.y: 3008
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3016
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3017
		Category: hPriv,
		//line sql.y: 3018
		Text: `SHOW ROLES
`,
		//line sql.y: 3019
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3109
	`PAUSE JOB`: {
		ShortDescription: `pause a background job`,
		//line sql.y: 3110
		Category: hMisc,
		//line sql.y: 3111
		Text: `PAUSE JOB <jobid>
`,
		//line sql.y: 3112
		SeeAlso: `SHOW JOBS, CANCEL JOB, RESUME JOB
`,
	},
	//line sql.y: 3120
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3121
		Category: hDDL,
		//line sql.y: 3122
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 3149
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE TABLE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 3660
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 3661
		Category: hDDL,
		//line sql.y: 3662
		Text: `
CREATE SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]

`,
		//line sql.y: 3671
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 3724
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 3725
		Category: hDML,
		//line sql.y: 3726
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 3727
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 3735
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 3736
		Category: hPriv,
		//line sql.y: 3737
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 3738
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 3760
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 3761
		Category: hPriv,
		//line sql.y: 3762
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 3763
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 3775
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 3776
		Category: hDDL,
		//line sql.y: 3777
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 3778
		SeeAlso: `CREATE TABLE, SHOW CREATE VIEW, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 3792
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 3793
		Category: hDDL,
		//line sql.y: 3794
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 3802
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE INDEX,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 3996
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 3997
		Category: hTxn,
		//line sql.y: 3998
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 3999
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4007
	`RESUME JOB`: {
		ShortDescription: `resume a background job`,
		//line sql.y: 4008
		Category: hMisc,
		//line sql.y: 4009
		Text: `RESUME JOB <jobid>
`,
		//line sql.y: 4010
		SeeAlso: `SHOW JOBS, CANCEL JOB, PAUSE JOB
`,
	},
	//line sql.y: 4018
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 4019
		Category: hTxn,
		//line sql.y: 4020
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 4021
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4036
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 4037
		Category: hTxn,
		//line sql.y: 4038
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 4046
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 4059
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 4060
		Category: hTxn,
		//line sql.y: 4061
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 4064
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 4088
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 4089
		Category: hTxn,
		//line sql.y: 4090
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 4091
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 4204
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 4205
		Category: hDDL,
		//line sql.y: 4206
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4207
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 4276
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 4277
		Category: hDML,
		//line sql.y: 4278
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4283
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 4302
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 4303
		Category: hDML,
		//line sql.y: 4304
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 4308
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 4413
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 4414
		Category: hDML,
		//line sql.y: 4415
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4422
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 4592
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 4593
		Category: hDML,
		//line sql.y: 4594
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 4605
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 4606
		Category: hDML,
		//line sql.y: 4607
		Text: `
SELECT [DISTINCT [ ON ( <expr> [ , ... ] ) ] ]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
`,
		//line sql.y: 4619
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 4694
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 4695
		Category: hDML,
		//line sql.y: 4696
		Text: `TABLE <tablename>
`,
		//line sql.y: 4697
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 4964
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 4965
		Category: hDML,
		//line sql.y: 4966
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 4967
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5068
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 5069
		Category: hDML,
		//line sql.y: 5070
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexhint> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> ON <expr>
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> USING ( <colnames...> )
  <source> NATURAL { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index hints:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 5088
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
