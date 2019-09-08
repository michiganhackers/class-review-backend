/* Remove is_anonymous from reviews and make userEmail not null */

ALTER TABLE reviews DROP is_anonymous;
ALTER TABLE reviews MODIFY userEmail varchar(255) NOT NULL;