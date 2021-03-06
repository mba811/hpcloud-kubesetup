// Copyright (c) 2014 Hewlett-Packard Development Company, L.P.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package compute_test

import (
	"testing"

	compute "git.openstack.org/stackforge/golang-client.git/compute/v2"
	identity "git.openstack.org/stackforge/golang-client.git/identity/v2"
)

func TestKeyPairScenarios(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	authenticator, err := identity.AuthenticateFromEnvVars()
	if err != nil {
		t.Fatal("Cannot authenticate from env vars:", err)
	}

	computeService := compute.NewService(authenticator)

	createdKeypair, err := computeService.CreateKeyPair("testkeyPairName", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDx8nkQv/zgGgB4rMYmIf+6A4l6Rr+o/6lHBQdW5aYd44bd8JttDCE/F/pNRr0lRE+PiqSPO8nDPHw0010JeMH9gYgnnFlyY3/OcJ02RhIPyyxYpv9FhY+2YiUkpwFOcLImyrxEsYXpD/0d3ac30bNH6Sw9JD9UZHYcpSxsIbECHw== Generated by Nova")
	if err != nil {
		t.Fatal("Cannot create keypair:", err)
	}

	queriedKeyPair, err := computeService.KeyPair(createdKeypair.Name)
	if err != nil {
		t.Fatal("Cannot requery keypair:", err)
	}

	keyPairs, err := computeService.KeyPairs()
	if err != nil {
		t.Fatal("Cannot access keypairs:", err)
	}

	foundKeyPair := false
	for _, keyPairValue := range keyPairs {
		if queriedKeyPair.Name == keyPairValue.Name {
			foundKeyPair = true
			break
		}
	}

	if !foundKeyPair {
		t.Fatal("Cannot find keypair that was created.")
	}

	err = computeService.DeleteKeyPair(queriedKeyPair.Name)
	if err != nil {
		t.Fatal("Cannot delete keypair:", err)
	}
}
