package gamestate

import (
	"context"
	"github.com/smartystreets/goconvey/convey"
	"rump/internal/codec/msgpack"
	"testing"
)

func TestNewVector(t *testing.T) {
	convey.Convey("Проверяем генерацию плеера", t, func() {
		p := NewVector(10.1, 2.5, 3.6)
		convey.So(p.X, convey.ShouldEqual, 10.1)
		convey.So(p.Y, convey.ShouldEqual, 2.5)
		convey.So(p.Z, convey.ShouldEqual, 3.6)
	})
}

func TestNewPlayer(t *testing.T) {
	convey.Convey("Проверяем генерацию плеера", t, func() {
		p := NewPlayer(10, NewVector(1, 1, 1))
		convey.So(p.ID, convey.ShouldEqual, 10)
		convey.So(p.Pos.X, convey.ShouldEqual, 1)
		convey.So(p.Pos.Y, convey.ShouldEqual, 1)
		convey.So(p.Pos.Z, convey.ShouldEqual, 1)
	})
}

func TestGeneratePlayer(t *testing.T) {
	convey.Convey("Проверяем генерацию плеера", t, func() {
		p := GeneratePlayer()
		convey.So(p.ID, convey.ShouldBeGreaterThan, 0)
		convey.So(p.Pos.X, convey.ShouldBeGreaterThan, 0)
		convey.So(p.Pos.Y, convey.ShouldBeGreaterThan, 0)
		convey.So(p.Pos.Z, convey.ShouldBeGreaterThan, 0)
	})
}

func TestGenerateVector(t *testing.T) {
	convey.Convey("Проверяем генерацию плеера", t, func() {
		v := GenerateVector()
		convey.So(v.X, convey.ShouldBeGreaterThan, 0)
		convey.So(v.Y, convey.ShouldBeGreaterThan, 0)
		convey.So(v.Z, convey.ShouldBeGreaterThan, 0)
	})
}

func TestSyncPos(t *testing.T) {
	player := GeneratePlayer()
	convey.Convey("Проверяем общие механики отправки игрока(позиции)", t, func() {
		state := NewState(context.TODO(), msgpack.New())
		state.SyncPlayer(player)
	})
}

func TestRcvPos(t *testing.T) {
	convey.Convey("Проверяем общие механики получения позиции", t, func() {
		var (
			decodedPlayer *Player
			fromDbPlayer  *Player
			err           error
		)
		_ = fromDbPlayer
		_ = decodedPlayer
		_ = err
		pack := msgpack.New()
		player := GeneratePlayer()
		bytes, err := pack.Encode(player)
		state := NewState(context.TODO(), msgpack.New())

		fromDbPlayer, err = state.RcvPlayer(player.ID)
		convey.Convey("Проверяем, что при юзера нет в записях сервиса", func() {
			convey.So(err, convey.ShouldEqual, ErrPlayerNotFound)
		})
		_ = bytes
		state.SyncPlayer(player)
		err = pack.Decode(bytes, &decodedPlayer)
		fromDbPlayer, err = state.RcvPlayer(player.ID)
		convey.Convey("Проверяем, что при получении позиции у нас нет ошибки", func() {
			convey.So(err, convey.ShouldEqual, nil)
		})
		convey.Convey("Проверяем, что добавленный юзер появился ", func() {
			convey.So(fromDbPlayer.ID, convey.ShouldEqual, player.ID)
			convey.So(fromDbPlayer.Pos.X, convey.ShouldEqual, player.Pos.X)
			convey.So(fromDbPlayer.Pos.Y, convey.ShouldEqual, player.Pos.Y)
			convey.So(fromDbPlayer.Pos.Z, convey.ShouldEqual, player.Pos.Z)
		})
	})
}
