module github.com/arnlen/converter

go 1.16

replace exchangeratesapi => ./exchangeratesapi

replace environments => ./environments

require (
	environments v0.0.0-00010101000000-000000000000
	exchangeratesapi v0.0.0-00010101000000-000000000000
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/jarcoal/httpmock v1.0.8
	github.com/tkanos/gonfig v0.0.0-20210106201359-53e13348de2f
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
