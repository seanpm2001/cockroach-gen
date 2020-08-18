// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package geo

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/cockroachdb/cockroach/pkg/geo/geopb"
	"github.com/cockroachdb/errors"
	"github.com/golang/geo/s2"
	"github.com/twpayne/go-geom"
)

// CartesianBoundingBox is the cartesian BoundingBox representation,
// meant for use for GEOMETRY types.
type CartesianBoundingBox struct {
	geopb.BoundingBox
}

// NewCartesianBoundingBox returns a properly initialized empty bounding box
// for carestian plane types.
func NewCartesianBoundingBox() *CartesianBoundingBox {
	return nil
}

// Repr is the string representation of the CartesianBoundingBox.
func (b *CartesianBoundingBox) Repr() string {
	// fmt.Sprintf with %f does not truncate leading zeroes, so use
	// FormatFloat instead.
	return fmt.Sprintf(
		"BOX(%s %s,%s %s)",
		strconv.FormatFloat(b.LoX, 'f', -1, 64),
		strconv.FormatFloat(b.LoY, 'f', -1, 64),
		strconv.FormatFloat(b.HiX, 'f', -1, 64),
		strconv.FormatFloat(b.HiY, 'f', -1, 64),
	)
}

// ParseCartesianBoundingBox parses a box2d string into a bounding box.
func ParseCartesianBoundingBox(s string) (*CartesianBoundingBox, error) {
	b := &CartesianBoundingBox{}
	var prefix string
	numScanned, err := fmt.Sscanf(s, "%3s(%f %f,%f %f)", &prefix, &b.LoX, &b.LoY, &b.HiX, &b.HiY)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing box2d")
	}
	if numScanned != 5 || strings.ToLower(prefix) != "box" {
		return nil, errors.Newf("expected format 'box(min_x min_y,max_x max_y)'")
	}
	return b, nil
}

// Compare returns the comparison between two bounding boxes.
// Compare lower dimensions before higher ones, i.e. X, then Y.
func (b *CartesianBoundingBox) Compare(o *CartesianBoundingBox) int {
	if b.LoX < o.LoX {
		return -1
	} else if b.LoX > o.LoX {
		return 1
	}

	if b.HiX < o.HiX {
		return -1
	} else if b.HiX > o.HiX {
		return 1
	}

	if b.LoY < o.LoY {
		return -1
	} else if b.LoY > o.LoY {
		return 1
	}

	if b.HiY < o.HiY {
		return -1
	} else if b.HiY > o.HiY {
		return 1
	}

	return 0
}

// withPoint includes a new point to the CartesianBoundingBox.
// It will edit any bounding box in place.
func (b *CartesianBoundingBox) withPoint(x, y float64) *CartesianBoundingBox {
	if b == nil {
		return &CartesianBoundingBox{
			BoundingBox: geopb.BoundingBox{
				LoX: x,
				HiX: x,
				LoY: y,
				HiY: y,
			},
		}
	}
	b.BoundingBox = geopb.BoundingBox{
		LoX: math.Min(b.LoX, x),
		HiX: math.Max(b.HiX, x),
		LoY: math.Min(b.LoY, y),
		HiY: math.Max(b.HiY, y),
	}
	return b
}

// AddPoint adds a point to the CartesianBoundingBox coordinates.
// Returns a copy of the CartesianBoundingBox.
func (b *CartesianBoundingBox) AddPoint(x, y float64) *CartesianBoundingBox {
	if b == nil {
		return &CartesianBoundingBox{
			BoundingBox: geopb.BoundingBox{
				LoX: x,
				HiX: x,
				LoY: y,
				HiY: y,
			},
		}
	}
	return &CartesianBoundingBox{
		BoundingBox: geopb.BoundingBox{
			LoX: math.Min(b.LoX, x),
			HiX: math.Max(b.HiX, x),
			LoY: math.Min(b.LoY, y),
			HiY: math.Max(b.HiY, y),
		},
	}
}

// Combine combines two bounding boxes together.
// Returns a copy of the CartesianBoundingBox.
func (b *CartesianBoundingBox) Combine(o *CartesianBoundingBox) *CartesianBoundingBox {
	if o == nil {
		return b
	}
	return b.AddPoint(o.LoX, o.LoY).AddPoint(o.HiX, o.HiY)
}

// Buffer adds n units to each side of the bounding box.
func (b *CartesianBoundingBox) Buffer(n float64) *CartesianBoundingBox {
	if b == nil {
		return nil
	}
	return &CartesianBoundingBox{
		BoundingBox: geopb.BoundingBox{
			LoX: b.LoX - n,
			HiX: b.HiX + n,
			LoY: b.LoY - n,
			HiY: b.HiY + n,
		},
	}
}

// Intersects returns whether the BoundingBoxes intersect.
// Empty bounding boxes never intersect.
func (b *CartesianBoundingBox) Intersects(o *CartesianBoundingBox) bool {
	// If either side is empty, they do not intersect.
	if b == nil || o == nil {
		return false
	}
	if b.LoY > o.HiY || o.LoY > b.HiY ||
		b.LoX > o.HiX || o.LoX > b.HiX {
		return false
	}
	return true
}

// Covers returns whether the BoundingBox covers the other bounding box.
// Empty bounding boxes never cover.
func (b *CartesianBoundingBox) Covers(o *CartesianBoundingBox) bool {
	if b == nil || o == nil {
		return false
	}
	return b.LoX <= o.LoX && o.LoX <= b.HiX &&
		b.LoX <= o.HiX && o.HiX <= b.HiX &&
		b.LoY <= o.LoY && o.LoY <= b.HiY &&
		b.LoY <= o.HiY && o.HiY <= b.HiY
}

// boundingBoxFromGeomT returns a bounding box from a given geom.T.
// Returns nil if no bounding box was found.
func boundingBoxFromGeomT(g geom.T, soType geopb.SpatialObjectType) (*geopb.BoundingBox, error) {
	switch soType {
	case geopb.SpatialObjectType_GeometryType:
		ret := BoundingBoxFromGeomTGeometryType(g)
		if ret == nil {
			return nil, nil
		}
		return &ret.BoundingBox, nil
	case geopb.SpatialObjectType_GeographyType:
		rect, err := boundingBoxFromGeomTGeographyType(g)
		if err != nil {
			return nil, err
		}
		if rect.IsEmpty() {
			return nil, nil
		}
		return &geopb.BoundingBox{
			LoX: rect.Lng.Lo,
			HiX: rect.Lng.Hi,
			LoY: rect.Lat.Lo,
			HiY: rect.Lat.Hi,
		}, nil
	}
	return nil, errors.Newf("unknown spatial type: %s", soType)
}

// BoundingBoxFromGeomTGeometryType returns an appropriate bounding box for a Geometry type.
func BoundingBoxFromGeomTGeometryType(g geom.T) *CartesianBoundingBox {
	if g.Empty() {
		return nil
	}
	bbox := NewCartesianBoundingBox()
	switch g := g.(type) {
	case *geom.GeometryCollection:
		for i := 0; i < g.NumGeoms(); i++ {
			shapeBBox := BoundingBoxFromGeomTGeometryType(g.Geom(i))
			if shapeBBox == nil {
				continue
			}
			bbox = bbox.withPoint(shapeBBox.LoX, shapeBBox.LoY).withPoint(shapeBBox.HiX, shapeBBox.HiY)
		}
	default:
		flatCoords := g.FlatCoords()
		for i := 0; i < len(flatCoords); i += g.Stride() {
			bbox = bbox.withPoint(flatCoords[i], flatCoords[i+1])
		}
	}
	return bbox
}

// boundingBoxFromGeomTGeographyType returns an appropriate bounding box for a
// Geography type. There are marginally invalid shapes for which we want
// bounding boxes that are correct regardless of the validity of the shape,
// since validity checks may return slightly different results in S2 and the
// other libraries we use. Therefore, instead of constructing s2.Region(s)
// from the shape, which will expose us to S2's validity checks, we use the
// points and lines directly to compute the bounding box.
func boundingBoxFromGeomTGeographyType(g geom.T) (s2.Rect, error) {
	if g.Empty() {
		return s2.EmptyRect(), nil
	}
	rect := s2.EmptyRect()
	switch g := g.(type) {
	case *geom.Point:
		return geogPointsBBox(g), nil
	case *geom.MultiPoint:
		return geogPointsBBox(g), nil
	case *geom.LineString:
		return geogLineBBox(g), nil
	case *geom.MultiLineString:
		for i := 0; i < g.NumLineStrings(); i++ {
			rect = rect.Union(geogLineBBox(g.LineString(i)))
		}
	case *geom.Polygon:
		for i := 0; i < g.NumLinearRings(); i++ {
			rect = rect.Union(geogLineBBox(g.LinearRing(i)))
		}
	case *geom.MultiPolygon:
		for i := 0; i < g.NumPolygons(); i++ {
			polyRect, err := boundingBoxFromGeomTGeographyType(g.Polygon(i))
			if err != nil {
				return s2.EmptyRect(), err
			}
			rect = rect.Union(polyRect)
		}
	case *geom.GeometryCollection:
		for i := 0; i < g.NumGeoms(); i++ {
			collRect, err := boundingBoxFromGeomTGeographyType(g.Geom(i))
			if err != nil {
				return s2.EmptyRect(), err
			}
			rect = rect.Union(collRect)
		}
	default:
		return s2.EmptyRect(), errors.Errorf("unknown type %T", g)
	}
	return rect, nil
}

// geogPointsBBox constructs a bounding box, represented as a s2.Rect, for the set
// of points contained in g.
func geogPointsBBox(g geom.T) s2.Rect {
	rect := s2.EmptyRect()
	flatCoords := g.FlatCoords()
	for i := 0; i < len(flatCoords); i += g.Stride() {
		rect = rect.AddPoint(s2.LatLngFromDegrees(flatCoords[i+1], flatCoords[i]))
	}
	return rect
}

// geogLineBBox constructs a bounding box, represented as a s2.Rect, for the line
// or ring/loop represented by g.
func geogLineBBox(g geom.T) s2.Rect {
	bounder := s2.NewRectBounder()
	flatCoords := g.FlatCoords()
	for i := 0; i < len(flatCoords); i += g.Stride() {
		bounder.AddPoint(s2.PointFromLatLng(s2.LatLngFromDegrees(flatCoords[i+1], flatCoords[i])))
	}
	return bounder.RectBound()
}
