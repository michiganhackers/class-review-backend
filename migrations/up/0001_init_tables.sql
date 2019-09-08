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
	difficulty INT NOT NULL,
    interest INT NOT NULL,
    courseId INT NOT NULL,
    review_date DATE NOT NULL,
    is_anonymous BOOLEAN NOT NULL,
	review_text TEXT, 
    professor_uniqname VARCHAR(255),
    helpfulCount INT NOT NULL DEFAULT 0,
    notHelpfulCount INT NOT NULL DEFAULT 0
);