package assets

import (
	"flag"
	"fmt"
	delivery "github.com/seams-cms/go-seams-cms-sdk/delivery"
)

func main() {
	workspace := flag.String("workspace", "", "Your workspace")
	apiKey := flag.String("api-key", "", "Your workspace API key")
	flag.Parse()

	config := delivery.Configuration{
		Workspace: *workspace,
		ApiKey:    *apiKey,
	}
	client := delivery.NewClientWithConfig(&config)

	col, err := client.GetAssetCollection(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("A total of %d assets found\n", col.Meta.Total)
	for i := range col.Entries {
		entry := col.Entries[i]
		fmt.Printf("%02d: %-20s  %-20s  %03d\n", i, entry.Asset.Title, entry.Asset.Link, entry.Asset.Size)
	}
}
