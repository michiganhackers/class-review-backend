/* Up migration*/

CREATE TABLE IF NOT EXISTS admins (
    admin_uniqname VARCHAR(255) PRIMARY KEY
);

/* Down migration -- comment out the up migration and uncomment this to run */
/*
DROP TABLE IF EXISTS admins;
*/