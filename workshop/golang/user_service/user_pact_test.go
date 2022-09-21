//go:build provider

package user_service_test

import (
	"demo/internal/web"
	"demo/user_service"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)

// The Provider verification
func TestPactProvider(t *testing.T) {
	go startInstrumentedProvider()

	pact := createPact()

	// Verify the Provider - Tag-based Published Pacts for any known consumers
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://127.0.0.1:%d", port),
		Tags:            []string{"master"},
		PactURLs:        []string{filepath.FromSlash(fmt.Sprintf("%s/login-service-user-service.json", pactDir))},
		ProviderVersion: "1.0.0",
	})

	if err != nil {
		t.Error(err)
	}
}

// Starts the provider API with hooks for provider states.
// This essentially mirrors the main.go file, with extra routes added.
func startInstrumentedProvider() {
	fmt.Print(pactDir, logDir);
	router := web.NewRouter()
	router.GET("/users/:id", user_service.GetAccountByIdHandler())
	web.ServeHttp(fmt.Sprintf(":%v", port), "user_service", router)
}

// Configuration / Test Data
var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../pacts", dir)
var logDir = fmt.Sprintf("%s/../logs", dir)
var port, _ = utils.GetFreePort()

// Setup the Pact client.
func createPact() dsl.Pact {
	return dsl.Pact{
		Provider: "user-service",
		PactDir:  pactDir,
		LogDir:   logDir,
		LogLevel: "INFO",
		DisableToolValidityCheck: true,
	}
}
