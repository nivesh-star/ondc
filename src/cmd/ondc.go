/*
 * Beckn Protocol Core
 *
 * Beckn Core Transaction API specification
 *
 * API version: 1.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log/slog"

	"github.com/nivesh-star/ondc/src/server"
)

func main() {
	slog.Info("Starting ONDC Backend engine")
	server := server.NewServer()
	server.Run()
}
