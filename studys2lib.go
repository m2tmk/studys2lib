package main

import (
	"github.com/golang/geo/r1"
	"github.com/golang/geo/r2"
	"github.com/golang/geo/r3"
	"github.com/golang/geo/s1"
	"log"
)

func main() {
  i := r1.Interval{1.0, 2.0}
	log.Printf("Interval: %v, %v\n", i.Lo, i.Hi)

  p := r2.Point{1.0, 2.0}
	log.Printf("Point: %v, %v\n", p.X, p.Y)

  pv := r3.NewPreciseVector(1.0, 2.0, 3.0)
	log.Printf("PreciseVector: %f, %f, %f\n", pv.X, pv.Y, pv.Z)

  d := s1.Degree*30
	log.Printf("Degree: %f\n", d)

}
