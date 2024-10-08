// Code generated by ent, DO NOT EDIT.

package destinypictures

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldLTE(FieldID, id))
}

// Picture applies equality check predicate on the "picture" field. It's identical to PictureEQ.
func Picture(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEQ(FieldPicture, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEQ(FieldPath, v))
}

// DestinyID applies equality check predicate on the "destiny_id" field. It's identical to DestinyIDEQ.
func DestinyID(v int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEQ(FieldDestinyID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEQ(FieldUpdatedAt, v))
}

// PictureEQ applies the EQ predicate on the "picture" field.
func PictureEQ(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEQ(FieldPicture, v))
}

// PictureNEQ applies the NEQ predicate on the "picture" field.
func PictureNEQ(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldNEQ(FieldPicture, v))
}

// PictureIn applies the In predicate on the "picture" field.
func PictureIn(vs ...string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldIn(FieldPicture, vs...))
}

// PictureNotIn applies the NotIn predicate on the "picture" field.
func PictureNotIn(vs ...string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldNotIn(FieldPicture, vs...))
}

// PictureGT applies the GT predicate on the "picture" field.
func PictureGT(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldGT(FieldPicture, v))
}

// PictureGTE applies the GTE predicate on the "picture" field.
func PictureGTE(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldGTE(FieldPicture, v))
}

// PictureLT applies the LT predicate on the "picture" field.
func PictureLT(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldLT(FieldPicture, v))
}

// PictureLTE applies the LTE predicate on the "picture" field.
func PictureLTE(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldLTE(FieldPicture, v))
}

// PictureContains applies the Contains predicate on the "picture" field.
func PictureContains(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldContains(FieldPicture, v))
}

// PictureHasPrefix applies the HasPrefix predicate on the "picture" field.
func PictureHasPrefix(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldHasPrefix(FieldPicture, v))
}

// PictureHasSuffix applies the HasSuffix predicate on the "picture" field.
func PictureHasSuffix(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldHasSuffix(FieldPicture, v))
}

// PictureEqualFold applies the EqualFold predicate on the "picture" field.
func PictureEqualFold(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEqualFold(FieldPicture, v))
}

// PictureContainsFold applies the ContainsFold predicate on the "picture" field.
func PictureContainsFold(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldContainsFold(FieldPicture, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldContainsFold(FieldPath, v))
}

// DestinyIDEQ applies the EQ predicate on the "destiny_id" field.
func DestinyIDEQ(v int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEQ(FieldDestinyID, v))
}

// DestinyIDNEQ applies the NEQ predicate on the "destiny_id" field.
func DestinyIDNEQ(v int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldNEQ(FieldDestinyID, v))
}

// DestinyIDIn applies the In predicate on the "destiny_id" field.
func DestinyIDIn(vs ...int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldIn(FieldDestinyID, vs...))
}

// DestinyIDNotIn applies the NotIn predicate on the "destiny_id" field.
func DestinyIDNotIn(vs ...int) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldNotIn(FieldDestinyID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasDestinies applies the HasEdge predicate on the "destinies" edge.
func HasDestinies() predicate.DestinyPictures {
	return predicate.DestinyPictures(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DestiniesTable, DestiniesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDestiniesWith applies the HasEdge predicate on the "destinies" edge with a given conditions (other predicates).
func HasDestiniesWith(preds ...predicate.Destinies) predicate.DestinyPictures {
	return predicate.DestinyPictures(func(s *sql.Selector) {
		step := newDestiniesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DestinyPictures) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DestinyPictures) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DestinyPictures) predicate.DestinyPictures {
	return predicate.DestinyPictures(sql.NotPredicates(p))
}
