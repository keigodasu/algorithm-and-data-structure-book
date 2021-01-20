package main

import "fmt"

type UnionFind struct {
	par, size []int
}

func NewUnionFind(n int) *UnionFind {
	par := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = -1
	}
	for i := 0; i < n; i++ {
		size[i] = 1
	}

	return &UnionFind{
		par:  par,
		size: size,
	}
}

func (u *UnionFind) Root(x int) int {
	if u.par[x] == -1 {
		return x
	} else {
		u.par[x] = u.Root(u.par[x])
		return u.par[x]
	}
}

func (u *UnionFind) IsSame(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

func (u *UnionFind) Unite(x, y int) {
	x = u.Root(x)
	y = u.Root(y)

	if x == y {
		return
	}

	if u.size[x] < u.size[y] {
		temp := u.size[y]
		u.size[y] = u.size[x]
		u.size[x] = temp
	}

	u.par[y] = x
	u.size[x] += u.size[y]

	return
}

func (u *UnionFind) Size(x int) int {
	return u.size[x]
}

func main() {
	uf := NewUnionFind(7)
	fmt.Println(uf)
	uf.Unite(1, 2)
	fmt.Println(uf)
	uf.Unite(2, 3)
	fmt.Println(uf)
	uf.Unite(5, 6)
	fmt.Println(uf)

	fmt.Println(uf.IsSame(1, 3))
	fmt.Println(uf.IsSame(2, 5))

	uf.Unite(1, 6)
	fmt.Println(uf)
	uf.IsSame(2, 5)
	fmt.Println(uf.IsSame(2, 5))
}
