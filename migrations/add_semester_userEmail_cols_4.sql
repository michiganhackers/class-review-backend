/* Add semester and userEmail fields to reviews table  */
/* Up migration*/


ALTER TABLE reviews ADD (semester char(6), userEmail varchar(255));

/* Down migration -- comment out the up migration and uncomment this to run */

/*
*/