package myLib

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type DBReader interface {
	Read(fileName string)
	Print()
}

type JsonReader struct {
	Cake []struct {
		Name        string `json:"name"`
		Time        string `json:"time"`
		Ingredients []struct {
			IngredientName  string `json:"ingredient_name"`
			IngredientCount string `json:"ingredient_count"`
			IngredientUnit  string `json:"ingredient_unit"`
		} `json:"ingredients"`
	} `json:"cake"`
}

type XmlReader struct {
	XMLName xml.Name `xml:"recipes"`
	Text    string   `xml:",chardata"`
	Cake    []struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name"`
		Stovetime   string `xml:"stovetime"`
		Ingredients struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text      string `xml:",chardata"`
				Itemname  string `xml:"itemname"`
				Itemcount string `xml:"itemcount"`
				Itemunit  string `xml:"itemunit"`
			} `xml:"item"`
		} `xml:"ingredients"`
	} `xml:"cake"`
}

func (jr *JsonReader) Read(fileName string) {

	// open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read file
	byteValue, _ := io.ReadAll(file)
	err = json.Unmarshal(byteValue, &jr)
	if err != nil {
		log.Fatal(err)
	}
}

func (xr *XmlReader) Read(fileName string) {

	// open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read file.
	byteValue, _ := io.ReadAll(file)
	err = xml.Unmarshal(byteValue, &xr)
	if err != nil {
		log.Fatal(err)
	}
}

func (xr *XmlReader) Print() {

	prnt := xr.XmlToJson()

	// print
	res, err := json.MarshalIndent(prnt, "", "    ")
	if err == nil {
		fmt.Println(string(res))
	} else {
		log.Fatalln(err)
	}
}

func (jr *JsonReader) Print() {

	prnt := jr.JsonToXml()

	// print
	res, err := xml.MarshalIndent(prnt, "", "    ")
	if err == nil {
		fmt.Println(string(res))
	} else {
		log.Fatalln(err)
	}
}

func (jr *JsonReader) JsonToXml() XmlReader {

	var xr XmlReader

	// reserve
	xr.Cake = make([]struct {
		Text        string "xml:\",chardata\""
		Name        string "xml:\"name\""
		Stovetime   string "xml:\"stovetime\""
		Ingredients struct {
			Text string "xml:\",chardata\""
			Item []struct {
				Text      string "xml:\",chardata\""
				Itemname  string "xml:\"itemname\""
				Itemcount string "xml:\"itemcount\""
				Itemunit  string "xml:\"itemunit\""
			} "xml:\"item\""
		} "xml:\"ingredients\""
	}, len(jr.Cake))

	// copy
	for i := range jr.Cake {
		xr.Cake[i].Name = jr.Cake[i].Name
		xr.Cake[i].Stovetime = jr.Cake[i].Time

		xr.Cake[i].Ingredients.Item = make([]struct {
			Text      string "xml:\",chardata\""
			Itemname  string "xml:\"itemname\""
			Itemcount string "xml:\"itemcount\""
			Itemunit  string "xml:\"itemunit\""
		}, len(jr.Cake[i].Ingredients))

		for j := range jr.Cake[i].Ingredients {
			xr.Cake[i].Ingredients.Item[j].Itemname = jr.Cake[i].Ingredients[j].IngredientName
			xr.Cake[i].Ingredients.Item[j].Itemcount = jr.Cake[i].Ingredients[j].IngredientCount
			xr.Cake[i].Ingredients.Item[j].Itemunit = jr.Cake[i].Ingredients[j].IngredientUnit
		}
	}

	return xr
}

func (xr *XmlReader) XmlToJson() JsonReader {

	var jr JsonReader

	// reserve
	jr.Cake = make([]struct {
		Name        string "json:\"name\""
		Time        string "json:\"time\""
		Ingredients []struct {
			IngredientName  string "json:\"ingredient_name\""
			IngredientCount string "json:\"ingredient_count\""
			IngredientUnit  string "json:\"ingredient_unit\""
		} "json:\"ingredients\""
	}, len(xr.Cake))

	// copy
	for i := range xr.Cake {
		jr.Cake[i].Name = xr.Cake[i].Name
		jr.Cake[i].Time = xr.Cake[i].Stovetime

		jr.Cake[i].Ingredients = make([]struct {
			IngredientName  string "json:\"ingredient_name\""
			IngredientCount string "json:\"ingredient_count\""
			IngredientUnit  string "json:\"ingredient_unit\""
		}, len(xr.Cake[i].Ingredients.Item))

		for j := range xr.Cake[i].Ingredients.Item {
			jr.Cake[i].Ingredients[j].IngredientName = xr.Cake[i].Ingredients.Item[j].Itemname
			jr.Cake[i].Ingredients[j].IngredientCount = xr.Cake[i].Ingredients.Item[j].Itemcount
			jr.Cake[i].Ingredients[j].IngredientUnit = xr.Cake[i].Ingredients.Item[j].Itemunit
		}
	}

	return jr
}
