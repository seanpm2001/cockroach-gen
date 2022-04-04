// Copyright 2022 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package backupccl

import (
	"sort"

	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfrapb"
	"github.com/cockroachdb/cockroach/pkg/util/interval"
)

type intervalSpan roachpb.Span

var _ interval.Interface = intervalSpan{}

// ID is part of `interval.Interface` but seemed unused by backupccl usage.
func (ie intervalSpan) ID() uintptr { return 0 }

// Range is part of `interval.Interface`.
func (ie intervalSpan) Range() interval.Range {
	return interval.Range{Start: []byte(ie.Key), End: []byte(ie.EndKey)}
}

// makeSimpleImportSpans partitions the spans of requiredSpans into a covering
// of RestoreSpanEntry's which each have all overlapping files from the passed
// backups assigned to them. The spans of requiredSpans are trimmed/removed
// based on the lowWaterMark before the covering for them is generated. Consider
// a chain of backups with files f1, f2… which cover spans as follows:
//
//  backup
//  0|     a___1___c c__2__e          h__3__i
//  1|         b___4___d           g____5___i
//  2|     a___________6______________h         j_7_k
//  3|                                  h_8_i              l_9_m
//   keys--a---b---c---d---e---f---g---h----i---j---k---l----m------p---->
// spans: |-------span1-------||---span2---|           |---span3---|
//
// The cover for those spans would look like:
//  [a, c): 1, 4, 6
//  [c, e): 2, 4, 6
//  [e, f): 6
//  [f, i): 3, 5, 6, 8
//  [l, m): 9
// This example is tested in TestRestoreEntryCoverExample.
func makeSimpleImportSpans(
	requiredSpans []roachpb.Span,
	backups []BackupManifest,
	backupLocalityMap map[int]storeByLocalityKV,
	lowWaterMark roachpb.Key,
) []execinfrapb.RestoreSpanEntry {
	if len(backups) < 1 {
		return nil
	}

	for i := range backups {
		sort.Sort(BackupFileDescriptors(backups[i].Files))
	}

	var cover []execinfrapb.RestoreSpanEntry
	for _, span := range requiredSpans {
		if span.EndKey.Compare(lowWaterMark) < 0 {
			continue
		}
		if span.Key.Compare(lowWaterMark) < 0 {
			span.Key = lowWaterMark
		}

		spanCoverStart := len(cover)

		for layer := range backups {
			covPos := spanCoverStart
			// TODO(dt): binary search to the first file in required span?
			for _, f := range backups[layer].Files {
				if sp := span.Intersect(f.Span); sp.Valid() {
					fileSpec := execinfrapb.RestoreFileSpec{Path: f.Path, Dir: backups[layer].Dir}
					if dir, ok := backupLocalityMap[layer][f.LocalityKV]; ok {
						fileSpec = execinfrapb.RestoreFileSpec{Path: f.Path, Dir: dir}
					}
					if len(cover) == spanCoverStart {
						cover = append(cover, makeEntry(span.Key, sp.EndKey, fileSpec))
					} else {
						// Add each file to every matching partition in the cover.
						for i := covPos; i < len(cover) && cover[i].Span.Key.Compare(sp.EndKey) < 0; i++ {
							if cover[i].Span.Overlaps(sp) {
								cover[i].Files = append(cover[i].Files, fileSpec)
							}
							// Later files start later, so this cover ends before this file,
							// it ends before them too and they can start searching after it.
							if cover[i].Span.EndKey.Compare(sp.Key) <= 0 {
								covPos = i + 1
							}
						}
						if covEnd := cover[len(cover)-1].Span.EndKey; sp.EndKey.Compare(covEnd) > 0 {
							cover = append(cover, makeEntry(covEnd, sp.EndKey, fileSpec))
						}
					}
				} else if span.EndKey.Compare(f.Span.Key) <= 0 {
					// This file is already after the end, so rest are too.
					break
				}
			}
		}
	}

	return cover
}

func makeEntry(start, end roachpb.Key, f execinfrapb.RestoreFileSpec) execinfrapb.RestoreSpanEntry {
	return execinfrapb.RestoreSpanEntry{
		Span: roachpb.Span{Key: start, EndKey: end}, Files: []execinfrapb.RestoreFileSpec{f},
	}
}
