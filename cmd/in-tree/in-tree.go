/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/Yuggupta27/migrate-to-ceph-csi/cmd"
	"github.com/spf13/cobra"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeConfig              string
	sourceStorageClass      string
	destinationStorageClass string
	rookNamespace           string
	cephClusterNamespace    string
)

// rootCmd represents the base command when called without any subcommands
var intree = &cobra.Command{
	Use:   "intree",
	Short: "Tool to migrate kubernetes ceph in-tree to CSI",
	Long:  `Tool to migrate kubernetes ceph in-tree to CSI`,
	Run: func(cmd *cobra.Command, args []string) {
		intreeToCSI()
	},
}

func init() {
	cmd.RootCmd.AddCommand(intree)
}

func intreeToCSI() {
	client, err := NewClient(kubeConfig)
	if err != nil {
		fmt.Printf(fmt.Sprintf("failed to create kubernetes client %v\n", err))
	}
	fmt.Print(&client)

	fmt.Printf("successfully migrated all the PVC from FlexVolume to CSI")
}

// NewClient create kubernetes client.
func NewClient(configPath string) (*k8s.Clientset, error) {
	var cfg *rest.Config
	var err error
	if configPath == "" {
		configPath = os.Getenv("KUBECONFIG")
		fmt.Print(configPath)
	}
	if configPath != "" {
		cfg, err = clientcmd.BuildConfigFromFlags("", configPath)
		if err != nil {
			return nil, fmt.Errorf("Failed to get cluster config with error: %v\n", err)
		}
	} else {
		cfg, err = rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("Failed to get cluster config with error: %v\n", err)
		}
		fmt.Print(cfg)
	}
	client, err := k8s.NewForConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("Failed to create client with error: %v\n", err)
	}
	return client, nil
}
