/*
Copyright 2019 The Jetstack cert-manager contributors.

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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/joshvanl/cert-manager-example-issuer/pkg/controller"
	"github.com/joshvanl/cert-manager-example-issuer/pkg/issuer"
)

func NewCommandStartCertManagerController(stopCh <-chan struct{}) *cobra.Command {
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("cert-manager-%s-controller", controller.ControllerName),
		Short: fmt.Sprintf("External cert-manager issuer cointroller for group %s", issuer.GroupName),
		Long: fmt.Sprintf(`
cert-manager-%s-controller is an issuer controller used to issue certificates
based of cert-manager's CertificateRequests. This is to be used in conjunction
with cert-manager itself. This controller will listen to CertificateRequests
that have issuer references in the %s Issuer Group.`,
			controller.ControllerName, issuer.GroupName),

		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
