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

// Digital Signal Processing

package dsp

// Mixes together two audio frames containing 160 samples. Uses saturation
// adding so you don't hear clipping if ppl talk loud. Uses 128-bit SIMD
// instructions so we can add eight numbers at the same time.
// func L16MixSat160(dst, src *int16)

// Compresses a PCM audio sample into a G.711 μ-Law sample. The BSR instruction
// is what makes this code fast.
//
// TODO(jart): How do I make assembly use proper types?
// func LinearToUlaw(linear int64) (ulaw int64)
func LinearToUlaw(pcm_val int16) uint8 { // 2's complement (16-bit range)

	var seg int
	var mask uint8
	var uval uint8

	// Get the sign and the magnitude of the value.
	pcm_val = pcm_val >> 2
	if pcm_val < 0 {
		pcm_val = -pcm_val
		mask = 0x7F
	} else {
		mask = 0xFF
	}
	if pcm_val > _CLIP {
		pcm_val = _CLIP // clip the magnitude
	}
	pcm_val += (_BIAS >> 2)

	/* Convert the scaled magnitude to segment number. */
	seg = search(pcm_val, seg_uend[:])

	// Combine the sign, segment, quantization bits; and complement the code word.
	if seg >= 8 {
		// out of range, return maximum value.
		uval = uint8(0x7F)
		return (uval ^ mask)
	} else {
		uval = uint8((int16(seg) << 4) | ((pcm_val >> (seg + 1)) & 0xF))
		return (uval ^ mask)
	}
}

// Turns a μ-Law byte back into an audio sample.
// func UlawToLinear(ulaw int64) (linear int64)
func UlawToLinear(u_val uint8) int16 {

	var t int16

	// Complement to obtain normal u-law value.
	u_val = ^u_val

	// Extract and bias the quantization bits. Then shift up
	// by the segment number and subtract out the bias.

	t = (int16(u_val&quantMask) << 3) + _BIAS
	t <<= ((uint(u_val) & segMask) >> segShift)

	if (u_val & signBit) != 0 {
		return (_BIAS - t)
	} else {
		return (t - _BIAS)
	}
}
