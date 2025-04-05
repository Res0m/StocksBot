package services

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func GenerateCurrencyChart(data map[string]float64, filename string) error {
	p := plot.New()

	p.Title.Text = "График курсов валют"
	p.X.Label.Text = "Дата"
	p.Y.Label.Text = "Курс"

	// Подготовка данных
	points := make(plotter.XYs, len(data))
	i := 0
	for _, rate := range data {
		points[i].X = float64(i)
		points[i].Y = rate
		i++
	}

	// Создание линии графика
	line, err := plotter.NewLine(points)
	if err != nil {
		return fmt.Errorf("failed to create line plot: %v", err)
	}
	line.Color = color.RGBA{R: 255, A: 255}

	p.Add(line)
	p.Legend.Add("Курс", line)

	// Сохранение графика в файл
	if err := p.Save(8*vg.Inch, 4*vg.Inch, filename); err != nil {
		return fmt.Errorf("failed to save chart: %v", err)
	}

	return nil
}
