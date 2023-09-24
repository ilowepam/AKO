DROP TABLE IF EXISTS tbl_Ada;

CREATE TABLE tbl_Ada (
                         id INT AUTO_INCREMENT PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         rank VARCHAR(25) NOT NULL,
                         companyId INT NOT NULL,
                         FOREIGN KEY (companyId) REFERENCES tbl_Company(id)
);

INSERT INTO tbl_Ada (name, rank, companyId) VALUES ('John Doe', 'Captain', 1);
INSERT INTO tbl_Ada (name, rank, companyId) VALUES ('Jane Smith', 'Lieutenant', 2);
INSERT INTO tbl_Ada (name, rank, companyId) VALUES ('Mike Johnson', 'Sergeant', 3);
INSERT INTO tbl_Ada (name, rank, companyId) VALUES ('Sam Brown', 'Corporal', 4);
INSERT INTO tbl_Ada (name, rank, companyId) VALUES ('Alex Turner', 'Major', 5);
