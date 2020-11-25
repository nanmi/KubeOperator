// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package i18n generated by go-bindata.// sources:
// locales/en-US/home.yml
// locales/zh-CN/home.yml
package i18n

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _localesEnUsHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x57\xcd\x92\xa3\x36\x10\xbe\xf3\x14\x1a\x7c\x4d\x6d\xe5\xec\x9b\x06\x34\x1e\xb2\x18\x5c\x80\x27\x99\x5c\x28\x19\xda\xb6\x32\x20\x51\x92\xd8\xc9\xec\x2d\xef\x95\x77\xca\x2b\xa4\x24\x64\xc0\xf6\x4c\xc6\x5b\x39\xed\x52\xa3\xaf\x7f\xbe\xee\xfe\xba\xbd\xa8\x44\xdb\x0a\xee\x25\x78\x4d\x4a\xf2\x5b\x94\x17\xf9\x12\xf9\x09\x6d\x01\xd1\x46\x02\xad\xdf\x10\xfc\xc9\x94\x56\xbe\x17\x6d\xca\x24\x2d\xa6\x47\x9b\x06\xa8\x02\xb4\x67\x4d\x83\x18\x47\xfa\x08\xa8\x11\x15\x6d\x50\xb4\x41\x62\xf8\x56\x6f\x4a\x43\x8b\x14\x68\xcd\xf8\x01\x75\xf4\x00\xbe\xe7\x2d\xaa\xa6\x57\x1a\xa4\x17\xc4\xdb\xbc\x20\x59\x19\x92\x98\x14\xa4\x7c\xc0\x51\x4c\xc2\x25\xf2\x2b\xca\x11\x17\x1a\xd5\xd0\x80\x06\xe4\x9e\x1b\x27\x55\x2f\x25\x70\x8d\x94\xa6\x1a\xfc\xd1\x40\x94\xdb\xd0\xb2\x6d\x92\x44\xc9\x6a\x89\xfc\xe2\x38\x83\x29\x6b\x4c\xf6\x9c\x33\x7e\xb8\x02\xc5\x69\x80\xe3\x25\xf2\xa3\xb6\x13\x52\x8f\xa8\x8a\x72\x83\xda\x01\xea\xbb\x83\xa4\x35\xd4\x36\x72\x09\x35\x70\xcd\x68\xe3\x9d\x05\x5d\x66\x24\x4f\xb7\x59\x40\x96\xc8\x7f\xa0\xac\x81\x1a\x69\xe1\xe2\xbf\x43\xc5\x11\x24\x98\x38\x28\x47\x54\x29\x51\x31\xaa\xa1\x46\x47\xa1\x34\xea\x79\x0d\x12\xe9\x23\x53\xe8\x05\xde\xfc\x0f\xcc\x96\xbf\xa7\xc9\x0f\xd9\xfe\x2e\x38\xbc\x63\xfb\x01\x6f\xe3\xa2\x0c\x32\x12\x92\xa4\x88\x70\x5c\x06\x38\xb1\x2c\x0c\x6e\x97\xc8\x0f\x61\x4f\xfb\x46\xa3\x29\xd3\x19\x15\x83\xd3\xda\x1f\xda\x25\x78\x24\xc1\xd7\xa9\x6a\x96\xf3\x09\xc5\x4d\x0f\x4d\x50\xdb\x0e\xb6\xad\x94\xfd\x7f\xaf\x40\xda\x37\xbe\xe7\x79\x0b\x43\x85\xf7\x98\xe6\x45\x89\xe3\x8c\xe0\xf0\x79\x6a\xb3\x47\xc3\xd2\x65\x2f\x3a\x96\x2c\x62\x0c\xe0\x5d\x72\x06\x96\x0d\x3f\xce\xc4\x8c\xa4\x57\xa6\x8f\x36\x18\x57\xf4\xf7\xec\x96\xf7\xcf\xe5\x26\x4b\x7f\x21\x41\xf1\xbf\x5c\x74\x52\xfc\x01\x95\xf6\xbd\xfc\x39\x2f\xc8\xba\x74\xd3\xf4\x90\x6e\x93\xf0\x93\x61\xda\x33\xa9\xf4\x3f\x7f\xff\xe5\x7b\x49\x6a\x70\xf8\x09\x47\x31\xbe\x8f\x4d\xb9\x12\x81\xa2\x0e\xd1\x6f\x94\x35\x74\xd7\x80\xef\x45\xf9\xd0\xd1\x36\x87\xd9\x2c\xb1\xa1\xbd\x07\xa3\x26\x60\xd3\xce\xa6\x08\xde\x06\xe7\xf9\xaf\x69\x16\xda\x78\xd6\xb8\x08\x1e\x5d\x31\x3b\xaa\xd4\xab\x90\xb5\x49\x8d\xf1\x4a\x48\x69\x13\x48\xb3\x68\x15\x25\x38\xbe\x7a\x2f\x24\x3b\x30\x4e\x9b\x8f\x80\xdb\x7c\x9a\x3a\x1c\x14\xd1\x13\x71\x40\xdb\x0b\x6e\x4c\x81\x9b\x3c\xea\x3b\xe4\x18\xa9\x04\xd7\xb4\xd2\x96\x11\x5a\xb7\x8c\x33\xa5\x25\xd5\x42\xde\x39\x83\x73\x16\x13\x81\x54\x5f\x1d\xad\x41\x4b\x18\x0e\xd7\x51\x72\xdd\xe4\xc6\x69\xed\x1a\xdd\x1a\x1d\x42\xb8\x6a\xf4\xbb\xf3\xa0\x33\x12\xe3\x82\x84\xb3\x7e\xd8\x1a\xd8\x91\x9a\xd0\xe7\x55\x77\xc5\xb6\x21\xc4\x21\xde\x8c\x11\x6c\x37\x21\x1e\x23\x68\x6a\xda\x5d\x3a\x86\x9a\x0d\x7e\x9f\x48\x16\x3d\x3c\x97\x41\x1a\xce\x84\xf1\x09\x24\xdb\xb3\x8a\x6a\x26\x38\xaa\x44\x0d\x08\xa4\x14\xd2\xf7\xc8\x1a\x47\x71\x19\x46\xb9\x6b\x8b\x35\x65\xcd\x49\x76\x15\xa2\x12\x50\xcd\xd4\x8d\xc4\x9e\xac\xcd\xcb\x4b\x5a\x63\xb0\xa5\xba\x3a\xa2\xbd\x9d\x81\xa1\x1f\x8d\x0c\x8c\xfd\x93\x9b\xaf\x31\x56\x43\xcd\x7f\x68\xc0\xa9\x47\x2e\x8d\x90\x2c\x4b\xb3\x25\xf2\x5f\xa5\xe0\x87\x49\x25\x90\x90\x33\x88\xb7\x90\x70\x60\x82\x9f\xe6\x35\x23\xab\x28\x4d\x6e\xd5\x62\x34\x80\x3f\x9b\x58\x23\xa1\xc6\x95\xf9\xf7\xe4\xc8\xc8\xf0\xcd\x6e\xac\x06\x7f\x26\x0b\x0d\xe5\xe7\x2b\x69\x1b\x9b\x6d\x14\x0c\xa4\x1d\x40\xcf\x05\x6a\x2a\xde\x11\xaa\x97\xe1\x2f\x82\xef\xd9\xa1\x97\xb6\x27\x6c\x51\xa2\x35\x5e\x91\x8f\x4d\xb1\x96\x1e\xe0\x26\x43\x9e\xb7\x10\x1d\x70\xa5\x69\xf5\xe2\xad\x48\x71\xa2\xf9\x54\xa2\x44\x9c\x98\xdc\x8b\x9e\xdb\xb2\xb8\xc6\x5f\x43\xbb\x03\x39\xce\x0e\x0e\xc3\xf9\xac\xec\x00\x38\xa2\xb5\x5b\xa9\x0e\x32\x2e\x27\x37\x5c\x1f\x6f\x26\x07\x78\x6f\x2d\x9d\xb0\x8f\x38\x2f\x1d\xa7\x46\x5b\x1d\x60\xc6\xff\xc8\x67\xf0\xce\xc8\x7b\x0b\x2e\x6a\xf0\x12\x33\x7a\xa7\x75\xe4\x4e\x8b\xb2\xc0\xf9\xd7\x25\xf2\x71\x5d\x23\xf3\xc8\xb4\xa5\xbb\x52\xec\xe7\xa9\xd4\xee\xd8\xf8\xa9\x1b\x58\x7e\xa5\x4c\x23\xa6\x51\x2d\x38\x7c\x31\x0e\x76\xb4\x7a\xe9\x3b\x5c\x55\xa2\xe7\xda\xdb\xe0\x0c\xaf\x4b\xb2\xde\x14\xcf\xe6\x10\xe1\xaa\xdf\xef\x59\xc5\xcc\xa1\xd3\x51\x49\x5b\xd0\x20\x95\x51\xff\xa2\xcc\xb7\x9b\x4d\x9a\x59\xe9\xe1\xaa\xef\x8c\xa6\x9b\xc6\x7b\xeb\xcc\x35\x74\xbe\x8f\xcf\xc4\x62\x18\xda\x71\xc1\xdd\xe3\xe0\xeb\x76\x53\xe2\x20\x48\xb7\xc9\x8f\xac\xba\xb3\xc0\x6f\xde\x79\xde\xc2\xf4\xf9\xc5\x6d\x73\x83\x37\x83\xfa\x01\x27\xae\xaa\xf7\x36\x46\xcf\xe5\xf8\x10\xc5\x64\x90\x6f\x37\x0f\xae\xf3\x5d\xd9\xf4\x98\x94\x59\xbe\x80\x76\xb0\x17\x12\x90\x7a\x65\xba\x3a\x9a\xa3\x75\xf6\x80\x0e\x69\x9f\xcd\xeb\xe0\xe5\xfa\xe4\xdc\x81\x01\x1b\x20\xd4\xa8\xef\xec\x40\xcd\x60\x19\xc9\x8b\x34\x23\xd7\x38\x09\x4a\x0b\xc9\xf8\x61\x18\xc1\xd3\x78\x64\xa0\x44\x2f\x2b\xb8\xe6\x70\x96\xe6\xa7\xc9\x4d\x97\xc4\xfb\x87\xce\x34\x33\xe3\x59\x73\xa2\x7e\x07\x8d\x30\xbb\x44\x8b\x73\x41\x2a\xcc\x75\x29\x3a\x90\x6e\x29\x8d\xf3\xd4\x81\xdc\x0b\xd9\xba\x89\x32\xbb\x6e\x58\x85\xe3\x42\x1d\xaa\x91\xbf\xf1\xea\x28\x05\x67\xdf\x0d\x4f\x0a\xe4\xb0\xae\x7e\x76\x8b\x33\x4e\x57\x51\x72\x89\xd9\xce\xef\x05\x73\x11\xde\xb9\xd7\xd3\x02\x2c\xa6\x5f\x1f\x9d\x14\x47\xb6\x63\x5a\x21\xf3\xc6\xf9\xd8\x4b\xd1\xa2\x46\x1c\x0e\xa6\x4a\x8c\x7f\xb9\xe5\xdc\xf0\x16\x15\x53\x5e\x10\xe5\x56\x05\x2e\xa5\xc1\xdc\x5a\x4c\x21\x4d\xd5\xcb\xa5\x0c\x18\xe8\xb7\x36\xb0\x02\xeb\x3d\xad\xcb\x20\x4d\x1e\xa2\xd5\x74\xe2\x06\x73\xe9\xbd\xba\x75\x27\xc0\xe5\xef\xa4\xe2\x52\xb6\x3f\xaa\x56\x0d\x5d\x23\xde\x5a\x2b\x28\x0d\xe5\xb7\x56\xcd\xf3\xfe\x0d\x00\x00\xff\xff\xfb\xcf\xa1\x3a\x1f\x0e\x00\x00")

func localesEnUsHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYml,
		"locales/en-US/home.yml",
	)
}

func localesEnUsHomeYml() (*asset, error) {
	bytes, err := localesEnUsHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yml", size: 3615, mode: os.FileMode(420), modTime: time.Unix(1606210898, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x4f\x57\xda\x58\x14\xdf\xbf\x4f\xc1\x81\xed\x9c\x39\xb3\xee\xee\x19\x9e\x9a\x69\x48\x38\x49\x70\xc6\xd9\xe4\x58\x64\x3a\x4e\x15\x3c\xfe\x99\xc5\xac\x8a\x56\x45\x0b\xc2\xb4\xe8\x68\x65\xaa\x58\xac\x9c\x5a\xc0\xd6\x56\x90\x88\xfd\x32\x79\x2f\x61\xd5\xaf\x30\xe7\xe5\x25\x21\xc0\xd0\x71\x49\xb8\xbf\xfb\xe7\x77\x7f\xf7\xde\x17\x8a\xa7\x16\x16\x52\x49\x20\xc2\x08\xd2\xd0\xcf\xbc\xa2\x2a\x0f\x02\x41\x5c\xc8\x99\xe7\x97\xb8\xf9\x11\xd7\x0e\x70\xa9\x1a\x04\x7c\x54\x13\x25\xb5\x67\x60\x35\x9a\xb8\x54\x35\xaf\x74\x53\x3f\xb6\xea\x77\x66\xa7\xde\x2d\x7f\xee\xfe\x73\x8a\xcb\x17\x78\xf3\x90\x7d\xe7\xa3\x41\x00\x42\xf1\xf9\xd5\xe5\x95\xc4\x12\xe0\x84\x98\xa2\x22\x59\x0b\x23\x01\xa9\x48\x1b\x87\xbc\x80\xc2\x0f\x02\x41\xf2\xf7\x09\xb9\xda\xc3\x99\x93\xee\x61\x05\x77\x5e\xe2\xed\x9c\xb9\x73\x4d\x9e\xa6\xcd\x57\xcf\xba\x47\x9b\xe6\x5d\x25\xe8\x41\x79\xc5\x4e\x42\x8e\x89\x22\x2f\x4e\x3c\x08\x04\x99\x81\xd1\xca\xe1\x52\xd5\xfa\x52\xb0\xca\x59\xa3\x55\xfb\x7a\x9b\x1e\x82\x08\x12\x07\x05\x5a\x57\xe3\x16\x6f\x9c\x31\x98\x13\x38\xb7\x65\xb6\xcf\xed\x44\x97\x12\xb3\x89\xe4\xca\xdc\xcc\x3c\xe8\xcb\x51\x93\x91\x22\xc5\x64\x0e\x51\x3c\x4b\xb3\xf2\xc1\xfa\x74\xf6\xf5\x36\x6d\x35\xce\xcc\xf3\x83\xee\x8b\x33\xa3\xf5\x9c\x94\xb6\xf1\xc6\x95\x95\x2e\x1a\x2d\x9d\x94\xda\xc1\x11\x4e\xb4\x5f\x24\xf1\xbe\x9e\x70\xbe\x61\x16\xab\x38\x6b\x3b\x1b\x87\x31\x41\xd5\x38\x19\x85\x91\xa8\xf2\x50\xd0\x38\x28\xda\xb5\xb1\x38\x94\x0d\xfd\xc0\xaa\x57\xf0\x56\x8d\xe4\xea\x46\x2b\x67\xad\x77\x58\x10\x9b\x10\xbb\xbf\xdc\x24\xe2\x1e\xf6\xa8\x67\x11\x59\xaf\x19\xc0\x68\xed\x9a\xc5\x2a\xc9\x34\xe9\xc7\xa3\x16\x2e\x64\x29\x33\xbf\xa5\x96\x57\xc0\xa4\xa4\xa8\x1a\x14\x64\x04\xc3\xd3\x3d\x1d\xb0\x62\x7d\x42\x71\xaa\xb6\xad\xbd\x40\xc3\xc5\x7a\x38\x53\xcf\xb3\x62\xdd\x66\x0f\x3b\xd0\xc6\xa6\xb5\xa8\x2c\xfd\x88\x38\xf5\xbe\xbe\xca\x37\xe6\x51\xdd\xae\x5b\x99\x56\x54\x14\xd1\x1c\xfd\x8e\x4b\x31\x31\xec\xc8\x77\x23\xc3\xc4\x4a\x4a\xef\x49\xa9\xcd\x47\x19\x4d\x12\x35\x85\x53\x90\x17\xe0\x98\x40\x59\xe5\xa3\x01\xeb\xf3\x33\xd2\x2e\x50\x86\xae\xaf\x82\x80\x57\x98\x9c\xec\x14\x7b\x02\xb6\x95\xc5\x7c\x05\x01\x00\xa1\xd5\xe5\xc4\x12\x88\x42\x45\xf9\x49\x92\xc3\x76\xf0\x08\x54\xb9\x49\x5b\x84\x9b\xe6\x49\xba\x5b\x3c\xb4\x1a\x8d\x20\x90\x64\x7e\x82\x17\xa1\xd0\x6f\xb2\x7b\xdc\x6f\x15\x53\x7a\x62\x86\x9c\xca\x4f\xd1\xd4\x58\xab\x48\xe9\x1d\x2e\x50\xa1\xd8\x6c\x34\xad\x74\x91\x8e\x5f\xbd\x6c\x16\x36\xf1\x5f\x07\x76\x55\x36\xda\x5f\x3f\xcd\xb9\x56\x61\x78\xdb\x02\x86\x23\xbc\x38\x4a\x52\x81\x99\xd9\x85\xb9\x64\x80\x99\x33\xa1\x58\xa7\x17\x3e\x71\xf9\xb3\x93\x91\x00\x55\x14\xf6\x75\xcc\x49\xf3\x63\xd9\x13\x36\xeb\x4f\x10\x08\x61\x18\xf5\x82\xc6\xa2\x61\x68\x07\xa5\x5f\xfb\x82\x19\x5f\xea\xa4\x78\x63\x47\x9a\x42\x32\x3f\x3e\xad\x71\x52\xd8\xb7\x42\xba\xef\xb2\x56\x23\xed\x63\x0b\x45\x20\x2f\x68\x61\x5e\x71\x7a\xd8\x5d\xab\x1b\xfa\x35\xdb\x53\xd6\xe9\x85\xf9\x36\x3d\x8a\x2e\x17\xeb\x6f\x06\x43\xe3\xec\x4d\x77\x23\xe7\xe9\xce\x99\x28\xaf\xc1\x0a\xfd\xd5\x9b\x2c\x77\x88\xbc\xb1\x62\xed\x74\x67\xaa\x1f\x8b\x64\x59\x92\xfd\x28\x92\xd9\xef\x6f\x3f\x00\xa1\xa5\xc4\xe3\xb9\x54\xd2\x1d\x0f\x19\x4d\xf0\x92\x78\xaf\x0d\x85\xb3\x6d\x7c\x7c\xec\x1f\x0f\xdf\x5e\x01\xa1\x3f\x53\xc9\x84\xeb\x95\xee\xa6\xfb\xf9\x74\x3d\xf4\x4d\xdd\x7a\xd5\xec\x7c\xb4\xea\x65\x9c\x79\xd1\xbf\x81\x63\x02\x5d\xbe\xd6\x6e\x13\xe7\xf7\xe9\xb6\xce\x5c\xb2\x69\x67\x2d\x20\x6f\x9e\x92\xe3\xb3\xee\x46\xce\xec\xb0\x99\xe5\x23\x70\x02\x8d\x02\xee\x95\xf0\x7a\x7e\x04\x10\x84\x52\x8b\x89\xe4\xf2\xca\x4c\xfc\x09\x98\x40\xaa\xcb\x92\xcb\x2f\xd9\xbe\x63\x4e\x18\x25\xd4\x7e\x71\x29\xf5\x7b\x22\xbe\x12\x49\x2c\x3c\x4a\x2c\x79\x3a\x86\x61\x5f\x13\xed\x0a\x49\x53\xc7\x3b\x27\x3e\x84\xb7\x91\x1d\x9d\x8f\x9a\x1d\x26\xf5\xa1\x75\xec\xa2\x26\xa1\xa2\x39\x4c\x51\x88\x6d\xec\xdf\x87\x5f\x6f\xd3\x43\x58\x10\x4a\xa6\x66\x13\x40\xa4\x33\xe0\x2e\x64\xe7\x26\x6a\x2a\x54\x1e\xd2\x8e\x6d\x5c\x1b\xfa\xbe\xb5\xb3\x66\xae\xdd\x90\xbd\xcb\xee\x56\x9e\xbc\xcc\x19\x9d\x12\xa9\xbd\xc1\xa5\x2a\xd9\x3e\xb7\xca\xd9\xef\x02\x56\xa3\x69\xd6\xb6\xf1\xdd\x06\xae\xaf\x1b\xfa\x7b\xf6\x19\xd7\xb3\xa4\xb1\xf7\x3d\x0d\xf3\x68\x26\xfe\x64\x75\x11\xc6\xe3\xa9\xd5\xe4\x0a\x88\x42\x19\x46\x34\x14\x89\xaa\xd3\x34\x42\x7e\x8d\xec\x5d\xba\x3b\x91\x16\xae\xc4\xa2\x51\x49\x56\xed\xa3\x90\x23\xc5\x06\xc9\xd2\x1b\x6e\x7e\xd0\xf1\xeb\xe7\x41\x30\x70\x7a\xc8\x49\xb9\xfb\x2e\xeb\x9b\x26\x47\x81\x63\x90\x7b\x18\x8b\x6a\x90\xe3\xa4\x98\x78\xdf\x03\x80\x2b\x5b\x86\xde\xb1\x3e\xbd\xc5\xf9\xe6\x88\x33\x00\x42\x8b\xf3\x33\xc9\x81\x9b\xfc\x3f\x6e\xfd\x7a\x1e\x76\xeb\x7b\xdb\x8c\xd9\x3c\x01\x27\xf7\x71\x5e\x40\x6c\x15\xba\xf2\x65\xc7\xc6\xf1\x4f\x53\x25\xfb\x5b\x74\xa3\x6c\xe6\x70\x66\x8b\xe4\x4e\xfd\xf9\xf7\x0d\x0e\xf3\xe8\x3d\x74\x58\xf7\x98\xf5\x7f\x3c\x74\x64\xa4\xa8\x92\x8c\x06\xcc\x49\xfa\x14\x57\x72\xae\xb9\xa7\x5f\x39\xb1\x9c\x5a\x5d\x8a\x27\x86\x29\xf1\x95\xf1\x8d\xe4\xfd\x4d\x1b\xb8\xd5\x3d\x3d\xf7\x5d\xe6\x0f\xaf\x8d\xf6\xee\x80\xaa\xad\x2f\x47\x56\x39\x4b\x6a\x15\x26\x4f\x27\xc5\xf9\xd9\x99\x45\x76\x1b\xbc\xc3\xc2\x98\xc4\x85\x2c\xa9\x9d\xe1\xcc\x25\x15\x96\x73\x20\xda\x3f\x38\x77\x44\x90\x26\x78\x71\x10\xe1\x9d\x11\xf6\x3e\xb1\x03\xd8\xd6\xbd\xe3\xc0\x1e\xab\xe6\xdb\x34\xa9\x9d\xd2\xbf\x18\xc4\x3c\xd4\xbb\x87\x9b\x81\x11\x07\x15\x84\xe2\x73\xcb\x80\xe3\x15\x7b\xe0\x06\xa7\x90\x66\xea\xbe\x88\xc8\xf6\x05\xc9\xe7\x8d\x56\xcd\x7c\xf5\xcc\xd0\x75\xbc\x53\xa6\xe8\x3f\x16\xb8\x54\xf2\xd7\xb9\xc7\x60\x2a\xa2\x71\x92\x38\xce\x4f\xf4\x9e\x54\x6c\x9f\xf9\x9e\x54\x3d\x9b\xc1\xb7\xb3\x67\xda\xe3\xd7\xa7\xd9\x6f\xb3\x0c\x00\xf8\x37\x00\x00\xff\xff\x8d\x37\xfa\x3b\xfc\x0b\x00\x00")

func localesZhCnHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYml,
		"locales/zh-CN/home.yml",
	)
}

func localesZhCnHomeYml() (*asset, error) {
	bytes, err := localesZhCnHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yml", size: 3068, mode: os.FileMode(420), modTime: time.Unix(1606210909, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"locales/en-US/home.yml": localesEnUsHomeYml,
	"locales/zh-CN/home.yml": localesZhCnHomeYml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesEnUsHomeYml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesZhCnHomeYml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
