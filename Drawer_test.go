package go3dplot

import "testing"

func Test(t *testing.T) {
	var drawer = GetGnuplotDrawer()
	drawer.SetPlotType(PM3D)
	drawer.SetTitle("this is title")
	drawer.SetWidth(800)
	drawer.SetHeight(300)
	drawer.SetFont("Times New Roman")
	drawer.SetFontSize(16)
	drawer.SetFormat(PNG)
	drawer.SetXLabel("xxx")
	drawer.SetYLabel("YY")
	drawer.SetZLabel("ZZzz..")
	drawer.SetXRange(Range{-1, 1})
	drawer.SetYRange(Range{-1, 2})
	drawer.SetZRange(Range{0, 2})
	err := drawer.Draw(
		[]float64{0, 1},
		[]float64{0, 1},
		[][]float64{
			{0, 1},
			{1, 2},
		},
		"test",
	)
	if err != nil {
		t.Error(err)
	}
}
