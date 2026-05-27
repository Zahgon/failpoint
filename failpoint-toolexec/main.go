// Copyright 2024 PingCAP, Inc.
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

package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pingcap/failpoint/code"
)

var logger = log.New(os.Stderr, "[failpoint-toolexec]", log.LstdFlags)

func main() {
	if len(os.Args) < 2 {
		return
	}
	goCmd, buildArgs := os.Args[1], os.Args[2:]
	goCmdBase := filepath.Base(goCmd)
	if runtime.GOOS == "windows" {
		goCmdBase = strings.TrimSuffix(goCmdBase, ".exe")
	}

	if strings.ToLower(goCmdBase) == "compile" {
		if err := injectFailpoint(&buildArgs); err != nil {
			logger.Println("failed to inject failpoint", err)
		}
	}

	cmd := exec.Command(goCmd, buildArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		logger.Println("failed to run command", err)
	}
}

func injectFailpoint(argsP *[]string) error { _ = "STUB: not implemented"; return nil }

// ref https://pkg.go.dev/cmd/compile#hdr-Command_Line

// find the golang source files of the caller's package

// ref https://github.com/golang/go/blob/bdd27c4debfb51fe42df0c0532c1c747777b7a32/src/cmd/go/internal/modload/init.go#L1511
func findCallersModule() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Look for enclosing go.mod.

var goModCache = os.Getenv("GOMODCACHE")
var goRoot = runtime.GOROOT()

func inSDKOrMod(path string) bool { _ = "STUB: not implemented"; return false }

var tmpFolder = filepath.Join(os.TempDir(), "failpoint-toolexec")

func injectFailpointForFile(w *code.Rewriter, file *string, module string) bool {
	_ = "STUB: not implemented"
	return false
}

func writeExtraFile(filePath, packageName, module string) error {
	_ = "STUB: not implemented"
	return nil
}
