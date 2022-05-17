package unionfind

import "fmt"

type UnionFind interface {
	Find(p int) int
	Union(i int, j int)
	Connected(i int, j int) bool
	ComponentCount() int
	NodeCount() int
	Members(p int) []int
	PrintMembers()
}

type unionFind struct {
	Size         []int
	Id           []int
	NumComponent int
	NodeNumber   int
}

// PrintMembers implements UnionFind
func (u *unionFind) PrintMembers() {
	// get parent first
	parents := make(map[int]int)
	for i := 0; i < u.NodeNumber; i++ {
		parents[u.Id[i]] = 1
	}

	for i, _ := range parents {
		members := u.Members(i)
		fmt.Println("Group ", i, members)
	}
}

// Members implements UnionFind
func (u *unionFind) Members(p int) []int {
	members := make([]int, 0)

	for v, i := range u.Id {
		if i == p {
			members = append(members, v)
		}
	}

	return members
}

// ComponentCount implements UnionFind
func (u *unionFind) ComponentCount() int {
	return u.NumComponent
}

// Connected implements UnionFind
func (u *unionFind) Connected(i int, j int) bool {
	return u.Find(i) == u.Find(j)
}

// Find implements UnionFind
func (u *unionFind) Find(p int) int {
	root := p
	for root != u.Id[root] {
		root = u.Id[root]
	}
	for root != p {
		next := u.Id[p]
		u.Id[p] = root
		p = next
	}
	return root
}

// NodeCount implements UnionFind
func (u *unionFind) NodeCount() int {
	return u.NodeNumber
}

// Union implements UnionFind
func (u *unionFind) Union(i int, j int) {
	if u.Connected(i, j) {
		return
	}
	root1 := u.Find(i)
	root2 := u.Find(j)
	if u.Size[root1] < u.Size[root2] {
		u.Size[root2] += u.Size[root1]
		u.Id[root1] = root2
		u.Size[root1] = 0
	} else {
		u.Size[root1] += u.Size[root2]
		u.Id[root2] = root1
		u.Size[root2] = 0
	}
	u.NumComponent--
}

func NewUnionFind(n int) UnionFind {
	size := make([]int, n)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		size[i] = 1
		id[i] = i
	}
	return &unionFind{
		Size:         size,
		Id:           id,
		NumComponent: n,
		NodeNumber:   n,
	}
}
