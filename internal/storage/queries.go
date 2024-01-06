package storage

// allPagedQuery needs the id subquery to handle pagination
const allPagedQuery = `
SELECT
    s.id,
    s.name,
    s.description,
    v.id,
    v.version
FROM
    versions v
    LEFT JOIN services s ON v.service_id = s.id
WHERE
    s.id IN (
        SELECT
            subid.id
        FROM
            services subid
        LIMIT $1 OFFSET $2)
ORDER BY s.sequence
`

const singleServiceQuery = `SELECT
    s.id,
    s.name,
    s.description,
    v.id,
    v.version
FROM
    versions v
    LEFT JOIN services s ON v.service_id = s.id
WHERE
	s.id = $1`

const searchByNameQuery = `
SELECT
    s.id,
    s.name,
    s.description,
		v.id,
		v.version
FROM
    versions v
    LEFT JOIN services s ON v.service_id = s.id
WHERE
    s.id IN (
        SELECT
            sub.id
        FROM
            services sub
        WHERE
            $1 % sub.name
        ORDER BY
            $2 <-> sub.name DESC)
`

const searchAllQuery = `SELECT
    s.id,
    s.name,
    s.description,
    v.id,
    v.version
FROM
    versions v
    LEFT JOIN services s ON v.service_id = s.id
WHERE
    $1 % s.name
UNION
SELECT
    s.id,
    s.name,
    s.description,
    v.id,
    v.version
FROM
    versions v
    LEFT JOIN services s ON v.service_id = s.id
WHERE
    s.id IN (
        SELECT
            sub.id
        FROM
            services sub
        WHERE
            tsvector (sub.description) @@ tsquery ($2)
        ORDER BY
            $3 <-> s.name DESC)`
