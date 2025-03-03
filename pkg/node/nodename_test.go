/*
Copyright 2023 k0s authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package node

import (
	"runtime"
	"testing"

	apitypes "k8s.io/apimachinery/pkg/types"
	nodeutil "k8s.io/component-helpers/node/util"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetNodeName(t *testing.T) {
	t.Run("should_always_return_override_if_given", func(t *testing.T) {
		name, err := GetNodeName("override")
		if assert.NoError(t, err) {
			assert.Equal(t, apitypes.NodeName("override"), name)
		}
	})

	if runtime.GOOS != "windows" {
		kubeHostname, err := nodeutil.GetHostname("")
		require.NoError(t, err)

		t.Run("should_call_kubernetes_hostname_helper_on_linux", func(t *testing.T) {
			name, err := GetNodeName("")
			if assert.NoError(t, err) {
				assert.Equal(t, apitypes.NodeName(kubeHostname), name)
			}
		})
	}
}
