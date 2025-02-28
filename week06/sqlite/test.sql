PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE dogs(id integer primary key, name text);
INSERT INTO dogs VALUES(1,'Rover');
INSERT INTO dogs VALUES(2,'Bailey');
CREATE TABLE words(id integer primary key, name text, count integer);
INSERT INTO words VALUES(1,'Romeo',7);
INSERT INTO words VALUES(2,'Juliet',8);
COMMIT;

