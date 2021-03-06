// Copyright (c) 2018 Baidu, Inc.
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

package web_monitor

import (
	"testing"
)

func TestIsValidForReload(t *testing.T) {
	if !isValidForReload("[::1]:8080") {
		t.Error("err in valid for reload, [::1]:8080 should allow reload")
	}
	if !isValidForReload("127.0.0.1:8080") {
		t.Error("err in valid for reload, 127.0.0.1:8080 should allow reload")
	}
}

func TestInitReloadACL(t *testing.T) {
	err := InitReloadACL("reload_src_conf/testdata/reload_src_conf_1.data")
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(RELOAD_SRC_ALLOWED) != 5 {
		t.Fatal("len(RELOAD_SRC_ALLOWED) != 5")
	}
}
