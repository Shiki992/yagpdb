// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testReputationConfigs(t *testing.T) {
	t.Parallel()

	query := ReputationConfigs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testReputationConfigsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ReputationConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReputationConfigsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ReputationConfigs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ReputationConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReputationConfigsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ReputationConfigSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ReputationConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReputationConfigsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ReputationConfigExists(ctx, tx, o.GuildID)
	if err != nil {
		t.Errorf("Unable to check if ReputationConfig exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ReputationConfigExists to return true, but got false.")
	}
}

func testReputationConfigsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	reputationConfigFound, err := FindReputationConfig(ctx, tx, o.GuildID)
	if err != nil {
		t.Error(err)
	}

	if reputationConfigFound == nil {
		t.Error("want a record, got nil")
	}
}

func testReputationConfigsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ReputationConfigs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testReputationConfigsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ReputationConfigs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testReputationConfigsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reputationConfigOne := &ReputationConfig{}
	reputationConfigTwo := &ReputationConfig{}
	if err = randomize.Struct(seed, reputationConfigOne, reputationConfigDBTypes, false, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}
	if err = randomize.Struct(seed, reputationConfigTwo, reputationConfigDBTypes, false, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = reputationConfigOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = reputationConfigTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ReputationConfigs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testReputationConfigsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	reputationConfigOne := &ReputationConfig{}
	reputationConfigTwo := &ReputationConfig{}
	if err = randomize.Struct(seed, reputationConfigOne, reputationConfigDBTypes, false, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}
	if err = randomize.Struct(seed, reputationConfigTwo, reputationConfigDBTypes, false, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = reputationConfigOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = reputationConfigTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ReputationConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testReputationConfigsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ReputationConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReputationConfigsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(reputationConfigColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ReputationConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReputationConfigsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testReputationConfigsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ReputationConfigSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testReputationConfigsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ReputationConfigs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	reputationConfigDBTypes = map[string]string{`GuildID`: `bigint`, `PointsName`: `character varying`, `Enabled`: `boolean`, `Cooldown`: `integer`, `MaxGiveAmount`: `bigint`, `RequiredGiveRole`: `character varying`, `RequiredReceiveRole`: `character varying`, `BlacklistedGiveRole`: `character varying`, `BlacklistedReceiveRole`: `character varying`, `AdminRole`: `character varying`, `DisableThanksDetection`: `boolean`, `MaxRemoveAmount`: `bigint`, `AdminRoles`: `ARRAYbigint`, `RequiredGiveRoles`: `ARRAYbigint`, `RequiredReceiveRoles`: `ARRAYbigint`, `BlacklistedGiveRoles`: `ARRAYbigint`, `BlacklistedReceiveRoles`: `ARRAYbigint`}
	_                       = bytes.MinRead
)

func testReputationConfigsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(reputationConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(reputationConfigColumns) == len(reputationConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ReputationConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testReputationConfigsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(reputationConfigColumns) == len(reputationConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ReputationConfig{}
	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ReputationConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, reputationConfigDBTypes, true, reputationConfigPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(reputationConfigColumns, reputationConfigPrimaryKeyColumns) {
		fields = reputationConfigColumns
	} else {
		fields = strmangle.SetComplement(
			reputationConfigColumns,
			reputationConfigPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ReputationConfigSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testReputationConfigsUpsert(t *testing.T) {
	t.Parallel()

	if len(reputationConfigColumns) == len(reputationConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ReputationConfig{}
	if err = randomize.Struct(seed, &o, reputationConfigDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ReputationConfig: %s", err)
	}

	count, err := ReputationConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, reputationConfigDBTypes, false, reputationConfigPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ReputationConfig struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ReputationConfig: %s", err)
	}

	count, err = ReputationConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
