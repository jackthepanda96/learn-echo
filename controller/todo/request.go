package todo

type ToDoRequest struct {
	Kegiatan string `json:"kegiatan" form:"kegiatan" validate:"required"`
}
