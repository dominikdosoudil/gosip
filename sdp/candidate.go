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
	"strconv"
)

type Candidate struct {
	Foundation  		string 	// 
	ComponentId   		string 	// 
	Transport 			string 	//
	Priority 			int	//
 	ConnectionAddress 	string  //
 	Port 				int    //
	CandType 			string  //
	RelAddr         	string  //
	RelPort         	int     //
}

func (candidate *Candidate) Append(b *bytes.Buffer) {
	b.WriteString("a=")
	b.WriteString("candidate")
	b.WriteString(":")
	b.WriteString(candidate.Foundation)
	b.WriteString(" ")
	b.WriteString(candidate.ComponentId)
	b.WriteString(" ")
	b.WriteString(candidate.Transport)
	b.WriteString(" ")
	b.WriteString(strconv.FormatInt(int64(candidate.Priority), 10))
	b.WriteString(" ")
	b.WriteString(candidate.ConnectionAddress)
	b.WriteString(" ")
	b.WriteString(strconv.FormatInt(int64(candidate.Port), 10))
	if candidate.CandType == "" {
	} else {
		b.WriteString(" typ ")
		b.WriteString(candidate.CandType)
	}
	if candidate.RelAddr == "" {
	} else {
		b.WriteString(" raddr ")
		b.WriteString(candidate.RelAddr)
	}
	if candidate.RelPort == 0 {
	} else {
		b.WriteString(" rport ")
		b.WriteString(strconv.FormatInt(int64(candidate.RelPort), 10))
	}
	b.WriteString("\r\n")
}
