/* Add is_anonymous to reviews and make userEmail null */

ALTER TABLE reviews ADD is_anonymous BOOLEAN NOT NULL;
ALTER TABLE reviews MODIFY userEmail varchar(255);