package main

import (
	"os/exec"
	"sync"

	"github.com/boltdb/bolt"
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

type Options struct {
	DiscordToken      string
	DiscordStatus     string
	DiscrodPrefix     string
	DiscrodPurgeTime  int64
	DiscordPlayStatus bool
	YoutubeToken      string
}

type TimeDuration struct {
	Day    int
	Hour   int
	Minute int
	Second int
}

type Song struct {
	ChannelId string
	User      string
	ID        string
	VidID     string
	Title     string
	Duration  string
	VideoURL  string
}

type PurgeMessage struct {
	ID, ChannelID string
	TimeSent      int64
}

type Channel struct {
	db *bolt.DB
}

type PkgSong struct {
	data Song
	v    *VoiceInstance
}

type PkgRadio struct {
	data string
	v    *VoiceInstance
}

type VoiceInstance struct {
	voice      *discordgo.VoiceConnection
	session    *discordgo.Session
	encoder    *dca.EncodeSession
	stream     *dca.StreamingSession
	run        *exec.Cmd
	queueMutex sync.Mutex
	audioMutex sync.Mutex
	nowPlaying Song
	queue      []Song
	recv       []int16
	guildID    string
	channelID  string
	speaking   bool
	pause      bool
	Stop       bool
	skip       bool
	radioFlag  bool
}
