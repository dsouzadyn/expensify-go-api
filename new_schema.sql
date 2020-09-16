CREATE USER IF NOT EXISTS 'expenseuser'@'localhost' IDENTIFIED BY 'r00tp4ssw0rd';
CREATE DATABASE IF NOT EXISTS expense;
GRANT CREATE ON expense.* TO 'expenseuser'@'localhost';
GRANT DELETE ON expense.* TO 'expenseuser'@'localhost';
GRANT INSERT ON expense.* TO 'expenseuser'@'localhost';
GRANT SELECT ON expense.* TO 'expenseuser'@'localhost';
GRANT UPDATE ON expense.* TO 'expenseuser'@'localhost';