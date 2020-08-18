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

type Group struct {
	Semantics  			string  	// 
	IdentificationTag   []string  	// 
}

func (group *Group) Append(b *bytes.Buffer) {
	b.WriteString("a=")
	b.WriteString("group")
	b.WriteString(":")
	b.WriteString(group.Semantics)
	for _, itag := range group.IdentificationTag {
		b.WriteString(" ")
		b.WriteString(itag)
	}
	b.WriteString("\r\n")

}