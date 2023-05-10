# pg-mini-cli
## Task
Write cli utility for database management.
Utility should be able to connect to DB and perform following operations:
- list databases
- delete selected database (or multiple)
- create db backup
- restore db from backup

## Install
```shell
go install github.com/vazik66/pg-mini-cli
```
## How to use
Default `--host` is localhost and `--port` is 5432 
Assume we have postgres running on localhost:5432 with user postgres and password mysecretpassword.
1. Show help
```shell
pg-mini-cli --help
```
Show help for individual command
```shell
pg-mini-cli list --help
pg-mini-cli backup --help
...
```

2. List databases
```shell
pg-mini-cli list -u postgres -p mysecretpassword
```

3. Delete databases
One database:
```shell
pg-mini-cli remove your_database_name -u postgres -p mysecretpassword 
```

Multiple databases:
```shell
pg-mini-cli remove your_database_name other_database and_another -u postgres -p mysecretpassword 
```

4. Create backup (requires psql installed)
```shell
pg-mini-cli backup your_database_name -u postgres -p mysecretpassword
```

Set custom backup filename
```shell
pg-mini-cli backup your_database_name -u postgres -p mysecretpassword -f mydump.dump
```

5. Restore from backup (requires psql installed)
```shell
pg-mini-cli restore mydump.dump -u postgres -p mysecretpassword -d database_name
```
