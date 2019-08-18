/* Remove is_anonymous from reviews and make userEmail not null */
/* Up migration*/

ALTER TABLE reviews DROP is_anonymous;
ALTER TABLE reviews MODIFY userEmail varchar(255) NOT NULL;


/* Down migration -- comment out the up migration and uncomment this to run */

/*
ALTER TABLE reviews ADD is_anonymous BOOLEAN NOT NULL;
ALTER TABLE reviews MODIFY userEmail varchar(255);
*/