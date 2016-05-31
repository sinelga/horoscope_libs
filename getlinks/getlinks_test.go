package getlinks

import (
	"testing"
	
)

func TestGet(t *testing.T) {

	result := GetLinks("http://anna.fi/kategoria/horoskoopit/paivahoroskoopit/")

	if len(result) < 10 {

		t.Error("Expected 10, got ", len(result))
	}

}
