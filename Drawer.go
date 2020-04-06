package go3dplot

//PlotType is a plot type
type PlotType string

const (
	//PM3D is a constant for "pm3d" plot type
	PM3D PlotType = "pm3d"
	//LINES is a constant for "lines" plot type
	LINES PlotType = "lines"
	//POINTS is a constant for "points" plot type
	POINTS PlotType = "points"
)

//Format is a output file format
type Format string

const (
	//PNG is a constant for "png" format
	PNG Format = "png"
	//PDF is a constant for "pdf" format
	PDF Format = "pdf"
)

//Range ...
type Range struct {
	start float64
	end   float64
}

//Drawer interface
type Drawer interface {
	Draw(x, y []float64, u [][]float64, filename string) error
	SetPlotType(plotType PlotType)
	SetFormat(format Format)
	SetWidth(width uint)
	SetHeight(height uint)
	SetFont(font string)
	SetFontSize(fontSize uint)
	SetTitle(title string)
	SetXLabel(label string)
	SetYLabel(label string)
	SetZLabel(label string)
	SetXRange(rng Range)
	SetYRange(rng Range)
	SetZRange(rng Range)
}
