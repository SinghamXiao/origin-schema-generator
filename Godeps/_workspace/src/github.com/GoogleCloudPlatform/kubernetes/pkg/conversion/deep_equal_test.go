/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package conversion

import (
	"testing"
)

func TestEqualities(t *testing.T) {
	e := Equalities{}
	type Bar struct {
		X int
	}
	type Baz struct {
		Y Bar
	}
	err := e.AddFuncs(
		func(a, b int) bool {
			return a+1 == b
		},
		func(a, b Bar) bool {
			return a.X*10 == b.X
		},
	)
	if err != nil {
		t.Fatalf("Unexpected: %v", err)
	}

	type Foo struct {
		X int
	}

	table := []struct {
		a, b  interface{}
		equal bool
	}{
		{1, 2, true},
		{2, 1, false},
		{"foo", "foo", true},
		{Foo{1}, Foo{2}, true},
		{Bar{1}, Bar{10}, true},
		{&Bar{1}, &Bar{10}, true},
		{Baz{Bar{1}}, Baz{Bar{10}}, true},
		{map[string]int{"foo": 1}, map[string]int{"foo": 2}, true},
		{map[string]int{}, map[string]int(nil), true},
		{[]int{}, []int(nil), true},
	}

	for _, item := range table {
		if e, a := item.equal, e.DeepEqual(item.a, item.b); e != a {
			t.Errorf("Expected (%+v == %+v) == %v, but got %v", item.a, item.b, e, a)
		}
	}
}
