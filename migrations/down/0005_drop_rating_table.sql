/* Drop ratings table and reviewId foreign key contstraint */

ALTER TABLE reviews ADD (helpfulCount INT NOT NULL DEFAULT 0, notHelpfulCount INT NOT NULL DEFAULT 0);
DROP TABLE ratings;