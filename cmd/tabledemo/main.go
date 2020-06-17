package main

import (
	"fmt"
	"fyne.io/fyne/app"
	"git.sr.ht/~charles/fynehax/table"
	"github.com/rocketlaunchr/dataframe-go"
)

func main() {

	s1 := dataframe.NewSeriesInt64("day", nil, 1, 2, 3, 4, 5, 6, 7, 8)
	s2 := dataframe.NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2, nil, nil, 84.2, 72, 89)
	s3 := dataframe.NewSeriesString("string!", nil, "foo", "bar", "three", "four", "five", "six", "seven", "eight")
	df := dataframe.NewDataFrame(s1, s2, s3)

	fmt.Print(df.Table())

	app := app.New()
	w := app.NewWindow("Table Demo")

	w.SetContent(table.NewTableWidget(df))

	w.ShowAndRun()

}
