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

SELECT * FROM device;

CREATE TABLE engineer (
    
)