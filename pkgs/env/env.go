package env

import (
	"log"
	"os"
)

func MustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("must set %s", key)
	}

	return v
}
