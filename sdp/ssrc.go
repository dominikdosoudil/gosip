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

type Ssrc struct {
	Id  			string  	// 
	Attribute    	string  	// 
	Value   		string  	// 
}

func (ssrc *Ssrc) Append(b *bytes.Buffer) {
	b.WriteString("a=")
	b.WriteString("ssrc")
	b.WriteString(":")
	b.WriteString(ssrc.Id)
	b.WriteString(" ")
	b.WriteString(ssrc.Attribute)
	if ssrc.Value != "" {
		b.WriteString(":")
		b.WriteString(ssrc.Value)
	}
	b.WriteString("\r\n")

}
