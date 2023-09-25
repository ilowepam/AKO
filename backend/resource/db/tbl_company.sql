DROP TABLE IF EXISTS tbl_company;

CREATE TABLE tbl_company (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

INSERT INTO tbl_company (name) VALUES ('Company A');
INSERT INTO tbl_company (name) VALUES ('Company B');
INSERT INTO tbl_company (name) VALUES ('Company C');
INSERT INTO tbl_company (name) VALUES ('Company D');
INSERT INTO tbl_company (name) VALUES ('Company E');