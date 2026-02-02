package config

import (
	"time"

	"github.com/spf13/viper"
	"github.com/yourorg/backoffice-go-service/internal/pkg/database"
)

// Config holds all application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	App      AppConfig
	Logging  LoggingConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port         string
	Host         string
	Mode         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	// Primary database (default)
	Primary DatabaseConnectionConfig `mapstructure:"primary"`
	
	// Additional databases can be added here
	Secondary DatabaseConnectionConfig `mapstructure:"secondary"`
	
	// Multiple databases support
	Databases map[string]DatabaseConnectionConfig `mapstructure:"databases"`
}

// DatabaseConnectionConfig holds configuration for a single database connection
type DatabaseConnectionConfig struct {
	Driver      string        `mapstructure:"driver"`       // postgresql, mysql, mongodb, sqlite
	Host        string        `mapstructure:"host"`
	Port        string        `mapstructure:"port"`
	User        string        `mapstructure:"user"`
	Password    string        `mapstructure:"password"`
	DBName      string        `mapstructure:"dbname"`
	SSLMode     string        `mapstructure:"sslmode"`      // For PostgreSQL
	Charset     string        `mapstructure:"charset"`      // For MySQL
	MaxOpenConns int          `mapstructure:"max_open_conns"`
	MaxIdleConns int          `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration  `mapstructure:"conn_max_idle_time"`
	UseGorm     bool          `mapstructure:"use_gorm"`
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret     string
	Expiration time.Duration
	Issuer     string
}

// AppConfig holds application-level configuration
type AppConfig struct {
	Name        string
	Version     string
	Environment string
	Debug       bool
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Channel     string // stdout, file, stack
	Level       string // debug, info, warn, error
	LogPath     string // Directory path for log files
	LogFileName string // Base name for log files
	MaxSize     int    // Maximum size in MB before rotation
	MaxBackups  int    // Maximum number of old log files to retain
	MaxAge      int    // Maximum number of days to retain old log files
	Compress    bool   // Whether to compress rotated log files
	DailyRotate bool   // Enable daily rotation
}

// LoadConfig loads configuration from environment variables and config files
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	
	// Enable environment variables
	viper.AutomaticEnv()
	
	// Set defaults
	setDefaults()
	
	// Try to read config file (optional)
	_ = viper.ReadInConfig()
	
	cfg := &Config{
		Server: ServerConfig{
			Port:         getString("SERVER_PORT", "8080"),
			Host:         getString("SERVER_HOST", "0.0.0.0"),
			Mode:         getString("GIN_MODE", "debug"),
			ReadTimeout:  getDuration("SERVER_READ_TIMEOUT", 15*time.Second),
			WriteTimeout: getDuration("SERVER_WRITE_TIMEOUT", 15*time.Second),
			IdleTimeout:  getDuration("SERVER_IDLE_TIMEOUT", 60*time.Second),
		},
		Database: DatabaseConfig{
			Primary: DatabaseConnectionConfig{
				Driver:         getString("DB_DRIVER", "postgresql"),
				Host:           getString("DB_HOST", "localhost"),
				Port:           getString("DB_PORT", "5432"),
				User:           getString("DB_USER", "postgres"),
				Password:       getString("DB_PASSWORD", ""),
				DBName:         getString("DB_NAME", "backoffice"),
				SSLMode:        getString("DB_SSLMODE", "disable"),
				Charset:        getString("DB_CHARSET", "utf8mb4"),
				MaxOpenConns:   getInt("DB_MAX_OPEN_CONNS", 25),
				MaxIdleConns:   getInt("DB_MAX_IDLE_CONNS", 5),
				ConnMaxLifetime: getDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
				ConnMaxIdleTime: getDuration("DB_CONN_MAX_IDLE_TIME", 10*time.Minute),
				UseGorm:        getBool("DB_USE_GORM", true),
			},
			Databases: make(map[string]DatabaseConnectionConfig),
		},
		JWT: JWTConfig{
			Secret:     getString("JWT_SECRET", "your-secret-key-change-in-production"),
			Expiration: getDuration("JWT_EXPIRATION", 24*time.Hour),
			Issuer:     getString("JWT_ISSUER", "backoffice-service"),
		},
		App: AppConfig{
			Name:        getString("APP_NAME", "Backoffice Service"),
			Version:     getString("APP_VERSION", "1.0.0"),
			Environment: getString("APP_ENV", "development"),
			Debug:       getBool("APP_DEBUG", true),
		},
		Logging: LoggingConfig{
			Channel:     getString("LOG_CHANNEL", "stdout"),
			Level:       getString("LOG_LEVEL", "debug"),
			LogPath:     getString("LOG_FILE_PATH", "./storage/logs"),
			LogFileName: getString("LOG_FILE_NAME", "app"),
			MaxSize:     getInt("LOG_MAX_SIZE", 10),
			MaxBackups:  getInt("LOG_MAX_BACKUPS", 5),
			MaxAge:      getInt("LOG_MAX_AGE", 28),
			Compress:    getBool("LOG_COMPRESS", true),
			DailyRotate: getBool("LOG_DAILY_ROTATE", true),
		},
	}
	
	return cfg, nil
}

// GetDatabaseDriverConfig converts DatabaseConnectionConfig to appropriate driver config
func (dbc *DatabaseConnectionConfig) GetDatabaseDriverConfig() (database.DriverType, interface{}, error) {
	driverType := database.DriverType(dbc.Driver)
	
	switch driverType {
	case database.DriverPostgreSQL:
		return driverType, &database.PostgresConfig{
			Host:            dbc.Host,
			Port:            dbc.Port,
			User:            dbc.User,
			Password:        dbc.Password,
			DBName:          dbc.DBName,
			SSLMode:         dbc.SSLMode,
			MaxOpenConns:    dbc.MaxOpenConns,
			MaxIdleConns:    dbc.MaxIdleConns,
			ConnMaxLifetime: dbc.ConnMaxLifetime,
			ConnMaxIdleTime: dbc.ConnMaxIdleTime,
			UseGorm:         dbc.UseGorm,
		}, nil
		
	case database.DriverMySQL:
		return driverType, &database.MySQLConfig{
			Host:            dbc.Host,
			Port:            dbc.Port,
			User:            dbc.User,
			Password:        dbc.Password,
			DBName:          dbc.DBName,
			Charset:         dbc.Charset,
			ParseTime:       true,
			Loc:             "Local",
			MaxOpenConns:    dbc.MaxOpenConns,
			MaxIdleConns:    dbc.MaxIdleConns,
			ConnMaxLifetime: dbc.ConnMaxLifetime,
			ConnMaxIdleTime: dbc.ConnMaxIdleTime,
			UseGorm:         dbc.UseGorm,
		}, nil
		
	default:
		return "", nil, database.ErrUnsupportedDriver
	}
}

// Helper functions
func setDefaults() {
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("GIN_MODE", "debug")
	viper.SetDefault("DB_DRIVER", "postgresql")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("APP_ENV", "development")
}

func getString(key, defaultValue string) string {
	if value := viper.GetString(key); value != "" {
		return value
	}
	return defaultValue
}

func getInt(key string, defaultValue int) int {
	if value := viper.GetInt(key); value != 0 {
		return value
	}
	return defaultValue
}

func getBool(key string, defaultValue bool) bool {
	if viper.IsSet(key) {
		return viper.GetBool(key)
	}
	return defaultValue
}

func getDuration(key string, defaultValue time.Duration) time.Duration {
	if value := viper.GetDuration(key); value != 0 {
		return value
	}
	return defaultValue
}
