package parser

import (
	"testing"

	"github.com/ystia/tdt2go/internal/pkg/model"

	"gotest.tools/v3/assert"
)

func TestParser_ParseTypes(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		p       *Parser
		args    args
		want    []model.DataType
		wantErr bool
	}{
		{"NoTOSCAFile", &Parser{}, args{"testdata/donotexists.yaml"}, nil, true},
		{"InvalidTOSCAFile", &Parser{}, args{"testdata/invalid.yaml"}, nil, true},
		{"InvalidIncludeFilter", &Parser{IncludePatterns: []string{`x{2,1}}`}}, args{"testdata/normative-light.yaml"}, nil, true},
		{"InvalidExcludeFilter", &Parser{ExcludePatterns: []string{`x{2,1}}`}}, args{"testdata/normative-light.yaml"}, nil, true},
		{"TestParseNormativeLight", &Parser{}, args{"testdata/normative-light.yaml"}, []model.DataType{
			{
				Name:        "Credential",
				FQDTN:       "tosca.datatypes.Credential",
				DerivedFrom: "Root",
				Fields: []model.Field{
					{
						Name:         "Keys",
						OriginalName: "keys",
						Type:         "map[string]string",
					},
					{
						Name:         "Protocol",
						OriginalName: "protocol",
						Type:         "string",
					},
					{
						Name:         "Token",
						OriginalName: "token",
						Type:         "string",
					},
					{
						Name:         "TokenType",
						OriginalName: "token_type",
						Type:         "string",
					},
					{
						Name:         "User",
						OriginalName: "user",
						Type:         "string",
					},
				},
			},
			{
				Name:   "Root",
				FQDTN:  "tosca.datatypes.Root",
				Fields: []model.Field{},
			},
			{
				Name:        "TimeInterval",
				FQDTN:       "tosca.datatypes.TimeInterval",
				DerivedFrom: "Root",
				Fields: []model.Field{
					{
						Name:         "EndTime",
						OriginalName: "end_time",
						Type:         "time.Time",
					},
					{
						Name:         "StartTime",
						OriginalName: "start_time",
						Type:         "time.Time",
					},
				},
			},
		}, false},
		{"TestExtraToscaTypes", &Parser{}, args{"testdata/extratypes.yaml"}, []model.DataType{
			{
				Name:        "SpecificTypes",
				FQDTN:       "tosca.datatypes.SpecificTypes",
				DerivedFrom: "Root",
				Fields: []model.Field{
					{
						Name:         "ANumber",
						OriginalName: "a_number",
						Type:         "float64",
					},
					{
						Name:         "AnotherType",
						OriginalName: "another_type",
						Type:         "Credential",
					},
					{
						Name:         "TestAList",
						OriginalName: "test_a_list",
						Type:         "[]int",
					},
					{
						Name:         "ValidBool",
						OriginalName: "valid_bool",
						Type:         "bool",
					},
				},
			},
		}, false},
		{"TestParseIncludeFilters", &Parser{
			IncludePatterns: []string{`tosca\.datatypes\.Cred.*`, `tosca.datatypes.Root`},
		}, args{"testdata/normative-light.yaml"}, []model.DataType{
			{
				Name:        "Credential",
				FQDTN:       "tosca.datatypes.Credential",
				DerivedFrom: "Root",
				Fields: []model.Field{
					{
						Name:         "Keys",
						OriginalName: "keys",
						Type:         "map[string]string",
					},
					{
						Name:         "Protocol",
						OriginalName: "protocol",
						Type:         "string",
					},
					{
						Name:         "Token",
						OriginalName: "token",
						Type:         "string",
					},
					{
						Name:         "TokenType",
						OriginalName: "token_type",
						Type:         "string",
					},
					{
						Name:         "User",
						OriginalName: "user",
						Type:         "string",
					},
				},
			},
			{
				Name:   "Root",
				FQDTN:  "tosca.datatypes.Root",
				Fields: []model.Field{},
			},
		}, false},
		{"TestParseExcludeFilters", &Parser{
			ExcludePatterns: []string{`tosca\.datatypes\.Cred.*`, `tosca\.datatypes.TimeInterval`},
		}, args{"testdata/normative-light.yaml"}, []model.DataType{
			{
				Name:   "Root",
				FQDTN:  "tosca.datatypes.Root",
				Fields: []model.Field{},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.p.ParseTypes(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.ParseTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				assert.DeepEqual(t, got, tt.want)
			}
		})
	}
}
