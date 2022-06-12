package unionfind

type UnionFind interface {
	Add(u int, v int)
	Find(u int) int
	IsConnected(u int, v int) bool
	FindMembers(u int) []int
}

type unionFind struct {
	nodeCount   int
	id          []int
	size        []int
	forestCount int
}

// Add implements UnionFind
func (unionFind *unionFind) Add(u int, v int) {
	uParent := unionFind.Find(u)
	vParent := unionFind.Find(v)
	if unionFind.IsConnected(uParent, vParent) {
		return
	}

	if unionFind.size[uParent] >= unionFind.size[vParent] {
		unionFind.id[vParent] = uParent
		unionFind.size[uParent] += unionFind.size[vParent]
		unionFind.size[vParent] = 0
	} else {
		unionFind.id[uParent] = vParent
		unionFind.size[vParent] += unionFind.size[uParent]
		unionFind.size[uParent] = 0
	}
	unionFind.forestCount--
}

// Find implements UnionFind
func (unionFind *unionFind) Find(u int) int {
	root := unionFind.id[u]
	for root != unionFind.id[root] {
		root = unionFind.id[root]
	}

	for u != root {
		next := unionFind.id[u]
		unionFind.id[u] = root
		u = next
	}

	return root
}

// FindMembers implements UnionFind
func (unionFind *unionFind) FindMembers(u int) []int {
	// find the members of the forest
	if u >= unionFind.forestCount {
		return []int{}
	}
	nums := make([]int, 0)
	for i := 0; i < unionFind.nodeCount; i++ {
		if unionFind.id[i] == u {
			nums = append(nums, i)
		}
	}
	return nums
}

// IsConnected implements UnionFind
func (unionFind *unionFind) IsConnected(u int, v int) bool {
	return unionFind.Find(u) == unionFind.Find(v)
}

func NewUnionFind(n int) UnionFind {
	id := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		size[i] = 1
		id[i] = i
	}
	return &unionFind{
		nodeCount:   n,
		forestCount: n,
		id:          id,
		size:        size,
	}
}
