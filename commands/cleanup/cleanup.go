//
// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package cleanup

import (
	"fmt"

	"github.com/apache/skywalking-infra-e2e/internal/components/cleanup"

	"github.com/spf13/cobra"

	"github.com/apache/skywalking-infra-e2e/internal/constant"
	"github.com/apache/skywalking-infra-e2e/internal/logger"

	"github.com/apache/skywalking-infra-e2e/internal/flags"
)

func init() {
	Cleanup.Flags().StringVar(&flags.Env, "env", "kind", "specify test environment")
	Cleanup.Flags().StringVar(&flags.File, "file", "kind.yaml", "specify configuration file")
}

var Cleanup = &cobra.Command{
	Use:   "cleanup",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		if flags.Env == constant.Compose {
			logger.Log.Info("env for docker-compose not implemented")
		} else if flags.Env == constant.Kind {
			if err := cleanup.KindCleanupInCommand(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("no such env for cleanup: [%s]. should use kind or compose instead", flags.Env)
		}

		return nil
	},
}
