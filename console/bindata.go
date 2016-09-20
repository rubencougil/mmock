// Code generated by go-bindata.
// sources:
// tmpl/css/style.css
// tmpl/index.html
// tmpl/js/script.js
// DO NOT EDIT!

package console

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

var _tmplCssStyleCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\xdd\x8e\x9b\x30\x10\x85\xaf\xcb\x53\x58\x5a\xf5\x66\x55\xef\x26\x69\xb4\xad\xe0\x6a\xfb\xb3\xaf\x51\xf9\x67\x00\x8b\xc1\x83\x26\xa6\x24\x8a\xf2\xee\x1d\x13\x52\x2a\xb5\x29\x57\x66\xce\x37\xc7\x47\x07\x2c\xf9\x93\x3a\x17\x4a\x1e\x0c\x11\x74\x0b\xa1\x69\x53\xa9\xb6\x4f\x2f\xd0\x57\xea\xf9\x51\xd5\xe1\xa8\x6a\x62\x95\x5a\x13\xe4\x10\xd3\xe3\x73\x71\x29\x8a\xa2\x4d\x3d\x7e\x50\x7f\xec\xff\x5e\xdd\x6c\xde\x57\x19\x79\x72\x42\x1b\x71\x65\x5d\xe3\x18\xfc\xb9\x78\xd7\x87\xa8\xff\xe6\x06\x13\x01\xf5\xcd\xea\xa6\x7f\xdc\x6d\x86\x63\xa5\x32\x51\x96\x7a\x02\xdb\x85\xa4\x0f\x8e\x09\xd1\x1a\x5e\x2e\x9d\x82\x4f\xad\x78\xed\x04\x15\xf2\x1f\xa0\x4e\x6c\x5c\xb7\xe0\x56\x8e\x0d\xd3\x18\xbd\x76\x84\xc4\xa5\x7a\xf8\xfa\xfa\x7d\xff\xf6\x5a\x5d\x65\x62\x2f\x69\x11\xea\x9c\x6f\x38\xaa\x03\x61\xf0\xea\xc1\x39\x77\xd7\xbd\x1d\x7b\x7b\xdf\x7d\xbf\xfb\x6c\x9d\xf9\xef\x72\xd9\xd2\x4f\xe0\xfb\x16\xdb\x6f\x9f\xf6\x5f\xde\xe6\xa6\x8a\x81\x21\x73\x4b\x45\xdb\x97\xdc\x90\xbc\x67\x83\x1a\x69\x2a\xcd\x98\x28\x0f\xae\xb5\xac\x0d\x1f\x00\xc1\x25\xf0\x3f\x98\x26\x75\x9e\x3f\xa3\xa4\xb9\xf6\x6c\x09\x7d\x75\x11\x26\x71\x88\x8d\xa8\xcb\xc5\x0d\x03\xc4\x4a\x89\x12\x25\x65\x4e\x78\x53\xbc\xe1\x8e\xd8\xc4\x06\x66\xd9\x12\x21\x98\xb8\xea\x16\x47\x58\x16\x11\xd7\x71\x6f\x1a\x90\x1f\x62\x56\x3a\x38\xad\x02\x83\xcf\xc3\x5f\x01\x00\x00\xff\xff\x43\xd7\x46\xd6\x8c\x02\x00\x00")

func tmplCssStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_tmplCssStyleCss,
		"tmpl/css/style.css",
	)
}

func tmplCssStyleCss() (*asset, error) {
	bytes, err := tmplCssStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/css/style.css", size: 652, mode: os.FileMode(420), modTime: time.Unix(1472417660, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x57\x6d\x53\xdb\xb8\x13\x7f\x5d\x3e\x85\xea\xff\xff\x66\x60\x38\xdb\x84\x94\x42\xc1\xc9\x0c\x85\xf4\x18\x86\x14\xda\xe4\xe0\xda\x77\x8a\xad\xd8\x4a\x64\xc9\x48\x72\x1e\x7a\xd3\xef\x7e\x2b\xcb\x8e\x1d\x08\x14\xe6\xe6\x18\x48\xe4\xd5\xee\x6f\x9f\x77\x4d\xf0\xf6\xfc\xfa\x6c\xf8\xed\xa6\x87\x2e\x86\xfd\xab\xee\x56\x90\xe8\x94\x99\x2f\x82\xa3\xee\xd6\x9b\x20\x25\x1a\xa3\x44\xeb\xcc\x25\xf7\x39\x9d\x75\x9c\x33\xc1\x35\xe1\xda\x1d\x2e\x33\xe2\xa0\xd0\x3e\x75\x1c\x4d\x16\xda\x37\xb2\x27\x28\x4c\xb0\x54\x44\x77\xfe\x1c\x7e\x72\x8f\x1c\xe4\x1b\x18\x4d\x35\x23\xdd\x7e\x5f\x84\x53\x04\x08\x4a\x30\x12\xf8\x96\xb8\x05\xd7\x2a\x94\x34\xd3\x48\xc9\xb0\xe3\xf8\x7e\x28\x22\xe2\x4d\xee\x73\x22\x97\x5e\x28\x52\xdf\x1e\xdd\x96\xd7\xda\xf7\xf6\xbc\x94\x72\x6f\xa2\x9c\x6e\xe0\x5b\xa9\xee\x4b\xe5\x53\x1a\x4b\xac\x09\xe0\xec\x7b\xad\x0d\x30\x80\xf3\xd6\x75\xd1\x15\xf0\x28\x0d\x9e\xa5\x19\x65\x24\x42\x98\x47\x08\x98\xe9\x98\xc2\xc3\xd9\x60\x80\x5c\xd7\xa8\x64\x94\x4f\x91\x24\xac\xe3\x28\xbd\x64\x44\x25\x84\x68\x07\x25\x92\x8c\x3b\x8e\x89\x97\x3a\xf6\xfd\x14\x2f\xc2\x88\x7b\x23\x21\xb4\xd2\x12\x67\xe6\xc1\x18\xb4\x22\xf8\x6d\xaf\xed\xbd\xf7\x43\xa5\x6a\x5a\x61\x19\x50\x1c\x44\x21\xb6\xb1\xa4\x7a\x09\x3a\x12\xdc\x3e\x7a\xe7\xb6\xee\x8f\xd2\xe1\xe5\xf5\xe9\x60\x71\x34\x69\x9d\xe6\xbb\xf8\xe0\xee\xfc\x96\xdf\xd0\x7d\x36\xfd\x34\x9e\xcf\x7b\xa7\xf8\x28\x39\x3f\x8f\x26\xdf\x59\x76\x45\xe2\x45\x32\xb9\xed\xf7\x5a\xe3\x78\x72\x77\xf3\x47\x3a\xfd\xa1\x0e\x21\x61\x52\x28\x25\x24\x8d\x29\xef\x38\x98\x0b\xbe\x4c\x45\x0e\x61\xa8\xbc\xbf\xce\x34\x15\x1c\x33\xa4\x13\x92\x92\xff\xda\x57\xb7\xd0\xf2\x9c\xc7\xe3\xab\xbb\xfd\xcf\x7b\x2d\xd6\xbf\x9f\xe0\xe9\xc7\xe9\xa2\xcd\xfc\xfe\x87\x1e\x4e\xf2\x79\x36\x18\x93\xcf\xb3\xdb\xf7\xed\xcb\x03\xf2\x83\xb7\xf3\xef\x3f\x70\x36\xdc\xcb\x0f\x7b\xdf\xd4\x5f\xfd\xc9\x97\xdb\xdd\xbd\x1e\x3f\x90\xbf\xf6\xf8\xd9\x7c\x5f\xe2\x19\x1e\xd8\xe2\xb2\xa1\x68\x56\xda\x6b\x5d\x9f\x3c\xcc\xf2\x64\xa3\xcb\x7b\xe9\x60\x74\x79\xde\xbb\xa0\x98\x8d\xd3\xfc\xe3\xc7\x2f\x37\xef\x4f\xdf\x7d\x91\x99\xbc\x3f\xb8\xbe\x1d\xdf\xb5\x0f\x6f\xbe\x7e\x6d\x4f\x0e\x7a\x57\xf7\x0b\xa5\x5a\xcb\xdb\xfb\x6b\xcd\x49\xc6\x2f\x6e\x6f\x3e\xe0\xcb\xc3\xc5\xe0\x69\x97\xd7\x6a\xbd\xf4\x44\x43\x1f\x97\xed\x5b\x3b\xeb\x58\x07\xc1\x60\xcb\xf5\xb8\x53\x36\x97\x44\x03\xac\xc8\xa7\x2d\x11\x93\xf2\x82\xab\x48\xf2\x6a\xaa\x70\x9c\x02\xf3\x8c\x92\x79\x26\xa4\x6e\xcc\x92\x39\x8d\x74\xd2\x89\xc8\x8c\x86\xc4\x2d\x1e\x7e\x87\x30\x51\x0d\x01\x71\x55\x88\x19\xe9\xb4\x0c\xca\x2f\x5c\x00\x8e\xff\x6f\xa3\x48\x84\x79\x0a\xa8\x68\xc7\x93\x30\xd2\x96\xdb\xe3\x9c\x87\xa6\xc4\xb7\x77\xd0\xdf\x80\x81\xd0\x0c\x4b\x34\x57\xa8\x83\x38\x99\xa3\x3b\x32\x1a\xc0\x8c\x22\x7a\xdb\x99\x9b\xd4\x3a\x68\x17\x31\x11\x62\x23\xe1\x25\x02\xea\x64\x17\x39\x3e\x09\x13\xe1\xec\x9c\x14\xe2\x73\xe5\x09\x9e\x12\xa5\x70\x4c\x00\x64\x05\x4f\x2a\xfc\x52\x45\xcd\x72\x39\xb8\xfe\xec\x65\x66\x46\x6e\x93\x19\x98\xe6\x45\x58\xe3\x12\x0d\x7e\x99\x88\xbf\xc2\xac\x85\x9a\xdc\x2e\x65\x76\x4e\xcc\x95\xf9\x29\x78\x7e\x02\xeb\xcf\x82\xbf\xce\x48\xe0\xdb\x81\xbd\xb5\x15\x8c\x44\xb4\xec\x96\xb5\x3d\xd0\x18\x06\x2e\x34\x19\x04\x97\xe5\x29\x57\x48\x70\x94\x8a\x11\x14\x3a\x1a\x2d\x51\x8a\xa7\x94\xc7\x40\x23\x60\x37\x63\x36\xd6\x45\x03\x18\x11\x01\x1f\x12\x25\x50\x85\xe5\x85\xed\x80\x88\xce\x50\xc8\xb0\x52\x90\x58\xc8\x18\xa6\x9c\x48\x77\xcc\x72\x1a\x55\x19\x84\xea\x31\xd9\xee\x38\x19\x8e\x22\x50\x70\x8c\x5a\x07\xd9\xe2\xc4\xe6\xac\x21\x2e\xc5\xbc\xa0\x3d\xc0\x64\x6e\x1a\xb9\xad\x7d\x63\xb2\xcb\x62\x7b\x2a\x36\x86\x65\x5e\xe3\xce\x30\x27\x0c\x15\x9f\x6e\x26\x69\x8a\xe5\xd2\x41\x96\xed\x31\x9f\x6b\x62\x04\xf6\x94\x38\xc0\x91\xb4\xd7\x19\x8a\x8d\xe4\x3c\xdc\x53\x49\xbb\x42\xf4\x01\xf2\x49\x74\x13\xf8\x95\xeb\x62\x46\xe4\x98\x89\xf9\x31\xce\xb5\x58\x29\x2c\xbf\xde\x04\x39\xab\x64\x19\x55\xda\x8d\xa5\xc8\x33\x98\x05\x51\xc7\x29\x8e\x67\x6b\x0e\xaf\x49\xfa\x39\x5b\x47\x7b\xde\xa8\x31\x0c\x1c\x22\x6b\x8f\x19\x1e\x41\xc0\xd6\xf3\xe3\x6a\x91\x41\x8e\x7e\x3b\xa9\xf5\x05\x94\x67\x79\xd5\x59\x61\x42\xc2\xe9\x48\x2c\xac\x81\x61\x32\x3d\x05\x9f\xa0\xcb\x04\x63\xd0\xb5\xe6\x92\x00\x5d\xcb\x1c\x5e\x08\xba\xc8\x5c\x22\x55\xdc\x56\x4a\xfd\x42\xeb\xca\x86\x51\xae\x35\xd4\xa1\x05\xb7\x0f\x4e\x65\xf5\x48\x73\x04\x7f\x75\x32\x35\x1e\x51\x1e\x91\x45\xc7\xd9\xb3\xfa\xe1\xf6\x8c\x11\x2c\xaf\x44\xbc\x8a\x36\x44\x1a\xeb\x63\x18\x78\x89\xae\x9d\x08\x14\x84\xa0\xc2\x8d\xd9\x32\x4b\x28\x14\x12\x5a\x9d\x5c\x49\x52\x48\x93\xab\x68\xcc\x8b\xe9\x06\xec\x5d\x54\x60\xd7\xb9\xb7\xe6\x3d\xce\x7f\x7d\xb4\xa7\x37\x5b\x0d\xda\x2f\x8b\x1c\x66\x98\xd4\xa8\xf8\x74\x29\x1f\x8b\x95\x23\x11\x55\x19\xc3\xcb\x63\x98\xd9\xc4\xba\x5b\x30\xf5\x61\x10\x10\x1e\x43\x41\xdc\x11\x06\xab\x85\x78\xb5\xae\xa7\x94\xa2\x17\xb5\x96\xd2\x58\x17\xdb\xf0\x25\x9d\xf5\xaf\x1b\xeb\x62\x38\xbc\x41\xe7\x30\xfc\x29\x53\xaf\xeb\x2b\xb3\x74\xea\x7b\x13\x18\x4d\x25\x09\x71\xf6\x44\xe8\x36\x04\x5a\xe5\x61\x48\xcc\x56\xaa\xac\x44\x28\x80\x5d\x2c\x78\xdc\xb5\xa5\x62\x51\xab\x88\xa0\xb2\x22\xe0\xcb\x32\xa1\x26\x97\xb4\x23\xba\x66\x5b\xd5\x7a\xed\x4a\xd5\xb4\x8d\x6e\xe7\x78\x06\x7b\x6f\xe6\x42\x55\x2b\xa7\x61\x07\xa3\x2b\x8b\x61\x7d\xcc\x4c\x23\x05\xb8\x5c\x9f\xff\x03\xe6\x4a\x1d\x32\xbb\x02\xfa\x35\x8e\x8d\xcb\x70\xe1\x74\xcb\x65\x11\xf8\x18\x4c\x65\x74\x0d\xf4\x21\x88\xca\xa0\xae\xc9\x66\x14\x7b\xf7\x02\x18\x66\x1a\xef\x31\x02\xf4\xe3\x43\x61\x3b\xac\xd6\x32\x57\x7a\x09\x12\x6e\xb5\x2d\x1a\xaa\x56\xb9\x6d\x78\xdc\x90\x30\xf5\x80\xc6\x38\x22\xf0\x3a\x80\xca\x40\xd5\xd2\x20\x9f\x49\x52\xc8\x27\xd1\x2a\x3f\x81\x0f\xc4\x86\x8a\x66\x7e\x1e\x6a\xac\xc2\xb3\x49\xe5\xd3\x8a\x4a\xa9\x57\x68\x2a\x22\xf8\x1a\x25\x46\x60\x23\x3e\x7a\xa6\xee\x4a\xd2\xc6\x71\xb5\xd5\x98\x1c\x15\x05\xde\x1d\x7c\xfb\xf2\x10\xf8\xf6\xbf\xc0\x7f\x02\x00\x00\xff\xff\x4b\x57\xc6\x8e\x1d\x0e\x00\x00")

func tmplIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplIndexHtml,
		"tmpl/index.html",
	)
}

func tmplIndexHtml() (*asset, error) {
	bytes, err := tmplIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/index.html", size: 3613, mode: os.FileMode(420), modTime: time.Unix(1472417660, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsScriptJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x56\x5b\x73\xe2\x36\x14\x7e\xe7\x57\x68\xdc\xcc\xda\x0e\xe1\x92\x4e\xfa\x50\x08\x49\xb7\x49\x3b\xd9\xce\x6e\x33\xb3\xec\x4b\x8b\x49\x47\xb1\x04\xa8\x31\x92\x2b\xc9\x4b\xd9\x65\xff\x7b\x8f\x2c\x1b\x64\x71\xd9\x6d\x32\x80\xad\xf3\x9d\x8b\x3e\x7d\x3a\xd2\x47\x2c\x51\x2a\x0a\xae\xd1\x08\xf5\x87\x2d\xf3\x2a\xe9\x3f\x05\x55\x5a\xc1\xc8\x64\x3a\x6c\xb5\xce\x22\x22\xd2\x62\x49\xb9\x8e\xbb\x92\x62\xb2\x8e\x66\x05\x4f\x35\x13\x3c\x8a\xd1\xe7\x16\x82\xbf\xb3\x28\xf8\xee\x59\xf3\xbb\x8c\x62\xf9\x56\xcc\x83\xb8\x9b\x66\x2c\x7d\xd9\x07\xd6\xe0\xb9\x14\x45\x7e\x27\xb8\x12\x19\x05\x34\x5d\xe6\x7a\x1d\xc5\xc3\x06\x46\x33\x49\x53\x9c\x83\x79\xc1\x08\xf5\xad\x0b\x52\xd5\x69\xec\x7a\x99\x45\x41\x70\x00\xa1\x72\xc8\x41\x4f\x40\xb2\xb2\xda\xa6\xf5\x0b\xfc\x9a\x4f\xab\x2e\x1f\x31\x9e\x4a\x6a\x18\xb8\x33\x54\x6d\x67\x23\xa9\x2e\x24\xb7\xfc\xb5\xdb\xe0\xe4\xb8\xa8\x85\x58\xdd\x53\x8d\x59\xa6\x22\x46\x6a\x0f\x45\xf5\x7b\xb1\x1a\xd3\x8c\xa6\x9a\x12\x63\xb0\x29\xa1\x8c\x1a\x5c\xd3\x3f\x61\x64\x1a\x7b\x31\xf7\xbc\x6d\x58\x98\x4a\x37\x63\x4a\x77\x4a\x5a\x3b\x4c\xd3\x65\x60\xd6\x6a\x29\x3e\xd2\xbb\x0c\x2b\x15\x85\xaa\x72\xfa\x4b\x8a\x55\x68\x93\x1a\x02\xe0\xad\x53\x25\xec\x04\x6d\x08\xd8\xc5\x84\x1c\x71\x71\x2b\x99\x53\xa0\x22\x13\xf2\xe7\xf5\x58\x63\x5d\xa8\x48\x95\x3f\x77\x82\x50\x53\x53\x39\x27\x36\x43\xce\x30\x1a\x8d\xd0\xf7\xfd\x3e\xda\x6c\x90\x3f\x78\xe9\xaa\xa3\xe2\x34\x50\x45\x9a\x52\xa5\x82\x6a\x49\x10\xcd\x14\x3d\x10\xf2\xaa\x7f\x75\xc8\x9b\x60\x3e\xa7\xb2\xe9\xbc\x8f\x5a\x61\xc9\x19\x9f\xd7\xb0\xc6\x0c\x9d\x15\xf9\x5b\x09\xee\x2a\x7d\x27\xcc\x19\x26\xf4\xb1\xd0\xd1\x65\xbf\x5f\x2d\xa4\xb3\x81\x60\xff\xfc\x36\x7e\xfc\xbd\xab\xb4\x84\x24\x6c\xb6\x2e\x03\x75\x2b\xeb\x05\x2a\x38\xa1\x33\xc6\x29\xb9\x40\x57\x0d\x6f\x2b\xda\xa3\xee\xd6\x7c\xd4\x1f\x0a\x3f\xee\x5a\x64\xc7\x13\x5b\x62\xc1\xb7\x91\xa7\xbb\xe3\x7b\x78\x80\x01\xac\xb5\x8c\xc2\xd4\x28\x26\xbc\x40\x21\xce\xa8\xd4\xa8\xfc\xee\x84\xa8\x7d\x4c\x27\x71\x7c\x28\x96\x61\xf3\x0d\x77\xc8\xb4\x66\xeb\x52\xef\xd2\x2a\x40\xd3\xbf\xd1\x08\x5c\x96\xbb\x4b\xaa\x17\x82\x40\x25\x01\xfc\xb7\x51\xc3\x96\x63\xbd\x70\x02\xed\x75\x14\xb5\xe6\x1a\xff\xfb\xc0\xe6\x8b\x0c\x3e\xba\xde\x99\xb1\xe7\xd3\xec\x31\xfb\x4e\xd6\xde\xf4\x72\xda\x8e\xef\x00\xa6\xd8\xee\xb6\x86\x18\xdf\xdb\xe4\x0d\x31\xfe\x8f\x55\x33\x50\x46\x00\xe6\xb7\xb2\x9d\x99\x60\x4d\x35\x5b\x1a\xe1\x99\x65\x2b\xa4\x04\xd4\x07\x18\x70\x41\xb3\x22\xcb\xde\x96\x0a\xdb\xc2\x0d\xb7\xd7\x9d\x7d\x7a\xbf\x46\xfd\xb0\xea\xa1\xbb\x76\x57\xcd\x62\x47\x94\x77\x4c\xe0\x3c\xa7\x9c\x44\xe1\x75\xc6\x60\x32\xa3\xc0\xed\x5e\x46\x6d\xcc\x64\x0b\x03\x54\xca\x71\x14\x78\x2d\x11\x79\xef\x27\x05\x6a\x03\x09\x5e\x9e\x63\xa3\xc0\xed\xe6\xbb\x4c\xf1\xb0\xea\x24\x33\x0c\x0d\x26\xb8\x31\x96\x9a\x20\x30\x5f\xf7\x32\x76\x53\x35\x5b\xff\x3c\x28\xc7\x52\x73\x5e\x3e\x66\x04\xf0\xaa\x26\x59\xa5\x52\x94\x11\xd4\xbd\x58\xf1\xc8\x57\x42\xd3\xc5\xd1\x01\x48\xe4\x0d\x4c\x6a\xcc\x3e\x99\x05\xf4\xc9\x83\xb9\x03\x7f\x0a\x8c\x75\x1e\xd3\x4a\x5d\x9f\x1b\xf4\x43\xbf\x79\x4a\x87\x7e\x84\xc1\x8c\x49\xa5\xc3\xfa\x64\xa9\x23\x35\xfb\xa6\x5f\x7f\x15\xd2\xa4\x33\x45\xa5\x8b\x97\xd7\x85\x16\xe3\x12\x05\x25\x31\xe0\x73\x90\x2e\x68\xfa\x42\x49\x18\x7f\xe5\x9a\x60\x63\x7f\x10\x79\x74\xc0\x08\x2b\x19\xf5\x6b\xcc\x03\x35\x7b\xe9\x60\x81\xbe\xb6\x1d\x0e\x53\x3b\x6e\x94\x0d\x1c\x72\xba\x42\xf7\xf0\x78\x64\x8f\x38\x60\x93\xdb\x22\x8d\xda\x7b\xa0\xf6\xed\x2c\x22\x0f\xf6\x4e\x70\xbd\x28\x71\x97\xfb\x60\x0f\xfb\x2b\x48\xe9\x0f\x58\x6e\x1b\x16\xfd\x84\x4e\x60\x1f\x44\x21\x95\x05\x0e\x4e\xc0\xde\x31\x5e\x68\xfa\x0d\xc0\x31\x4d\x05\x27\x5b\x55\x56\x3a\xaf\xa7\xef\x5d\x4b\xbc\x0e\xe6\xb6\x28\xf3\xbc\x6b\x4e\x79\x86\x53\x1a\xf5\x5e\xf5\xe6\x70\x56\xbc\xc2\xcb\x7c\x58\xaa\xa9\x1a\xbe\xb6\xc3\x99\x6e\x8c\xde\xd8\xd1\xb9\x19\x6d\x14\xd3\x8c\x19\x05\x51\x92\x14\x13\xdc\xf9\xf4\xba\xf3\x67\xbf\xf3\xe3\xf4\xf3\xd5\x97\x4d\x92\x4c\x9e\x8a\xe9\x66\xf2\x94\x24\xc1\x34\x3e\x07\x88\x3a\x1f\xc4\xb7\x9b\xe4\x39\xd2\xb2\xa0\x9b\x72\xdf\x6e\x38\xf0\x1c\x27\xcf\x9b\xce\x6d\x42\xda\xd1\xed\x20\xe9\x26\xe4\x3c\xbe\x85\xa7\x09\xfd\x65\x3a\x69\x27\x9d\xa9\xb1\xc4\xb7\xb1\x29\x65\x7b\xa7\x5d\x62\x9d\x2e\x5c\xc5\x96\x12\xca\x4c\x2f\x0e\x79\xb1\x7c\xa6\x32\xdc\xdd\x36\x8d\xfe\x7b\x4f\x41\xaf\xab\x4d\x13\xb7\xae\xae\xef\x16\x33\x38\x3b\x89\xb1\x6d\xa3\xcc\xf1\x42\xd7\x4e\x82\x52\xe8\xfe\x45\xc7\xf3\xb0\xf7\x01\xdf\xa9\xe5\xb9\x97\x65\xec\xe8\x39\x59\x4e\x15\xf8\x59\xc0\x1e\xc4\xdc\x89\xec\xc6\x32\xfc\x7e\x4b\x14\x83\x73\x43\xf8\xf7\xb5\xf0\x5a\xe5\x98\xd7\xdd\xdd\x34\x5b\xe3\x68\x1a\x75\xd9\x79\xcb\xd8\xb6\xef\x1a\xdc\x4d\xe8\xdc\xe6\xff\x0b\x00\x00\xff\xff\x8c\xe9\x3f\x24\xe7\x0c\x00\x00")

func tmplJsScriptJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsScriptJs,
		"tmpl/js/script.js",
	)
}

func tmplJsScriptJs() (*asset, error) {
	bytes, err := tmplJsScriptJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/script.js", size: 3303, mode: os.FileMode(420), modTime: time.Unix(1472417660, 0)}
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
	"tmpl/css/style.css": tmplCssStyleCss,
	"tmpl/index.html": tmplIndexHtml,
	"tmpl/js/script.js": tmplJsScriptJs,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"style.css": &bintree{tmplCssStyleCss, map[string]*bintree{}},
		}},
		"index.html": &bintree{tmplIndexHtml, map[string]*bintree{}},
		"js": &bintree{nil, map[string]*bintree{
			"script.js": &bintree{tmplJsScriptJs, map[string]*bintree{}},
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

