//go:build consumer

package login_service_test

import (
	"demo/login_service"
	"fmt"
	"os"
	"testing"

	"net/url"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/stretchr/testify/assert"
)

var commonHeaders = dsl.MapMatcher{
	"Content-Type": term("application/json; charset=utf-8", `application\/json`),
}

var u *url.URL
var client *login_service.Gateway

func TestMain(m *testing.M) {
	var exitCode int

	// Setup Pact and related test stuff
	setup()

	// Run all the tests
	exitCode = m.Run()

	// Shutdown the Mock Service and Write pact files to disk
	if err := pact.WritePact(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pact.Teardown()
	os.Exit(exitCode)
}

func TestClientPact_GetUserById(t *testing.T) {
	t.Run("user exists", func(t *testing.T) {
		pact.
			AddInteraction().
			Given("User somkiat exists").
			UponReceiving("A request to get user id = 1 as Somkiat").
			WithRequest(request{
				Method: "GET",
				Path:   term("/users/1", "/users/[0-9]+"),
			}).
			WillRespondWith(dsl.Response{
				Status:  200,
				Body:    dsl.Match(login_service.User{}),
				Headers: commonHeaders,
			})
		id := 1
		err := pact.Verify(func() error {
			user, err := client.GetUser(id)

			// Assert basic fact
			if user.User != "somkiat" {
				return fmt.Errorf("wanted user somkiat but got %v", user.User)
			}

			return err
		})

		if err != nil {
			t.Fatalf("Error on Verify: %v", err)
		}
	})

	t.Run("the user does not exist", func(t *testing.T) {
		pact.
			AddInteraction().
			Given("User does not exist").
			UponReceiving("A request to get user id = 2").
			WithRequest(request{
				Method: "GET",
				Path:   term("/users/2", "/users/[0-9]+"),
			}).
			WillRespondWith(dsl.Response{
				Status:  404,
			})

		err := pact.Verify(func() error {
			_, err := client.GetUser(2)

			return err
		})

		assert.Equal(t, login_service.ErrNotFound, err)
	})

}

// Common test data
var pact dsl.Pact

// Aliases
var term = dsl.Term

type request = dsl.Request

func setup() {
	pact = createPact()

	// Proactively start service to get access to the port
	pact.Setup(true)

	u, _ = url.Parse(fmt.Sprintf("http://localhost:%d", pact.Server.Port))

	client = &login_service.Gateway{
		BaseURL: u,
	}
}

func createPact() dsl.Pact {
	return dsl.Pact{
		Consumer:                 "login-service",
		Provider:                 "user-service",
		LogDir:                   "../logs",
		PactDir:                  "../pacts",
		LogLevel:                 "INFO",
		DisableToolValidityCheck: true,
	}
}
