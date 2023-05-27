package core

type Config struct {
	Package                     string   `json:"namespace"`
	EmitExactTableNames         bool     `json:"emit_exact_table_names"`
	Async                       bool     `json:"async"`
	EmitClasses                 bool     `json:"emit_classes"`
	TypeAffinity                bool     `json:"type_affinity" default:"true"`
	InflectionExcludeTableNames []string `json:"inflection_exclude_table_names"`
}
