package buff

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"time"
)

const ExpirationTask = "buff_expiration_task"

type expirationTask struct {
	l logrus.FieldLogger
}

func ExpireTask(l logrus.FieldLogger) *expirationTask {
	return &expirationTask{l: l}
}

func (r *expirationTask) Run() {
	span := opentracing.StartSpan(ExpirationTask)
	for _, bs := range GetRegistry().GetAll() {
		if bs.Expired() {
			GetRegistry().Expire(bs.CharacterId(), bs.SourceId())
			Cancel(r.l, span)(bs.CharacterId(), bs.Buff().Stats())
		}
	}
	span.Finish()
}

func (r *expirationTask) SleepTime() time.Duration {
	return time.Millisecond * time.Duration(50)
}
