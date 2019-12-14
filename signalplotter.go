package splot

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// SignalPlotter represents a plotter to plot signals.
type SignalPlotter struct {
	Title      string
	XAxisTitle string
	YAxisTitle string
}

// NewSignalPlotter creates and returns a new signal plotter.
func NewSignalPlotter(title string, xAxisTitle string, yAxisTitle string) *SignalPlotter {
	return &SignalPlotter{
		Title:      title,
		XAxisTitle: xAxisTitle,
		YAxisTitle: yAxisTitle,
	}
}

// PlotSignals plots signals and saves the image to a PNG file.
func (p *SignalPlotter) PlotSignals(signals [][]float64, titles []string, filename string) error {
	pe, err := plot.New()
	if err != nil {
		return err
	}

	pe.Title.Text = p.Title
	pe.X.Label.Text = p.XAxisTitle
	pe.Y.Label.Text = p.YAxisTitle

	var plots []plotter.XYs
	for _, signal := range signals {
		plotValues := make(plotter.XYs, len(signal))

		for i, v := range signal {
			plotValues[i].X = float64(i)
			plotValues[i].Y = v
		}

		plots = append(plots, plotValues)
	}

	for i, plot := range plots {
		if err := plotutil.AddLines(pe, titles[i], plot); err != nil {
			return err
		}
	}

	// Save the plot to a PNG file.
	if err := pe.Save(16*vg.Inch, 8*vg.Inch, filename); err != nil {
		return err
	}

	return nil
}
