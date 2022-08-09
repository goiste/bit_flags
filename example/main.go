package main

import (
	"fmt"
)

func main() {
	ntf := New()
	fmt.Println(ntf.Flags.Get())     // 39
	fmt.Println(ntf.FlagsToString()) // Messages, Replies, Likes, By email

	ntf.Flags.Add(News | BySms)
	ntf.Flags.Remove(Replies | Likes)
	fmt.Println(ntf.FlagsToString())   // Messages, News, By email, By sms
	fmt.Println(ntf.HasFlag(Messages)) // true
	fmt.Println(ntf.HasFlag(Likes))    // false

	ntf.Flags.Reset()
	fmt.Println(ntf.Flags.Get())     // 0
	fmt.Println(ntf.GetFlagsNames()) // []

	ntf.SetAll()
	fmt.Println(ntf.Flags.Get())     // 255
	fmt.Println(ntf.FlagsToString()) // Messages, Replies, Likes, New articles, News, By email, By sms, By telegram
}
