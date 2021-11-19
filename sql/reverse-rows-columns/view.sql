-- view vith categories data
CREATE OR REPLACE VIEW r AS
SELECT
	SUM(BEER) as BEER,
	SUM(MILK) as MILK,
	SUM(PAPER) as PAPER,
	SUM(TEA) as TEA,
	SUM(BEER) + SUM(MILK) + SUM(PAPER) + SUM(TEA) AS TOTAL,
	empl_id
FROM (
SELECT 
 IF(STRCMP(category,'BEER'),0,count(*)) as BEER,
 IF(STRCMP(category,'MILK'),0,count(*)) as MILK,
 IF(STRCMP(category,'PAPER'),0,count(*)) as PAPER,
 IF(STRCMP(category,'TEA'),0,count(*)) as TEA,
 empl_id
FROM products
WHERE active = 1
GROUP BY category, empl_id) r
GROUP BY empl_id;

-- final result
CREATE OR REPLACE VIEW employess_with_metadata AS
SELECT e.*, 
IFNULL(r.BEER,0) beer,
IFNULL(r.TEA,0) tea,
IFNULL(r.PAPER,0) paper,
IFNULL(r.MILK,0) milk,
IFNULL(TOTAL,0) total
FROM employees e
LEFT JOIN r ON e.id = r.empl_id;