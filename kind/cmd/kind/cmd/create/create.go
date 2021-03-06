/*
Copyright 2018 The Kubernetes Authors.

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

// Package create implements the `create` command
package create

import (
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"k8s.io/test-infra/kind/pkg/cluster"
)

type flags struct {
	Name string
}

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand() *cobra.Command {
	flags := &flags{}
	cmd := &cobra.Command{
		// TODO(bentheelder): more detailed usage
		Use:   "create",
		Short: "Creates a cluster",
		Long:  "Creates a Kubernetes cluster",
		Run: func(cmd *cobra.Command, args []string) {
			run(flags, cmd, args)
		},
	}
	cmd.Flags().StringVar(&flags.Name, "name", "", "the cluster name")
	return cmd
}

func run(flags *flags, cmd *cobra.Command, args []string) {
	// TODO(bentheelder): make this more configurable
	config := cluster.NewConfig(flags.Name)
	ctx := cluster.NewContext(config)
	err := ctx.Create()
	if err != nil {
		glog.Errorf("Failed to create cluster: %v", err)
		os.Exit(-1)
	}
}
