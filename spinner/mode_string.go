// Code generated by "enumer -type=Mode -yaml"; DO NOT EDIT

package Spinner

import (
	"fmt"
)

const _Mode_name = "LEFT_TO_RIGHTUP_TO_DOWNFIX"

var _Mode_index = [...]uint8{0, 13, 23, 26}

func (i Mode) String() string {
	if i >= Mode(len(_Mode_index)-1) {
		return fmt.Sprintf("Mode(%d)", i)
	}
	return _Mode_name[_Mode_index[i]:_Mode_index[i+1]]
}

var _ModeNameToValue_map = map[string]Mode{
	_Mode_name[0:13]:  0,
	_Mode_name[13:23]: 1,
	_Mode_name[23:26]: 2,
}

func ModeString(s string) (Mode, error) {
	if val, ok := _ModeNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Mode values", s)
}

func (i Mode) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Mode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = ModeString(s)
	return err
}
