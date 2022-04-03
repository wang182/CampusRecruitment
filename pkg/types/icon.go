package types

type IconForm struct {
	baseForm

	Path string `form:"path" binding:"required"`
}
