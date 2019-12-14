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

// PlotSignal plots a signal and saves the image to a PNG file.
func (p *SignalPlotter) PlotSignal(signal []float64, seriesTitle string, filename string) error {
	pe, err := plot.New()
	if err != nil {
		return err
	}

	pe.Title.Text = p.Title
	pe.X.Label.Text = p.XAxisTitle
	pe.Y.Label.Text = p.YAxisTitle

	plotValues := make(plotter.XYs, len(signal))

	for i, v := range signal {
		plotValues[i].X = float64(i)
		plotValues[i].Y = v
	}

	if err := plotutil.AddLines(pe, seriesTitle, plotValues); err != nil {
		return err
	}

	// Save the plot to a PNG file.
	if err := pe.Save(16*vg.Inch, 8*vg.Inch, filename); err != nil {
		return err
	}

	return nil
}
