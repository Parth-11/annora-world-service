package config

import (
	"net"
	"os"
	"strconv"
	"time"
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}

func mustGetEnv(key string) string {
	v := os.Getenv(key)

	if v == "" {
		panic("missing env: " + key)
	}

	return v
}

func getDuration(key string, fallback time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		dur, err := time.ParseDuration(v)
		if err != nil {
			panic("Invalid duration for " + key)
		}

		return dur
	}

	return fallback
}

func getInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic("Invalid int for " + key)
		}

		return val
	}

	return fallback
}

func getNetAddr(key string, fallback net.Addr) net.Addr {
	if v := os.Getenv(key); v != "" {
		addr, err := net.ResolveTCPAddr("tcp", v)
		if err != nil {
			panic("Invalid net addr for " + key)
		}

		return addr
	}

	return fallback
}
