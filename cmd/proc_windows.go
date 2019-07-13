// Copyright 2015 Matthew Holt and The Caddy Authors
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

package caddycmd

import (
    "fmt"
    "os/exec"
    "strconv"
)

func gracefullyStopProcess(pid int) error {
    return tryStopProcess(pid, false)
}

func tryStopProcess(pid int, force bool) error {
    extraparam := ""
    if force {
        extraparam = "/f"
    }
    cmd := exec.Command("taskkill", "/pid", strconv.Itoa(pid), extraparam)
    if err := cmd.Run(); err != nil {
        // if taskkill fails try again to force.
        if err.Error() == "exit status 1" && !force {
            trygracefullyStopProcess(pid, true)
        } else {
            return fmt.Errorf("taskkill: %v", err)
        }
    }
    return nil
}
