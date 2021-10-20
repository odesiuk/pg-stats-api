# pg-stats-api

---

Run to use git hooks `git config core.hooksPath .githooks`

---

To show stats, the connected DB must support the `pg_stat_statements` extention.
---

Create your `.env` file using the `.env.dist` template

```dotenv
# ----- optional -----

# skip if you will use docker-compose.yaml.
PORT=8080

# service will return queries that were slower than MIN_QUERY_DURATION
# default is 2000.
MIN_QUERY_DURATION=5000

# db host & port default values.
PG_HOST=localhost
PG_PORT=5432

# ----- required -----

# db credentials.
PG_USER=test
PG_PASSWORD=testpass
PG_DATABASE=test_db

```

Use `/storage/script.sql` to add some test queries to DB.

Run the service

```
docker-compose build --no-cache

docker-compose --env-file .env up -d
```
