package legos

type Lego struct {
	Color      string
	Dimensions LegoDimension
}

type LegoDimension struct {
	Size   string
	Height int
}

func NewLegoDimension(size string, height int) LegoDimension {
	return LegoDimension{
		Size:   size,
		Height: height,
	}
}
func (d LegoDimension) Equals(comparable LegoDimension) bool {
	return d.Height == comparable.Height && d.Size == comparable.Size
}
