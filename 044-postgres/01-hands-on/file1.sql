/*Big
Comment
*/

-- psql DBNAME USERNAME
-- psql matthew postgres on iMac
-- psql postgres postgres in laptop
-- psql bookstore bond


-- todds build…
CREATE TABLE employees (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           TEXT    NOT NULL,
   SCORE          INT     DEFAULT 10 NOT NULL,
   SALARY         REAL);

-- my build…
CREATE TABLE employees (
   ID SERIAL PRIMARY KEY  NOT NULL,
   NAME           TEXT    NOT NULL,
   SCORE          INT     DEFAULT 10 NOT NULL,
   SALARY         INT);


-- my insert…
INSERT INTO employees (NAME,SCORE,SALARY) VALUES 
('Daniel', 23, 55000),
('Arin', 25, 65000),
('Juan', 24, 72000),
('Shen', 26, 64000),
('Myke', 27, 58000),
('McLeod', 26, 72000),
('James', 32, 35000);

--todds insert…
INSERT INTO employees (NAME,SCORE,SALARY) VALUES 
('Daniel', 23, 55000.00), 
('Arin', 25, 65000.00), 
('Juan', 24, 72000.00), 
('Shen', 26, 64000.00), 
('Myke', 27, 58000.00), 
('McLeod', 26, 72000.00), 
('James', 32, 35000.00);