// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// UdtPrivilegeTable is the database name for the table.
const UdtPrivilegeTable = "information_schema.udt_privileges"

// UdtPrivilege represents a row from 'information_schema.udt_privileges'.
type UdtPrivilege struct {
	Grantor       SQLIdentifier `yaml:"grantor,omitempty"`        // grantor
	Grantee       SQLIdentifier `yaml:"grantee,omitempty"`        // grantee
	UdtCatalog    SQLIdentifier `yaml:"udt_catalog,omitempty"`    // udt_catalog
	UdtSchema     SQLIdentifier `yaml:"udt_schema,omitempty"`     // udt_schema
	UdtName       SQLIdentifier `yaml:"udt_name,omitempty"`       // udt_name
	PrivilegeType CharacterData `yaml:"privilege_type,omitempty"` // privilege_type
	IsGrantable   YesOrNo       `yaml:"is_grantable,omitempty"`   // is_grantable
}

// Constants defining each column in the table.
const (
	UdtPrivilegeGrantorField       = "grantor"
	UdtPrivilegeGranteeField       = "grantee"
	UdtPrivilegeUdtCatalogField    = "udt_catalog"
	UdtPrivilegeUdtSchemaField     = "udt_schema"
	UdtPrivilegeUdtNameField       = "udt_name"
	UdtPrivilegePrivilegeTypeField = "privilege_type"
	UdtPrivilegeIsGrantableField   = "is_grantable"
)

// WhereClauses for every type in UdtPrivilege.
var (
	UdtPrivilegeGrantorWhere       SQLIdentifierField = "grantor"
	UdtPrivilegeGranteeWhere       SQLIdentifierField = "grantee"
	UdtPrivilegeUdtCatalogWhere    SQLIdentifierField = "udt_catalog"
	UdtPrivilegeUdtSchemaWhere     SQLIdentifierField = "udt_schema"
	UdtPrivilegeUdtNameWhere       SQLIdentifierField = "udt_name"
	UdtPrivilegePrivilegeTypeWhere CharacterDataField = "privilege_type"
	UdtPrivilegeIsGrantableWhere   YesOrNoField       = "is_grantable"
)

// QueryOneUdtPrivilege retrieves a row from 'information_schema.udt_privileges' as a UdtPrivilege.
func QueryOneUdtPrivilege(db XODB, where WhereClause, order OrderBy) (*UdtPrivilege, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, udt_catalog, udt_schema, udt_name, privilege_type, is_grantable ` +
		`FROM information_schema.udt_privileges WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	up := &UdtPrivilege{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&up.Grantor, &up.Grantee, &up.UdtCatalog, &up.UdtSchema, &up.UdtName, &up.PrivilegeType, &up.IsGrantable)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return up, nil
}

// QueryUdtPrivilege retrieves rows from 'information_schema.udt_privileges' as a slice of UdtPrivilege.
func QueryUdtPrivilege(db XODB, where WhereClause, order OrderBy) ([]*UdtPrivilege, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, udt_catalog, udt_schema, udt_name, privilege_type, is_grantable ` +
		`FROM information_schema.udt_privileges WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*UdtPrivilege
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		up := UdtPrivilege{}

		err = q.Scan(&up.Grantor, &up.Grantee, &up.UdtCatalog, &up.UdtSchema, &up.UdtName, &up.PrivilegeType, &up.IsGrantable)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &up)
	}
	return vals, nil
}

// AllUdtPrivilege retrieves all rows from 'information_schema.udt_privileges' as a slice of UdtPrivilege.
func AllUdtPrivilege(db XODB, order OrderBy) ([]*UdtPrivilege, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, udt_catalog, udt_schema, udt_name, privilege_type, is_grantable ` +
		`FROM information_schema.udt_privileges`

	sqlstr := origsqlstr + order.String()

	var vals []*UdtPrivilege
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		up := UdtPrivilege{}

		err = q.Scan(&up.Grantor, &up.Grantee, &up.UdtCatalog, &up.UdtSchema, &up.UdtName, &up.PrivilegeType, &up.IsGrantable)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &up)
	}
	return vals, nil
}