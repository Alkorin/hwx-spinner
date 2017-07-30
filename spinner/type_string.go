// Code generated by "enumer -type=Type -yaml"; DO NOT EDIT

package Spinner

import (
	"fmt"
)

const _Type_name = "SPEEDVOLTVERSION"

var _Type_index = [...]uint8{0, 5, 9, 16}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}

var _TypeNameToValue_map = map[string]Type{
	_Type_name[0:5]:  0,
	_Type_name[5:9]:  1,
	_Type_name[9:16]: 2,
}

func TypeString(s string) (Type, error) {
	if val, ok := _TypeNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Type values", s)
}

func (i Type) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Type) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = TypeString(s)
	return err
}
