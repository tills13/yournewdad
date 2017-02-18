package kaa

import (
	"container/heap"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math"
	"math/rand"
	"os"
	"time"
)

func getMyHead(data *MoveRequest) (Point, error) {
	for _, snake := range data.Snakes {
		if snake.Id == data.You && len(data.You) > 0 {
			return snake.Head(), nil
		}
	}
	return Point{}, errors.New("Could not get head")
}

func getStaticData(data *MoveRequest, direc string) ([]*StaticData, error) {
	head, err := getMyHead(data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to get head of your snake"))
	}

	switch direc {
	case UP:
		return graphSearch(head.Up(data), data), nil
	case DOWN:
		return graphSearch(head.Down(data), data), nil
	case LEFT:
		return graphSearch(head.Left(data), data), nil
	case RIGHT:
		return graphSearch(head.Right(data), data), nil
	}
	return nil, errors.New(fmt.Sprintf("invalid direction", direc))
}

// returns an array of static data, the final static data is
// the maximum depth and the other depths, are defined in moves_to_depth
// in data.go. Will search from the point pos to the maximum depth provided
// a depth of any positive integer will max out at that integer and a depth of
// any negative integer will allow any negative number

func pushOntoPQ(
	p *Point,
	seen map[string]bool,
	pq *PriorityQueue,
	priority int) {

	if p != nil {
		if !seen[p.String()] {
			seen[p.String()] = true

			heap.Push(pq, &Item{
				value:    *p,
				priority: priority + 1,
			})
		}
	}
}

func graphSearch(pos *Point, data *MoveRequest) []*StaticData {
	ret := []*StaticData{}
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	// start with priority 0
	priority := 0
	seen := make(map[string]bool)

	// make the first item priority 1 so that if statement
	// at the end of the loop is executed
	pushOntoPQ(pos, seen, &pq, 1)

	accumulator := &StaticData{}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		p := item.value
		//fmt.Printf("%v", p)

		pushOntoPQ(p.Up(data), seen, &pq, item.priority)
		pushOntoPQ(p.Down(data), seen, &pq, item.priority)
		pushOntoPQ(p.Left(data), seen, &pq, item.priority)
		pushOntoPQ(p.Right(data), seen, &pq, item.priority)

		if data.FoodMap[p.String()] {
			accumulator.Food += 1
		}

		accumulator.Moves += 1

		if item.priority > priority {
			priority = item.priority
			ret = append(ret, accumulator)
			// copy accumulator
			tmp := *accumulator
			accumulator = &tmp
			//fmt.Printf("\n")
		}
	}
	return ret
}

func ClosestFood(data []*StaticData) int {
	for i, staticData := range data {
		//fmt.Printf("direction : %v\ndata for move %v: %#v\n", direc, i, staticData)
		if staticData.Food > 0 {
			return i + 1
		}
	}
	return math.MaxInt64
}

func FilterPossibleMoves(metaD map[string]*MetaData) []string {
	directions := []string{UP, DOWN, LEFT, RIGHT}
	ret := []string{}
	for _, direc := range directions {
		if len(metaD[direc].MovesAway) == 0 {
			ret = append(ret, direc)
		}
	}
	return directions
}

func ClosestFoodDirections(metaD map[string]*MetaData, moves []string) []string {
	directions := []string{}
	min := math.MaxInt64
	for _, direc := range moves {
		if metaD[direc].ClosestFood < min {
			directions = []string{}
			directions = append(directions, direc)
			min = metaD[direc].ClosestFood
		} else if metaD[direc].ClosestFood == min {
			directions = append(directions, direc)
		}
	}
	return directions
}

// not necessairily the best move but the move that we are going with
func bestMoves(metaD map[string]*MetaData) []string {
	moves := FilterPossibleMoves(metaD)
	rand.Seed(time.Now().Unix()) // initialize global pseudorandom generator
	moves = ClosestFoodDirections(metaD, moves)
	return moves
}

func bestMove(metaD map[string]*MetaData) string {
	moves := bestMoves(metaD)
	return moves[rand.Intn(len(moves))]
}

func GenerateMetaData(data *MoveRequest) (map[string]*MetaData, error) {
	metad := make(map[string]*MetaData)
	metad["up"] = &MetaData{}
	metad["down"] = &MetaData{}
	metad["right"] = &MetaData{}
	metad["left"] = &MetaData{}

	for direc, direcMD := range metad {
		sd, err := getStaticData(data, direc)
		if err != nil {
			return metad, err
		}

		direcMD.MovesAway = sd
		direcMD.ClosestFood = ClosestFood(sd)
	}
	return metad, nil
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		//log.Errorf("%s environment variable not set.", k)
	}
	return v
}
