// Code generated by go-bindata.
// sources:
// static/index.html
// static/main.js
// DO NOT EDIT!

package bindata

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x59\xdf\x6f\xdb\xb6\x13\x7f\xef\x5f\xc1\xaa\xf8\xbe\x55\xa2\xfd\x4d\xba\x15\x9d\x6c\xa0\x48\xdb\x61\xc0\xba\x15\x6d\x37\xac\x4f\xc3\x59\x3c\x45\x4c\x29\x52\x23\x4f\x89\x8d\xc0\xff\xfb\x40\xfd\xb0\x15\x59\xb2\x93\xd4\x68\xb2\xbc\x48\xa4\x8e\x77\xbc\x0f\x3f\x77\x3e\x5e\xe2\xa7\x6f\x7e\x3f\xfb\xfc\xe5\xc3\x5b\x96\x51\xae\xe6\x4f\xe2\xfa\xc1\x18\x63\x71\x86\x20\xea\xd7\x6a\x48\x92\x14\xce\x63\x5e\x3f\xb7\xf3\x39\x12\xb0\x24\x03\xeb\x90\x66\x41\x49\x69\xf8\x32\x60\xbc\x23\xa0\xa4\xfe\xca\x32\x8b\xe9\x2c\xc8\x88\x0a\xf7\x8a\xf3\x44\xe8\x0b\x17\x25\xca\x94\x22\x55\x60\x31\x4a\x4c\xce\xe1\x02\x96\x5c\xc9\x85\xe3\x8b\x52\xe5\xc0\x27\xd1\xff\xa3\x13\x9e\xb8\x66\x1c\xe5\x52\x47\x89\x73\x01\xb3\xa8\x66\x81\xa3\x95\x42\x97\x21\x52\x30\x8f\xb9\xb7\xf1\x4d\x26\x53\xa3\x29\x84\x2b\x74\x26\x47\x7e\x1a\xfd\x18\x4d\x2a\xcb\xdd\xe9\xbb\x6c\xc0\x25\x56\x16\xc4\x9c\x4d\x6e\xbd\x83\x8b\x7f\x4a\xb4\x2b\x7e\x12\x4d\xa3\x69\x33\xa8\x2c\x5e\x38\xaf\xbf\x56\xf8\x6d\x16\x32\xd0\x42\xe1\x02\xac\x8b\x2e\x1c\x3f\x8d\x26\xd1\x0f\xdd\xb9\xe3\x1a\x83\xa5\x34\x8e\x4f\xa2\xe9\x8b\xe8\xa4\x1e\x1c\xd9\xc0\x02\x41\xf3\x69\xe4\x2d\x54\xef\x47\x56\xff\xc7\xc7\x5f\x3c\x4a\xd3\x68\xfa\x32\x3a\xad\x46\xc7\xd5\xff\x4e\x2a\xfc\x04\x97\x68\x6b\x2b\x27\xd1\x49\x67\xea\xee\xa6\xa2\x42\x19\x8a\xd4\x8a\xfb\xa7\x5a\x85\x0a\x08\x1d\x8d\x28\x8a\x79\x1d\xdb\xf5\x60\x61\xc4\xaa\x79\xaf\xc6\x42\x5e\xb2\x44\x81\x73\xb3\x20\x43\x6b\x98\x74\xa1\x92\xe7\x19\x05\xdb\x5d\x0c\xc9\x85\x5e\x4f\x4f\xa6\x2f\x97\x18\x4d\x20\x35\x5a\xaf\x34\x55\xa5\x14\x03\x0b\xaa\x45\x4f\xc3\x90\x9d\x65\x60\x89\xbd\xb6\x08\x2c\x0c\x47\xe4\x3a\xca\x17\x66\x39\xa2\x6d\x23\x29\xc5\x2c\xf0\xa9\xaa\xef\xcb\x0d\x41\x2e\xe4\xe5\x88\xb5\xe1\x4f\x03\xd3\xbd\xa9\x66\x38\x88\xf1\x06\x93\xe0\xc9\x08\x02\x67\x46\x93\x35\xca\xed\xa0\x10\x6b\xd8\xa8\x51\x78\x89\xea\x00\xfc\x95\x4c\xa8\x30\x1d\xf3\x7f\x57\x58\x12\xe6\xfb\xc0\x2a\xba\x6e\x58\xa3\x58\x06\x2e\x04\x21\x8c\x76\x7b\x96\x55\x4b\xa5\x2e\x4a\xaa\x8e\xa4\x7a\x0b\x53\xa9\x08\x6d\xc0\xaa\x1f\x98\x59\xf0\xae\x1a\xb2\xc5\x8a\xe5\x48\x56\x26\x4c\x43\x8e\xcf\x19\x24\x09\x16\xe4\x98\xc5\xf3\x52\x81\x65\xb8\x2c\x2c\x3a\x27\xbd\xc1\xbd\xf6\x9a\xbf\x66\xbb\x95\xc9\x80\xd1\xaa\xc0\x59\x40\xb8\xa4\x80\x15\x0a\x12\xcc\x8c\x12\x68\x37\xd6\x6b\xd3\x2e\x8a\xa2\x43\xde\xc0\x86\x86\x25\x91\xd1\x41\xe5\xd8\x82\x74\x98\x28\x04\xdb\x3a\xb7\x5f\x49\xa5\xc8\x15\xa0\x37\xbb\x4c\x8c\xf6\xa1\xe2\x72\x50\x43\x87\x3b\x8c\x6b\xbb\x3a\x05\x96\x42\x98\x28\xe3\xd0\x67\x00\x79\x0b\xe3\xdc\x5b\x3f\xe0\x29\x87\x7d\xc1\x53\xdc\x29\x74\xd8\xbd\x48\x57\x41\xe4\xf1\xad\x51\x45\x11\x26\xa6\xd4\x14\xcc\xff\x6a\x1c\x60\x26\xed\x48\x35\x87\xd8\x0a\x7d\xd9\xe7\xe5\xed\x43\x7c\x78\xeb\xb6\xc9\x94\x87\xb2\x9a\x27\x1e\xab\xa9\xb2\x1b\xd7\xf7\x47\xe6\x1b\xc2\xb1\x47\xe0\xaa\xa4\xa3\xb0\x19\x48\x17\x42\x42\xf2\x12\xb7\xbc\x56\x52\x63\x78\x28\x99\x6e\xb4\x1f\x9b\xd5\x5d\xeb\xdf\x87\xda\xec\x10\x48\x5b\x68\xc0\x22\x3c\x1c\x34\x5d\xeb\x8f\x0e\x9a\x05\xd8\x87\x43\xa6\x63\xfc\x71\xa4\xc3\xf1\x2c\xf1\x89\xc0\x12\x77\x64\x0a\x06\x5a\x30\x8b\xa9\x45\x97\xb1\x02\xad\x34\xe2\x51\x26\x8c\xed\x19\x17\x50\xfa\x5f\x9c\xef\x7f\xbe\x8d\xe1\xef\x48\xfa\x7e\x0d\x53\x9f\xcf\xa6\x86\xf9\x78\xe3\xd8\x9e\x33\xa9\x99\xc3\xc4\x68\x71\xff\x5a\x45\x97\xf9\xc2\x57\x49\xb9\xd4\xb3\x60\x1a\xb0\x4b\x50\x25\xce\x82\x17\x7b\x2b\xda\xfb\xb2\xf0\xbd\x74\x09\xf3\x89\xff\x11\xfd\x4a\xb5\x2c\xc3\x65\x61\x2c\x85\x85\x3e\x0f\x7a\x44\x7c\x00\xe6\x09\x73\xa5\x95\x01\x71\x44\xf2\x6d\xf6\x39\xff\xf0\xdb\xcf\xc7\x4a\xd1\x2d\x78\xa6\x40\x1d\x5a\xb8\xea\x43\xd7\x34\x4d\xb8\xc0\x45\x79\xce\x9b\xc2\x29\x60\x04\xf6\x1c\x69\x16\xfc\xbd\x50\xa0\xbf\x3e\x48\x15\x6b\xc4\x31\x23\x7b\xb3\xc7\xf9\x47\xb8\xfa\x8f\x54\xbd\x47\x89\x9a\xfa\x36\x22\x80\x60\x27\x6a\xce\xfc\x27\xf6\x06\x08\xee\xc2\xa4\x4a\x61\x4f\x97\x3f\x6c\x01\xfa\xbc\xbe\xe7\x9b\x92\x7c\xb5\x26\x5a\x0b\xaf\x95\x3a\x2a\x9a\x83\xd7\x6f\x0d\x97\xf3\x3a\x87\xbd\xd5\x62\xe8\x1e\xbd\x7b\xd9\x7e\x5f\xb3\x9d\xfd\x2a\x1d\xed\x5e\xb5\xdb\xf6\x41\x1b\x13\x87\x1a\x1d\xaa\xcc\x47\xcf\x65\x57\xf2\x50\xe3\xe2\x56\x2d\x8e\x81\x7d\x86\x04\x0b\x55\x05\xce\x28\x11\x47\x20\x3c\xf0\x69\xb4\xe7\xb1\x05\x7d\x07\xcf\x7e\x3b\x64\x33\xec\xf6\xb4\x72\xd8\xd3\xfb\xa2\x55\xb7\xf5\xcc\x9e\x55\xc5\x1c\xbb\xbe\xb9\xbb\x2b\x29\x28\x7b\xc5\xa6\x93\xc9\xff\x7e\xba\xf9\x25\x43\x7f\x3b\x7b\xc5\x5e\x4c\x26\xc5\xb2\xf3\x6d\xbd\x7d\x25\x1b\x35\xc8\x59\x73\x15\xf9\xa4\x85\x0a\x13\x42\xc1\x48\xf4\x0d\x55\xcd\xe1\xab\x46\xe7\xc2\x28\x31\xac\xf2\x59\xb7\x34\x18\xd9\xec\x29\xe6\x83\x8b\x63\xde\xf1\x39\xe6\x9d\x56\x5d\x05\xf3\x67\xcc\x8b\xaa\xd3\xb7\xe1\x6b\x8b\x65\x9f\x02\x21\x35\xa2\xdd\x9e\x07\x5f\x86\xdb\xfe\xef\x56\xa2\xdb\xf4\xf7\x6b\x5b\xf6\xd5\x03\x8f\x09\x59\x59\x60\xbf\x77\x17\xd3\xcd\xff\x18\x6c\xe7\xed\x08\xb5\x28\x9b\xd7\x24\x89\x39\x65\xe3\x32\x67\xa5\xb5\xa8\x89\xfd\xe9\x4b\x9d\x61\xd1\x98\xf7\x8d\x78\xb9\x9d\xed\xc4\x54\x23\xd8\x5f\x7e\x7d\xfd\x0c\x21\xc9\x98\xcf\x8b\xeb\xf5\x90\x07\x2d\x06\x5b\x72\xf8\x45\x32\x65\x2d\x3f\xd6\xeb\x0e\x59\xae\xaf\xb9\x4c\xd7\xeb\xb1\xf0\x27\xd1\xaa\xd3\x90\x63\x30\xbf\xbe\xf6\xcf\xf5\x3a\xe6\x34\x80\x5f\x6f\x49\x55\xf0\xf9\x35\xd5\xcb\xd8\xa2\x5d\x40\x6a\x3f\xb9\xf7\xb3\xe7\x62\xcc\x07\x60\x89\x29\x35\x86\x1e\xd1\x69\xde\xdc\x4e\xcc\x2b\x36\xb6\x81\xd1\xa6\x8a\x27\x31\xaf\xff\x83\xf5\x6f\x00\x00\x00\xff\xff\xdf\x71\xbc\xa3\xd9\x1a\x00\x00")

func staticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticIndexHtml,
		"static/index.html",
	)
}

func staticIndexHtml() (*asset, error) {
	bytes, err := staticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/index.html", size: 6873, mode: os.FileMode(420), modTime: time.Unix(1481054189, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticMainJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x59\x5d\x8f\xdb\xb6\xd2\xbe\xae\x7f\xc5\x14\x2d\x22\x69\x63\xcb\x0e\xf0\xbe\x07\x07\xde\x3a\x40\x92\xa6\x68\x81\x16\xed\x49\xd2\xab\xc5\x62\x41\x4b\x63\x8b\x89\x4c\x2a\x24\xb5\x5e\x37\xf1\x7f\x3f\xe0\x97\x44\xca\x72\x76\xdb\xb3\xc9\x45\xa2\x90\xc3\x99\x67\x86\xf3\xc5\xf1\xfc\xe2\x62\x02\x17\xf0\x92\x48\x84\xa2\x26\x52\xc2\x86\x0b\xe0\xeb\xf7\x58\x28\x09\xfb\x8a\x16\x15\x10\x81\x80\xb7\xc8\x14\x48\xde\x8a\x02\x65\x3e\x81\x8b\xf9\xc4\x92\xff\xbe\x96\x28\x6e\xc9\xba\x46\xf8\x34\x01\xe0\x2c\x35\xa4\x53\xa8\x08\x2b\x6b\x14\x99\x59\x06\x58\x23\x61\x39\x67\xa9\xaa\xa8\x9c\xc2\x80\x66\x02\x70\x9c\x4c\x00\x36\x54\xa0\x3f\x9f\xe7\x79\x49\x14\x89\xce\x9b\xfd\x88\x83\xa7\x32\x1c\x8e\x93\x89\xd3\xe7\x37\x54\x82\x16\x6f\x51\x01\xb2\x82\x34\xb2\xad\x89\x42\x09\xaa\x42\x90\xa8\x80\x6f\x40\x62\x8d\x85\xc2\x12\x76\x86\x54\x82\xe2\xb0\x46\xd8\x0a\xd2\x54\x58\x6a\x1e\x84\x95\xe6\x40\xc1\x6b\x47\xaa\x25\x19\xfb\xa8\x8a\x28\x4b\x1a\x9a\x22\x10\x7a\xa7\x90\x95\x27\xc6\x99\x5f\x5c\x40\x43\x84\xb6\x24\x95\x40\xdc\x01\x09\x94\x49\x45\x58\x81\xce\xde\x5a\xc1\x80\x99\x21\x95\xed\xda\x01\xbf\x98\x4f\x00\x0a\xce\xa4\x12\x6d\xa1\xb8\x48\x2d\x47\x6f\x27\xd9\x36\x28\xd2\xcc\x7c\x6b\x46\xb9\x13\xb8\x72\x92\xfb\x0d\xaf\xf8\x0a\x18\xee\xe1\x2d\xaa\xf0\x94\x51\x75\x05\x9f\xcc\xb5\x44\x9c\xf4\x25\x26\x6d\x53\x12\x85\xc9\x14\xd2\x0c\x56\xcf\x9d\x68\x80\xf9\x1c\x5e\x91\x46\xb5\x02\x8d\xad\x1a\x4e\x99\xb2\x1e\x45\xd8\xe1\xc4\xe4\xee\x90\xde\x4e\x6b\x54\xc0\xc8\x0e\xb5\x86\x21\xba\xac\xe3\xed\x40\x48\xc5\x05\xde\x18\xce\xa9\x3e\x90\xb9\x6d\x07\xd4\x51\x19\x47\xe9\x51\xea\x35\x4b\x78\xcc\x02\x85\xc6\x34\x89\x54\x15\xc8\x4a\x6d\x4d\x38\x42\xe7\xa3\xa4\x2c\x53\x0b\xce\x63\xa3\x1b\x48\xbf\x35\xc7\x2a\x22\xfd\x5e\x0f\x3c\x54\x27\x0f\x4e\x4f\xce\x68\x35\xb2\x6d\xd5\x21\x65\x99\x4c\xe1\xec\xb6\xd3\xc4\xe9\xe9\xe1\x0a\xdc\xf1\x5b\x1c\x41\xfc\x40\xc0\x25\xd6\xa8\x70\x00\xca\x2e\xf6\x9e\x72\x65\xb7\xaf\x4f\x41\x59\xf1\x7f\x1f\xb6\xe2\xdb\x6d\xfd\x0f\x60\xc7\xea\x5a\x9e\x80\xb5\xc4\x98\x6c\x78\x0d\x9d\xdc\xa2\x46\xa2\x2f\xfc\xd3\x49\x9c\xe4\x6e\x6b\x2c\x46\x06\x2a\x19\xca\x24\x1b\x2e\x07\x9a\xf6\xb2\x6e\x34\x97\x4e\x60\x17\x0b\x56\xe8\xf9\x68\x38\x31\x3d\xac\xe0\xea\xba\xd3\xe5\x8b\x72\x03\xeb\x39\x7e\x02\x55\x2b\x58\xac\x6e\x40\xe5\x0f\x8e\x78\xa9\x63\xa0\x21\xbb\x68\x5f\x8d\x60\xfb\xfc\xd9\x83\xb3\x44\x79\xd3\xca\x2a\xed\x74\xa1\x3b\x94\x8a\xec\x9a\x65\x14\x7c\xdd\xf2\xd4\xd1\xdd\x92\xba\xc5\x98\xc6\x81\x8d\xfc\xef\x38\xb8\xa2\xc0\x42\x56\xba\x57\xa7\xa9\xb9\xaa\x0f\xf1\x05\x38\x4b\xfc\x6e\x0a\x60\xfe\x01\x0f\x32\xed\x18\x65\xf9\x8e\x34\x26\xe9\x84\x39\x4f\xab\x2e\xc8\x3e\xd2\x5b\xd3\xf8\x60\x70\x1c\xfb\x34\x76\xb7\xd4\xf4\x86\x97\xc1\xa3\x99\x99\x8f\x5e\xe3\x6c\xda\x51\x1f\xce\x52\x1b\x73\x04\x94\xea\xd0\xe0\x12\x12\x59\x10\xa5\x50\x24\xfd\x86\x46\xb3\x34\x7f\x77\xb9\xb2\xb3\xd3\x69\xd1\x94\x67\x4b\x26\xa9\xeb\xae\x5a\xba\xe2\x48\x05\xec\xb8\x54\x20\xb0\x30\xb5\xe5\xc2\x5e\x92\x9c\xc2\x06\x55\x51\x61\x09\x1b\xc1\x77\x86\xc9\xbc\xc4\x75\xbb\x9d\xef\x3a\x21\xa5\x55\xe3\xa4\x80\xca\x73\xe5\x33\x2c\x7a\x67\xcb\x5d\x5f\xd5\xc2\xb8\xd4\xca\xcb\x3e\x44\xcc\x5a\x67\x6d\x5d\x01\xdb\xba\x0e\x0a\x1d\x0a\xca\x4b\x58\xc1\xff\x2f\x16\x8b\x30\x9a\x6a\x85\xa2\xa7\x36\x4e\x64\xf4\xbd\x3f\x9c\x02\x0f\x35\xc7\x6c\x4c\x76\x6a\x90\x3b\xca\x65\xbe\x45\x95\x26\xb1\x9d\x12\x9f\x31\x73\x55\x21\x4b\x53\x21\x9b\xa8\xe2\x9e\x68\x2d\x64\x63\x7c\x30\x8f\x8b\xec\x98\xce\xb8\x87\x1f\x0d\x88\x8e\xa6\x4b\x3f\x52\x11\x73\xe5\x61\x1c\x04\x9c\x77\x7a\x5f\x66\x61\x71\xb6\x81\x60\xac\x11\xa3\xb0\xb4\x57\xfa\xef\xeb\x80\xda\xe4\xf2\x43\x83\x7c\x93\x5a\x47\x86\xd5\x0a\x12\xd6\xee\xd6\x28\x92\x98\xb1\x65\x6d\xe3\x0e\x92\x2d\xef\x98\xe6\x09\x3c\x35\x50\x23\xe2\xc8\xec\x26\x12\x61\x65\x81\x05\x64\xc7\xc9\xe9\x57\xe4\x29\x27\x19\xc0\xa7\x60\x1f\x43\x83\x8b\x19\xbb\x15\xdf\x3e\xc4\xab\x67\x7b\x93\x2e\x24\x4d\xaa\x25\x42\xf5\x5e\xae\x78\x13\x3a\x39\x65\x4a\x07\x46\x0d\x2b\xd8\x53\x56\xf2\x7d\x2e\x51\xfd\xe2\x16\x1d\x16\x43\xe8\xdd\x6c\x1a\x3a\x76\x98\xcf\x9b\xf4\xa4\xb8\x7a\xe6\xfd\x25\x38\x19\xa6\x5e\x75\x52\x62\xda\xb0\x2c\x05\xe8\xba\xb8\xea\x0a\xac\x44\x75\x63\x61\xa4\x0e\x4d\x58\x6b\xbb\xc8\xb3\x1f\x03\x60\xf2\x46\xb4\x8c\x51\xb6\x4d\x87\x65\x3f\x30\x50\xb7\x62\x0c\x18\x0b\x0f\x39\xc4\xb1\xfa\x92\xf3\x1a\x09\x3b\xd5\xca\x83\xfe\xc9\x04\x7f\x6a\x73\x40\x84\xb9\x4b\x0b\xf6\x63\x32\x7e\xfd\x5a\xef\x56\xd4\x69\xd6\x37\x66\xae\xb7\xec\x2b\xa8\x40\x17\x96\x6f\x70\xfb\xfa\xae\x49\x43\xf6\x9f\x3f\x43\x92\x5f\xb8\x74\xd0\x77\x08\x5d\xbd\x35\x7e\x3b\x5a\xa0\x46\x4a\x50\x58\x13\xa6\xc1\x6a\x58\x61\x6d\x62\x33\x2d\x76\x48\xe2\x9b\xf8\xa5\xce\x1d\x24\xf7\xff\x35\xed\x42\xd8\x8f\xf7\x81\x75\xcc\x9c\x0e\x2e\x4b\x6a\x6c\x02\x73\xbc\xc3\xc2\xad\x18\xf0\x99\xeb\xcc\xbf\x4f\x93\xef\x9c\x6e\xb3\x82\xb7\x4c\x25\x59\xae\xf0\x4e\xa5\x81\x9a\x35\xb2\xad\xaa\xb2\x8e\xde\xf2\xc7\x32\x3e\xe0\x3b\x19\x4f\x3d\x60\x3f\x53\xba\xb2\x24\x59\x5e\xa9\x5d\x9d\x2a\xdc\x35\xa6\xe2\xb9\x40\xbf\x31\xbb\x37\x7e\xb9\x6b\x56\x74\x56\x5b\x46\x2f\x98\xa3\x47\xae\x6f\x85\xb2\xa6\xd5\x8f\x2d\x2d\xc6\x7c\xcf\x2c\x36\x77\x71\xda\x97\xcd\xb2\x36\x6f\x9a\xc1\xb7\xab\xd0\x87\x7a\xa7\xee\x69\xc2\xed\xc0\x99\x4d\xd5\x86\xd9\xa3\xfc\xd1\x6c\xff\xa8\xb9\x52\x94\x6d\xf5\xf7\xe3\x70\xd5\x65\x7d\xb2\x69\x59\xa1\x28\x67\xa6\xd7\xba\xa9\x29\xc3\x9b\xa2\xea\x33\x9b\xb6\x98\xeb\xa2\x63\x6f\x8a\x5a\xb3\x49\xe7\xc1\x7f\x98\xe5\x9c\xe1\x5e\x7f\xa5\x89\xe1\x95\x4c\x0d\x8f\x29\x7c\x3a\x4e\x2d\x1b\xde\x68\x99\xd2\x71\xc9\xb4\xb1\x62\x20\x44\x20\x39\x01\x22\x15\x29\x3e\x60\xf9\x42\xa0\xc6\xe3\x0f\xa4\x4a\x90\x02\x65\x94\x25\xcd\x0a\x3c\x79\x02\xf6\xcb\x3b\x58\x9f\x92\xcc\xf2\xd5\xe2\x5a\x5f\x9c\x4e\x83\x89\xe2\x7f\xa1\xe0\x87\x24\x22\xc8\x37\x5c\xbc\x26\x45\x95\x4a\x14\x54\x17\x9d\xe7\xf0\xc9\x7e\xe6\x87\x5c\xe0\x2d\x0a\x89\x69\x76\x09\x6e\xed\xae\x5f\xd3\x85\xe2\x1b\x53\xae\xd3\x5b\x22\x80\xae\x9e\x5d\x02\xfd\x21\x42\x73\x09\xf4\xe9\xd3\xe8\xf5\x6c\x41\xd1\x10\x14\xc3\x3b\x75\x48\x34\xab\x9e\xd9\xfb\xd5\xe2\x12\xde\xff\x90\xfe\x46\x54\x95\xef\xa8\x37\xc0\x15\xbd\xbe\x4a\x0e\xc9\xb5\x63\x3f\xed\xf8\xcd\x9e\x45\x1b\x59\x76\x09\xef\xad\xe4\x6f\xbe\x09\xa4\x1a\x9a\xab\xf7\xd7\xf0\x74\x75\x72\xf4\xea\xfd\xf5\xa5\x05\x71\xd4\xff\x1c\x1f\xd1\x4a\x8e\xa9\xef\xc8\x0c\xc7\xcb\x89\x95\x11\x78\x5f\x70\xf7\xe9\x17\x3c\xf1\x2b\xb8\xe2\x9a\x88\x7f\x16\x12\x7d\xaf\x66\xed\xc2\x37\x10\xce\xcb\x9c\x39\x74\x8b\xa5\xef\x7a\x4d\x44\xe2\xd4\xfe\x9f\x15\xd0\x5e\x62\xb0\x6f\x98\xae\xd4\x71\x64\x0f\xf4\x73\x5a\x39\x99\xee\x54\x9a\x3d\x7a\xfe\xfa\xf3\xcd\xaf\xf0\x33\x61\x65\xfd\x15\x73\x58\x57\xc2\xbb\x8b\x6a\x5d\xbd\xfe\xf3\xcd\x2f\xe6\x4e\x5a\x37\x7f\xf8\x4f\x8b\xe2\x90\x26\xb6\x4a\x24\x23\x3b\x41\x51\xd0\x19\x25\xbe\x69\x5f\xb7\x24\xfd\x0b\xe1\x39\x2c\xfc\x95\xb6\x39\x29\xcb\x98\xf5\x14\x5e\x08\x41\x0e\xb9\x7e\x65\x9d\xe1\x62\x1f\x64\x69\x96\x65\xee\xfe\x3b\x81\x9e\x22\x2e\x3e\xa1\x14\x07\xd3\xb9\xc2\x80\xde\xb1\xab\xa8\x54\x5c\x1c\xcc\x93\xfe\xad\x32\xa5\xf2\x38\x85\x24\x99\x42\x9b\x2b\xfe\x56\x09\xdb\xaf\xc5\x9e\x4f\x84\xc4\x81\x2d\x3f\x86\xb6\xcc\x3f\x1a\x00\x4a\xb4\xe8\x6d\xf4\x31\x1f\x19\x06\xd9\x07\x44\xbf\xa5\xdf\x10\xd2\x88\x0c\xde\x10\xb1\x5d\x48\x59\xf6\x07\xc6\x06\x44\x23\x33\x98\xa1\xec\x71\xae\xf1\x88\xeb\x18\x54\x6b\xaf\x41\x6c\xe9\xc8\xa8\x7d\x8b\xf9\x31\x32\xf0\x63\x07\xca\x5b\x54\x6d\x03\x4f\xe0\xb5\x19\xe8\xff\x6c\x07\xf0\xf2\x6b\x85\x0c\x65\xd4\xa7\x80\xae\xc1\xd2\x4f\x73\xa3\xff\x68\xa3\xb5\x74\x98\xd6\x44\xc8\xbc\xe0\xbb\x86\xd6\x98\x0e\x9b\xb6\x99\xa7\xf6\xdd\x5b\xe7\xdc\x3f\xba\x99\x9c\x11\xe0\x72\xd7\xb2\xbb\x35\x9b\xc4\x96\xc1\x2d\xca\x8a\xef\x7f\xa5\xec\xc3\x12\x36\xa4\x96\x41\x33\xbc\xe3\x25\xbe\x24\xe2\x65\xab\x14\x67\xf2\x1d\x7f\x63\xc2\x77\x09\x57\x89\xe2\xbf\xec\xc8\x56\x3f\xdd\x12\x89\xac\xd4\x12\xdf\xf1\x57\x35\x6f\xcb\xe4\xba\x3f\x5f\x52\xd9\xd4\xe4\x50\xf3\x2d\x77\xbc\x63\xbf\x98\x06\x26\x90\x4b\xe3\xf9\x6e\xfc\x91\x86\xba\x74\x1e\xe6\xa2\xa3\xfb\x59\x20\x0a\xe1\x6c\x48\x1d\x8f\xb6\xb5\xda\xf7\x90\xb8\xdc\xd6\x51\x75\x6e\xd9\x3d\xa3\xa2\x65\xff\xae\x9c\x40\x18\xcb\xda\x59\xbf\x4f\x4b\x5e\xb4\x3b\x64\x2a\xcb\x05\x92\xf2\x10\x3c\x8d\xad\x37\xe8\x66\x79\x3e\x87\x77\x66\xc2\xeb\x83\xcc\xa2\xd2\x2e\xb3\xaf\x90\x41\x51\xd3\xe2\x03\x65\x5b\xe0\x0c\x54\x85\x3b\xa0\xe6\xdf\xa0\xef\x9e\xcf\xc1\xb8\xc2\x44\xb7\xf4\xbd\x48\xad\x94\x39\xac\xaf\x47\x09\x07\x78\x26\xf8\x1e\x54\x69\x5e\x0f\xc9\x14\x52\x0c\x5e\xeb\xb1\x51\xdc\xdc\x19\x73\x45\xc4\x16\x95\x79\x48\xbc\xe2\x4c\x69\xe6\x13\xff\xf3\xc1\x7c\xae\xc3\x28\x04\x04\x35\x95\xaa\x7f\x00\x0e\x11\x99\x76\x5e\x23\x8a\xdf\x04\x63\x48\x4e\xf3\x01\xe6\x45\x2b\x04\x32\xf5\xce\x42\xb2\xe3\x92\x31\x30\x02\x37\x02\x65\xd5\x3f\x9e\xef\xc3\x61\x09\xef\xc3\xe1\xdf\xeb\xa3\x40\xe0\x02\x9e\x2d\x16\x8b\x08\xcf\xab\x1a\x75\x23\xd9\x4a\x0b\xcb\xf4\x05\xae\xaf\xa8\x91\xdc\xea\x7b\xed\x6d\x17\x5c\x3d\x65\x8a\x14\xea\x4b\x37\xfa\xdd\x5a\xb1\x99\x99\x43\xcc\x34\xbf\x2f\xdf\x65\x38\x5f\x3f\x85\xd7\x23\x23\xac\x3c\xc1\xf2\x30\x10\x0f\x90\x1f\x8b\x76\x4e\x6f\x1a\x3c\xf7\xbb\xe2\x43\x84\x35\xa4\x95\x27\x8e\x6b\x9e\x9b\x05\x67\xee\xb5\xd9\xd1\x01\x0d\x1e\x9b\xd1\x5d\x8e\x0f\x50\x06\xf1\x1e\x0c\x52\x34\x77\xd7\xb7\xbc\xaa\x89\x94\x69\xb2\x21\x0e\x4b\x44\x42\xca\x32\xd8\xaf\xc9\x21\x19\x2d\xac\x67\x12\xcb\x79\x49\x3d\xa7\x51\x41\x01\x90\xe3\xe9\xfd\xfe\x74\x2e\x18\xc7\x9c\x69\x3c\x20\x4f\x1f\xf1\xe6\x49\x9e\x38\xa1\x67\xe2\x95\xb5\x75\x1d\x5d\xfa\xeb\xbb\x86\x6b\x37\x33\xa3\x4a\x5a\x00\xd5\xe5\xe3\x5e\x60\x68\x4e\xcd\x1a\xb6\x1d\xc2\xb2\x7d\x75\x34\x7e\xdc\x96\xb0\x7a\x1e\x4d\x41\x5d\x7b\xef\x8a\x55\xba\x2d\xa7\x83\x99\xaa\x69\x74\x76\x44\x2d\x13\x23\x62\xb0\x57\x21\xdd\x56\x6a\x09\xff\x5a\x2c\x86\x5b\x7b\x5a\xaa\x6a\x09\xcf\xfe\xcf\xcd\xc6\xfd\x9f\x63\x36\x18\x89\x6a\x2f\x37\x45\x25\x9e\x46\x05\x2f\x1c\x37\x58\x24\x8a\xaf\x3b\xea\x5c\x36\x35\x55\x69\x32\x4d\xb2\xab\x67\xd7\x59\x74\x6c\xdd\x6e\xc0\x15\x42\xd3\xf7\xbe\x6c\x37\x1b\x14\xe6\x68\x34\x13\xf2\xf4\xb7\x14\xf7\xbe\xad\xa4\x4c\xfd\xdb\x1c\x4a\xd7\xed\x66\x30\xf2\xb6\xb1\x04\x2b\x58\x5c\x02\x85\x1f\x20\x60\x68\xde\xd1\x10\x0f\xa4\x35\xdb\x2b\x7a\x0d\x2b\x4b\xa8\xb3\xc8\x2b\x5e\xe2\x0b\x95\xd2\x0c\x9e\xc0\xe2\x6e\xb3\x19\x19\x2f\x1b\x05\x6a\xbe\x76\x80\x5e\xd6\x7c\x9d\x5e\x69\x56\xd7\x53\xf8\x64\x7e\xbc\x49\x8c\x67\xcc\xf5\x7d\x1c\x7b\x84\x7e\xc2\x4b\x6e\xf1\x85\x4c\x35\x87\x29\x24\x5b\x3e\xf3\x9e\xa7\xc9\xe3\x01\x72\x57\x17\xf6\x54\x15\x15\x28\x0e\xfa\x7d\x66\xb3\xdd\xbd\x7e\xa7\x49\x67\xfe\x31\x38\x0c\x07\xa3\xab\x9a\xad\x4d\x57\x94\x64\x71\xc8\x52\x39\x23\x85\xa2\xb7\x3e\x28\x7d\x56\x0a\x38\x66\x41\x08\x0f\xc9\xcf\x3f\x29\x23\xa7\x1f\xd1\x8e\x08\x24\x0f\xd4\x4e\x93\x3e\xae\x76\x01\xc7\xbf\xa1\x5d\x3f\x81\xba\x57\xbb\x35\x11\x0f\x54\x6e\x4d\xc4\xe3\xea\xd6\x33\xfc\x1b\xaa\x75\x13\x8d\x33\x9a\xbd\xc1\x52\x90\xbd\xa9\xbf\x7a\x53\xb7\x78\xd6\xc5\x41\xa0\x7e\xf2\x1a\x35\xed\x8a\x86\xa9\x97\x52\xdb\xc1\x1e\xb3\xc9\x7f\x03\x00\x00\xff\xff\x5b\xb8\xf6\x15\xb0\x24\x00\x00")

func staticMainJsBytes() ([]byte, error) {
	return bindataRead(
		_staticMainJs,
		"static/main.js",
	)
}

func staticMainJs() (*asset, error) {
	bytes, err := staticMainJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/main.js", size: 9392, mode: os.FileMode(420), modTime: time.Unix(1481052264, 0)}
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
	"static/index.html": staticIndexHtml,
	"static/main.js": staticMainJs,
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
	"static": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{staticIndexHtml, map[string]*bintree{}},
		"main.js": &bintree{staticMainJs, map[string]*bintree{}},
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

