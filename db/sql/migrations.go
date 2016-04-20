// Code generated by go-bindata.
// sources:
// migrations/0001-initial.sql
// DO NOT EDIT!

package sql

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

var _migrations0001InitialSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x93\x4f\x6f\xb2\x40\x10\xc6\xcf\x2f\x9f\x62\x8e\x98\xb7\x7e\x02\x4e\xa0\x23\x21\xc5\xc5\xac\xcb\xc1\x13\x59\x65\x63\x68\x04\x1a\xa0\x51\xfc\xf4\xdd\xf2\xa7\x65\x37\x22\xa1\xe5\x38\xbf\x67\x1e\x9e\x19\x86\xe5\x12\xfe\xa7\xc9\xb9\xe0\x95\x80\xf0\xdd\x58\x51\xb4\x19\x02\xb3\x1d\x1f\xe1\x9a\x17\x97\xb8\x34\x0d\x90\x4f\x12\x83\x47\x18\xba\x48\x61\x47\xbd\xad\x4d\x0f\xf0\x8a\x07\xb0\x43\x16\x78\x44\x76\x6d\x91\xb0\x97\x46\x99\xf1\x54\x80\xb4\xf1\x81\x04\x0c\x48\xe8\xfb\x0b\x4b\xf5\xbd\xf0\x5a\x14\xb3\x7d\x9b\x34\xd1\x8f\x7e\xec\x6d\x6d\x7d\x13\x50\xf4\x5c\xf2\x65\x66\xf6\x9d\x0b\x29\xdc\x20\x45\xb2\xc2\x7d\x3f\x9c\xac\x1a\x7a\xbe\xf2\x24\x32\x31\x3b\x5f\x33\xd5\x6f\xf2\x81\xd9\xb7\x2a\x01\xbb\x2d\x3d\x0a\x78\x16\x79\x2a\xaa\xa2\x8e\xf2\xe3\x9b\x38\x55\x7f\x5e\xa5\x96\x4d\x9f\x44\xc3\xcd\x7a\xc6\xf1\x31\xff\xc8\xe2\x32\xba\x45\x69\x92\x3d\x9a\xbc\xe3\xf5\x04\xbf\x4f\x70\xe9\xcf\x6f\xcf\xfd\x9f\xf2\xfb\x18\xff\x5e\x6e\xcc\x2b\x0e\x8e\x1f\x38\x9a\x40\x52\xde\xb0\x3d\xa3\x1e\x71\x07\xf4\xdf\x8c\xab\xd3\xd4\x13\x27\xa0\xa9\xfb\x4f\xa0\xa8\xbb\xb3\x95\x55\x79\x2f\xc6\xf0\xc7\x5e\xe7\xd7\xcc\x58\xd3\x60\x37\x72\x3f\xd6\x10\xb6\x36\x4a\xa9\xcd\xa1\x94\xda\x41\xac\xcf\x00\x00\x00\xff\xff\xc1\xbc\xb4\x9b\x3f\x04\x00\x00")

func migrations0001InitialSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations0001InitialSql,
		"migrations/0001-initial.sql",
	)
}

func migrations0001InitialSql() (*asset, error) {
	bytes, err := migrations0001InitialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/0001-initial.sql", size: 1087, mode: os.FileMode(420), modTime: time.Unix(1458425772, 0)}
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
	"migrations/0001-initial.sql": migrations0001InitialSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"0001-initial.sql": &bintree{migrations0001InitialSql, map[string]*bintree{}},
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

