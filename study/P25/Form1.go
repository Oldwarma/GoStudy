// Automatically generated by the res2go, do not edit.

package main

import (
	"github.com/ying32/govcl/vcl"
)

type TForm1 struct {
	*vcl.TForm
	FirstWeekStart  *vcl.TComboBox
	FirstWeekEnd    *vcl.TComboBox
	SecondWeekStart *vcl.TComboBox
	SecondWeekEnd   *vcl.TComboBox
	ThirdWeekStart  *vcl.TComboBox
	ThirdWeekEnd    *vcl.TComboBox
	Button1         *vcl.TButton

	//::private::
	TForm1Fields
}

var Form1 *TForm1

// Loaded in bytes.
// vcl.Application.CreateForm(&Form1)

func NewForm1(owner vcl.IComponent) (root *TForm1) {
	vcl.CreateResForm(owner, &root)
	return
}

var Form1Bytes = []byte("\x54\x50\x46\x30\x0B\x54\x44\x65\x73\x69\x67\x6E\x46\x6F\x72\x6D\x05\x46\x6F\x72\x6D\x31\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\x90\x01\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\x58\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x05\x46\x6F\x72\x6D\x31\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x90\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x58\x02\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x0E\x46\x69\x72\x73\x74\x57\x65\x65\x6B\x53\x74\x61\x72\x74\x04\x4C\x65\x66\x74\x02\x56\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x80\x00\x05\x57\x69\x64\x74\x68\x02\x64\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x11\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x0C\x46\x69\x72\x73\x74\x57\x65\x65\x6B\x45\x6E\x64\x04\x4C\x65\x66\x74\x03\xE9\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x80\x00\x05\x57\x69\x64\x74\x68\x02\x64\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x11\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x0F\x53\x65\x63\x6F\x6E\x64\x57\x65\x65\x6B\x53\x74\x61\x72\x74\x04\x4C\x65\x66\x74\x02\x56\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xCB\x00\x05\x57\x69\x64\x74\x68\x02\x64\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x11\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x0D\x53\x65\x63\x6F\x6E\x64\x57\x65\x65\x6B\x45\x6E\x64\x04\x4C\x65\x66\x74\x03\xEA\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xCE\x00\x05\x57\x69\x64\x74\x68\x02\x64\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x11\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x0E\x54\x68\x69\x72\x64\x57\x65\x65\x6B\x53\x74\x61\x72\x74\x04\x4C\x65\x66\x74\x02\x57\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x22\x01\x05\x57\x69\x64\x74\x68\x02\x64\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x11\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x0C\x54\x68\x69\x72\x64\x57\x65\x65\x6B\x45\x6E\x64\x04\x4C\x65\x66\x74\x03\xE6\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x21\x01\x05\x57\x69\x64\x74\x68\x02\x64\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x11\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x31\x04\x4C\x65\x66\x74\x03\x7C\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xCB\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x07\x42\x75\x74\x74\x6F\x6E\x31\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x06\x00\x00\x00")

// 注册窗口资源
var _ = vcl.RegisterFormResource(Form1, &Form1Bytes)
