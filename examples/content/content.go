package content

import (
	"flag"
	"fmt"
	"github.com/seams-cms/go-seams-cms-sdk/delivery"
	"strings"
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

	col, err := client.GetContentTypeCollection(nil)
	if err != nil {
		panic(err)
	}
	for i := range col.Entries {
		displayContentTypeDetails(client, col.Entries[i].ApiId)
	}
}

func displayContentTypeDetails(client *delivery.DeliveryApi, contentType string) {
	ct, err := client.GetContentType(contentType, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s\n", ct.Name)
	fmt.Printf("Fields %d\n", len(ct.Fields))
	for i := range ct.Fields {
		fmt.Printf("%02d: %-20s  %-20s\n", i, ct.Fields[i].FieldType, ct.Fields[i].Name)
	}

	fmt.Println("\n\n\n** Member content **")
	coll4, err := client.GetContentCollection("members", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("A total of %d entries found\n", coll4.Meta.Total)
	for idx := range coll4.Entries {
		entry := coll4.Entries[idx]

		printEntryData(entry, 0)

		fmt.Printf("\n\n\n")
		break
	}
}

func max(s string, l int) string {
	if len(s) < l {
		l = len(s)
	}

	return s[:l]
}

func printEntryData(entry delivery.ContentEntry, indent int) {
	prefix := strings.Repeat(" ", indent)

	fmt.Printf("%s", prefix)
	fmt.Println("--------------------------------------------")
	fmt.Printf("%s", prefix)
	fmt.Printf("Entry ID   : %s\n", entry.Meta.EntryId)
	fmt.Printf("%s", prefix)
	fmt.Printf("Created By : %s\n", entry.Meta.CreatedBy)

	for k, _ := range entry.Content {
		content := entry.Content[k]
		fmt.Printf("\n")
		fmt.Printf("%s", prefix)
		fmt.Printf("Key : %s\n", k)
		fmt.Printf("%s", prefix)
		fmt.Printf(" - Is Localized : %t\n", content.IsLocalized())
		fmt.Printf("%s", prefix)
		fmt.Printf(" - Is Content   : %t\n", content.IsContent())
		fmt.Printf("%s", prefix)
		fmt.Printf(" - Is Reference : %t\n", content.IsReference())

		if content.IsLocalized() {
			for locale, _ := range content.Locales {
				fmt.Printf("%s", prefix)
				fmt.Printf(" - LOCALE : %s\n", locale)

				if content.IsReference() {
					e, _ := content.GetLocaleReferences(locale)
					for k, _ := range e {
						printEntryData(e[k], indent+4)
					}
				} else {
					s, _ := content.GetLocaleString(locale)
					fmt.Printf("%s", prefix)
					fmt.Printf(" - LOC S : %s\n", max(s, 30))
					i, _ := content.GetLocaleInt(locale)
					fmt.Printf("%s", prefix)
					fmt.Printf(" - LOC I : %d\n", i)
					b, _ := content.GetLocaleBool(locale)
					fmt.Printf("%s", prefix)
					fmt.Printf(" - LOC B : %t\n", b)

					sl, _ := content.GetLocaleSlice(locale)
					for k, _ := range sl {
						fmt.Printf("%s", prefix)
						fmt.Printf(" - LOC SL %d : %s\n", k, sl[k])
					}
				}
			}
		}

		if content.IsReference() {
			e, _ := content.GetReferences()
			for k, _ := range e {
				printEntryData(e[k], indent+8)
			}

		} else {
			s, _ := content.GetString()
			fmt.Printf("%s", prefix)
			fmt.Printf(" - S : %s\n", max(s, 30))
			i, _ := content.GetInt()
			fmt.Printf("%s", prefix)
			fmt.Printf(" - I : %d\n", i)
			b, _ := content.GetBool()
			fmt.Printf("%s", prefix)
			fmt.Printf(" - B : %t\n", b)

			sl, _ := content.GetSlice()
			for k, _ := range sl {
				fmt.Printf("%s", prefix)
				fmt.Printf(" - SL %d : %s\n", k, sl[k])
			}
		}
	}
}
