package main

type FileItem struct {
	Path   string
	Tmpl   string
	IsDir  bool
	Childs []*FileItem
}

func (f *FileItem) Add(fi *FileItem) {
	f.Childs = append(f.Childs, fi)
}