package main

import (
	"fmt"
	"log"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

func main() {
	mainLogger := setupLogger()

	specificLogger := withProjectOrgLogger(mainLogger, "GetOrganization", "project-id-1", "org-id-3")
	specificLogger.Info("event received")

	fmt.Println("DONE")
}

func withProjectOrgLogger(logger logr.Logger, event string, projectId string, organizationId string, keysAndValues ...any) logr.Logger {
	const logKeyOrganization = "organization-id"
	return withProjectLogger(logger, event, projectId, logKeyOrganization, organizationId, keysAndValues)
}

func withProjectLogger(logger logr.Logger, event string, projectId string, keysAndValues ...any) logr.Logger {
	const logKeyProject = "project-id"
	return logger.WithValues("event", event, logKeyProject, projectId).WithValues(keysAndValues...)
}

func setupLogger() logr.Logger {
	zapLog, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create zap logger: %v", err)
	}

	return zapr.NewLogger(zapLog)
}
