package value_object

type PictureType string

const (
	PictureTypeFront PictureType = "front"
	PictureTypeSide  PictureType = "side"
	PictureTypeBack  PictureType = "back"
)

func NewPictureType(s string) (PictureType, error) {
	switch PictureType(s) {
	case PictureTypeFront, PictureTypeSide, PictureTypeBack:
		return PictureType(s), nil
	}
	return "", ErrInvalidPictureType
}

func (p PictureType) String() string {
	return string(p)
}

var ErrInvalidPictureType = &ValidationError{Field: "pictureType", Message: "must be one of: front, side, back"}
