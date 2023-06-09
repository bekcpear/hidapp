package hidapp

import (
	"fmt"
	"testing"
)

func ExampleProcessor_AppendRegexp() {
	processor := NewProcessor()
	_ = processor.AppendRegexp(
		"ghp_[[:graph:]]+",
		"glpat-[[:graph:]]+",
		"(?:ABC)(testg_[a-z]+)(?:BBB)([a-z]*)")
	fmt.Println(processor.Process("plaintextCCKglpat-djksaljkl"))
	fmt.Println(processor.Process("plaintexghp_jkdlqjklda"))
	fmt.Println(processor.Process("plaintextABCtestg_jdklsajklBBBdjklKK"))
	// Output:
	// plaintextCCK**********
	// plaintex**********
	// plaintextABC**********BBB**********KK
}

func TestProcess(t *testing.T) {
	processor := NewProcessor()
	err := processor.AppendRegexp(
		"ghp_[[:graph:]]+",
		"glpat-[[:graph:]]+",
		"(?:CCK)glpat-[[:graph:]]+",
		"(?:ABC)(testg_[a-z]+)(?:BBB)([a-z]*)")
	if err != nil {
		t.Fatal(err)
	}

	// TODO: more tests
	tc := map[string]string{
		"jdkaCCKglpat-djksaljkl":                        "jdkaCCK**********",
		"jdkaglpat-djksaljkl":                           "jdka**********",
		"dkakl_dsjkaABCtestg_djkajskldBBB":              "dkakl_dsjkaABC**********BBB",
		"jdkakl_dsjkaABCtestg_djkajskldBBBjdklasjklJKS": "jdkakl_dsjkaABC**********BBB**********JKS",
	}

	for s, d := range tc {
		fmt.Printf(">>> src: %#v, dst: %#v\n", s, d)
		dd := processor.Process(s)
		if dd != d {
			t.Fatalf("ERR: %#v\n", dd)
		}
	}
}
