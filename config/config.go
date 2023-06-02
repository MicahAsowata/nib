package config

import "os"

var Secret = os.Getenv("SECRET")
