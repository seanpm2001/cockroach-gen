// Code generated by "stringer -type=Attribute -trimprefix=Attribute"; DO NOT EDIT.

package scpb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AttributeType-0]
	_ = x[AttributeDescID-1]
	_ = x[AttributeDepID-2]
	_ = x[AttributeColumnID-3]
	_ = x[AttributeElementName-4]
	_ = x[AttributeIndexID-5]
}

const _Attribute_name = "TypeDescIDDepIDColumnIDElementNameIndexID"

var _Attribute_index = [...]uint8{0, 4, 10, 15, 23, 34, 41}

func (i Attribute) String() string {
	if i < 0 || i >= Attribute(len(_Attribute_index)-1) {
		return "Attribute(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Attribute_name[_Attribute_index[i]:_Attribute_index[i+1]]
}
