package run

import "text/template"

// DBType defines the list of supported databases.
type DBType int

// Supported DB Types.
const (
	Postgres DBType = iota
	Mysql
)

// Config holds the schema that is expected to exist in the gnorm.toml file.
type Config struct {
	// ConnStr is the connection string for the database.  Environment variables
	// in $FOO form will be expanded.
	ConnStr string

	// The type of DB you're connecting to.  Currently the possible values are
	// "postgres" or "mysql".
	DBType DBType

	// Schemas holds the names of schemas to generate code for.
	Schemas []string

	// IncludeTables is a map of schema names to table names. It is whitelist of
	// tables to generate data for. Tables not in this list will not be included
	// in data generated by gnorm. You cannot set IncludeTables if ExcludeTables
	// is set.
	IncludeTables map[string][]string

	// ExcludeTables is a map of schema names to table names.  It is a blacklist
	// of tables to ignore while generating data. All tables in a schema that
	// are not in this list will be used for generation. You cannot set
	// ExcludeTables if IncludeTables is set.
	ExcludeTables map[string][]string

	// TemplateDir contains the relative path to the directory where gnorm
	// expects to find templates to render.  The default is the current
	// directory where gnorm is running.
	TemplateDir string

	// PostRun is a command with arguments that is run after each file is
	// generated by GNORM.  It is generally used to reformat the file, but it
	// can be for any use. Environment variables will be expanded, and the
	// special $GNORMFILE environment variable may be used, which will expand to
	// the name of the file that was just generated.
	PostRun []string

	// TablePath is a relative path for tables to be rendered.  The table
	// template will be rendered with each table in turn. If the path is empty,
	// tables will not be rendered.
	//
	// The table path may be a template, in which case the values .Schema and
	// .Table may be referenced, containing the name of the current schema and
	// table being rendered.  For example, "{{.Schema}}/{{.Table}}/{{.Table}}.go" would render
	// the "public.users" table to ./public/users/users.go.
	TablePath *template.Template

	// SchemaPath is a relative path for schemas to be rendered.  The schema
	// template will be rendered with each schema in turn. If the path is empty,
	// schema will not be rendered.
	//
	// The schema path may be a template, in which case the value .Schema may be
	// referenced, containing the name of the current schema being rendered. For
	// example, "schemas/{{.Schema}}/{{.Schema}}.go" would render the "public"
	// schema to ./schemas/public/public.go
	SchemaPath *template.Template

	// EnumPath is a relative path for enums to be rendered.  The enum.tpl template
	// will be rendered with each enum in turn. If the path is empty, enums will not
	// be rendered this way (thought you could render them via the schemas template).
	//
	// The enum path may be a template, in which case the values .Schema and .Enum
	// may be referenced, containing the name of the current schema and Enum being
	// rendered.  For example, "gnorm/{{.Schema}}/enums/{{.Enum}}.go" would render
	// the "public.book_type" enum to ./gnorm/public/enums/users.go.
	EnumPath *template.Template

	// NameConversion defines how the DBName of tables, schemas, and enums are
	// converted into their Name value.  This is a template that may use all the
	// regular functions.  The "." value is the DB name of the item. Thus, to
	// make an item's Name the same as its DBName, you'd use a template of
	// "{{.}}". To make the Name the PascalCase version of DBName, you'd use
	// "{{pascal .}}".
	NameConversion *template.Template

	// TypeMap is a mapping of database type names to replacement type names
	// (generally types from your language for deserialization).  Types not in
	// this list will remain in their database form.  In the data sent to your
	// template, this is the Column.Type, and the original type is in
	// Column.OrigType.  Note that because of the way tables in TOML work,
	// TypeMap and NullableTypeMap must be at the end of your configuration
	// file.
	TypeMap map[string]string

	// NullableTypeMap is a mapping of database type names to replacement type
	// names (generally types from your language for deserialization)
	// specifically for database columns that are nullable.  Types not in this
	// list will remain in their database form.  In the data sent to your
	// template, this is the Column.Type, and the original type is in
	// Column.OrigType.   Note that because of the way tables in TOML work,
	// TypeMap and NullableTypeMap must be at the end of your configuration
	// file.
	NullableTypeMap map[string]string
}