package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"dakhl/app/facades"
)

type M20260423141955CreateTanentsTable struct{}

// Signature The unique signature for the migration.
func (r *M20260423141955CreateTanentsTable) Signature() string {
	return "20260423141955_create_tanents_table"
}

// Up Run the migrations.
func (r *M20260423141955CreateTanentsTable) Up() error {
	if !facades.Schema().HasTable("tanents") {
		return facades.Schema().Create("tanents", func(table schema.Blueprint) {
			table.ID()
			table.String("name").Default("")
			table.String("photo").Default("")
			table.String("hashed_password").Default("")
			table.String("address").Default("")
			table.String("phone").Default("")
			table.String("license").Default("")
			table.String("description").Default("")
			table.SoftDeletes()
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20260423141955CreateTanentsTable) Down() error {
	return facades.Schema().DropIfExists("tanents")
}
