package methods

type UserLevel struct {
	UserID  int64
	Level   uint32
	Exp     uint32
	NextExp uint32
}

func (level *UserLevel) AddExp(exp uint32) {
	level.Exp += exp
	if level.Exp >= level.NextExp {
		level.Exp = level.Exp - level.NextExp
		level.Level += 1
		level.NextExp = (level.Level + 1) * (level.Level + 5) * 10
	}
}
