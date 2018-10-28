
type DisjointSet struct {
	parent []int
	rank []int
}

func NewDisjointSet(n int) *DisjointSet {
	dj := DisjointSet{
		parent: make([]int, n),
		rank: make([]int, n),
	}
	for i := 0; i < n; i++ {
		dj.parent[i] = i
	}
	return &dj
}

func (dj *DisjointSet) find(x int) int {
	if dj.parent[x] != x {
		dj.parent[x] = dj.find(dj.parent[x])
	}
	return dj.parent[x]
}

func (dj *DisjointSet) union(x, y int) bool {
	xr := dj.find(x)
	yr := dj.find(y)

	if xr == yr {
		return false
	}

	if dj.rank[xr] < dj.rank[yr] {
		dj.parent[xr] = yr
	} else if dj.rank[xr] > dj.rank[yr] {
		dj.parent[yr] = xr
	} else {
		dj.parent[xr] = yr
		dj.rank[yr]++
	}
	return true
}
