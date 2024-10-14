package seed

import (
	"fmt"

	"github.com/ad9311/renio-go/internal/model"
)

var entryClassesSeed = []model.EntryClass{
	{
		UID:   "banking",
		Name:  "Banking",
		Group: 0,
	},
	{
		UID:   "clothing",
		Name:  "Clothing",
		Group: 0,
	},
	{
		UID:   "entertainment",
		Name:  "Entertainment",
		Group: 0,
	},
	{
		UID:   "extra",
		Name:  "Extra",
		Group: 1,
	},
	{
		UID:   "foodDelivery",
		Name:  "Food Delivery",
		Group: 0,
	},
	{
		UID:   "groceries",
		Name:  "Groceries",
		Group: 0,
	},
	{
		UID:   "home",
		Name:  "Home",
		Group: 0,
	},
	{
		UID:   "insurance",
		Name:  "Insurance",
		Group: 0,
	},
	{
		UID:   "onlineShopping",
		Name:  "Online Shopping",
		Group: 0,
	},
	{
		UID:   "other",
		Name:  "Other",
		Group: 0,
	},
	{
		UID:   "restaurants",
		Name:  "Restaurants",
		Group: 0,
	},
	{
		UID:   "savings",
		Name:  "Savings",
		Group: 0,
	},
	{
		UID:   "subscriptions",
		Name:  "Subscriptions",
		Group: 0,
	},
	{
		UID:   "transportation",
		Name:  "Transportation",
		Group: 0,
	},
	{
		UID:   "utilities",
		Name:  "Utilities",
		Group: 0,
	},
	{
		UID:   "wages",
		Name:  "Wages",
		Group: 1,
	},
}

func seedEntryClasses() {
	for _, entryClass := range entryClassesSeed {
		if err := entryClass.Insert(); err != nil {
			fmt.Printf("could not insert entry class: %s", err.Error())
		}
	}
}
