package files

import "os"

type ByName []os.FileInfo

func (s ByName) Len() int {
	return len(s)
}

func (s ByName) Less(i, j int) bool {
	return s[i].Name() < s[j].Name()
}

func (s ByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

