/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package discovery

import (
	"sync"

	"github.com/mysteriumnetwork/node/identity"
	identity_registry "github.com/mysteriumnetwork/node/identity/registry"
	"github.com/mysteriumnetwork/node/server"
	dto_discovery "github.com/mysteriumnetwork/node/service_discovery/dto"
)

// Discovery structure holds discovery service state
type Discovery struct {
	identityRegistry            identity_registry.IdentityRegistry
	ownIdentity                 identity.Identity
	identityRegistration        identity_registry.RegistrationDataProvider
	mysteriumClient             server.Client
	signerCreate                identity.SignerFactory
	signer                      identity.Signer
	proposal                    dto_discovery.ServiceProposal
	statusChan                  chan Status
	status                      Status
	proposalAnnouncementStopped *sync.WaitGroup
	unsubscribe                 func()
	stop                        func()

	sync.RWMutex
}

// NewService creates new discovery service
func NewService(
	identityRegistry identity_registry.IdentityRegistry,
	identityRegistration identity_registry.RegistrationDataProvider,
	mysteriumClient server.Client,
	signerCreate identity.SignerFactory,
) *Discovery {
	return &Discovery{
		identityRegistry:     identityRegistry,
		identityRegistration: identityRegistration,
		mysteriumClient:      mysteriumClient,
		signerCreate:         signerCreate,
		statusChan:           make(chan Status),
		status:               StatusUndefined,
		proposalAnnouncementStopped: &sync.WaitGroup{},
		unsubscribe:                 func() {},
		stop:                        func() {},
		RWMutex:                     sync.RWMutex{},
	}
}
