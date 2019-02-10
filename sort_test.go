// Copyright 2013-2014 Canonical Ltd.
// Copyright 2019 Patrick Gaskin

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestVersionOrder(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	versions := []string{
		"1.2",
		"1.2rc5",
		"1.2rc4",
		"1.2rc3",
		"1.2rc2",
		"1.2rc1",
		"1.1.2",
		"1.1.1",
		"1.1",
		"1.1rc3",
		"1.1rc2",
		"1.1rc1",
		"1.1beta2",
		"1.1beta1",
		"1.0.3",
		"1.0.2",
		"1.0.1",
	}

	for perm := 0; perm < 1000; perm++ {
		vs := make([]string, len(versions))
		ts := make(tarballSlice, len(versions))
		pi := rand.Perm(len(vs))
		for i := range versions {
			ts[i] = &Tarball{Version: versions[pi[i]]}
		}
		sort.Sort(ts)
		for i := range versions {
			vs[i] = ts[i].Version
		}
		actual, expected := fmt.Sprintf("%#v", vs), fmt.Sprintf("%#v", versions)
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}
}
