/*
Copyright 2023 The Bestchains Authors.

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

package common

import (
	"os"
	"path/filepath"

	"k8s.io/cli-runtime/pkg/genericclioptions"
)

const (
	CreateDepository          = "/basic/putValue"
	CreateUntrustedDepository = "/basic/putUntrustValue"
	GetDepository             = "/basic/depositories/%s"
	ListDepository            = "/basic/depositories"
	CurrentNonce              = "/basic/currentNonce"
	WalletHomeDir             = ".bestchains/wallet"
)

var (
	WalletConfigDir = filepath.Join(os.Getenv("HOME"), WalletHomeDir)
)

type WalletConfig struct {
	Address    string `json:"address"`
	PrivateKey []byte `json:"privKey"`
}

type Options struct {
	genericclioptions.IOStreams
}
