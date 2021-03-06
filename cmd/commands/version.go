/*
 * Copyright 2018 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	versionNumberOfArgs = iota
)

var cli_version = "unknown"
var cli_gitsha = "unknown sha"
var cli_gitdirty = ""

func Version() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information about riff",
		Args: cobra.ExactArgs(versionNumberOfArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			if cli_gitdirty != "" {
				cli_gitdirty = ", with local modifications"
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Version\n  riff cli: %s (%s%s)\n", cli_version, cli_gitsha, cli_gitdirty)
			return nil
		},
	}
}
