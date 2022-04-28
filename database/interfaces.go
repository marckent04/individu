package database


type BaseEntity interface {
	create()
	update()
	delete()
}