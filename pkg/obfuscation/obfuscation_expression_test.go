//go:build unit || integration
// +build unit integration

package obfuscation

import (
	"log"
	"os"

	"github.com/newrelic/newrelic-client-go/v2/pkg/config"
)

func Example_basic() {
	// Initialize the client configuration.  A Personal API key is required to
	// communicate with the backend API.
	cfg := config.New()
	cfg.PersonalAPIKey = os.Getenv("NEW_RELIC_API_KEY")

	// Initialize the client.
	client := New(cfg)

	accountID := 12345678

	// List the obfuscation expression for a given account.
	_, err := client.GetList(accountID)
	if err != nil {
		log.Fatal("error listing NRQL drop rules: ", err)
	}

	// Create a new expression.
	createInput := LogConfigurationsCreateObfuscationExpressionInput{
		Description: "Brief desc",
		Name:        "some thing",
		Regex:       "log.*",
	}

	created, err := client.LogConfigurationsCreateObfuscationExpression(accountID, createInput)
	if err != nil {
		log.Fatal("error creating Obfuscation expression: ", err)
	}

	// Update an existing obfuscation expression.
	expressionName := "Example updated workload"
	updateInput := LogConfigurationsUpdateObfuscationExpressionInput{
		Name: workloadName,
	}

	updated, err := client.WorkloadUpdate(created.ID, updateInput)
	if err != nil {
		log.Fatal("error updating obfuscation expression: ", err)
	}

	// Delete an existing events to metrics rule.
	deleteInput := updateInput.ID
	_, err = client.LogConfigurationsDeleteObfuscationExpression(accountID, deleteInput)
	if err != nil {
		log.Fatal("error deleting obfuscation expression: ", err)
	}
}
