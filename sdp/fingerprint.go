// Copyright 2020 Justine Alexandra Roberts Tunney
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

package sdp

import (
	"bytes"
)

type Fingerprint struct {
	HashFunc  string  // "sha-1", "sha-224", "sha-256" etc.
	Fingerprint   string  // Fingerprint string
}

func (fingerprint *Fingerprint) Append(b *bytes.Buffer) {
	b.WriteString("a=")
	b.WriteString("fingerprint")
	b.WriteString(":")
	b.WriteString(fingerprint.HashFunc)
	b.WriteString(" ")
	b.WriteString(fingerprint.Fingerprint)
	b.WriteString("\r\n")

}
