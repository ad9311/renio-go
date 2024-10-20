package seed

import (
	"fmt"

	"github.com/ad9311/renio-go/internal/model"
)

var entryClassesSeed = []model.EntryClass{
	{
		UID:  "banking",
		Name: "Banking",
		Type: 0,
	},
	{
		UID:  "clothing",
		Name: "Clothing",
		Type: 0,
	},
	{
		UID:  "entertainment",
		Name: "Entertainment",
		Type: 0,
	},
	{
		UID:  "extra",
		Name: "Extra",
		Type: 1,
	},
	{
		UID:  "foodDelivery",
		Name: "Food Delivery",
		Type: 0,
	},
	{
		UID:  "groceries",
		Name: "Groceries",
		Type: 0,
	},
	{
		UID:  "home",
		Name: "Home",
		Type: 0,
	},
	{
		UID:  "insurance",
		Name: "Insurance",
		Type: 0,
	},
	{
		UID:  "onlineShopping",
		Name: "Online Shopping",
		Type: 0,
	},
	{
		UID:  "other",
		Name: "Other",
		Type: 0,
	},
	{
		UID:  "restaurants",
		Name: "Restaurants",
		Type: 0,
	},
	{
		UID:  "savings",
		Name: "Savings",
		Type: 0,
	},
	{
		UID:  "subscriptions",
		Name: "Subscriptions",
		Type: 0,
	},
	{
		UID:  "transportation",
		Name: "Transportation",
		Type: 0,
	},
	{
		UID:  "utilities",
		Name: "Utilities",
		Type: 0,
	},
	{
		UID:  "wages",
		Name: "Wages",
		Type: 1,
	},
}

func seedEntryClasses() {
	for _, entryClass := range entryClassesSeed {
		if err := entryClass.InsertIfNotExists(); err != nil {
			fmt.Printf("could not insert entry class: %s", err.Error())
		}
	}
}
