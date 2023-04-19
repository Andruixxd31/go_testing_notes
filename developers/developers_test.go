package developers_test

import (
	"testing"

	"github.com/andruixxd31/go-testing-bible/developers"
	"github.com/stretchr/testify/assert"
)

func TestFilterUnique(t *testing.T) {
    input := []developers.Developer{
		{"Eliot", 25},
        {"Andres", 24},
        {"Eliot", 25},
        {"Andres", 24},
	}

    expected := []string{
        "Eliot",
        "Andres",
    }

    result := developers.FilterUnique(input)
    assert.Equal(t, expected, result)
    
    
    // Making test only using go standar libraries
    // result := developers.FilterUnique(input)
    // 
    // if !reflect.DeepEqual(result, expected){
    //     t.Fail()
    // }
}

func TestNegativeFilterUnique(t *testing.T){
    input := []developers.Developer{
		{"Eliot", 25},
        {"Andres", 24},
        {"Eliot", 25},
        {"Andres", 24},
	}

    expected := []string{
        "Eliot",
    }

    result := developers.FilterUnique(input)
    assert.NotEqual(t, expected, result)
}
