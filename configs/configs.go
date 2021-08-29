package configs

const (
	// DB configs
	DBName     = "test-db"
	DBUser     = "postgres"
	DBPassword = "shubham"
	DBHost     = "localhost"
	DBPort     = 5432
	DBSSLMode  = "disable"

	// DB Connection Pool configs
	DBPoolMaxConns = 10
	DBMaxIdleConns = 2
	DBConnLifetime = 2 // in minutes

	// Cart table Queries
	DropCartTableQuery   = "DROP TABLE IF EXISTS cart"
	CreateCartTableQuery = "CREATE TABLE IF NOT EXISTS cart (cart_id uuid not null primary key, product_details jsonb not null)"

	// HTTP Router settings
	HTTPRouterPort = ":3000"
)
