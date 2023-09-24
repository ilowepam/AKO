DROP TABLE IF EXISTS tbl_Company;

CREATE TABLE tbl_Company (
                             id INT AUTO_INCREMENT PRIMARY KEY,
                             name VARCHAR(255) NOT NULL,
                             UNIQUE (id)
);


INSERT INTO tbl_Company (name) VALUES ('Company A');
INSERT INTO tbl_Company (name) VALUES ('Company B');
INSERT INTO tbl_Company (name) VALUES ('Company C');
INSERT INTO tbl_Company (name) VALUES ('Company D');
INSERT INTO tbl_Company (name) VALUES ('Company E');
