package value_object

import "testing"

func TestNewPictureType_ShouldAcceptValidTypes(t *testing.T) {
	validTypes := []string{"front", "side", "back"}
	for _, s := range validTypes {
		pt, err := NewPictureType(s)
		if err != nil {
			t.Errorf("Expected no error for %q, got %v", s, err)
		}
		if pt.String() != s {
			t.Errorf("Expected %q, got %q", s, pt.String())
		}
	}
}

func TestNewPictureType_ShouldRejectInvalidTypes(t *testing.T) {
	invalidTypes := []string{"", "left", "top", "FRONT", "Front"}
	for _, s := range invalidTypes {
		_, err := NewPictureType(s)
		if err == nil {
			t.Errorf("Expected error for %q, got nil", s)
		}
	}
}

func TestPictureTypeConstants(t *testing.T) {
	if PictureTypeFront.String() != "front" {
		t.Errorf("Expected front, got %s", PictureTypeFront.String())
	}
	if PictureTypeSide.String() != "side" {
		t.Errorf("Expected side, got %s", PictureTypeSide.String())
	}
	if PictureTypeBack.String() != "back" {
		t.Errorf("Expected back, got %s", PictureTypeBack.String())
	}
}
