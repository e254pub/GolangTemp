package xmlFeedParser

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"main/configs"
	conf "main/configs/connects"
	"main/internal/pkg/xmlFeedParser/Attributes"
	"main/internal/pkg/xmlFeedParser/Cards"
	"main/internal/pkg/xmlFeedParser/Offers"
	"os"
	"path/filepath"
)

func Parse() {
	conn := conf.PostgresConnect()

	shopFeed := filepath.Dir(configs.GetDefaultPath()) + "/assets/xml/ShopFeeds/file6.xml"

	xmlFile, err := os.Open(shopFeed)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var c CatalogDef
	if err := xml.Unmarshal(byteValue, &c); err != nil {
		panic(err)
	}

	attributes := Attributes.AttributeParse(c, err, conn)
	fmt.Println("Attributes: ", attributes)

	cards := Cards.CardsParse(c, err, conn)
	fmt.Println("Cards: ", cards)

	offer := Offers.OffersParse(c, err, conn)
	fmt.Println("Offers: ", offer)

	defer conn.Close(context.Background())
}
