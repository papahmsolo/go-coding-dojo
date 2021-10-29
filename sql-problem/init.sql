CREATE TABLE IF NOT EXISTS employees (
    id INT PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    reports_to text,
    position text NOT NULL,
    age int NOT NULL
);

INSERT INTO employees
(id, first_name, last_name, reports_to, position, age)
VALUES 
(1, 'Daniel', 'Smith', 'Bob Boss', 'Engineer', 25),
(2, 'Mike', 'White', 'Bob Boss', 'Contractor', 22),
(3, 'Jenny', 'Richards', null, 'CEO', 45),
(4, 'Robert', 'Black', 'Daniel Smith', 'Sales', 22),
(5, 'Noah', 'Fritz', 'Jenny Richards', 'Assistant', 30),
(6, 'David', 'S', 'Jenny Richards', 'Director', 32),
(7, 'Ashley', 'Wells', 'David S', 'Assistant', 25),
(8, 'Ashley', 'Johnson', null,'Intern', 25)