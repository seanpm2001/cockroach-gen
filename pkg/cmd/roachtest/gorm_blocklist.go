// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package main

var gormBlocklists = blocklistsForVersion{
	{"v20.1", "gormBlocklist20_1", gormBlocklist20_1, "gormIgnorelist20_1", gormIgnorelist20_1},
	{"v20.2", "gormBlocklist20_2", gormBlocklist20_2, "gormIgnorelist20_2", gormIgnorelist20_2},
	{"v21.1", "gormBlocklist21_1", gormBlocklist21_1, "gormIgnorelist21_1", gormIgnorelist21_1},
}

var gormBlocklist21_1 = gormBlocklist20_2

var gormBlocklist20_2 = gormBlocklist20_1

var gormBlocklist20_1 = blocklist{}

var gormIgnorelist21_1 = gormIgnorelist20_2

var gormIgnorelist20_2 = gormIgnorelist20_1

var gormIgnorelist20_1 = blocklist{}
