DROP TABLE IF EXISTS tbl_lesson;

CREATE TABLE tbl_lesson (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO tbl_lesson (name) VALUES ('Lesson 1');
INSERT INTO tbl_lesson (name) VALUES ('Lesson 2');
INSERT INTO tbl_lesson (name) VALUES ('Lesson 3');
INSERT INTO tbl_lesson (name) VALUES ('Lesson 4');
INSERT INTO tbl_lesson (name) VALUES ('Lesson 5');

