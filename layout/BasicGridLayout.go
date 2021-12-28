package layout

import (
	"github.com/bringdesk/bringdesk/widgets"
	"log"
)

type GridItem struct {
	widgetName string
	x          int
	y          int
	width      int
	height     int
}

type GridRow struct {
	item   []*GridItem
	x      int
	y      int
	width  int
	height int
}

type BasicGridLayout struct {
	row      []*GridRow
	width    int
	height   int
	paddingX int
	paddingY int
}

func (self *BasicGridLayout) inititlize() {
	/* Row 1. Provide 6 cells */
	row1 := NewGridRow()
	for range []byte{0, 1, 2, 3, 4, 5} {
		row1item := NewGridItem()
		row1.item = append(row1.item, row1item)
	}
	self.row = append(self.row, row1)
	/* Row 2. Provide 2 wide cells */
	row2 := NewGridRow()
	row2wide1 := NewGridItem()
	row2.item = append(row2.item, row2wide1)
	row2wide2 := NewGridItem()
	row2.item = append(row2.item, row2wide2)
	self.row = append(self.row, row2)
	/* Row 3. Provide 1 cells */
	row3 := NewGridRow()
	row3wide := NewGridItem()
	row3.item = append(row3.item, row3wide)
	self.row = append(self.row, row3)
}

func NewGridItem() *GridItem {
	return new(GridItem)
}

func NewGridRow() *GridRow {
	return new(GridRow)
}

func NewBasicGridLayout() *BasicGridLayout {
	newBasicGridLayout := new(BasicGridLayout)
	newBasicGridLayout.inititlize()
	return newBasicGridLayout
}

func (self *BasicGridLayout) SetSize(width int, height int) {
	self.width = width
	self.height = height
	self.paddingX = 10
	self.paddingY = 10
	//
	self.updateRows()
}

func (self *BasicGridLayout) updateRows() {
	//
	rowWidth := self.width
	rowHeight := self.height / len(self.row)
	//
	for rowIndex, row := range self.row {
		/* Setup row height */
		row.x = 0
		row.y = rowIndex*rowHeight + self.paddingY
		row.height = rowHeight - 2*self.paddingY
		row.width = rowWidth
		/* Setup cell width */
		itemWidth := rowWidth / len(row.item)
		for index, item := range row.item {
			item.x = index*itemWidth + self.paddingX
			item.y = row.y
			item.width = itemWidth - 2*self.paddingX
			item.height = row.height
		}
	}
}

func (self *BasicGridLayout) UpdatePos(rowIndex int, index int, widget widgets.IWidget) {
	for curRow, row := range self.row {
		for curIndex, item := range row.item {
			if curRow == rowIndex && curIndex == index {
				log.Printf("Update widget: row = %d col = %d: x = %d y = %d width = %d height = %d",
					rowIndex, index,
					item.x, item.y,
					item.width, item.height,
				)
				widget.SetRect(item.x, item.y, item.width, item.height)
			}
		}
	}
}
