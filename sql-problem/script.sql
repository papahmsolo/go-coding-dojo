SELECT reports_to, COUNT(id), ROUND(AVG(age)) as COUNT 
FROM employees
WHERE reports_to IS NOT NULL
GROUP BY reports_to;