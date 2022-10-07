package cache

type Stat struct {
	Count     int64
	KeySize   int64
	ValueSize int64
}

func (s *Stat) statAdd(k string, v []byte) {
	s.Count++
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
}

func (s *Stat) statDel(k string, v []byte) {
	s.Count--
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
}
