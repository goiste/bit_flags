# Bit Flags

A simple package to store up to 64 boolean flags in one uint field.

Requires Go version 1.18+ (uses generics)

For usage example see [example dir](example)

```go
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

func New() *Notification { ... }

func (n *Notification) SetDefaultFlags() {
	n.Flags.Set(Messages | Replies | Likes | ByEmail)
}

func (n *Notification) SetAll() {
	n.Flags.Set(AllNotifications | AllMethods)
}

func (n *Notification) SetNone() {
	n.Flags.Reset()
}

...
```
[full notification.go](example/notification.go)
```go
func main() {
	ntf := notification.New()
	fmt.Println(ntf.Flags.Get())     // 39
	fmt.Println(ntf.FlagsToString()) // Messages, Replies, Likes, By email

	ntf.Flags.Add(notification.News | notification.BySms)
	ntf.Flags.Remove(notification.Replies | notification.Likes)
	fmt.Println(ntf.FlagsToString())                // Messages, News, By email, By sms
	fmt.Println(ntf.HasFlag(notification.Messages)) // true
	fmt.Println(ntf.HasFlag(notification.Likes))    // false

	ntf.Flags.Reset()
	fmt.Println(ntf.Flags.Get())     // 0
	fmt.Println(ntf.GetFlagsNames()) // []

	ntf.SetAll()
	fmt.Println(ntf.Flags.Get())     // 255
	fmt.Println(ntf.FlagsToString()) // Messages, Replies, Likes, New articles, News, By email, By sms, By telegram
}
```
[full main.go](example/main.go)