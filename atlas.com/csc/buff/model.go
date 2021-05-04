package buff

func NewModel(expiration int64, stats []Stat) *Model {
	return &Model{
		expiration: expiration,
		stats:      stats,
	}
}

type Model struct {
	expiration int64
	stats      []Stat
}

func (m Model) Stats() []Stat {
	return m.stats
}

type Stat struct {
	First  bool
	Mask   uint64
	Amount uint16
}
