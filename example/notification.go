package main

import (
	"strings"

	bf "github.com/goiste/bit_flags"
)

const (
	Messages    = 1 << iota // 1
	Replies                 // 2
	Likes                   // 4
	NewArticles             // ... powers of 2
	News

	ByEmail
	BySms
	ByTelegram

	AllNotifications = Messages | Replies | Likes | NewArticles | News
	AllMethods       = ByEmail | BySms | ByTelegram
)

type Notification struct {
	SomeOtherFields string
	Flags           bf.BitFlags[uint8]
}

func New() *Notification {
	n := &Notification{}
	n.SetDefaultFlags()
	return n
}

func (n *Notification) SetDefaultFlags() {
	n.Flags.Set(Messages | Replies | Likes | ByEmail)
}

func (n *Notification) SetAll() {
	n.Flags.Set(AllNotifications | AllMethods)
}

func (n *Notification) SetNone() {
	n.Flags.Reset()
}

func (n Notification) HasFlag(flag uint8) bool {
	return n.Flags.Has(flag)
}

func (n Notification) FlagsToString() string {
	return strings.Join(n.GetFlagsNames(), ", ")
}

func (n Notification) GetFlagsNames() (names []string) {
	for _, p := range n.Flags.List() {
		name := getFlagName(p)
		if name == "" {
			continue
		}
		names = append(names, name)
	}
	return
}

func getFlagName(flag uint8) string {
	switch flag {
	case Messages:
		return "Messages"
	case Replies:
		return "Replies"
	case Likes:
		return "Likes"
	case NewArticles:
		return "New articles"
	case News:
		return "News"
	case ByEmail:
		return "By email"
	case BySms:
		return "By sms"
	case ByTelegram:
		return "By telegram"
	}
	return ""
}
