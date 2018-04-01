package topology

import (
    "math"
)

type SubnetGate struct {
    coord Coord
}

func Dist(from *SubnetGate, to *SubnetGate) CoordVal {
    deltaX := to.coord.X - from.coord.X
    deltaY := to.coord.Y - from.coord.Y
    return math.Sqrt64(deltaX * deltaX + deltaY * deltaY)
}

