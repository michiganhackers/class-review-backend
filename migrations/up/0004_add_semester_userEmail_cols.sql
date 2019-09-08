/* Add semester and userEmail fields to reviews table  */

ALTER TABLE reviews ADD (semester char(6), userEmail varchar(255));