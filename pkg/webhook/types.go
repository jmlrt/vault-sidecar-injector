// Copyright © 2019 Talend
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook

import (
	"net/http"
	"talend/vault-sidecar-injector/pkg/config"
)

var vaultInjectorAnnotationKeys = [...]string{
	vaultInjectorAnnotationInjectKey,
	vaultInjectorAnnotationAuthMethodKey,
	vaultInjectorAnnotationModeKey,
	vaultInjectorAnnotationProxyPortKey,
	vaultInjectorAnnotationRoleKey,
	vaultInjectorAnnotationSATokenKey,
	vaultInjectorAnnotationSecretsPathKey,
	vaultInjectorAnnotationSecretsTemplateKey,
	vaultInjectorAnnotationTemplateDestKey,
	vaultInjectorAnnotationLifecycleHookKey,
	vaultInjectorAnnotationTemplateCmdKey,
	vaultInjectorAnnotationWorkloadKey,
	vaultInjectorAnnotationStatusKey,
}

var vaultInjectorModes = [...]string{
	vaultInjectorModeSecrets,
	vaultInjectorModeProxy,
}

// VaultInjector : Webhook Server entity
type VaultInjector struct {
	*config.InjectionConfig
	Server *http.Server
}

// Struct to carry computed placeholders' values and context info for current injection
type sidecarContext struct {
	modes                          map[string]bool
	k8sDefaultSATokenVolumeName    string
	vaultInjectorSATokenVolumeName string
	vaultAuthMethod                string
	vaultRole                      string
	proxy                          string
	templates                      string
}

// This struct represents a JSON Patch operation (see http://jsonpatch.com/)
type patchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}
