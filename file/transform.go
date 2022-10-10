package file

func Meta2File(m *MetaData) File {
	var f File
	f.Name = m.Name
	f.Description = m.Description
	f.Tags = m.Tags
	return f
}
