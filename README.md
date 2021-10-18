# pg-stats-api

---

Run to use git hooks `git config core.hooksPath .githooks`

---

To show stats, the connected DB must support the `pg_stat_statements` extention.
---

Create your `.env` file using the `.env.dist` template and run the service

```
docker-compose build --no-cache

docker-compose --env-file .env up -d
```
