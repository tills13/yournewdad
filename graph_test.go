package main

import (
	"reflect"
	"testing"
)

func Test_FullStats(t *testing.T) {
	data, err := NewMoveRequest(gameString9)

	if err != nil {
		t.Errorf("%v", err)
	}

	stats := quickStats2(data, "")
	if stats.Food != 0 {
		t.Errorf("There is no food within 5 moves got %v", stats.Food)
	}
	if stats.ClosestFood != nil {
		t.Errorf("closest food should be nil within 5 moves got %v", stats.ClosestFood)
	}

	if reflect.DeepEqual(stats.ClosestFood, &Point{X: 0, Y: 8}) {
		t.Errorf("closest food should be nil within 5 moves got %v", stats.ClosestFood)
	}
}
