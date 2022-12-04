package vrmf

import "testing"

func TestParse(t *testing.T) {
	test := "1.1.1.1_test"
	vrmf, err := Parse(test)
	if err != nil {
		t.Fatalf("Failed to parse %s", test)
	}
	if vrmf.Version != 1 || vrmf.Release != 1 || vrmf.Modification != 1 || vrmf.FixPack != 1 {
		t.Fatalf("Failed to parse %s", test)
	}
	if vrmf.fix != "_test" {
		t.Fatalf("Failed to parse %s", test)
	}
}

func TestCompare(t *testing.T) {
	v := []string{"1.1.1.1", "1.1.1.1_iFix0001"}
	vrmfs := make([]*VRMF, len(v))
	for i, s := range v {
		vrmf, err := Parse(s)
		if err != nil {
			t.Fatalf("Failed to parse %s (%d, %s)", s, i, s)
		}
		vrmfs[i] = vrmf
	}
	if !vrmfs[0].IsEqual(vrmfs[1]) {
		t.Fatalf("Failed to compare %s and %s", v[0], v[1])
	}
	if vrmfs[0].IsLessThan(vrmfs[1]) {
		t.Fatalf("Failed to compare %s and %s", v[0], v[1])
	}
	if vrmfs[0].IsGreaterThan(vrmfs[1]) {
		t.Fatalf("Failed to compare %s and %s", v[0], v[1])
	}
}

func TestIsInRangeInclusive(t *testing.T) {
	v := []string{"11.5.6.0", "11.5.0.0", "11.5.9.0"}
	vrmfs := make([]*VRMF, len(v))
	for i, s := range v {
		vrmf, err := Parse(s)
		if err != nil {
			t.Fatalf("Failed to parse %s", s)
		}
		vrmfs[i] = vrmf
	}
	if !vrmfs[0].IsInRangeInclusive(vrmfs[1], vrmfs[2]) {
		t.Fatalf("Failed %s IsInRangeInclusive ( %s , %s )", v[0], v[1], v[2])
	}
}

func TestIsInRangeExclusive(t *testing.T) {
	v := []string{"10.2.0.0", "11.5.0.0", "11.5.9.0"}
	vrmfs := make([]*VRMF, len(v))
	for i, s := range v {
		vrmf, err := Parse(s)
		if err != nil {
			t.Fatalf("Failed to parse %s", s)
		}
		vrmfs[i] = vrmf
	}
	if vrmfs[0].IsInRangeExclusive(vrmfs[1], vrmfs[2]) {
		t.Fatalf("Failed %s IsInRangeExclusive ( %s , %s )", v[0], v[1], v[2])
	}
}
