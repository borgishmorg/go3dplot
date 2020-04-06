package go3dplot

import (
	"fmt"
	"os/exec"
)

//GetGnuplotDrawer returns drawer that uses gnuplot to draw
func GetGnuplotDrawer() Drawer {
	var drawer = gnuplotDrawer{
		config:   make(map[string]string),
		format:   "png",
		plotType: PM3D,
		font:     "Verdana",
		fontSize: 14,
		width:    640,
		height:   480,
		density:  100,
	}
	return &drawer
}

type gnuplotDrawer struct {
	config   map[string]string
	plotType PlotType
	format   Format
	width    uint
	height   uint
	font     string
	fontSize uint
	title    string
	xLabel   string
	yLabel   string
	zLabel   string
	xRange   *Range
	yRange   *Range
	zRange   *Range
	density  uint
}

func (drawer *gnuplotDrawer) Draw(x, y []float64, u [][]float64, filename string) error {
	err := drawer.writeData(x, y, u, filename)
	if err != nil {
		return err
	}
	err = drawer.writeScript(filename)
	if err != nil {
		return err
	}

	cmd := exec.Command("gnuplot", filename+".script")
	err = cmd.Run()
	if err != nil {
		return err
	}

	err = removeFile(filename + ".script")
	if err != nil {
		return err
	}
	err = removeFile(filename)
	if err != nil {
		return err
	}
	return nil
}

func (drawer *gnuplotDrawer) writeData(x, y []float64, u [][]float64, filename string) error {
	var data string
	for i, layer := range u {
		for j, v := range layer {
			data += fmt.Sprintf("%f %f %f\n", x[i], y[j], v)
		}
		data += fmt.Sprintf("\n")
	}

	err := writeToFile(filename, data)
	if err != nil {
		removeFile(filename)
		return err
	}
	return nil
}

func (drawer *gnuplotDrawer) writeScript(filename string) error {
	var script string

	script += fmt.Sprintf("#This is auto generated script\n")
	script += fmt.Sprintf("set terminal %s size %d,%d enhanced font '%s,%d'\n",
		drawer.format,
		drawer.width,
		drawer.height,
		drawer.font,
		drawer.fontSize)
	script += fmt.Sprintf("set output '%s.%s'\n", filename, drawer.format)
	script += fmt.Sprintf("set title '%s'\n", drawer.title)
	script += fmt.Sprintf("set xlabel '%s'\n", drawer.xLabel)
	script += fmt.Sprintf("set ylabel '%s'\n", drawer.yLabel)
	script += fmt.Sprintf("set zlabel '%s'\n", drawer.zLabel)

	if drawer.xRange != nil {
		script += fmt.Sprintf("set xrange [%f:%f]\n", drawer.xRange.start, drawer.xRange.end)
	}
	if drawer.yRange != nil {
		script += fmt.Sprintf("set yrange [%f:%f]\n", drawer.yRange.start, drawer.yRange.end)
	}
	if drawer.zRange != nil {
		script += fmt.Sprintf("set zrange [%f:%f]\n", drawer.zRange.start, drawer.zRange.end)
	}

	script += fmt.Sprintf("splot '%s' w %s\n", filename, drawer.plotType)

	err := writeToFile(filename+".script", script)
	if err != nil {
		removeFile(filename + ".script")
		return err
	}
	return nil
}

func (drawer *gnuplotDrawer) SetPlotType(plotType PlotType) {
	drawer.plotType = plotType
}

func (drawer *gnuplotDrawer) SetFormat(format Format) {
	drawer.format = format
}

func (drawer *gnuplotDrawer) SetWidth(width uint) {
	drawer.width = width
}

func (drawer *gnuplotDrawer) SetHeight(height uint) {
	drawer.height = height
}

func (drawer *gnuplotDrawer) SetFont(font string) {
	drawer.font = font
}

func (drawer *gnuplotDrawer) SetFontSize(fontSize uint) {
	drawer.fontSize = fontSize
}

func (drawer *gnuplotDrawer) SetDensity(density uint) {
	drawer.density = density
}

func (drawer *gnuplotDrawer) SetTitle(title string) {
	drawer.title = title
}

func (drawer *gnuplotDrawer) SetXLabel(label string) {
	drawer.xLabel = label
}

func (drawer *gnuplotDrawer) SetYLabel(label string) {
	drawer.yLabel = label
}

func (drawer *gnuplotDrawer) SetZLabel(label string) {
	drawer.zLabel = label
}

func (drawer *gnuplotDrawer) SetXRange(rng Range) {
	drawer.xRange = &rng
}

func (drawer *gnuplotDrawer) SetYRange(rng Range) {
	drawer.yRange = &rng
}

func (drawer *gnuplotDrawer) SetZRange(rng Range) {
	drawer.zRange = &rng
}
