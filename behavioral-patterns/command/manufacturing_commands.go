package command

import "strings"

type ManufacturingContext struct {
	parts []string
}

func NewManufacturingContext() *ManufacturingContext {
	return &ManufacturingContext{}
}

func (c *ManufacturingContext) Snapshot() string {
	return strings.Join(c.parts, ", ")
}

type AddWheelsCommand struct {
	ctx *ManufacturingContext
}

func NewAddWheelsCommand(ctx *ManufacturingContext) *AddWheelsCommand {
	return &AddWheelsCommand{ctx: ctx}
}

func (c *AddWheelsCommand) Execute() string {
	c.ctx.parts = append(c.ctx.parts, "wheels")
	return "Added wheels"
}

func (c *AddWheelsCommand) Undo() string {
	if len(c.ctx.parts) == 0 {
		return "Nothing to undo"
	}
	c.ctx.parts = c.ctx.parts[:len(c.ctx.parts)-1]
	return "Removed wheels"
}

type PaintFrameCommand struct {
	ctx *ManufacturingContext
}

func NewPaintFrameCommand(ctx *ManufacturingContext) *PaintFrameCommand {
	return &PaintFrameCommand{ctx: ctx}
}

func (c *PaintFrameCommand) Execute() string {
	c.ctx.parts = append(c.ctx.parts, "painted frame")
	return "Painted frame"
}

func (c *PaintFrameCommand) Undo() string {
	if len(c.ctx.parts) == 0 {
		return "Nothing to undo"
	}
	c.ctx.parts = c.ctx.parts[:len(c.ctx.parts)-1]
	return "Removed paint"
}
