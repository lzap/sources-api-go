package statuslistener

import (
	"os"
	"testing"

	"github.com/RedHatInsights/sources-api-go/internal/testutils"
)

// runningIntegration is used to skip integration tests if we're just running unit tests.
var runningIntegration = false

func TestMain(t *testing.M) {
	flags := testutils.ParseFlags()

	if flags.CreateDb {
		testutils.CreateTestDB()
	} else if flags.Integration {
		runningIntegration = true
		testutils.ConnectAndMigrateDB("status_listener")
		testutils.CreateFixtures()
	}

	code := t.Run()

	if flags.Integration {
		testutils.DropSchema("status_listener")
	}

	os.Exit(code)
}
