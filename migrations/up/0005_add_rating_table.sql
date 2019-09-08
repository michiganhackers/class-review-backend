/* Add ratings table and reviewId foreign key contstraint */

ALTER TABLE reviews DROP helpfulCount, DROP notHelpfulCount;
CREATE TABLE IF NOT EXISTS ratings(
     userEmail VARCHAR(255) NOT NULL,
     reviewId INT NOT NULL,
     helpful BOOLEAN NOT NULL,
     PRIMARY KEY (userEmail, reviewId)
);

ALTER TABLE ratings ADD CONSTRAINT fk_reviewId FOREIGN KEY (reviewId) REFERENCES reviews(id) ON DELETE CASCADE;