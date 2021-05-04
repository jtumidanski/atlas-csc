package buff

import (
	"github.com/sirupsen/logrus"
	"time"
)

type expirationTask struct {
	l logrus.FieldLogger
}

func ExpireTask(l logrus.FieldLogger) *expirationTask {
	return &expirationTask{l: l}
}

func (r *expirationTask) Run() {
	for _, bs := range GetRegistry().GetAll() {
		if bs.Expired() {
			GetRegistry().Expire(bs.CharacterId(), bs.SourceId())
			Cancel(r.l)(bs.CharacterId(), bs.Buff().Stats())
		}
	}
}

func (r *expirationTask) SleepTime() time.Duration {
	return time.Millisecond * time.Duration(50)
}
