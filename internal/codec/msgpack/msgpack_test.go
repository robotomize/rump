package msgpack

import (
	"github.com/smartystreets/goconvey/convey"
	"rump/internal/gamestate"
	"testing"
)

func TestEncodePositive(t *testing.T) {
	convey.Convey("Проверяем encoding", t, func() {
		player := gamestate.GeneratePlayer()
		pack := New()
		_, err := pack.Encode(player)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestEncodeNegative(t *testing.T) {
	convey.Convey("Проверяем encoding", t, func() {

	})
}

func TestDecodeNegative(t *testing.T) {
	convey.Convey("Проверяем decoding", t, func() {
	})
}

func TestDecodePositive(t *testing.T) {
	convey.Convey("Проверяем decoding", t, func() {
		var (
			decodedPlayer *gamestate.Player
			err           error
		)
		_ = decodedPlayer
		_ = err
		player := gamestate.GeneratePlayer()
		pack := New()
		bytes, err := pack.Encode(player)
		err = pack.Decode(bytes, &decodedPlayer)
		convey.Convey("Проверяем, что при декодировании нет ошибок", func() {
			convey.So(err, convey.ShouldEqual, nil)
		})
		convey.Convey("Проверяем, что при декодировании все корректно", func() {
			convey.So(decodedPlayer.ID, convey.ShouldEqual, player.ID)
			convey.So(decodedPlayer.Pos.X, convey.ShouldEqual, player.Pos.X)
			convey.So(decodedPlayer.Pos.Y, convey.ShouldEqual, player.Pos.Y)
			convey.So(decodedPlayer.Pos.Z, convey.ShouldEqual, player.Pos.Z)
		})
	})
}
