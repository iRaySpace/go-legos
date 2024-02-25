# learn-sqlc-goose
---
Experimenting sqlc and goose

## Instructions
```
cd data/sqlc
sqlc generate
cd migrations
goose postgres "postgres://postgres:123@localhost:5432/postgres" up
```