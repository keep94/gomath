package gomath

import (
  "sort"
)

const (
  kXNotIncreasing = "X values must be strictly increasing"
  kZeroValueSpline = "Operation not allowed on zero value spline"
)

// Point represents a single (x, y) point
type Point struct {
  X float64
  Y float64
}
// Spline represents a cubic spline.
type Spline struct {
  points []Point
  polys []polyType
}

// NewSpline returns a new cubic spline going through each point in points.
// The second derivative of the spline at the first and last point is 0.
// The x values in points must be strictly increasing.
func NewSpline(points []Point) *Spline {
  points = checkAndCopySplinePoints(points)
  return &Spline{points: points, polys: splineNormal(points)}
}

// NewSplineWithSlopes returns a new cubic spline going through each point in
// points. The x values in points must be strictly increasing. beginSlope and
// endSlope specify the slope of the spline at the first point and last point
// respectively.
func NewSplineWithSlopes(
    points []Point, beginSlope, endSlope float64) *Spline {
  points = checkAndCopySplinePoints(points)
  return &Spline{
      points: points,
      polys: splineSlopes(points, beginSlope, endSlope)}
}


// Eval evaluates this cubic spline at x.  Eval panics if x doesn't fall
// between what MinX and MaxX return. Eval also panics if called on the zero
// value Spline.
func (s *Spline) Eval(x float64) float64 {
  if s.points == nil {
    panic(kZeroValueSpline)
  }
  if x < s.MinX() || x > s.MaxX() {
    panic("x value out of range for spline")
  }
  idx := sort.Search(
      len(s.points), func(i int) bool { return x < s.points[i].X }) - 1
  return s.polys[idx].eval(x - s.points[idx].X)
}

// MinX returns the minimum value of x for this cubic spline. MinX panics if
// called on the zero value Spline.
func (s *Spline) MinX() float64 {
  if s.points == nil {
    panic(kZeroValueSpline)
  }
  return s.points[0].X
}

// MaxX returns the maximum value of x for this cubic spline. MaxX panics if
// called on the zero value Spline.
func (s *Spline) MaxX() float64 {
  if s.points == nil {
    panic(kZeroValueSpline)
  }
  return s.points[len(s.points)-1].X
}

func checkAndCopySplinePoints(points []Point) []Point {
  if len(points) < 2 {
    panic("points must have length of at least 2")
  }
  for i := 1; i < len(points); i++ {
    if points[i].X <= points[i-1].X {
      panic(kXNotIncreasing)
    }
  }
  result := make([]Point, len(points))
  copy(result, points)
  return result
}

func computeSpline(points []Point, xcoef, x2coef float64) []polyType {
  var result []polyType
  cubic := polyType{points[0].Y, xcoef, x2coef, 0.0}
  for i := 0; i < len(points) - 1; i++ {
    xdiff := points[i+1].X - points[i].X
    cubic.fit(xdiff, points[i+1].Y)
    result = append(result, cubic)
    cubic = cubic.shift(xdiff)
  }
  result = append(result, cubic)
  return result
}

func splineNormal(points []Point) []polyType {
  spline0 := computeSpline(points, 0.0, 0.0)
  spline1 := computeSpline(points, 1.0, 0.0)
  plen := len(points)
  end2nd0 := spline0[plen-1][2]
  end2nd1 := spline1[plen-1][2]
  start1st := -end2nd0 / (end2nd1-end2nd0)
  return computeSpline(points, start1st, 0.0)
}

func splineSlopes(points []Point, beginSlope, endSlope float64) []polyType {
  spline0 := computeSpline(points, beginSlope, 0.0)
  spline1 := computeSpline(points, beginSlope, 1.0)
  plen := len(points)
  end1st0 := spline0[plen-1][1]
  end1st1 := spline1[plen-1][1]
  start2nd := (endSlope - end1st0) / (end1st1 - end1st0)
  return computeSpline(points, beginSlope, start2nd)
}

type polyType [4]float64

func (p *polyType) eval(x float64) float64 {
  sum := 0.0
  for i := 3; i >= 0; i-- {
    sum = sum*x + p[i]
  }
  return sum
}

func (p *polyType) fit(x, y float64) {
  p[3] = 0.0
  actual := p.eval(x)
  p[3] = (y - actual) / (x*x*x)
}

func (p *polyType) shift(x float64) polyType {
  // p'(x)
  p1 := polyType{p[1], 2.0*p[2], 3.0*p[3], 0.0}
  // p''(x) / 2.0
  p2 := polyType{p[2], 3.0*p[3], 0.0, 0.0}
  return polyType{p.eval(x), p1.eval(x), p2.eval(x), 0.0}
}
