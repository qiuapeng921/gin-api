package server

type Frame struct {
	fd   int
	data string
}

func (f *Frame) SetFd(fd int) {
	f.fd = fd
}

func (f *Frame) GetFd() int {
	return f.fd
}

func (f *Frame) SetData(data string) {
	f.data = data
}

func (f *Frame) GetData() string {
	return f.data
}
