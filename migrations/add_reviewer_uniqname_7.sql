/* Up migration */
ALTER TABLE reviews ADD uniqname VARCHAR(255);

/* Down migration */
/* 
ALTER TABLE reviews DROP uniqname; 
*/