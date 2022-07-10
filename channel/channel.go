// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package channel

import "sync"

var once sync.Once
var lock = make(chan struct{})

func Global() {
	once.Do(func() {
		defer close(lock)
	})

	<-lock
}
