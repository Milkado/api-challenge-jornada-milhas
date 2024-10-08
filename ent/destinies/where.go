// Code generated by ent, DO NOT EDIT.

package destinies

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Destinies {
	return predicate.Destinies(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Destinies {
	return predicate.Destinies(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Destinies {
	return predicate.Destinies(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Destinies {
	return predicate.Destinies(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Destinies {
	return predicate.Destinies(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Destinies {
	return predicate.Destinies(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Destinies {
	return predicate.Destinies(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldName, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldPrice, v))
}

// Meta applies equality check predicate on the "meta" field. It's identical to MetaEQ.
func Meta(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldMeta, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldDescription, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Destinies {
	return predicate.Destinies(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Destinies {
	return predicate.Destinies(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldContainsFold(FieldName, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Destinies {
	return predicate.Destinies(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Destinies {
	return predicate.Destinies(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Destinies {
	return predicate.Destinies(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Destinies {
	return predicate.Destinies(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Destinies {
	return predicate.Destinies(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Destinies {
	return predicate.Destinies(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Destinies {
	return predicate.Destinies(sql.FieldLTE(FieldPrice, v))
}

// MetaEQ applies the EQ predicate on the "meta" field.
func MetaEQ(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldMeta, v))
}

// MetaNEQ applies the NEQ predicate on the "meta" field.
func MetaNEQ(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldNEQ(FieldMeta, v))
}

// MetaIn applies the In predicate on the "meta" field.
func MetaIn(vs ...string) predicate.Destinies {
	return predicate.Destinies(sql.FieldIn(FieldMeta, vs...))
}

// MetaNotIn applies the NotIn predicate on the "meta" field.
func MetaNotIn(vs ...string) predicate.Destinies {
	return predicate.Destinies(sql.FieldNotIn(FieldMeta, vs...))
}

// MetaGT applies the GT predicate on the "meta" field.
func MetaGT(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldGT(FieldMeta, v))
}

// MetaGTE applies the GTE predicate on the "meta" field.
func MetaGTE(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldGTE(FieldMeta, v))
}

// MetaLT applies the LT predicate on the "meta" field.
func MetaLT(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldLT(FieldMeta, v))
}

// MetaLTE applies the LTE predicate on the "meta" field.
func MetaLTE(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldLTE(FieldMeta, v))
}

// MetaContains applies the Contains predicate on the "meta" field.
func MetaContains(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldContains(FieldMeta, v))
}

// MetaHasPrefix applies the HasPrefix predicate on the "meta" field.
func MetaHasPrefix(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldHasPrefix(FieldMeta, v))
}

// MetaHasSuffix applies the HasSuffix predicate on the "meta" field.
func MetaHasSuffix(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldHasSuffix(FieldMeta, v))
}

// MetaEqualFold applies the EqualFold predicate on the "meta" field.
func MetaEqualFold(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldEqualFold(FieldMeta, v))
}

// MetaContainsFold applies the ContainsFold predicate on the "meta" field.
func MetaContainsFold(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldContainsFold(FieldMeta, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Destinies {
	return predicate.Destinies(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Destinies {
	return predicate.Destinies(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Destinies {
	return predicate.Destinies(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Destinies {
	return predicate.Destinies(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Destinies {
	return predicate.Destinies(sql.FieldContainsFold(FieldDescription, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Destinies {
	return predicate.Destinies(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasTestimonies applies the HasEdge predicate on the "testimonies" edge.
func HasTestimonies() predicate.Destinies {
	return predicate.Destinies(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestimoniesTable, TestimoniesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestimoniesWith applies the HasEdge predicate on the "testimonies" edge with a given conditions (other predicates).
func HasTestimoniesWith(preds ...predicate.Testimonies) predicate.Destinies {
	return predicate.Destinies(func(s *sql.Selector) {
		step := newTestimoniesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDestinyPictures applies the HasEdge predicate on the "destiny_pictures" edge.
func HasDestinyPictures() predicate.Destinies {
	return predicate.Destinies(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DestinyPicturesTable, DestinyPicturesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDestinyPicturesWith applies the HasEdge predicate on the "destiny_pictures" edge with a given conditions (other predicates).
func HasDestinyPicturesWith(preds ...predicate.DestinyPictures) predicate.Destinies {
	return predicate.Destinies(func(s *sql.Selector) {
		step := newDestinyPicturesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Destinies) predicate.Destinies {
	return predicate.Destinies(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Destinies) predicate.Destinies {
	return predicate.Destinies(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Destinies) predicate.Destinies {
	return predicate.Destinies(sql.NotPredicates(p))
}
