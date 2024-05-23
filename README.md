# Create env variables
| Variable | Value |
|---|---|
ENV | staging
VERSION | 1.0
HOST | localhost
ORIGINS | *
DB_CONNECTION | postgresql://postgres:olEK98s8n9Z8hSKH12f@db:5432/protec
DB_PASSWORD | olEK98s8n9Z8hSKH12f
DB_NAME | protec
|  |  |

# Build and install

- Run the next command in docker compose to run the project in local:

`docker-compose build && docker-compose up -d`

# Install golang migrate

- You can see all the documentation of the library here:

https://github.com/golang-migrate/migrate

- For Mac:

`brew install golang-migrate`

# Migration

- Execute migration, run the next commant inside the root directory of the app:

`make migrate_up`

# Go to the API URI

http://localhost:3000/swagger/index.html

