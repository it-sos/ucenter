package core

import "os"

func GetEnviron() string {
	env, ok := os.LookupEnv("GOENVIRON")
	if ok == false {
		env = "dev"
	}
	return env
}

func IsProudctEnv() bool {
	return GetEnviron() == "product"
}
