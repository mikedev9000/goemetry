package goemetry

type BoundingBox struct {
	BottomLeft Point
	Height     uint
	Width      uint
}

func (receiver *BoundingBox) IsAboveish(other BoundingBox) bool {

	yDistance := receiver.BottomLeft.Y - other.BottomLeft.Y

	if yDistance <= 0 {
		// receiver is actually below or level with other
		return false
	}

	if yDistance > int(receiver.Height)*2 && yDistance > int(other.Height)*2 {
		return false
	}

	return true

}
