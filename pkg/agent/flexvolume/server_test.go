/*
Copyright 2017 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Some of the code below came from https://github.com/coreos/etcd-operator
which also has the apache 2.0 license.
*/

// Package flexvolume to manage Kubernetes storage attach events.
package flexvolume

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigureFlexVolume(t *testing.T) {
	driverDir, _ := ioutil.TempDir("", "")
	defer os.RemoveAll(driverDir)
	driver, _ := os.OpenFile(path.Join(driverDir, "flexvolume"), os.O_RDONLY|os.O_CREATE, 0755)

	err := configureFlexVolume(driver.Name(), driverDir)
	assert.Nil(t, err)
	_, err = os.Stat(path.Join(driverDir, "rook"))
	assert.False(t, os.IsNotExist(err))

}
