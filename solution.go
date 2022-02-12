package solution

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	fmt.Println(emoji.Sprint("Hello :world_map:!"))
	return ""
}
