package packager

import (
	"reflect"
	"testing"
)

func TestGoModule_Parse(t *testing.T) {
	type fields struct {
		vcs     string
		owner   string
		name    string
		version string
	}
	type args struct {
		path    string
		version string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    GoModule
		wantErr bool
	}{
		{
			"success",
			fields{"", "", "", ""},
			args{"github.com/doddi/gopackager", "v0.0.1"},
			GoModule{"github.com", "doddi", "gopackager", "v0.0.1"},
			false,
		},
		{
			"no vcs supplied",
			fields{"", "", "", ""},
			args{"doddi/gopackager", "v0.0.1"},
			GoModule{},
			true,
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
			got, err := goModule.Parse(tt.args.path, tt.args.version)
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
