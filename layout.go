package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const Width = 800

type fysionLayout struct { 
	top, left, right, content fyne.CanvasObject
	dividers [3]fyne.CanvasObject
 }

func newFysionLayout(top, left, right, content fyne.CanvasObject, dividers[3] fyne.CanvasObject) fyne.Layout {
	return &fysionLayout{top: top, left: left, right: right, content: content, dividers: dividers}
}

func (l *fysionLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	topHeight := l.top.MinSize().Height
	l.top.Resize(fyne.NewSize(size.Width, topHeight))

	sideWidth := fyne.NewSize(size.Width/6, size.Height-topHeight)
	l.left.Move(fyne.NewPos(0, topHeight))
	l.left.Resize(sideWidth)
	l.right.Move(fyne.NewPos(size.Width-sideWidth.Width, topHeight))
	l.right.Resize(sideWidth)

	contentSize := fyne.NewSize(size.Width-(sideWidth.Width*2), size.Height-topHeight)
	l.content.Move(fyne.NewPos(sideWidth.Width, topHeight))
	l.content.Resize(contentSize)

	dividerThickness := theme.SeparatorThicknessSize()
	l.dividers[0].Move(fyne.NewPos(0, topHeight))
	l.dividers[0].Resize(fyne.NewSize(size.Width, dividerThickness))
	l.dividers[1].Move(fyne.NewPos(sideWidth.Width, topHeight))
	l.dividers[1].Resize(fyne.NewSize(dividerThickness, size.Height-topHeight))
	l.dividers[2].Move(fyne.NewPos(size.Width-sideWidth.Width, topHeight))
	l.dividers[2].Resize(fyne.NewSize(dividerThickness, size.Height-topHeight))
}

func (l *fysionLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	topSize := l.top.MinSize()
	leftSize := l.left.MinSize()
	rightSize := l.right.MinSize()
	contentSize := l.content.MinSize()
	return fyne.NewSize(
		max(leftSize.Width, rightSize.Width, contentSize.Width),
		topSize.Height+max(leftSize.Height, rightSize.Height, contentSize.Height),
	)
}