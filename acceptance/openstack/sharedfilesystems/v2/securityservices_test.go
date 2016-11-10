package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
)

func TestSecurityServiceCreate(t *testing.T) {
	client, err := clients.NewSharedFileSystemV2Client()
	if err != nil {
		t.Fatalf("Unable to create shared file system client: %v", err)
	}

	securityService, err := CreateSecurityService(t, client)
	if err != nil {
		t.Fatalf("Unable to create security service: %v", err)
	}

	// TODO: Delete the security service once Delete is supported

	PrintSecurityService(t, securityService)
}
