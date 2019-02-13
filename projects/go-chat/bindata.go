// Code generated by go-bindata.
// sources:
// public/app.js
// public/index.html
// public/style.css
// DO NOT EDIT!

package main

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

var _publicAppJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\x51\x6f\xdb\x36\x10\x7e\x0f\x90\xff\x70\xcd\x06\x50\x86\x3d\x29\x18\xb0\x17\x3b\x0a\xd0\x75\x01\xda\xa2\x59\x87\xa6\xdb\xb0\x47\x5a\x3a\x4b\xec\x28\x52\xe0\x9d\xac\x65\x43\xfe\xfb\x40\x4a\x96\x64\x5b\xc9\x30\x54\x4f\xd4\xe9\xee\xe3\xc7\xef\x3b\x1d\x0d\xb6\xf0\x5b\x83\xd1\x3f\x97\x17\x00\x00\xa8\xd7\x20\xbe\x91\x75\x2d\x56\x97\x17\x5d\x28\x97\x2c\xd7\xd0\x7f\xf7\x4f\x4b\x6b\x30\x8d\xd6\x2b\x48\x12\xf8\xd8\x38\x68\x71\x4b\x36\xfb\x13\x79\xcc\x31\xd8\xde\x53\xb1\x06\x21\x42\xd6\x5b\xab\x73\xf2\x41\xa8\x90\x48\x16\x48\xc0\x16\xb6\x08\x84\x86\xfd\x92\x4b\xbf\x76\x7b\x74\x23\x46\x56\x4a\x7e\x63\x0d\xa3\xe1\x01\xe8\x35\xb8\xc6\x18\x65\x0a\xd0\x8a\x18\xec\x2e\x64\x8d\xa8\xb9\xa2\x5a\xcb\x47\xcc\xc1\x9a\x0e\x34\x73\x88\x66\x04\xc5\x4a\x2a\x3d\xe1\x7f\xe7\xdf\x41\xe6\xb9\x43\x22\x68\x08\x73\xd8\x59\x07\x85\x93\xdb\xad\xdf\x46\x1a\x90\x7b\xc9\x72\xc2\xab\x21\x74\x46\x56\x78\xa2\xc2\x21\x3c\x26\x7e\xb1\xca\x60\xbe\x86\x9d\xd4\x84\x3e\xed\xb3\x6b\x10\xd4\xae\x23\x01\xd2\xe4\x43\x11\x94\x72\x8f\xb0\x45\x34\xb0\x53\x5a\x63\x0e\xaa\x27\xfd\x34\x18\x91\x39\x94\x1c\xf0\x1a\x93\xb1\xb2\x26\x5a\x4c\x6d\xd9\x4b\x07\x84\x7a\x07\x29\x70\xa9\x68\x33\x7e\xf1\xaf\x71\x4b\x90\x06\x07\x7e\xc7\xed\x43\x70\x2b\x12\x2d\xad\x93\x44\xc0\x12\x5a\x65\x72\xdb\xc6\xda\x66\xd2\x03\xc7\xa5\x25\x86\x25\x88\xa4\x25\xb1\x38\x47\x8a\x65\x9e\xdf\xed\xd1\xf0\x07\x45\x8c\x06\x5d\x24\x7a\x07\xc4\x6a\x64\x87\x47\xf4\x0e\x14\x2b\x2a\x20\x85\xf7\x0f\x1f\x7f\x8e\x6b\xe9\x08\x23\x8c\x7d\x87\x4d\x77\xf1\x8f\x3f\x49\x3c\xe9\x00\x58\xa6\x20\x6e\x72\xb5\x87\x4c\x4b\xa2\xf4\x2a\x2b\x55\x7d\x75\x2b\x8e\xab\x0e\xcf\x12\xc4\x8d\xaa\x0a\x20\x97\xa5\x57\xfe\x80\x01\xae\x70\x9d\x95\xbf\x7e\xfa\x10\x55\x54\xc4\xc1\x86\x85\x4f\xbe\xba\x15\xa1\xbb\x4e\x9c\x3e\x86\xf4\x25\xe7\x26\x4f\xb7\x4c\x72\xb5\x9f\xa3\xb4\x04\xac\xec\x17\x65\x0d\xc6\x6c\xdf\x55\xb2\xc0\xb0\x7f\xaf\x59\x60\x70\xb3\x75\xc9\xad\xd8\x78\x16\xbf\x78\x59\xba\x0a\x3a\x78\x3f\x55\x10\x35\x56\x5e\x91\x14\x72\x9b\x35\x7e\x19\x17\xc8\x77\x5d\xf4\xc7\xc7\x77\x79\x24\xbc\x70\xdf\x1d\x7e\x0a\x71\xaa\x6d\x0f\x10\x53\xe6\xac\xd6\x9f\x6d\x0d\xe9\x49\xec\x2d\xaa\xa2\xe4\xc0\xe6\x75\xc3\x16\xba\xe8\xe1\x47\xdd\x5a\x66\x5b\x8d\x98\x4f\x87\x0d\xc6\x5e\xad\x90\x4b\x9b\xd3\xd1\xdc\x20\x34\x93\xe6\x85\xe8\xac\x3d\xd4\x0e\xa2\xd0\x61\xdd\xf4\x80\x57\x29\x08\x71\x96\x05\x93\x3e\xf4\x90\xd1\xbc\x5f\xa1\xc3\x88\x9d\x32\x85\xda\x3d\x46\x33\x20\x83\x1c\xdd\x48\x08\x98\x61\xbd\x7a\x3e\x77\xfc\xf7\x43\xfa\xe1\xf5\x85\x8a\xde\x85\x35\x7c\x1b\x89\x9b\xfa\x56\x2c\xe2\x92\x2b\x3d\x3d\xe7\x22\x66\xfc\x8b\xa3\x85\x57\xfb\x81\x9d\xaa\xc1\x36\x0c\x3e\x6b\x1e\xf5\xe9\x3c\xbc\x38\xb5\x78\x10\xa9\x97\xd2\x2b\x19\xec\xfc\x84\x84\xdc\x8f\xe7\xe3\x92\x09\xec\xe8\x23\xf4\x63\xec\x3f\x7d\x7b\x35\xca\x37\xeb\xd8\xbd\x64\x74\x4a\x6a\xf5\xb7\xff\x07\x24\x71\x24\xfe\xb0\x0d\x54\x0d\x31\xa0\x61\x74\x7e\xd2\x86\x72\xb1\x82\xef\xaf\xaf\xaf\xe7\x4e\xe4\x90\x1b\x67\x9e\xa5\x7d\xcc\xe5\xe0\xcd\xff\xa6\x93\x95\xd6\x12\x82\x1c\xcc\xfe\x3a\x46\xa3\x30\x90\xce\x35\x41\xa7\x59\xdf\x03\x9b\x99\xda\xe1\x8e\x98\x2d\x1f\x8e\xf9\x02\x42\x77\x11\xf9\x8b\xc1\x35\xb8\x79\xce\xe7\xc9\x78\x9c\xdc\x31\xf3\x96\x76\xe7\x06\x51\x32\xd7\xeb\x24\x69\xdb\x76\x98\xae\x71\x66\xab\xa4\x5b\x86\xbb\xe5\x8d\x7b\xac\xd9\xbe\x7f\x88\xef\x7f\xfa\xa1\x47\x9b\x52\xe8\x27\xc7\xe5\xc5\xd3\x62\xf3\x6f\x00\x00\x00\xff\xff\x95\xb0\xe5\x08\x8d\x08\x00\x00")

func publicAppJsBytes() ([]byte, error) {
	return bindataRead(
		_publicAppJs,
		"public/app.js",
	)
}

func publicAppJs() (*asset, error) {
	bytes, err := publicAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/app.js", size: 2189, mode: os.FileMode(438), modTime: time.Unix(1548846541, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdf\x6f\xdb\x36\x10\x7e\x0f\x90\xff\x81\xe5\xd3\x86\x41\xe4\xe2\xee\x47\x36\x48\x59\x80\xac\x03\x36\xa0\xd8\x80\xb6\x0f\x7b\xa4\xc9\x93\x44\x85\x22\x39\xde\x49\xa9\xfb\xd7\x0f\xa2\x2c\xd7\xb1\x93\x2c\x31\x0a\xec\xc9\xc7\xd3\xf1\xbe\xef\xf8\x1d\x7d\x2c\x5f\xfd\xfa\xe7\xcd\xfb\xbf\xff\x7a\xc3\x5a\xea\xdd\xd5\xf9\x59\x39\xfd\x32\xa7\x7c\x53\x71\xf0\x3c\x7b\x40\x99\xab\xf3\x33\xc6\x18\x2b\x7b\x20\xc5\x74\xab\x12\x02\x55\xfc\xc3\xfb\xdf\x8a\x4b\xbe\x7c\x23\x4b\x0e\xae\xde\xd9\x3e\x3a\x60\x37\xad\xa2\x52\xce\xae\xf3\xb3\x6d\x84\xb3\xfe\x96\x25\x70\x15\x47\xda\x38\xc0\x16\x80\x38\x6b\x13\xd4\x15\x6f\x89\x22\xfe\x2c\xa5\x36\xbe\x43\xa1\x5d\x18\x4c\xed\x54\x02\xa1\x43\x2f\x55\xa7\x3e\x4a\x67\xd7\x28\x7b\x45\x90\xac\x72\xf6\x13\xc8\x6f\xc5\x4f\x3f\x8a\x4b\xa9\xf1\x9e\x5b\xf4\xd6\x0b\x8d\xb8\xe3\xf5\x1c\xd4\x3a\x78\x42\xd1\x84\xd0\x38\x50\xd1\x62\x46\xb5\x3a\xf8\x5f\x6a\xd5\x5b\xb7\xa9\xde\x6e\x01\xbe\xf9\x5d\x07\xff\xb2\xdc\xda\x78\xd1\xa1\x01\x67\xc7\x24\x3c\x90\x84\x3e\x74\x36\x78\x90\x2b\xb1\x12\x3f\x48\x85\x08\x84\xb9\x8c\xe5\xcb\xae\x06\xf9\x1c\x20\x99\x5d\x4b\xcd\xe7\x67\xa5\xdc\x4a\x56\xae\x83\xd9\x2c\x12\x42\x5a\x72\x79\x35\x6e\xcd\xbc\x34\x76\x64\xda\x29\xc4\x8a\x7b\x35\x16\x77\x49\xc5\x08\x89\xef\x85\xe4\x30\xb5\xa0\xf1\x25\x7a\x9d\x94\x37\x85\x0b\x4d\x60\xc9\x36\x2d\xf1\xfb\xe2\xab\x7d\x10\x69\xec\x02\x5a\xca\x99\xc0\x4c\x33\xd3\x2a\x7b\x65\x3d\xb3\xa6\xe2\x2a\xc6\xdd\xe1\xee\x11\x4b\xe1\x8e\x3f\xc2\x59\x07\xc7\xf0\x62\x75\xc4\x77\x3f\x44\x25\xc3\xda\x90\xec\xa7\xe0\x49\xb9\xc3\xd0\x5d\xf8\x44\x40\xb7\x8a\x8a\x1e\x10\x55\x03\xc8\xf7\x33\x14\x3a\x78\x02\x4f\x9c\x8d\xc5\x74\x4b\xe6\xd8\x9b\xad\xf3\xa1\x9c\x7b\x45\x3f\xe6\xbb\x7f\x30\x7b\xf6\x41\xf1\x6c\x2c\x6c\x5d\xf1\x2e\x58\x0f\xe6\xb1\xa3\xb0\x3e\x0e\x54\xd4\x16\x9c\x61\xf9\x58\x2e\x8f\x4e\x25\x87\x30\xda\x44\xa8\x38\xc1\xc7\x5c\x4c\x1f\xcc\xd4\x58\x1e\xee\xde\x62\xc3\xd9\xf5\x2d\x6c\x86\x28\xc0\x13\xa4\x8a\x23\xf8\xfb\x78\x87\x05\x3c\x85\xff\xdd\x11\xfe\x7a\x20\x0a\x7e\xd9\x70\xa7\x46\xc0\x02\xea\x1a\x34\xb1\x79\xe1\xa6\x4e\x62\x6b\xf2\x9c\x5d\x6b\x67\xf5\xed\x31\x85\xcf\xc5\x2c\x79\x96\xeb\x5f\x4c\x57\x16\x97\x6e\xd4\xb9\x0d\xed\x03\x3b\xdf\x81\x37\x87\xba\xcc\xcc\x4e\x97\xe6\xd5\x17\xd5\x06\x7a\x65\xdd\x4e\x1c\x41\xc9\xf6\x3b\x67\x74\x4a\x43\x1b\x9c\x99\xf4\x79\x93\x7d\xa7\x0a\xf4\x82\x06\xd9\x72\x18\x10\x92\x57\x3d\x1c\xd0\xf8\xb0\xb8\xff\xdf\x56\x99\x34\xf8\xea\xeb\x53\x9a\xc5\x04\x0f\x0f\x37\xcb\x1f\xc1\xfa\x93\x9a\xa5\x94\xd3\xdf\xda\x64\xd4\x21\x10\xa4\x85\x40\x54\x0d\x14\xb3\x2b\x0f\x58\x39\xdb\x93\x89\x3a\xd9\x48\x0c\x93\xfe\x3c\x3f\x06\x1f\x6f\x9b\x3c\x90\xc6\x01\xae\x57\xe2\x42\xbc\x96\xc6\x22\x4d\xcb\x3c\x28\x3a\xe4\x57\xa5\x9c\xb7\x3e\x96\xe4\xbf\x86\x90\xb3\x6b\xd9\x1d\x0c\xa0\x67\xe5\x0d\x06\x44\xf7\xcf\x00\x69\x93\x29\xce\x66\x31\xb1\xbc\x78\x09\xb9\xa7\x66\xbe\x4e\x9b\x48\xa1\xe8\x50\xbe\x16\x17\x62\x25\x53\x70\x6e\x88\x28\x7b\xf3\xfd\x97\x48\xff\xc0\x93\xa2\x3b\x7e\x51\x3c\x05\x24\x55\x8c\x47\x01\x72\x19\xbe\x32\x3f\xad\xfe\x0d\x00\x00\xff\xff\x95\x90\xac\x90\x6a\x09\x00\x00")

func publicIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_publicIndexHtml,
		"public/index.html",
	)
}

func publicIndexHtml() (*asset, error) {
	bytes, err := publicIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/index.html", size: 2410, mode: os.FileMode(438), modTime: time.Unix(1548846541, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicStyleCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\xc1\x0a\xc3\x20\x0c\xc6\xf1\xbb\xe0\x3b\x04\xc6\x8e\x42\x7b\xd9\xc1\x3e\x8d\x53\x5b\x03\xd1\x8c\x6a\xdb\x95\xb1\x77\x1f\x6e\x75\x63\xd7\x1f\x7f\xf2\xe5\xca\x6e\x87\x87\x14\x00\x00\x0e\xf3\x8d\xcc\xae\x61\x24\x7f\x1f\x3e\x16\x31\xa9\xe0\x71\x0a\x45\x43\xdf\x75\x6b\x38\xbc\x26\xca\xe1\xec\x6d\x41\x4e\x1a\x2c\xd3\x12\xd3\x20\xc5\x53\x0a\x29\xa2\xc1\xd4\xae\xd6\x52\x43\x0f\x1d\x98\xa5\x70\x2b\x4e\x36\x98\xa2\xa2\xcf\xd9\x4c\x3e\xb7\xf6\x7f\xed\x3b\xd6\xe8\xf2\xa3\x0d\x5d\x09\xef\x97\xce\x87\xf0\xea\xe7\x91\x78\x53\xbb\x86\x6c\x67\x26\xaa\x5b\xaf\x00\x00\x00\xff\xff\xea\x87\x4f\x49\xe0\x00\x00\x00")

func publicStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_publicStyleCss,
		"public/style.css",
	)
}

func publicStyleCss() (*asset, error) {
	bytes, err := publicStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/style.css", size: 224, mode: os.FileMode(438), modTime: time.Unix(1548846541, 0)}
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
	"public/app.js": publicAppJs,
	"public/index.html": publicIndexHtml,
	"public/style.css": publicStyleCss,
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
	"public": &bintree{nil, map[string]*bintree{
		"app.js": &bintree{publicAppJs, map[string]*bintree{}},
		"index.html": &bintree{publicIndexHtml, map[string]*bintree{}},
		"style.css": &bintree{publicStyleCss, map[string]*bintree{}},
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
