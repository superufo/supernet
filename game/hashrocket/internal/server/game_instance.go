package server

import (
	"go.uber.org/zap"
	"time"

	"github.com/supernet/game/hashrocket/internal/server/entity"
	"github.com/supernet/game/hashrocket/pkg/log"
	"github.com/supernet/game/hashrocket/pkg/mysql"
)

type GameInstanceService struct {
}

func NewGameInstanceService() *GameInstanceService {
	return &GameInstanceService{}
}

func (us GameInstanceService) Create() (egi *entity.GameInstance, err error) {
	var (
		affected int64
		uuid     string
	)

	res, err := mysql.M().QueryString("select hex(uuid_to_bin(uuid(),1)) as uuid")
	uuid = res[0]["uuid"]
	if err != nil {
		return nil, err
	}

	egi = &entity.GameInstance{
		GameSid:   uuid,
		Starttime: time.Now(),
		Explosion: 0,
	}

	if affected, err = mysql.M().Table(egi.TableName()).Insert(egi); err == nil {
		if affected < 1 {
			log.ZapLog.With(zap.Any("affected", affected)).Warn("数据库插入数据为空")
		}
	} else {
		log.ZapLog.With(zap.Any("err", err)).Warn("数据库插入错误")
	}
	log.ZapLog.With(zap.Any("egi:", egi)).Info("Create")

	return egi, err
}
