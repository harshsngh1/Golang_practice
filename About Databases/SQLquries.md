## Most Asked MySql queries

1. Find Employees and Their Managers

![Question1](/About%20Databases/table%20images/ques1.png)  

```
SELECT e.name AS Employee, m.name AS Manager
FROM employees e
LEFT JOIN employees m ON e.manager_id = m.id
```
JOIN Used: LEFT JOIN

2. List All Departments and Employees

![Question2](/About%20Databases/table%20images/ques2.png)

```
SELECT d.department_name, e.name AS Employee
FROM departments d
LEFT JOIN employees e ON d.id = e.department_id;
```
JOIN Used: LEFT JOIN

3. Find Employees Without Managers

![Question3](/About%20Databases/table%20images/ques3.png)

```
SELECT name
FROM employees
WHERE manager_id IS NULL;
```

4. Find Employees with More Than One Manager

![Question4](/About%20Databases/table%20images/ques4.png)

```
SELECT em.employee_id, COUNT(DISTINCT em.manager_id) AS manager_count
FROM employee_managers em
GROUP BY em.employee_id
HAVING COUNT(DISTINCT em.manager_id) > 1;
```
JOIN Used: No JOIN (aggregated query)

5. Find Departments with No Employees

![Question5](/About%20Databases/table%20images/ques5.png)

```
SELECT d.department_name
FROM departments d
LEFT JOIN employees e ON d.id = e.department_id
WHERE e.id IS NULL;
```
JOIN Used: LEFT JOIN

6. Find the Second Highest Salary

![Question6](/About%20Databases/table%20images/ques6.png)

```
SELECT MAX(salary) AS second_highest_salary
FROM employees
WHERE salary < (SELECT MAX(salary) FROM employees);
```
JOIN Used: No JOIN (subquery)

7. Find Duplicate Entries

![Question7](/About%20Databases/table%20images/ques7.png)

```
SELECT e1.name, e1.department_id
FROM employees e1
INNER JOIN employees e2
ON e1.name = e2.name AND e1.department_id = e2.department_id
WHERE e1.id <> e2.id;
```
JOIN Used: INNER JOIN

8. Find Employees with the Same Salary as Their Manager

![Question8](/About%20Databases/table%20images/ques8.png)

```
SELECT e.name AS Employee, m.name AS Manager, e.salary
FROM employees e
INNER JOIN employees m ON e.manager_id = m.id
WHERE e.salary = m.salary;
```
JOIN Used: INNER JOIN

9. List Departments and Count of Employees

![Question9](/About%20Databases/table%20images/ques9.png)

```
SELECT d.department_name, COUNT(e.id) AS employee_count
FROM departments d
LEFT JOIN employees e ON d.id = e.department_id
GROUP BY d.department_name;
```
JOIN Used: LEFT JOIN

10. Find Employees Who Report to Themselves

![Question10](/About%20Databases/table%20images/ques10.png)

```
SELECT name
FROM employees
WHERE id = manager_id;
```
JOIN Used: No JOIN (simple query)

