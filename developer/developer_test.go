package developer_test

import (
	"testing"

	"github.com/Delaram-Gholampoor-Sagha/testing_in_go/developer"
	"github.com/stretchr/testify/assert"
)

func TestFilterUnique(t *testing.T) {
	input := []developer.Developer{
		developer.Developer{Name: "Elliot"},
		developer.Developer{Name: "Elliot"},
		developer.Developer{Name: "David"},
		developer.Developer{Name: "Alexandra"},
		developer.Developer{Name: "Eva"},
		developer.Developer{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"David",
		"Alexandra",
		"Eva",
		"Alan",
	}
	assert.Equal(t, expected, developer.FilterUnique(input))
}
