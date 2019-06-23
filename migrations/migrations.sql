/*Just copying over the data from our struct models*/

CREATE TABLE IF NOT EXISTS courses (
    id INT PRIMARY KEY AUTO_INCREMENT,
    department VARCHAR(255) NOT NULL,
    course_number INT NOT NULL,
    subsection INT NOT NULL, 
    title VARCHAR(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS reviews (
    id INT PRIMARY KEY AUTO_INCREMENT,
    rating INT NOT NULL,
    
	/*Does it make sense to keep these as optional inputs?*/
    difficulty INT NOT NULL,
    interest INT NOT NULL,
    
    courseId INT NOT NULL,
    review_date DATE NOT NULL,
    is_anonymous BOOLEAN NOT NULL,
    
    /*Would this be null in the event of a blank review? Do we want to allow that?*/
    review_text TEXT, 
    professor_name VARCHAR(255),
    helpfulCount INT NOT NULL DEFAULT 0,
    notHelpfulCount INT NOT NULL DEFAULT 0
);

/*
DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS courses;
*/