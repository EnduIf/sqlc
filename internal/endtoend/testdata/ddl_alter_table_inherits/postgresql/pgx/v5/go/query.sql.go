// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package querytest

import (
	"context"
)

const getAllOrganisations = `-- name: GetAllOrganisations :many
SELECT party_id, name, joined_on, rank, legal_name, region FROM organisation
`

func (q *Queries) GetAllOrganisations(ctx context.Context) ([]Organisation, error) {
	rows, err := q.db.Query(ctx, getAllOrganisations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Organisation
	for rows.Next() {
		var i Organisation
		if err := rows.Scan(
			&i.PartyID,
			&i.Name,
			&i.JoinedOn,
			&i.Rank,
			&i.LegalName,
			&i.Region,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllParties = `-- name: GetAllParties :many
SELECT party_id, name, joined_on, rank, region FROM party
`

func (q *Queries) GetAllParties(ctx context.Context) ([]Party, error) {
	rows, err := q.db.Query(ctx, getAllParties)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Party
	for rows.Next() {
		var i Party
		if err := rows.Scan(
			&i.PartyID,
			&i.Name,
			&i.JoinedOn,
			&i.Rank,
			&i.Region,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllPeople = `-- name: GetAllPeople :many
SELECT party_id, name, joined_on, rank, first_name, last_name, region FROM person
`

func (q *Queries) GetAllPeople(ctx context.Context) ([]Person, error) {
	rows, err := q.db.Query(ctx, getAllPeople)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Person
	for rows.Next() {
		var i Person
		if err := rows.Scan(
			&i.PartyID,
			&i.Name,
			&i.JoinedOn,
			&i.Rank,
			&i.FirstName,
			&i.LastName,
			&i.Region,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrganizationsByRegion = `-- name: GetOrganizationsByRegion :many
SELECT party_id, name, joined_on, rank, legal_name, region
FROM organisation
WHERE
	region = 'us' AND
	rank = 'ensign'
`

func (q *Queries) GetOrganizationsByRegion(ctx context.Context) ([]Organisation, error) {
	rows, err := q.db.Query(ctx, getOrganizationsByRegion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Organisation
	for rows.Next() {
		var i Organisation
		if err := rows.Scan(
			&i.PartyID,
			&i.Name,
			&i.JoinedOn,
			&i.Rank,
			&i.LegalName,
			&i.Region,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
