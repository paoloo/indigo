package main

type State struct {
	MemoryHash string
}

func (s *State) Init(newhash string) {
	s.MemoryHash = newhash
}

func (s *State) Check(newhash string) bool {
	return s.MemoryHash == newhash
}

func (s *State) Set(newhash string) {
	s.MemoryHash = newhash
}

func (s *State) CallIfChanged(newhash string, callback func()) {
	if !s.Check(newhash) {
		callback()
		s.Set(newhash)
	}
}
