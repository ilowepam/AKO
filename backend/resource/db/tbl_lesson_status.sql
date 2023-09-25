DROP TABLE IF EXISTS tbl_lesson_status;

CREATE TABLE tbl_lesson_status (
    id SERIAL PRIMARY KEY,
    AdaId INT,
    LessonId INT,
    Status BOOLEAN
);

INSERT INTO tbl_lesson_status (AdaId, LessonId, Status) VALUES (1, 1, true);
INSERT INTO tbl_lesson_status (AdaId, LessonId, Status) VALUES (3, 2, true);
INSERT INTO tbl_lesson_status (AdaId, LessonId, Status) VALUES (4, 2, true);
