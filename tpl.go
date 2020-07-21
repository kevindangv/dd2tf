// Code generated by go-bindata.
// sources:
// tmpl/dashboard.tmpl
// tmpl/monitor.tmpl
// tmpl/screenboard.tmpl
// tmpl/timeboard.tmpl
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

var _tmplDashboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1a\xcb\x6e\xe3\x36\xf0\xee\xaf\x20\x8c\x3d\x6e\xfc\x01\x05\x72\x68\x93\xa6\x0d\xb0\x69\xf7\x91\xee\xa2\x2d\x0a\x81\x31\x69\x87\x28\xf5\x88\x48\x6d\xec\x0a\xfa\xf7\x82\x43\x52\x22\x25\x51\x92\x23\xb9\x58\x14\xc9\xc5\xe2\x3c\xc9\xe1\xcc\x70\x66\x90\x9c\x8a\xb4\xc8\xb7\x14\xad\x09\x96\x98\xa4\xfb\x88\x60\xf1\xf8\x90\xe2\x9c\xac\xd1\x9a\x90\xa8\x2c\xd1\xe6\x96\xa0\xaa\x5a\x23\x54\xae\x10\x92\x4c\x72\x8a\xf4\xdf\x25\x5a\x2b\xf4\x3d\x80\xaa\x6a\xbd\x42\x88\x50\xb1\xcd\x59\x26\x59\x9a\x58\xf4\xb5\x03\xd2\x44\x1c\x1f\xd3\x42\x7a\x32\xde\x01\xe8\xfe\x98\x59\x41\x39\xc5\x24\x4a\x13\x7e\x6c\x68\x6e\xc5\x47\x8a\xc9\xaf\x0a\x76\xa1\x89\xca\xf2\x99\xc9\x47\xb4\xf9\xc2\xc8\x9e\x4a\xa1\xc0\x00\xcd\x71\xb2\xa7\x68\x83\x60\xf9\x0c\x48\xd8\x3c\x42\x8c\xd4\xd2\x88\x95\xa2\x38\x2e\x90\x96\x74\x4d\x77\x2c\x61\xb0\x57\x2d\xcc\x45\xfe\x94\xa7\x45\xa6\x75\x35\x74\x86\x6a\xaf\x70\x11\x69\xd8\xb5\x3e\x6b\xb0\x1e\x53\xf9\xb2\xfd\x23\x68\x9c\x77\x0c\xf5\xe7\x1d\x25\x7c\x9c\xd1\x23\xf9\x04\xf7\x69\xc6\x99\x90\xed\x83\x21\x8f\x18\x6d\xcc\xe5\xf4\x9d\x71\xf4\x9c\xbe\xc2\x8f\xf4\xa9\xa0\x42\x1f\xb6\x3e\xa5\xa3\x4d\xdd\x3e\x50\x78\x0a\x10\x7a\xaa\x5d\xe1\x43\x4b\xb6\x96\xce\x76\x68\xf3\x49\x1e\x39\xd5\x82\xb5\x32\x0d\xf0\xa4\x23\x24\x00\xe8\x4b\x57\x32\x94\x84\xf7\x98\x53\x29\x15\x4b\x66\xbe\x8c\xd2\x06\xb1\x56\xda\x68\x42\x5a\x52\xad\x84\x77\x2c\xa1\xc6\x5a\x9c\x25\x34\x92\xea\xdb\x3a\x7a\x83\x9b\x20\xe5\x0b\x23\xf2\xd1\x8a\x79\x86\x85\x23\xc7\x62\x03\x82\xaa\xb2\xa4\x09\xd1\xb6\xe8\x62\xad\xc1\xae\xd2\x84\xc0\x6d\x62\x7e\x93\xe6\x31\x96\x02\x29\x06\x7b\x57\x41\x74\xc7\x39\xf5\xdf\xb6\xa1\x8f\x76\x86\xa1\x6b\x67\x74\xa2\xa1\xdb\x5a\x94\x9e\x38\xc3\x39\x96\x69\x6e\xb9\xae\x1a\x48\xdb\x39\x6a\x95\x9f\x31\x2f\x94\xdc\xaf\xf0\x6b\x18\x2d\x30\xac\xcc\x70\x5f\x15\x42\xa6\xf1\x0f\xfb\xab\x94\x83\x92\x2d\xac\xa3\x87\x7d\xb4\x05\x88\xdd\x48\x8b\x6c\xa2\xdc\x9b\xb6\xdc\x5d\xaf\xdc\x9b\x51\xb9\xf6\xba\xfd\x9f\x55\x90\x60\xe5\xc0\xdb\x7e\xe4\x64\x09\x16\x53\x41\x73\x46\xc5\x7f\x9f\x28\x4e\x4d\x13\x03\x59\x42\xd9\xfb\x9a\x89\x8c\xe3\xa3\xd9\x27\xd1\x2b\x2f\x4a\x7d\x8a\x40\x7c\xf9\x29\xe7\x35\xe3\x84\x1d\xcb\xb5\xd6\x1d\x95\x58\xd5\x1a\x8e\xc1\x3c\x50\xdf\x3d\x23\x14\x5b\x92\x7e\x0b\xfe\x78\xc8\x72\x2a\x84\x76\x46\xda\x2c\xcc\xd6\x3d\xf4\x88\x11\xbe\xe7\x0c\x8b\x5f\x70\xac\xec\x85\xd5\x77\x94\xa8\x85\x91\xe4\x62\xa7\x19\x61\x7a\xf8\x39\x36\xc2\xf9\xdf\x34\x17\xae\x89\x1c\x48\x9f\x85\x62\x83\x2f\xcf\xe9\xee\xe0\x23\xf8\x81\x72\xe5\x1f\xf0\x5b\x57\x6f\x1a\x38\xc8\x39\x29\xf7\x0e\x19\x2a\x6c\xaf\xdf\xf1\x81\x09\x3f\xe2\x8e\x00\xea\x33\x07\x10\x6f\x6e\x93\x2d\x2f\x08\xfd\x83\xe6\x29\xaa\x2a\xa6\x57\xd1\x3f\x6a\x69\xb6\xd6\x47\x37\x78\x40\xcd\xd0\x6b\x20\x1f\x35\x41\xca\x1d\x3e\xa0\xaa\x8a\xf1\xc1\x97\xa0\xc1\x53\xf8\x99\x72\xf5\x98\x25\x2d\x7e\x36\x10\x01\x2e\xff\xa7\x2d\x86\x1c\x26\xe0\xd7\x93\x61\x51\x83\x57\xb6\x0a\x02\xac\x47\x7f\x28\x68\x7e\x84\xeb\x1f\x7a\x4f\x9e\x14\x55\x04\x2e\x33\xeb\x3d\x81\xc8\x2e\x64\xaa\xcf\x73\x51\x55\xb8\x5e\xd8\xc8\xae\x01\x81\xa3\xd5\x05\x13\x3c\xc3\xbf\x25\x4c\x2a\x39\xe6\xad\x2e\xd4\xd2\x7b\xa7\x81\xa0\xea\x79\x52\x1b\x49\xef\x73\xba\x65\xc2\x14\xe6\x59\xbd\xb0\xc9\xbf\x06\x0c\x6c\xe8\x9c\x05\x35\xd8\x6c\xbf\xcf\xe9\xde\x14\x55\xb8\x59\x58\xa3\xb9\xe8\x91\x67\xf2\xb5\xd0\x1c\x2b\x34\xe7\x14\x6e\xa1\xb2\xed\xe7\x54\xc8\x3b\xdc\xe9\x5a\x5d\xd5\x8f\xa9\x90\x31\xee\x6d\x5e\xd1\xb4\xf8\x6a\xbb\xe1\x04\xd7\x6b\x76\x78\xc3\x38\x6f\x5b\x62\xa7\x60\xed\xdb\x1c\xa8\xeb\x06\x6c\xb5\xea\xdc\x0d\xb4\xf1\x2a\x54\xa0\x67\x47\x97\xe8\xcf\xb2\x7c\xa3\xbf\xbf\xbb\x34\xe8\x26\x8e\xde\xb0\x84\xd0\xc3\x5b\xf4\x86\x72\x1a\xfb\x04\x6c\x67\xb0\x55\xf5\xd6\x68\x5b\x97\x25\x10\xc2\x17\x40\xfe\xea\x0f\x5d\x5d\x39\x6e\xd3\x0c\xb2\x91\x80\x0f\xbd\x13\xfd\xad\x14\x01\x7a\x68\x27\x35\xc1\xfc\x9d\xd8\xb6\xd9\x41\xe9\x92\xf5\xb2\x75\x0f\x36\x84\xe0\xcf\x5c\x08\x70\xbb\xd1\xd4\xc7\x11\xed\x38\xcb\x7a\x39\x6e\x14\xc2\xe7\xea\xc9\x9b\xad\x0b\xed\x04\x8b\x4b\xdf\xb4\xbe\xad\x11\xce\xc8\xac\x63\xa4\x7d\x99\x38\xcb\x19\x4e\xc7\xdd\x78\x08\xa5\xe1\xe9\x33\x8d\x6e\x77\x31\xaf\xb3\x98\xdf\x55\x2c\xd0\x51\x84\xe7\x17\x67\x78\x52\xc6\x9e\x93\xb9\x4f\xc9\x09\xcf\xc8\xcb\x9f\x90\x73\xcc\x29\x96\x9f\x51\x0c\x3e\x73\x7d\xd0\x56\x70\x9f\x32\x93\x58\x36\xa0\xa7\x86\x73\x20\x9a\x17\x6a\xc6\x26\xcf\x1d\xfe\xc7\x59\x21\x9c\x13\x4e\x9f\x2f\xf4\xce\x16\x16\x9a\x2b\x2c\x30\x53\x38\x39\x5c\x5e\x3a\x47\xe8\xce\x10\x16\x74\xd9\x97\xcc\x0e\x5e\x32\x37\x18\xb4\x56\x60\x5e\xd0\x9e\x15\x9c\x65\x4e\x30\x7f\x46\x30\x67\x3e\x30\x67\x36\x30\x77\x2e\xd0\xce\xe5\x6e\x26\x9f\x3a\x0b\x18\x99\x03\x8c\xe6\xf1\x99\xfd\xff\x32\xbd\xff\x02\x7d\xff\x39\x8a\xcc\x05\xfa\xfc\xd7\x82\x6c\xa9\x42\xa7\xaf\xcc\x19\xe9\xe1\x07\xfa\xf7\x91\xb8\x08\xf6\xed\x5d\xf7\x09\xf7\xeb\x9d\x5e\x3d\x50\xff\x04\xce\xdf\x78\xf6\xb7\xd0\x9c\x7f\x1b\x8d\x79\xa0\x29\xef\x36\xe4\xa7\x35\xe3\xa7\x36\xe2\xbd\x4d\x35\x2c\x06\x9a\x6f\xa7\x3a\xa7\x71\xc6\xb1\xa4\x9f\x71\xce\xf0\x03\xa7\xdd\xaa\x43\x1a\x8a\xe8\xab\x21\x31\x47\x73\xab\x23\x5b\x18\x01\x22\xcb\xe9\x8e\x1d\x9c\x4c\xa9\x56\x16\x49\xe8\x0e\x17\x5c\x36\xff\x77\xa2\x97\x1a\xdd\x0e\xb7\x6a\xf5\x6f\x00\x00\x00\xff\xff\x0d\xc3\x3d\xde\xfb\x22\x00\x00")

func tmplDashboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplDashboardTmpl,
		"tmpl/dashboard.tmpl",
	)
}

func tmplDashboardTmpl() (*asset, error) {
	bytes, err := tmplDashboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/dashboard.tmpl", size: 8955, mode: os.FileMode(436), modTime: time.Unix(1595276227, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplMonitorTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\xcf\x6e\xe2\x30\x10\xc6\xef\x79\x8a\x51\xce\xbb\x3c\x41\x39\xac\x5a\x56\xe5\xb0\xa0\xad\x90\x7a\x58\xad\x2c\x2b\x1e\x88\xd5\x60\x53\xc7\x21\x42\x5e\xbf\xfb\xca\xff\x20\x4e\x4d\xcb\x29\xfe\xbe\x6f\x7e\x99\x4c\x26\x28\xec\xe5\xa0\x1a\x84\x9a\x51\x4d\x99\x3c\x90\xa3\x14\x5c\x4b\x55\x43\xcd\x18\x31\x06\x16\x6b\x06\xd6\xd6\x60\x2a\x00\x41\x8f\x08\xf9\x6f\x09\xb5\x0b\x6d\x9c\x63\x6d\x5d\x01\xe8\xcb\xe9\x4e\x68\xe7\x9c\x10\x32\xe6\x3b\xf0\x3d\x2c\x76\xf4\xd0\x83\xb5\xae\xcc\x5d\xcd\xcb\xfe\x18\xa3\xa8\x38\x60\x08\x5a\x5b\x1b\xb3\xb0\xb6\xfe\x66\x0c\x0a\x66\xed\xdf\x48\x42\xc1\x02\xe4\x88\x7d\x4f\x0f\x98\x43\x1e\x1e\x56\xdb\x5d\xe5\x1a\xf8\x15\x6d\x6b\x2b\x27\x01\x60\xdf\xd0\x8e\x6a\x2e\x05\x49\xa5\xd3\xfc\xf6\xe4\xac\x7e\xb1\xba\xc6\x66\x84\x0a\xe0\x7d\x40\x75\x81\x25\xb8\xfc\x6f\x7f\xfd\xcf\x63\x4f\xf8\xd8\x52\x45\x1b\x8d\xca\x3f\x61\x6c\x75\xe4\xba\xbd\x82\x43\xd3\x69\x16\x1b\xa9\xf9\xfe\xb2\x91\x4f\x54\xd3\xe0\x08\xaf\x10\x21\x89\x7b\x3b\xe1\x69\xfc\xb4\x3f\x24\xb3\x29\x24\xe0\x0b\x06\xc0\x5a\x68\x54\x67\xda\x05\x57\x45\x95\xf0\x24\x07\x68\x39\x5d\x04\x87\xfb\xff\x18\x18\xd7\x59\xa3\xd4\x2b\xd3\x1e\x27\x99\x22\x69\xc7\x8f\x28\x07\xfd\x1c\x77\x20\x9c\x48\x1b\x19\xb9\x5b\x04\xac\x45\xd3\x0d\x0c\x6f\x7b\xc4\x83\x40\xfc\x3e\x05\xcc\x34\xf3\xe9\xb8\xde\x07\xae\xf0\xe7\xd0\x75\xaf\x5c\x30\x39\xa6\x79\x79\x99\xec\x87\xae\x23\x63\x30\xd2\xc4\x8a\x05\xe5\x91\xe1\xf8\x2c\x7b\xfd\x84\x1d\xbd\xc4\x99\xe1\x48\x5a\xd9\x6b\xc2\xbc\x16\xa7\xf6\x21\x56\xa4\xad\xce\xb4\x1b\xfc\x42\x4e\x92\x78\x15\x33\x64\x31\x3b\xa1\x4e\xdf\x46\xab\xb0\x6f\x65\xc7\x26\x9b\x19\x16\x76\xee\xe8\xdb\x79\xe9\xff\x1a\x6e\x90\xed\x5b\x88\x00\xc8\xb7\xd8\xc2\x4d\xca\x9e\xe6\x56\xf3\x4a\x95\xe0\xe2\x90\xe4\x31\x1e\x43\xf5\xcc\xfc\x1c\xf1\x82\x8d\x3c\xbb\xaf\x30\x47\x11\x95\xf4\x8c\x39\x4f\xdf\x61\x3f\x2a\xae\x79\x93\x3e\x09\x80\x26\x9d\x03\x6c\x6e\x7f\x41\x99\xdf\x34\xd1\xe6\x3d\xde\xcb\x67\xf8\xe2\x8a\x94\x0e\x95\xad\xfe\x07\x00\x00\xff\xff\x41\xd6\x1b\xf0\xed\x05\x00\x00")

func tmplMonitorTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplMonitorTmpl,
		"tmpl/monitor.tmpl",
	)
}

func tmplMonitorTmpl() (*asset, error) {
	bytes, err := tmplMonitorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/monitor.tmpl", size: 1517, mode: os.FileMode(436), modTime: time.Unix(1594064150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplScreenboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x59\xd9\x52\xe4\x3a\x12\x7d\xe7\x2b\x14\xf5\x3c\x4d\x7f\x01\x0f\x34\xf4\x42\x04\x4c\x33\x14\x4d\xcf\x12\x13\x0e\x61\x67\xb9\x14\xd8\x56\xb5\x24\x17\x14\x35\xfe\xf7\x09\xed\x29\x97\xb7\xbe\x71\x2f\x2f\x94\xce\x49\x9d\xd4\x92\x4e\x6d\x02\x24\x6f\x45\x0e\x64\x55\x50\x45\x0b\x5e\x66\x32\x17\x00\xcd\x33\xa7\xa2\x58\x91\x55\x51\x64\xc7\x23\x39\xbf\x29\x48\xd7\xad\xc8\xf1\x8c\x10\xc5\x54\x05\xe4\x82\xac\x34\xfe\x68\x0a\x5d\xb7\x3a\x23\xe4\x78\xfc\x40\xd8\x86\x9c\x3f\x00\x2d\xbe\x37\xd5\x81\x74\xdd\x19\x21\x02\x68\x91\x71\x5d\xbc\x20\xba\x46\xca\xea\x3a\xd0\x14\xb1\xa0\x05\xd6\x5b\x2a\xc0\x61\xd2\xfe\xb6\x75\x31\x31\x58\xf3\x11\xea\x5d\x45\x15\x3c\x51\xc1\xe8\x73\x05\x32\xd2\xaf\x4c\x6d\x27\x0d\x04\x6d\x4a\x20\xe7\x16\x50\xce\x2e\xdb\x3b\x43\xd3\x75\x42\x1a\x5a\x83\xfe\xef\xba\xff\x77\x5d\xb4\xbd\x27\x64\x27\x60\xc3\xde\x02\x77\x6f\x8b\x9e\x2d\x60\x43\xdb\x4a\x79\xf6\xda\x15\x2d\x3d\xd8\xa1\xa1\x42\xd2\x99\x9f\xac\x28\x41\x8d\x75\xe1\xd5\xb0\xae\xdd\xea\xb0\x8b\x73\xa6\x7f\xfb\x66\xbd\x79\xf4\x9f\x01\x3a\x78\xe8\x5f\x01\x0a\xe3\xeb\xa6\xdb\x6a\x9e\x04\xc2\x23\xbc\xa9\xa4\x4e\xe8\x42\x4f\xe2\xb2\x62\x65\x93\xe8\x64\xd4\x40\x58\xcd\x1b\xcd\xcb\xad\xd9\x7b\xda\xaa\x4c\x6a\x04\x8b\x39\x93\x49\xad\x6f\xc0\xca\xad\xf2\xe8\xd6\x96\x9c\x48\xe0\x26\x15\x7e\xb2\x42\x6d\x3d\xf8\x6a\x0a\xae\xbe\x67\xa6\x3b\x63\xc7\xcf\xf6\x43\xff\xf6\x3d\x58\x30\xae\x57\xbc\xe2\xc2\x83\xb9\x29\xb8\xda\x9e\x99\x19\xc7\x1a\x30\xe6\x3e\x17\x84\x2a\xfd\xdb\x46\x13\x21\x15\xdb\x43\x26\x77\x34\xcc\xd8\x2d\xdb\xc3\x5a\x97\xbd\x9f\x6e\xcc\xdb\xa0\xf3\x0a\xae\x61\x33\xe4\x3f\x21\x14\xab\x20\x2b\x60\x13\x9a\xe1\x05\x9e\xd8\xbb\xb7\x21\x64\xcf\xde\x7d\xab\x2c\xbe\x42\xd6\xc8\x3b\x1a\xbb\x56\x2a\x5e\xff\x68\x98\x8a\x5c\x6e\xb0\xac\xd5\xa0\x1f\x48\x6c\x36\x2b\x7a\xd9\x2a\x2e\x73\x1a\xbf\x17\x42\x68\x80\x9c\x22\xb6\x99\x15\xd4\x51\x90\x7c\x38\x36\x48\x7a\x5f\x0e\x32\x9a\x55\x7c\x80\x5f\x2d\x48\x9f\x41\xd2\xa1\x1f\xe6\x92\x0c\xa3\xff\x84\x35\x0b\x33\x12\xc5\xff\xd1\x82\x38\x44\x43\x42\x7e\xf9\x36\x7a\x66\x95\xd4\x49\x5a\x88\x7a\x6d\xb3\x55\xc0\x47\x33\xd9\x8c\x8e\x71\xda\x17\xfb\xa5\xc1\x0c\x4b\x62\xb3\x45\xba\x77\xa0\x04\xcb\x31\x53\x5b\xc4\x09\x06\x7e\x59\x6f\xe1\x4d\x7d\x61\x95\x02\x91\xf4\x59\x4f\xf3\xc6\xc2\x68\x9e\x83\xe1\x22\xe9\x5b\x56\xe3\xf8\xd6\xdf\x70\x1d\x63\xdb\xb3\x8b\xa4\x2e\xcb\x52\x40\x49\x15\x4f\x5a\x49\x23\xea\xc3\x1b\xdb\x2d\x52\xbe\xe2\xf5\x8e\x0a\x78\xe4\x98\xcc\x2d\x98\x29\x1e\x53\x5a\x34\x5b\xa6\xbb\xd5\x81\xdb\x9f\xfe\xdc\xa0\xc9\xfc\x27\x86\x8b\xa4\xbf\x8b\x02\xc4\xa7\x24\xd2\xb9\x86\xb2\xe7\xb0\x90\x46\x93\xe5\x8a\xd7\x4c\x9c\x4a\x16\x4c\x24\x9a\xd6\x68\x91\xe8\xe7\x37\x25\xe8\x15\xaf\x30\x07\x1a\xcb\x72\x5e\x79\x51\x64\xb4\x48\xf4\xa6\xc9\x05\x50\x09\x5f\x39\x4f\x78\xe6\xf0\xac\xd4\x84\x13\xef\x19\x2f\x8c\x87\xa6\x60\x8a\xf1\x86\x56\x5f\xb8\xa8\x29\xce\x47\x38\x5b\x0d\xd8\x7d\xe8\x19\x9e\xa4\x2e\x1d\x58\xa1\x56\xb6\x31\xd5\x50\x1e\x1b\x59\x5a\x7d\xcd\xd1\x05\x76\xb4\x4f\x51\xf0\x9e\x56\xa0\x14\xa4\xe4\xce\x81\x7e\xfb\x18\x6c\x16\xcb\xda\xaf\xa2\xff\x59\xfa\xef\x07\x7f\x98\x89\xe5\x62\xfd\x9b\x66\x0f\x42\xa5\x1c\xb3\x58\x98\x62\x67\xb1\x58\xf3\x89\x56\x6d\x6f\x20\xf6\x06\xf2\x2b\xb8\xe3\x97\x37\xb2\xa6\x25\xfc\x78\xb8\xed\x35\x53\xa3\x59\x2b\x42\xa4\x23\xb3\x19\xe9\x6e\x2e\x4e\x97\x41\xe6\x60\xa3\x0e\x15\x0c\x07\xf0\x09\x25\x0d\x70\x31\x18\x90\x7f\x51\xfc\xf4\xd3\xe3\xcc\x62\x3b\xa3\x96\xec\x85\xed\xdf\xc4\x8e\x78\x7c\xf8\xc9\xc7\x8f\x66\x2c\xfe\xe0\x98\x3f\xd2\xd2\xae\x93\x49\xe6\x50\xb4\x74\x0b\xaa\x24\x17\xe4\x3f\xc7\xa3\xcb\x0e\xd1\xba\xeb\x56\xc7\xe3\x79\xd7\xad\xfe\x76\x3c\x42\x53\x74\xdd\x7f\xc7\xbd\xe9\x36\xba\x8d\xd0\xd4\x96\x6b\x1e\x30\x89\x7a\x0f\x4d\x92\xbe\x62\x90\x0c\x52\x27\x79\x0d\xb4\x15\x8a\x9b\xd1\x7d\x97\x6e\xb6\x31\xee\x35\xe9\xc3\x49\x23\x87\x91\x5e\xc3\xef\xa8\x78\x49\x86\x39\x36\x7c\x90\x3a\x69\x78\x6d\xac\x06\xb6\x92\x7f\xd6\x1e\xf0\x96\x3e\x43\xb2\x00\x56\x06\xf0\x5b\x20\xc7\x2e\x92\x3a\xc9\x5b\x33\x59\x6b\x30\x66\x86\x06\x7f\x71\xc4\x68\x01\x7f\x20\x3a\x1b\xe0\x27\xce\x98\xc9\x79\x59\xef\x2d\x93\xe3\x72\x34\x98\x3c\x31\xe2\xe3\x12\x3e\x25\x25\xe7\xa3\x91\xba\xf7\x02\x72\x26\x19\x0f\x47\x99\x5d\x00\xe2\xed\x49\xb0\x98\x3d\x33\xa7\xd7\x09\x0b\xce\x44\x23\x5a\x5f\x78\x93\x8c\xcd\x86\x37\xe9\xd8\x20\x83\x49\xa1\xcb\x0a\x84\xba\xb9\xf6\x30\xd5\xc5\x8c\x85\x9d\x50\xa4\xa7\x55\x5a\xc5\x1f\x60\x23\x40\x86\x3c\xaa\x4f\x90\x99\x70\x18\x3a\x44\x46\xb3\x49\xc5\x5b\x28\x11\x5a\xd9\x92\x0f\x7f\xcf\x2d\x50\xc0\x83\x64\x55\x92\x61\x4a\x8c\x26\xe5\x92\x43\xa2\x39\x8d\x0d\x27\xab\xb1\x18\x8c\xeb\x3c\x5a\xdb\xf1\xb2\x3e\x76\x5b\xb2\x85\xfc\x25\xdc\x96\x98\x42\x38\x00\x58\x66\xb2\xfa\x57\xc1\xdb\x1d\x6b\x4a\x8f\x97\xbe\xec\x44\x10\x3f\xaf\x93\x88\x24\x0a\xf3\x97\x36\xf9\xcb\x3d\x97\xf1\x7a\x24\x7f\xc9\x76\x5c\xc6\x8b\x2f\x4f\xcf\xaa\x7c\x2e\x4a\x48\x64\x40\x03\x48\xc7\x19\x4c\xdf\x9f\x3d\xde\x85\xe9\xd8\xaa\x3a\xcc\x87\xc3\x67\x1b\x81\x1b\x80\x7d\xcf\xd5\xfd\x54\xe6\x78\x87\xfe\x5c\x26\xbb\xf3\xc8\x4e\x8a\x98\xb5\x15\xc7\xb5\x49\xce\x49\x58\x63\x93\x49\xad\x35\x7b\x47\xc1\x21\x6d\xc9\x89\x04\x6e\x52\xe1\x8e\x8a\x92\x85\x94\x56\xdb\x92\xbf\x51\xf0\xdc\x74\x7f\x9a\x7d\xe8\x49\xb3\x0f\x5d\x30\xe8\x74\xe3\x41\xec\x59\x0e\xee\x5f\xe8\x84\x2d\x66\xee\x7f\xe8\x4d\xdf\x78\x89\xb4\xbb\x35\x4f\x74\xcd\xc5\x7a\x2a\x9a\x5c\xae\x8f\x8f\x34\x3c\x81\xc0\x4b\x89\x9e\xb1\x6c\xef\xb0\x38\xe8\xc8\x6c\x3a\xc1\xd1\x03\x6f\x55\x4f\xb3\x32\x60\x5f\xb5\x6f\x3a\x3d\xa3\xad\x54\xeb\x2d\x7f\xfd\xc6\xe2\x19\xb6\x6e\xa5\xca\xe4\x96\xbf\x66\x5b\x8d\xfa\xf9\x4d\x2d\x17\xa9\x7e\x16\x82\x8b\x01\x5d\xb0\x78\x4f\x39\x58\x2f\xd2\xbe\xa5\x0a\x9a\xfc\x70\x2a\x5e\x39\xa2\xa7\x1e\xed\x17\xc9\x7f\x12\x40\x5f\x0a\xfe\xda\x9c\x3a\x78\x0e\x54\xcf\x05\xae\xb3\xc8\xc9\x35\x93\x4a\xb0\xe7\x56\xa1\x49\x8d\x7e\x0a\xcc\xf6\x5c\xf5\x6a\x2e\xf2\xf6\xe0\x1e\xd5\x6e\x99\x54\xa7\xde\xfc\x93\x5b\x56\x69\xba\xe7\xae\x57\x75\xd2\xdd\x35\x93\xbb\x8a\x1e\xec\x85\x87\x27\x0b\x0b\xfa\xfb\x0c\xff\xdc\xd4\x33\x9d\x7f\x4a\xb8\x17\xb0\x01\x01\x4d\x4c\x01\x26\x8d\x66\xbb\x88\xe3\xeb\x8f\xc4\x7c\x7a\x95\x60\x05\xfc\x1b\x04\xbf\xe2\x6d\x13\x3f\x85\x2d\x2b\x20\x7b\x07\xc1\xb3\xdc\xe2\x7e\xed\xe8\x5b\xcf\xa4\xcd\x86\x96\xb0\x56\x54\xb5\x52\x8f\x66\xf2\x5e\x55\x1b\x32\x93\x86\xb5\x53\x91\x3c\x61\x8d\x55\x5e\xec\x11\xbf\x81\x0d\x78\xb4\x2f\x53\xf8\x5d\x67\xac\xf2\xef\x79\xc4\x8b\xd6\x90\x47\xbc\x84\x8d\x55\xfe\x3d\x8f\xc9\x7e\x7b\xc8\x65\xb2\xfd\x1e\xad\x3e\x7d\x44\xa0\x82\xd6\x12\xa3\xf6\x28\x99\xe2\x3b\x5b\xea\xbf\x07\xad\x39\xbe\xa7\x92\x3c\xde\x50\x39\x66\xd1\x63\x4b\x44\x47\x1f\xe3\x26\x14\x4c\xc4\x46\xd8\x84\x75\xfc\x64\x2c\x37\x2b\xb2\x56\x34\xe9\x89\x29\xfa\xae\x38\x6e\x44\x64\xf4\x14\x38\xfc\xc5\xb7\x75\x23\xd1\x97\x6e\x8a\xf6\xd5\xdd\x93\xff\x23\x20\x73\xba\x83\xab\x2d\x15\x34\xc7\x37\x2a\x63\x4b\x29\x2f\x25\x84\xd6\x57\xb6\xe4\xd7\x4e\xcf\x0d\x04\x81\x3e\xd9\xda\xd7\xeb\x89\x97\xf0\xee\xec\xff\x01\x00\x00\xff\xff\x8a\x1e\x80\x8f\xbb\x20\x00\x00")

func tmplScreenboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplScreenboardTmpl,
		"tmpl/screenboard.tmpl",
	)
}

func tmplScreenboardTmpl() (*asset, error) {
	bytes, err := tmplScreenboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/screenboard.tmpl", size: 8379, mode: os.FileMode(436), modTime: time.Unix(1594064150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplTimeboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xc9\x6e\xdb\x3c\x10\xbe\xfb\x29\x06\x42\x0e\xff\x0f\xc4\x7e\x80\x02\x3e\xa4\x09\x1c\x14\xe8\x92\x26\x41\x7a\x28\x0a\x81\x96\x46\x0a\x51\x6a\x09\x45\x25\x71\x08\xbe\x7b\xc1\xe1\x26\xdb\x8a\x7b\xa8\x4f\x9c\xf9\xbe\x59\x39\x43\x4b\xe2\xd0\x8d\xb2\x40\xc8\x4a\xa6\x58\xd9\xd5\xb9\xe2\x0d\x6e\x3b\x26\xcb\x57\xf7\xcb\x20\x2b\xcb\x5c\x6b\x58\x7d\x2a\xc1\x98\x0c\xf4\x02\x40\x71\x25\x10\xdc\x6f\x0d\x99\x45\xef\x49\x65\x4c\xb6\x00\x28\x71\x28\x24\xef\x15\xef\xda\x00\x5f\x4d\x54\x8e\x24\x91\x95\x79\xd7\x8a\x1d\xf9\xb0\x9c\x5b\x64\xe5\x37\xab\x58\x1a\xb3\x00\xd0\xfa\x85\xab\x47\x58\x5d\x4b\xd6\x3f\x0e\x51\x29\x59\x5b\x23\xac\x80\xc4\xda\x62\x94\x52\x48\x6a\x26\x9d\xe4\xea\x0a\x2b\xde\x72\x4a\xc2\xb9\x03\x78\xe6\x6f\xa9\x88\x07\xfe\x66\x81\x60\xb4\x04\x5e\xc1\xea\x62\x54\xdd\x50\x30\x81\x16\x62\x51\xf0\x26\x09\x35\x26\xb3\x26\xd8\x96\xde\x75\x70\x70\x23\xb1\xe0\x83\x0f\xda\x47\xc1\x3b\x48\xe8\xbb\x0e\xae\x65\x37\xf6\xd4\x81\xda\x9e\x60\x0d\x3f\xb5\x3e\xab\x9d\xf6\xc3\x3a\x10\x8c\x09\xdd\x39\xe3\x6d\x89\xaf\xe7\x70\x86\x02\x9b\x03\x06\xaf\x3c\x6c\xcc\xb9\xd6\x14\x2c\xd3\x9a\x98\x74\x22\xcd\xaf\xf9\x44\xee\x8a\xae\x47\x4a\x64\xb0\x27\x9f\xc8\xe0\xb4\x36\x8c\x23\x9c\x4a\x24\x31\xfe\x29\x11\xb5\x73\xf7\x41\xda\x81\xa4\xb5\x9f\x03\x80\x9e\x09\x54\x0a\xf7\xe6\x93\x2c\x56\x37\x1e\x09\x77\x1c\xb9\x79\x25\x78\x3f\xcb\xdd\x58\x20\xf0\xcd\x7c\x3a\x5f\x98\xfc\x8d\x92\x1a\x63\x55\x6e\xd8\xf6\x94\x7e\x6a\xbd\x5d\x43\x50\xcc\x57\xed\x7a\x4c\x8b\x64\x85\x94\xdf\x33\x13\x63\x1c\xb6\x07\x12\x12\xa8\x35\x85\xff\xcc\xb6\x28\x6c\x1c\x41\x07\x4f\x76\xda\xa3\xa1\x4a\x25\x1c\x1f\x62\x51\xae\x82\x5b\x7c\x1a\x71\x50\xb3\x25\x48\x87\xc5\x1a\x9e\x26\xbd\xfe\x3e\xa2\xdc\xa5\x45\x8a\x79\x52\x69\x4b\x63\xa8\xde\x83\x72\xb5\xb6\x29\x80\xf7\x1e\x4d\x2e\xea\x5a\x62\xcd\x54\x27\x5d\x12\x56\xd9\x22\x64\x19\xfc\x77\x85\xb7\x58\xdd\x29\xc9\xdb\x7a\xca\xfb\x9f\xd6\x34\x99\x85\x3d\x4d\x9a\x18\x8d\x1c\xda\xa8\xcb\xc3\xb0\x71\xbe\xb4\xf6\xbd\x70\x9a\xc8\x3b\x1c\x39\x6b\x49\xbb\xee\x07\xcc\x6e\xba\x3f\x86\x3d\x4f\xa3\xb7\xdf\xef\x64\xfc\x83\x97\xea\xd1\x9a\xbe\xd0\xc1\x1b\x3a\xed\x09\xb3\xd3\x6d\xdd\xb7\x99\xd4\xe9\x0b\x0f\xad\x88\x80\x1b\xe9\xcb\xae\x2d\xe9\xa5\x64\x62\xd3\xc9\x86\xa9\x01\xa6\xc3\xfd\x2e\x1c\x9e\xe7\xd4\xd2\x22\x51\xf3\x8a\xb8\x7b\x5d\x83\x69\xdb\x4e\x77\x6d\x7f\x42\xac\xe7\xa6\x67\x72\x7a\xcb\x97\x49\x93\xa6\x2f\x15\x15\xf6\x67\x7e\xa9\x8e\x03\xc4\x66\x8c\x83\xea\x9a\x8f\xf5\x65\x27\xc8\x73\x41\x72\xbe\xad\xf3\x82\x34\x21\xfa\x01\xed\xaf\x1e\x37\x87\x1e\xab\x59\x8f\x9b\xf7\x3d\x4e\xaf\xcd\x01\x73\xa7\x45\xa2\x1e\xef\xfa\xbc\x64\xfc\x03\x3c\xd5\xa5\xeb\xbf\xc7\xa6\x17\x4c\xe1\x03\x93\x9c\x6d\x05\xc6\xa7\x6f\xf2\xef\x6c\xbf\x13\x3c\x2d\x7f\xf6\x3c\x7f\xf3\x2d\x6b\x70\xf2\x5e\x7c\xb5\x62\xb8\xaf\x5e\x62\xc5\x5f\x61\xf2\xf7\x68\xc5\x80\x96\x58\xb1\x51\xa8\xf4\x59\xe1\x44\x07\x1f\xe7\x6c\x16\x8b\x3f\x01\x00\x00\xff\xff\x17\xc0\x0c\xf4\xe3\x08\x00\x00")

func tmplTimeboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplTimeboardTmpl,
		"tmpl/timeboard.tmpl",
	)
}

func tmplTimeboardTmpl() (*asset, error) {
	bytes, err := tmplTimeboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/timeboard.tmpl", size: 2275, mode: os.FileMode(436), modTime: time.Unix(1594175251, 0)}
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
	"tmpl/dashboard.tmpl": tmplDashboardTmpl,
	"tmpl/monitor.tmpl": tmplMonitorTmpl,
	"tmpl/screenboard.tmpl": tmplScreenboardTmpl,
	"tmpl/timeboard.tmpl": tmplTimeboardTmpl,
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
		"dashboard.tmpl": &bintree{tmplDashboardTmpl, map[string]*bintree{}},
		"monitor.tmpl": &bintree{tmplMonitorTmpl, map[string]*bintree{}},
		"screenboard.tmpl": &bintree{tmplScreenboardTmpl, map[string]*bintree{}},
		"timeboard.tmpl": &bintree{tmplTimeboardTmpl, map[string]*bintree{}},
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

