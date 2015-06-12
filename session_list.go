package main

type SessionList []*Session

func (l SessionList) Len() int {
    return len(l)
}

func (l SessionList) Swap(i, j int) {
    l[i], l[j] = l[j], l[i]
}

func (l SessionList) Less(i, j int) bool {
    return l[i].Number < l[j].Number
}
