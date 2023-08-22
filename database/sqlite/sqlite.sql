-- SQLite

.echo on
.mode column
.headers on
.nullvalue NULL

CREATE TABLE device (
    id integer PRIMARY KEY,
    code text NOT NULL,
    name text NOT NULL,
    categorie text,
    serial text,
    commissioning integer,
    model text,
    port text,
    aggregat text
);

CREATE TABLE engineer (
    id integer PRIMARY KEY,
    name TEXT
);

SELECT categorie, count(*), 'categorie' FROM device GROUP BY categorie;

SELECT * FROM device WHERE categorie LIKE '';

CREATE TABLE categorie (
    id integer PRIMARY KEY AUTOINCREMENT,
    name TEXT
);

INSERT INTO categorie (name) VALUES('jklkjlk');

INSERT INTO categorie (name)
    SELECT categorie FROM device GROUP BY categorie;

SELECT
        device.id,
        device.code,
        device.name,
        device.categorie,
        categorie.id
    FROM device, categorie
    WHERE device.categorie = categorie.name;