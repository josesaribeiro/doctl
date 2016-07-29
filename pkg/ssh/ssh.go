/*
Copyright 2016 The Doctl Authors All rights reserved.
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

package ssh

import (
	"runtime"

	"github.com/digitalocean/doctl/pkg/runner"
)

// Options is the type used to specify options passed to the SSH command
type Options map[string]interface{}

// Runner runs ssh commands.
type Runner struct {
	User            string
	Host            string
	KeyPath         string
	Port            int
	AgentForwarding bool
}

var _ runner.Runner = &Runner{}

// Run ssh.
func (r *Runner) Run() error {
	if runtime.GOOS == "windows" {
		return runInternalSSH(r)
	}

	return runExternalSSH(r)
}
