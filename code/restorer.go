// Copyright 2019 PingCAP, Inc.
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

package code

const (
	failpointStashFileSuffix = "__failpoint_stash__"
	failpointBindingFileName = "binding__failpoint_binding__.go"
)

// Restorer represents a manager to restore currentFile tree which has been modified by
// `failpoint-ctl enable`, e.g:
/*
// ├── foo
// │   ├── foo.go
// │   └── foo.go__failpoint_stash__
// ├── bar
// │   ├── bar.go
// │   └── bar.go__failpoint_stash__
// └── foobar
//     ├── foobar.go
//     └── foobar.go__failpoint_stash__
// Which will be restored as below:
// ├── foo
// │   └── foo.go <- foo.go__failpoint_stash__
// ├── bar
// │   └── bar.go <- bar.go__failpoint_stash__
// └── foobar
//     └── foobar.go <- foobar.go__failpoint_stash__
*/
type Restorer struct {
	path string
}

// NewRestorer returns a non-nil restorer which is used to clean the workspace
// of the specified path
func NewRestorer(path string) *Restorer { _ = "STUB: not implemented"; return nil }

// Restore restores the currentFile tree which will delete all files generated
// by `failpoint-ctl enable` and replace it by fail point stashed currentFile
func (r Restorer) Restore() error { _ = "STUB: not implemented"; return nil }

// Rewrite original file

// Merge modifications after `failpoint-ctl enable`

func failpointBindingPath(path string) string { _ = "STUB: not implemented"; return "" }

func isBindingFileExists(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func writeBindingFile(path, pak string) error { _ = "STUB: not implemented"; return nil }
