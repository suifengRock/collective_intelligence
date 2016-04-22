package main

import (
	dist "collective_intelligence/distance"
	"fmt"
	"sort"
)

type Similarity struct {
	Name  string
	Score float64
}

type SimilarityList []*Similarity

func (s SimilarityList) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}
func (s SimilarityList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SimilarityList) Len() int {
	return len(s)
}

var critics map[string]map[string]float64 = map[string]map[string]float64{
	"Lisa Rose": {
		"Lady in the Water":   2.5,
		"Snakes on a Plane":   3.5,
		"Just My Luck":        3.0,
		"Superman Returns":    3.5,
		"You , Me and Dupree": 2.5,
		"The Night Listener":  3.0,
	},
	"Gene Seymour": {
		"Lady in the Water":   3.0,
		"Snakes on a Plane":   3.5,
		"Just My Luck":        1.5,
		"Superman Returns":    5.0,
		"You , Me and Dupree": 3.5,
		"The Night Listener":  3.0,
	},
	"Claudia Pui g": {
		"Snakes on a Plane":   3.5,
		"Just My Luck":        3.0,
		"Superman Returns":    5.0,
		"You , Me and Dupree": 2.5,
		"The Night Listener":  4.5,
	},
	"Mick LaSalle": {
		"Lady in the Water":   3.0,
		"Snakes on a Plane":   4.0,
		"Just My Luck":        2.0,
		"Superman Returns":    3.0,
		"You , Me and Dupree": 2.0,
		"The Night Listener":  3.0,
	},
	"Jack Ma tthews": {
		"Lady in the Water":   3.0,
		"Snakes on a Plane":   4.0,
		"Superman Returns":    5.0,
		"You , Me and Dupree": 3.5,
		"The Night Listener":  3.0,
	},
	"Toby": {
		"Snakes on a Plane":   4.5,
		"Superman Returns":    4.0,
		"You , Me and Dupree": 1.0,
	},
}

var SimilaritySet map[string]SimilarityList = map[string]SimilarityList{}

func countEuclidSimilarity() {
	for key, _ := range critics {
		for tmp, _ := range critics {
			if key == tmp {
				continue
			}
			euclid := dist.EuclideanDistance(critics[key], critics[tmp])
			sim := new(Similarity)
			sim.Name = tmp
			sim.Score = euclid
			if list, ok := SimilaritySet[key]; ok {
				list = append(list, sim)
				SimilaritySet[key] = list
			} else {
				list := make([]*Similarity, 0)
				list = append(list, sim)
				SimilaritySet[key] = list
			}
		}
		sort.Sort(SimilaritySet[key])
	}
}
func countPearsonSimilarity() {
	for key, _ := range critics {
		for tmp, _ := range critics {
			if key == tmp {
				continue
			}
			euclid := dist.PearsonScore(critics[key], critics[tmp])
			sim := new(Similarity)
			sim.Name = tmp
			sim.Score = euclid
			if list, ok := SimilaritySet[key]; ok {
				list = append(list, sim)
				SimilaritySet[key] = list
			} else {
				list := make([]*Similarity, 0)
				list = append(list, sim)
				SimilaritySet[key] = list
			}
		}
		sort.Sort(SimilaritySet[key])
	}
}

func PrintSimilarity() {
	fmt.Println("-----------------------------")
	for key, obj := range SimilaritySet {
		fmt.Printf("%s\t", key)
		for _, score := range obj {
			fmt.Printf("%s:%f\t", score.Name, score.Score)
		}
		fmt.Println("  |")
	}
	fmt.Println("-----------------------------")
}

func recommend(name string) {

}
func main() {
	countPearsonSimilarity()
	PrintSimilarity()
}
