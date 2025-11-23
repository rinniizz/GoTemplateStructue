# Database Migrations

This directory contains database migration files for the project using [golang-migrate](https://github.com/golang-migrate/migrate).

## Migration Files

Migration files are named with the pattern: `{version}_{description}.{up|down}.sql`

- `up` files contain changes to apply
- `down` files contain changes to rollback

## Available Migrations

### 000001_create_users_table
Creates the users table with:
- Basic user fields (email, username, password, first_name, last_name)
- Timestamps (created_at, updated_at, deleted_at for soft delete)
- Indexes on email and username for fast lookups
- Auto-update trigger for updated_at timestamp

## Commands

### Install migrate CLI

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Create New Migration

```bash
migrate create -ext sql -dir migrations -seq <migration_name>
```

Example:
```bash
migrate create -ext sql -dir migrations -seq add_user_avatar
```

### Run Migrations

```bash
# Apply all pending migrations
migrate -path migrations -database "postgresql://user:password@localhost:5432/database?sslmode=disable" up

# Apply specific number of migrations
migrate -path migrations -database "postgresql://user:password@localhost:5432/database?sslmode=disable" up 1

# Rollback last migration
migrate -path migrations -database "postgresql://user:password@localhost:5432/database?sslmode=disable" down 1

# Rollback all migrations
migrate -path migrations -database "postgresql://user:password@localhost:5432/database?sslmode=disable" down

# Check migration version
migrate -path migrations -database "postgresql://user:password@localhost:5432/database?sslmode=disable" version

# Force set version (use with caution!)
migrate -path migrations -database "postgresql://user:password@localhost:5432/database?sslmode=disable" force <version>
```

### Windows Batch Scripts

For convenience on Windows, you can create batch files:

**migrate-up.bat:**
```batch
@echo off
set DB_URL=postgresql://postgres:a2l80MV4asdaEsmvrYF2@184.82.113.222:5432/gotemplate?sslmode=disable
migrate -path migrations -database "%DB_URL%" up
```

**migrate-down.bat:**
```batch
@echo off
set DB_URL=postgresql://postgres:a2l80MV4asdaEsmvrYF2@184.82.113.222:5432/gotemplate?sslmode=disable
migrate -path migrations -database "%DB_URL%" down 1
```

**migrate-create.bat:**
```batch
@echo off
if "%1"=="" (
    echo Usage: migrate-create.bat migration_name
    exit /b 1
)
migrate create -ext sql -dir migrations -seq %1
```

## Environment Variables

You can use environment variables for the database URL:

```bash
export DATABASE_URL="postgresql://user:password@localhost:5432/database?sslmode=disable"
migrate -path migrations -database "$DATABASE_URL" up
```

## Best Practices

1. **Always test migrations** on development database first
2. **Keep migrations small** and focused on one change
3. **Write reversible migrations** - ensure down migrations work correctly
4. **Never modify existing migrations** that have been applied to production
5. **Use transactions** when possible (add `-- +migrate StatementBegin` and `-- +migrate StatementEnd`)
6. **Backup production database** before running migrations

## Troubleshooting

### Dirty Database State

If you see "Dirty database version" error:

```bash
# Check current version
migrate -path migrations -database "$DATABASE_URL" version

# Force to previous version (if migration failed halfway)
migrate -path migrations -database "$DATABASE_URL" force <version>

# Then fix the issue and try again
migrate -path migrations -database "$DATABASE_URL" up
```

### Migration Failed

1. Check the error message carefully
2. Rollback the migration: `migrate down 1`
3. Fix the migration file
4. Run again: `migrate up`

## Integration with Application

The application uses GORM's AutoMigrate for development convenience, but for production:

1. Run migrations manually using migrate CLI
2. Or integrate migrations into deployment pipeline
3. Or use the migration package in code (see `pkg/database/migrate.go` if implemented)

## References

- [golang-migrate Documentation](https://github.com/golang-migrate/migrate)
- [Migration Best Practices](https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md)
