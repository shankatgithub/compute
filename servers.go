// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package compute

import (
	p "github.com/gcloud/compute/providers"
	"github.com/gcloud/identity"
)

// The Servers type interacts with Compute services.
type Servers struct {
	account  identity.Account
	Provider string
}

// List servers available on the account.
func (s *Servers) List() string {
	return p.Providers[s.Provider].Servers.List()
}

// Show server information for a given id.
func (s *Servers) Show(id string) {}

// Create a server.
func (s *Servers) Create() {}

// Destroy a server.
func (s *Servers) Destroy() {}

// Reboot a server.
func (s *Servers) Reboot() {}

// Start a server that is stopped.
func (s *Servers) Start() {}

// Stop a server that is running.
func (s *Servers) Stop() {}
