package sound

import (
	"io/ioutil"

	"github.com/yohamta/godanmaku/danmaku/internal/resources/audios"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

const (
	sampleRate = 22050
)

type BgmKind int

const (
	BgmKindBattle BgmKind = iota
)

type SeKind int

const (
	SeKindHit SeKind = iota
	SeKindHit2
	SeKindShot
	SeKindBomb
	SeKindJump
	SeKindItemGet
)

var (
	audioContext *audio.Context
	seDic        = map[SeKind]*audio.Player{}
	bgmDic       = map[BgmKind]*audio.Player{}
	bgmVolume128 = 128
	seVolume128  = 128
)

func Load() {
	audioContext, _ = audio.NewContext(sampleRate)

	bgmDic[BgmKindBattle] = loadMp3(audioContext, &audios.BGM_MAOUDAMASHII_8BIT18)

	seDic[SeKindHit] = loadMp3NoLoop(audioContext, &audios.TM2_HIT000)
	seDic[SeKindHit2] = loadMp3NoLoop(audioContext, &audios.TM2_BOM001)
	seDic[SeKindShot] = loadMp3NoLoop(audioContext, &audios.SILENCER)
	seDic[SeKindBomb] = loadMp3NoLoop(audioContext, &audios.BAKUHA)
	seDic[SeKindJump] = loadWav(audioContext, &audios.JUMP)
	seDic[SeKindItemGet] = loadMp3NoLoop(audioContext, &audios.SE_MAOUDAMASHII_BATTLE02)

}

func PlayBgm(kind BgmKind) {
	bgmDic[kind].Rewind()
	bgmDic[kind].SetVolume(float64(bgmVolume128) / 128)
	bgmDic[kind].Play()
}

func PlaySe(kind SeKind) {
	seDic[kind].Rewind()
	seDic[kind].SetVolume(float64(seVolume128) / 128)
	seDic[kind].Play()
}

func loadWav(c *audio.Context, wavBytes *[]byte) *audio.Player {
	s, _ := wav.Decode(c, audio.BytesReadSeekCloser(*wavBytes))
	b, _ := ioutil.ReadAll(s)
	player, _ := audio.NewPlayerFromBytes(audioContext, b)
	return player
}

func loadMp3NoLoop(c *audio.Context, mp3Bytes *[]byte) *audio.Player {
	s, _ := mp3.Decode(audioContext, audio.BytesReadSeekCloser(*mp3Bytes))
	player, _ := audio.NewPlayer(audioContext, s)
	return player
}

func loadMp3(c *audio.Context, mp3Bytes *[]byte) *audio.Player {
	s, _ := mp3.Decode(audioContext, audio.BytesReadSeekCloser(*mp3Bytes))
	infiniteStream := audio.NewInfiniteLoop(s, s.Length())
	player, _ := audio.NewPlayer(audioContext, infiniteStream)
	return player
}
