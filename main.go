package main

import (
	"fmt"
	"os"
	"time"

	"github.com/guimpo/pdfgen/data"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	fmt.Println("init")
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	buildFruitList(m)

	path := "./pdf-out/" + time.Now().UTC().Format(time.UnixDate) + ".pdf"
	err := m.OutputFileAndClose(path)
	if err != nil {
		fmt.Println("PDF Not Created", err)
		os.Exit(1)
	}
	fmt.Println("PDF Created!")
}

func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("img/path236.png", props.Rect{
					Center:  false,
					Percent: 45,
				})

				if err != nil {
					fmt.Println("Image not found", err)
				}

				m.Text("Company Address ", props.Text{
					Top:   30,
					Style: consts.Bold,
					Align: consts.Left,
					Color: getDarkPurpleColor(),
				})

				m.Text("Invoice Number xxxx", props.Text{
					Top:   3,
					Style: consts.Bold,
					Align: consts.Right,
					Color: getDarkPurpleColor(),
				})

				currentTime := time.Now().Local().Format("2006-01-02 15:04:05")
				m.Text(fmt.Sprintf("Date: %s", currentTime), props.Text{
					Top:   8,
					Style: consts.Bold,
					Align: consts.Right,
					Color: getDarkPurpleColor(),
				})
			})
		})
	})
}

func buildFruitList(m pdf.Maroto) {
	tableHeadings := []string{"Fruit", "Description", "Price"}
	contents, total := data.FruitList(20)
	lightGrayishBlue := getLightGrayishBlue()

	m.SetBackgroundColor(color.NewWhite())
	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      12,
			GridSizes: []uint{3, 7, 2},
			Style:     consts.Normal,
			Color:     getDarkPurpleColor(),
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 7, 2},
		},
		Align:                  consts.Left,
		AlternatedBackground:   &lightGrayishBlue,
		HeaderContentSpace:     1,
		Line:                   false,
		VerticalContentPadding: 3,
	})

	m.SetBackgroundColor(getPastelBlue0())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("Total: %.2f ", total), props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewBlack(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Right,
			})
		})
	})
}

// func getLightPorpleColor() color.Color {
// 	return color.Color{
// 		Red:   210,
// 		Green: 200,
// 		Blue:  230,
// 	}
// }

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

// func getTealColor() color.Color {
// 	return color.Color{
// 		Red:   3,
// 		Green: 166,
// 		Blue:  166,
// 	}
// }

func getPastelBlue0() color.Color {
	return color.Color{
		Red:   174,
		Green: 198,
		Blue:  207,
	}
}

// func getPastelBlue4() color.Color {
// 	return color.Color{
// 		Red:   56,
// 		Green: 85,
// 		Blue:  95,
// 	}
// }

func getLightGrayishBlue() color.Color {
	return color.Color{
		Red:   234,
		Green: 238,
		Blue:  244,
	}
}
