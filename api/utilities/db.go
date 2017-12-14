package utilities

// User for the production/development database
var DB_USER = GetVar("DB_USER", "gradeportal")

// Port for the production/development database
var DB_PORT = GetVar("DB_PORT", "5432")

// Passowrd for the production/development database
var DB_PASSWORD = GetVar("DB_PASSWORD", "cs130gp")

// Host for the production/development database
var DB_HOST = GetVar("DB_HOST", "localhost")

// Name for the production/development database
var DB_NAME = GetVar("DB_NAME", "gp_development")
