package image

import (
	"fmt"
	"github.com/seams-cms/go-seams-cms-sdk/image"
)

func main() {
	url := image.Build("foobar", "path.jpg").SkipCDN().Blur().Boxed(10, 20).Url()

	fmt.Printf("URL: %s\n", url)
}
