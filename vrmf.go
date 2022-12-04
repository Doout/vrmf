package vrmf

import (
	"fmt"
	"strings"
)

type VRMF struct {
	Version      int // 1-99
	Release      int // 1-99
	Modification int // 0-9999
	FixPack      int // 0-9999
	fix          string
}

func Parse(v string) (*VRMF, error) {
	//The format of VRMF is: <Version>.<Release>.<Modification>.<FixPack>_<iFix>
	// iFix is optional
	vrmf := &VRMF{}
	//Check if we have an iFix
	if strings.LastIndex(v, "_") != -1 {
		vrmf.fix = v[strings.LastIndex(v, "_"):]
		v = v[:strings.LastIndex(v, "_")]
	}
	_, err := fmt.Sscanf(v, "%d.%d.%d.%d", &vrmf.Version, &vrmf.Release, &vrmf.Modification, &vrmf.FixPack)
	if err != nil {
		return nil, err
	}
	return vrmf, nil
}

func (v *VRMF) String() string {
	return fmt.Sprintf("%d.%d.%d.%d%s", v.Version, v.Release, v.Modification, v.FixPack, v.fix)
}

func (v *VRMF) Compare(v2 *VRMF) int {
	if v.Version != v2.Version {
		return v.Version - v2.Version
	}
	if v.Release != v2.Release {
		return v.Release - v2.Release
	}
	if v.Modification != v2.Modification {
		return v.Modification - v2.Modification
	}
	if v.FixPack != v2.FixPack {
		return v.FixPack - v2.FixPack
	}
	return 0
}

func (v *VRMF) IsEqual(v2 *VRMF) bool {
	return v.Compare(v2) == 0
}

func (v *VRMF) IsLessThan(v2 *VRMF) bool {
	return v.Compare(v2) < 0
}

func (v *VRMF) IsGreaterThan(v2 *VRMF) bool {
	return v.Compare(v2) > 0
}

func (v *VRMF) IsGreaterThanOrEqual(v2 *VRMF) bool {
	return v.Compare(v2) >= 0
}

func (v *VRMF) IsLessThanOrEqual(v2 *VRMF) bool {
	return v.Compare(v2) <= 0
}

func (v *VRMF) IsInRangeExclusive(v2, v3 *VRMF) bool {
	return v.IsGreaterThan(v2) && v.IsLessThan(v3)
}

func (v *VRMF) IsInRangeInclusive(v2, v3 *VRMF) bool {
	return v.IsGreaterThanOrEqual(v2) && v.IsLessThanOrEqual(v3)
}
