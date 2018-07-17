package backup

// Archiver represents type capable of archiving
type Archiver interface {
	Archive(src, dest string) error
}
