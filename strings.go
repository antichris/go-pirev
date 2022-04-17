// Code generated by "stringer -output=strings.go -type=Manufacturer,Processor,Model -linecomment"; DO NOT EDIT.

package pirev

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SonyUK-0]
	_ = x[Egoman-1]
	_ = x[Embest-2]
	_ = x[SonyJapan-3]
	_ = x[Embest2-4]
	_ = x[Stadium-5]
}

const _Manufacturer_name = "Sony UKEgomanEmbestSony JapanEmbestStadium"

var _Manufacturer_index = [...]uint8{0, 7, 13, 19, 29, 35, 42}

func (i Manufacturer) String() string {
	if i >= Manufacturer(len(_Manufacturer_index)-1) {
		return "Manufacturer(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Manufacturer_name[_Manufacturer_index[i]:_Manufacturer_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BCM2835-0]
	_ = x[BCM2836-1]
	_ = x[BCM2837-2]
	_ = x[BCM2711-3]
}

const _Processor_name = "BCM2835BCM2836BCM2837BCM2711"

var _Processor_index = [...]uint8{0, 7, 14, 21, 28}

func (i Processor) String() string {
	if i >= Processor(len(_Processor_index)-1) {
		return "Processor(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Processor_name[_Processor_index[i]:_Processor_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[A-0]
	_ = x[B-1]
	_ = x[APlus-2]
	_ = x[BPlus-3]
	_ = x[TwoB-4]
	_ = x[Alpha-5]
	_ = x[CM1-6]
	_ = x[_Unspecified0-7]
	_ = x[ThreeB-8]
	_ = x[Zero-9]
	_ = x[CM3-10]
	_ = x[_Unspecified1-11]
	_ = x[ZeroW-12]
	_ = x[ThreeBPlus-13]
	_ = x[ThreeAPlus-14]
	_ = x[Internal-15]
	_ = x[CM3Plus-16]
	_ = x[FourB-17]
	_ = x[Zero2W-18]
	_ = x[FourHundred-19]
	_ = x[CM4-20]
}

const _Model_name = "ABA+B+2BAlpha (early prototype)CM1_Unspecified03BZeroCM3_Unspecified1Zero W3B+3A+Internal use onlyCM3+4BZero 2 W400CM4"

var _Model_index = [...]uint8{0, 1, 2, 4, 6, 8, 31, 34, 47, 49, 53, 56, 69, 75, 78, 81, 98, 102, 104, 112, 115, 118}

func (i Model) String() string {
	if i >= Model(len(_Model_index)-1) {
		return "Model(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Model_name[_Model_index[i]:_Model_index[i+1]]
}
