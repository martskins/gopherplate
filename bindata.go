// Code generated by go-bindata.
// sources:
// out.tmpl
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

var _outTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x54\xc1\x6e\xdc\x36\x10\xbd\xeb\x2b\x5e\x8c\xa0\x90\x16\xae\x74\x0f\xb0\x87\xa0\xdd\x26\x0e\xda\x8d\x91\x5d\xa7\x87\xa2\x80\xb9\xab\xd1\x8a\x31\x97\xdc\x92\xa3\x3a\x2e\xa1\x7f\x2f\x28\x4a\xb6\x22\xcb\x8d\xdb\xa2\xa8\x4e\x84\x66\xf8\xe6\xbd\x37\xc3\x29\x16\xc9\x02\xdb\x5a\x3a\x54\x52\x11\x6e\x85\x83\x68\xd8\x7c\x7b\x20\x4d\x56\x30\x95\xd8\xdd\xc1\xfd\xa6\x0e\xa4\x3f\x23\xad\x99\x4f\xee\x55\x51\x1c\x24\xd7\xcd\x2e\xdf\x9b\x63\x71\x14\x96\xdd\x8d\xd4\xae\x38\x98\x53\x4d\xf6\xa4\x04\x53\x96\x27\x0b\x5c\x30\xa4\x83\x28\x7f\x97\x4e\xec\x14\x81\x0d\x6e\x88\x4e\xe0\xfb\x6a\x8d\xde\xd7\x42\x1f\xa8\xcc\x93\x45\x91\x24\x27\xb1\xbf\x11\x07\x82\xf7\xc8\x2f\xe3\x79\x2d\x8e\x84\xb6\x4d\x12\xef\x61\x43\x2a\x5e\x3a\xb6\xcd\x9e\xf1\x6a\x89\x7c\xd3\x1d\x5d\x97\x80\x70\x4d\x56\x78\x99\x6f\x6a\xd3\xa8\x72\xf5\xf9\x64\x2c\x87\x10\x00\x14\x05\x2e\xb4\x23\xcb\xde\xf7\x00\x79\x80\x6e\xdb\x40\x91\x6b\x0a\x12\xe1\x58\x30\x1d\x49\x73\xa0\x2a\xbb\x74\x08\x4c\x6f\x74\x78\xde\x77\xa5\x22\x66\xdb\xee\x8d\x76\xfc\x54\x85\x25\xae\x2f\xd6\x9b\xd5\x87\x2d\x2e\xd6\xdb\xf7\xf0\x3e\xdf\x06\x3f\xfa\x68\xea\x7d\x0f\x93\xff\x20\x49\x95\xae\x6d\x33\x7c\x7c\xfd\xe3\xd5\x6a\x33\x8e\x5d\x2a\xb1\xa7\xb7\x46\x95\x64\x43\xc6\xb5\xf7\xa4\xcb\x07\x71\x57\xa7\x52\x30\x3d\x5b\x5c\xd3\xa5\x7f\x45\x5c\xc4\x1c\xc4\x3d\x51\x61\x89\xeb\xab\xcb\xef\x5f\x6f\x57\x53\x5d\x9b\xd5\x36\xfc\x8a\xd7\xf2\x0d\x31\x77\xcc\xf1\xf3\xdb\xd5\x87\x15\xa6\xfc\xdf\x19\xa9\x9f\xcd\xfe\x93\x91\x1a\xce\x1c\x09\xdc\x0d\xd6\xad\xe4\xba\xcb\x7c\x84\xd0\xc5\xc7\x9a\x42\x9d\x41\xd1\x6c\xcd\x65\xa0\xd6\xa5\xdd\xb7\x63\xc4\xd5\x7b\x90\x72\x84\xbf\x9a\x02\xf9\x3f\x4d\xc1\x6c\xd7\x9a\xff\xae\x6b\x33\x8e\x7e\xfa\x67\x8e\xea\xb2\x7f\xc1\x45\x81\x75\xa3\xd4\xdc\x20\x08\xe8\x46\xa9\xae\xdd\xd3\xf0\x39\x8e\xc6\xb1\xba\x43\xe3\xa8\x6a\x14\x2a\x63\xa1\xa8\x8a\x74\x5c\x02\xf0\xdd\x89\xe6\x81\xfb\x5d\xe2\x3b\x22\xfd\x7e\xa9\x02\xc9\xb0\x5e\x86\xdc\x81\x75\xaf\x3a\x26\x0c\x08\x0b\xef\x37\x42\x4b\x96\x7f\xd0\x36\x94\xe9\xa3\xe1\xdc\xb6\xc3\x4a\xea\xff\x89\x43\x58\x55\x0f\x10\xe1\x47\x4c\x8a\x16\x3c\xb8\x32\xb8\xf1\x51\x28\x59\x42\xea\x52\xee\x05\x93\xc3\x6d\x4d\x5c\x93\x0d\xc3\x6e\x29\xda\xf2\xd8\x2c\x1d\x37\xec\x9c\xe0\x04\xa8\x1a\xbd\x47\xaa\xb1\x98\x8b\x67\xb1\x62\x9a\x61\x67\x8c\xfa\x3b\xbe\xc8\x0a\x3a\x9f\x9a\xf3\x62\x09\x2d\x23\x4c\xf8\x2c\x71\x63\x35\xd8\x36\xf1\x4d\xde\x0b\x4e\x46\xd1\x4a\x28\x47\x23\x0b\xde\x10\xf7\x21\x37\xff\xc6\xff\x9d\xde\x37\xc4\x69\x16\xba\x38\x41\xf5\x83\xaa\x17\x3a\x1f\x3c\x99\xe8\xd0\x52\x45\x19\x5f\xf0\xff\x66\x0a\xe5\x67\x2c\x9c\xce\x54\x67\x9f\x61\xa4\xdf\x19\xcd\x42\x6a\x37\x9e\x23\x9c\x2d\xce\xb2\x6e\x3c\xbe\x70\xf7\x15\x16\x8f\x1c\x3f\xf7\x3e\xec\xa7\x99\xdc\xd9\xd4\x60\xfe\xf8\x49\x8f\x47\xef\x27\x61\x5d\x2d\xd4\xbb\xcd\xfb\x35\x8e\xf1\x1c\xc6\x6d\xf6\x19\x85\x9d\xec\x8c\xfe\xba\xdb\x23\xd0\x34\x43\xfa\xcb\xaf\xbb\x3b\xa6\x73\x90\xb5\xc6\x66\xcf\xf1\x3c\xde\x48\xcf\xc2\x32\x38\xcb\xce\x9f\xe8\x41\x20\x93\xf7\xb5\x52\x9d\x77\x3d\xce\xa2\xb6\x87\x7d\xf3\x67\x00\x00\x00\xff\xff\x87\x4c\xeb\xdd\xfa\x08\x00\x00")

func outTmplBytes() ([]byte, error) {
	return bindataRead(
		_outTmpl,
		"out.tmpl",
	)
}

func outTmpl() (*asset, error) {
	bytes, err := outTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "out.tmpl", size: 2298, mode: os.FileMode(420), modTime: time.Unix(1555504091, 0)}
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
	"out.tmpl": outTmpl,
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
	"out.tmpl": &bintree{outTmpl, map[string]*bintree{}},
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

