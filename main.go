package main

import (
	"container/heap"
	"encoding/json"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Point struct {
	Name     string
	Lat      float64
	Lng      float64
	Capacity int // sức chứa tối đa
	Occupied int // số xe hiện tại (sẽ được cập nhật qua API)
}

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

var points = map[string]Point{
	"Nha xe D9":   {"Nha xe D9", 21.004061, 105.844573, 300, 0},
	"Nha xe D3-5": {"Nha xe D3-5", 21.004972, 105.845431, 500, 0},
	"Nha xe C7":   {"Nha xe C7", 21.005054, 105.844911, 300, 0},
	"Nha xe C5":   {"Nha xe C5", 21.005863, 105.844629, 200, 0},
	"Nha xe D4-6": {"Nha xe D4-6", 21.004592, 105.842322, 300, 0},
	"Nha xe B1":   {"Nha xe B1", 21.005002, 105.846058, 100, 0},
	"Nha xe TC":   {"Nha xe TC", 21.002553, 105.847055, 500, 0},
	"Nha xe B13":  {"Nha xe B13", 21.006460, 105.847312, 100, 0},
	"Nha xe B6":   {"Nha xe B6", 21.006319, 105.8465455, 100, 0},

	// Các ngã rẽ
	"Intersection D4-D6": {"Intersection D4-D6", 21.005079, 105.842333, 0, 0},

	"Intersection TDN": {"Intersection TDN", 21.005032, 105.845634, 0, 0},

	"Intersection D9-C5": {"Intersection D9-C5", 21.005027, 105.844605, 0, 0},
	"Intersection B6":    {"Intersection B6", 21.005022, 105.8465080, 0, 0},

	"Intersection B13TC":     {"Intersection B13TC", 21.004927, 105.846958, 0, 0},
	"Intersection TQB-TC":    {"Intersection TQB-TC", 21.003244, 105.847993, 0, 0},
	"Intersection TQB1":      {"Intersection TQB1", 21.004501, 105.847210, 0, 0},
	"Intersection TQB2":      {"Intersection TQB2", 21.004081, 105.847231, 0, 0},
	"Intersection quaydauTC": {"Intersection quaydauTC", 21.001862, 105.846331, 0, 0},
}

type Edge struct {
	To     string
	Weight int
}

type Graph map[string][]Edge

var graph = Graph{
	// Trục chính
	"Intersection D4-D6":     {{"Nha xe D4-6", 1}, {"Intersection D9-C5", 1}},
	"Intersection TDN":       {{"Nha xe C7", 1}, {"Nha xe B1", 1}},
	"Intersection D9-C5":     {{"Intersection D4-D6", 1}, {"Nha xe D9", 1}, {"Nha xe C5", 1}, {"Nha xe C7", 1}, {"Nha xe D3-5", 1}},
	"Intersection B6":        {{"Nha xe B1", 1}, {"Nha xe B6", 1}, {"Intersection B13TC", 1}},
	"Intersection B13TC":     {{"Intersection B6", 1}, {"Intersection TQB1", 1}, {"Nha xe B13", 1}},
	"Intersection TQB-TC":    {{"Intersection TQB2", 1}, {"Nha xe TC", 1}},
	"Intersection TQB1":      {{"Intersection B13TC", 1}, {"Intersection TQB2", 1}},
	"Intersection TQB2":      {{"Intersection TQB1", 1}, {"Intersection TQB-TC", 1}},
	"Intersection quaydauTC": {{"Intersection TQB-TC", 1}},

	// Các nhà xe
	"Nha xe C7":   {{"Intersection D9-C5", 1}},
	"Nha xe C5":   {{"Intersection D9-C5", 1}},
	"Nha xe D3-5": {{"Intersection TDN", 1}},

	"Nha xe D9":   {{"Intersection D9-C5", 1}},
	"Nha xe D4-6": {{"Intersection D4-D6", 1}},
	"Nha xe B1":   {{"Intersection TDN", 1}, {"Intersection B6", 1}},
	"Nha xe B6":   {{"Intersection B6", 1}},
	"Nha xe TC":   {{"Intersection quaydauTC", 1}},
	"Nha xe B13":  {{"Intersection B13TC", 1}},
}

// Dijkstra
type Item struct {
	Node     string
	Distance int
	Index    int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Distance < pq[j].Distance }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i]; pq[i].Index = i; pq[j].Index = j }
func (pq *PriorityQueue) Push(x any)        { *pq = append(*pq, x.(*Item)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(graph Graph, start, end string) ([]string, int) {
	dist := make(map[string]int)
	prev := make(map[string]string)
	pq := &PriorityQueue{}
	heap.Init(pq)

	for node := range graph {
		dist[node] = math.MaxInt
	}
	dist[start] = 0
	heap.Push(pq, &Item{Node: start, Distance: 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(*Item)
		if u.Node == end {
			break
		}
		for _, edge := range graph[u.Node] {
			alt := dist[u.Node] + edge.Weight
			if alt < dist[edge.To] {
				dist[edge.To] = alt
				prev[edge.To] = u.Node
				heap.Push(pq, &Item{Node: edge.To, Distance: alt})
			}
		}
	}

	path := []string{}
	for u := end; u != ""; u = prev[u] {
		path = append([]string{u}, path...)
		if u == start {
			break
		}
	}
	return path, dist[end]
}
func updateOccupiedHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID       string `json:"id"`
		Occupied int    `json:"occupied"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if point, ok := points[req.ID]; ok {
		point.Occupied = req.Occupied
		points[req.ID] = point
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Updated"))
	} else {
		http.Error(w, "Point not found", http.StatusNotFound)
	}
}

func main() {
	router := gin.Default()

	router.GET("/path", func(c *gin.Context) {
		from := c.Query("from")
		to := c.Query("to")
		path, _ := Dijkstra(graph, from, to)
		result := []Coordinate{}
		for _, id := range path {
			p := points[id]
			result = append(result, Coordinate{p.Lat, p.Lng})
		}
		c.JSON(http.StatusOK, result)
	})

	router.StaticFile("/", "./index.html") // serve frontend
	router.Run(":8080")
}
