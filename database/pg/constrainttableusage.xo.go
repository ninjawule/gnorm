// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

// ConstraintTableUsageTable is the database name for the table.
const ConstraintTableUsageTable = "information_schema.constraint_table_usage"

// ConstraintTableUsage represents a row from 'information_schema.constraint_table_usage'.
type ConstraintTableUsage struct {
	TableCatalog      SQLIdentifier `json:"table_catalog"`      // table_catalog
	TableSchema       SQLIdentifier `json:"table_schema"`       // table_schema
	TableName         SQLIdentifier `json:"table_name"`         // table_name
	ConstraintCatalog SQLIdentifier `json:"constraint_catalog"` // constraint_catalog
	ConstraintSchema  SQLIdentifier `json:"constraint_schema"`  // constraint_schema
	ConstraintName    SQLIdentifier `json:"constraint_name"`    // constraint_name
}

// Constants defining each column in the table.
const (
	ConstraintTableUsageTableCatalogField      = "table_catalog"
	ConstraintTableUsageTableSchemaField       = "table_schema"
	ConstraintTableUsageTableNameField         = "table_name"
	ConstraintTableUsageConstraintCatalogField = "constraint_catalog"
	ConstraintTableUsageConstraintSchemaField  = "constraint_schema"
	ConstraintTableUsageConstraintNameField    = "constraint_name"
)

// WhereClauses for every type in ConstraintTableUsage.
var (
	ConstraintTableUsageTableCatalogWhere      SQLIdentifierField = "table_catalog"
	ConstraintTableUsageTableSchemaWhere       SQLIdentifierField = "table_schema"
	ConstraintTableUsageTableNameWhere         SQLIdentifierField = "table_name"
	ConstraintTableUsageConstraintCatalogWhere SQLIdentifierField = "constraint_catalog"
	ConstraintTableUsageConstraintSchemaWhere  SQLIdentifierField = "constraint_schema"
	ConstraintTableUsageConstraintNameWhere    SQLIdentifierField = "constraint_name"
)

// QueryOneConstraintTableUsage retrieves a row from 'information_schema.constraint_table_usage' as a ConstraintTableUsage.
func QueryOneConstraintTableUsage(db XODB, where WhereClause, order OrderBy) (*ConstraintTableUsage, error) {
	const origsqlstr = `SELECT ` +
		`table_catalog, table_schema, table_name, constraint_catalog, constraint_schema, constraint_name ` +
		`FROM information_schema.constraint_table_usage WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	ctu := &ConstraintTableUsage{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&ctu.TableCatalog, &ctu.TableSchema, &ctu.TableName, &ctu.ConstraintCatalog, &ctu.ConstraintSchema, &ctu.ConstraintName)
	if err != nil {
		return nil, err
	}
	return ctu, nil
}

// QueryConstraintTableUsage retrieves rows from 'information_schema.constraint_table_usage' as a slice of ConstraintTableUsage.
func QueryConstraintTableUsage(db XODB, where WhereClause, order OrderBy) ([]*ConstraintTableUsage, error) {
	const origsqlstr = `SELECT ` +
		`table_catalog, table_schema, table_name, constraint_catalog, constraint_schema, constraint_name ` +
		`FROM information_schema.constraint_table_usage WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*ConstraintTableUsage
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		ctu := ConstraintTableUsage{}

		err = q.Scan(&ctu.TableCatalog, &ctu.TableSchema, &ctu.TableName, &ctu.ConstraintCatalog, &ctu.ConstraintSchema, &ctu.ConstraintName)
		if err != nil {
			return nil, err
		}

		vals = append(vals, &ctu)
	}
	return vals, nil
}