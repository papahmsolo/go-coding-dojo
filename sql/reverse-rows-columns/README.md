# go-coding-dojo

## The scope of the task is to revert rows to columns in SQL aggregation:
_____
FROM:
category | count 
--- | --- 
BEER | 2
FRUIT | 1
MILK | 5
PAPER | 2
TEA | 4
___
TO:
BEER | MILK |PAPER | TEA | TOTAL 
--- | --- | --- | --- |--- 
2 | 1 | 1 | 1 | 5 | 1
0 | 2 | 1 | 2 | 5 | 2
0 | 2 | 0 | 0 | 2 | 3


## Task with *. Create a view that represents employees amount of products in each category, with total:

id | name |email | beer | milk | paper | tea | total
--- | --- | --- | --- |--- |--- |--- |--- 
1 | George B | gb@email.com | 2 | 1 | 1 | 1 | 5
2 | Johny B | john@email.com | 0 | 2 | 1 | 2 | 5
3 | Oleg Pupkin | obleg@email.com | 0 | 0 | 0 | 2 | 2
4 | Bohdan | b@email.com | 0 | 0 | 0 | 0 | 0

~~~~
* Please note that final solution should contain all employees 