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

package storage

import (
	"path/filepath"

	"github.com/mysteriumnetwork/node/core/promise/storage/boltdb"
)

// Storage stores persistent objects for future usage
type Storage interface {
	Store(issuer string, data interface{}) error
	Delete(issuer string, data interface{}) error
	Close() error
}

// NewStorage creates a new BoltDB storage for service promises
func NewStorage(path string) (Storage, error) {
	return boltdb.OpenDB(filepath.Join(path, "myst.db"))
}