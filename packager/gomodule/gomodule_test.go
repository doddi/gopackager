package gomodule

import (
	"reflect"
	"testing"
)

func TestGoModule_GetModuleZipName(t *testing.T) {
	type fields struct {
		vcs     string
		owner   string
		name    string
		version string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"succes",
			fields{"", "", "gopackager", "v0.0.1"},
			"gopackager@v0.0.1.zip",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			goModule := GoModule{
				vcs:     tt.fields.vcs,
				owner:   tt.fields.owner,
				name:    tt.fields.name,
				version: tt.fields.version,
			}
			if got := goModule.GetModuleZipName(); got != tt.want {
				t.Errorf("GetModuleZipName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		path    string
		version string
	}
	tests := []struct {
		name    string
		args    args
		want    GoModule
		wantErr bool
	}{
		{
			"success",
			args{"github.com/doddi/gopackager", "v0.0.1"},
			GoModule{"github.com", "doddi", "gopackager", "v0.0.1"},
			false,
		},
		{
			"no vcs supplied",
			args{"doddi/gopackager", "v0.0.1"},
			GoModule{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.path, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoModule_GetProjectPath(t *testing.T) {
	type fields struct {
		vcs     string
		owner   string
		name    string
		version string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"success",
			fields{"github.com", "doddi", "gopackager", "v1.0.0"},
			"github.com/doddi/gopackager",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			goModule := GoModule{
				vcs:     tt.fields.vcs,
				owner:   tt.fields.owner,
				name:    tt.fields.name,
				version: tt.fields.version,
			}
			if got := goModule.GetProjectPath(); got != tt.want {
				t.Errorf("GetProjectPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
