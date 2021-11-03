/*
Copyright 2020 The Flux authors

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

package main

import (
	"fmt"
	"time"
)

func main() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	for {
		t := time.Now()

		if t.Second() == 0 {
			println("BEEP")
			break
		} else if t.Second() == 54 {
			t = t.In(loc)
			_, _ = fmt.Printf("At the tone, %d hour(s), %d minute(s), Coordinated Universal Time\n", t.Hour(), t.Minute())
		} else {
			println("tick")
		}

		time.Sleep(time.Second)
	}
}
