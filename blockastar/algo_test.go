package blockastar

import (
	"fmt"
	"math"
	"testing"
)

func TestRandom(t *testing.T) {
	for K := 4; K <= 4; K *= 2 {
		for percentage := 10; percentage <= 10; percentage += 10 {
			for tc := 0; tc <= 0; tc++ {
				t.Run(fmt.Sprintf("K%d-%d-%d", K, percentage, tc), func(t *testing.T) {
					grid := BuildGridFromMaps(fmt.Sprintf("../maps/random512-%d-%d.map", percentage, tc), K)
					lddb := GenerateLDDB(grid)
					for _, td := range BuildTestDataFromScen(fmt.Sprintf("../maps/random512-%d-%d.map.scen", percentage, tc)) {
						len := Run(&grid, grid.Cells[td.x1][td.y1], grid.Cells[td.x2][td.y2], lddb)
						if math.Abs(len-td.len) > 0.000001 {
							t.Error(td.x1, td.y1, td.x2, td.y2, "Result", len, "differ from", td.len)
						}
					}
				})
			}
		}
	}
}

func TestAren2(t *testing.T) {
	grid := BuildGridFromMaps("../maps/arena2.map", 4)
	lddb := GenerateLDDB(grid)
	for _, td := range BuildTestDataFromScen("../maps/arena2.map.scen") {
		len := Run(&grid, grid.Cells[td.x1][td.y1], grid.Cells[td.x2][td.y2], lddb)
		if math.Abs(len-td.len) > 0.000001 {
			t.Error(td.x1, td.y1, td.x2, td.y2, "Result", len, "differ from", td.len)
		}
	}
}

func BenchmarkRandom(b *testing.B) {
	for K := 1; K <= 512; K *= 2 {
		for percentage := 10; percentage <= 10; percentage += 10 {
			for tc := 0; tc <= 0; tc++ {
				b.Run(fmt.Sprintf("K%d-%d-%d", K, percentage, tc), func(b *testing.B) {
					grid := BuildGridFromMaps(fmt.Sprintf("../maps/random512-%d-%d.map", percentage, tc), K)
					lddb := GenerateLDDB(grid)
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						for _, td := range BuildTestDataFromScen(fmt.Sprintf("../maps/random512-%d-%d.map.scen", percentage, tc)) {
							Run(&grid, grid.Cells[td.x1][td.y1], grid.Cells[td.x2][td.y2], lddb)
						}
					}
				})
			}
		}
	}
}
