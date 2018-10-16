/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
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

package location

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationCacheFirstCall(t *testing.T) {
	locationDetector := NewDetectorFake("100.100.100.100", "country")
	locationCache := NewLocationCache(locationDetector)

	location := locationCache.Get()
	assert.Equal(t, Location{}, location)
}

func TestLocationCacheFirstSecondCalls(t *testing.T) {
	locationDetector := NewDetectorFake("100.100.100.100", "country")
	locationCache := NewLocationCache(locationDetector)

	location, err := locationCache.RefreshAndGet()
	assert.Equal(t, "country", location.Country)
	assert.Equal(t, "100.100.100.100", location.IP)
	assert.NoError(t, err)

	locationSecondCall := locationCache.Get()
	assert.Equal(t, location, locationSecondCall)
}

func TestLocationCacheWithError(t *testing.T) {
	locationErr := errors.New("location DbResolver error")
	locationDetector := NewDetectorFakeFailing(locationErr)
	locationCache := NewLocationCache(locationDetector)

	location, err := locationCache.RefreshAndGet()
	assert.EqualError(t, locationErr, err.Error())
	assert.Equal(t, Location{}, location)
}
