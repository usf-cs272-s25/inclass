PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE dogs(id integer primary key, name text);
INSERT INTO dogs VALUES(1,'Spot');
INSERT INTO dogs VALUES(2,'Scooby');
CREATE TABLE vetbills(id integer primary key, pet_id integer,
foreign key(pet_id) references pets(id));
COMMIT;

