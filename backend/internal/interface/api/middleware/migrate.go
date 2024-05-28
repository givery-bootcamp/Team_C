package middleware

type DBMigrator interface {
	Migrate() error
}
