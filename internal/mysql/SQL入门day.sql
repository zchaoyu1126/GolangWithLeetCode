-- ****************day1****************
-- 595
SELECT name, population, area
From World
WHERE area >= 3000000 OR population >= 25000000;

-- 1757
SELECT product_id
FROM Products
WHERE low_fats = 'Y' AND recyclable = 'Y';

-- 584
-- 这题考察 referee_id != 2 得到的结果不包含NULL
-- 可以使用这种方式，但效率很低 <=> 中包含键值为NULL的记录
-- SELECT name
-- FROM customer
-- WHERE NOT referee_id <=> 2;
SELECT name
FROM customer
WHERE id NOT IN (
    SELECT id
    FROM customer
    WHERE referee_id = 2
);

-- 183
SELECT name AS Customers
FROM Customers 
WHERE id NOT IN (
    SELECT CustomerId 
    FROM Orders 
)

-- ****************day2****************
-- 1873
SELECT employee_id, IF(employee_id % 2 != 0 AND LEFT(name, 1) != 'M', salary, 0) AS bonus
FROM employees
ORDER BY employee_id

-- 627
UPDATE salary SET sex=IF(sex='m', 'f', 'm')
-- 将'm'+'f'相加再减去当前的，即可完成交换
UPDATE salary SET sex=char(211-ascii(sex))
UPDATE salary 
SET sex = CASE sex
        WHEN 'm' THEN 'f'
        ELSE 'm'
    END;

-- 196
-- 考察自连接
DELETE p1
FROM person p1, person p2
WHERE p1.email = p2.email AND p1.id > p2.id

-- ****************day3****************
-- 1667
-- 考察拼接字段的使用
SELECT user_id,
CONCAT(UPPER(LEFT(name, 1)), LOWER(SUBSTR(name,2))) AS name
FROM users
ORDER BY user_id

-- 1484
-- 默认升序
SELECT sell_date,
COUNT(DISTINCT product) AS num_sold,
GROUP_CONCAT(DISTINCT product ORDER BY product asc) AS products
FROM activities
GROUP BY sell_date
ORDER BY sell_date

-- 1527
-- 模糊匹配
SELECT patient_id, patient_name, conditions
FROM patients
WHERE conditions LIKE 'DIAB1%' OR conditions LIKE '% DIAB1%'

-- ****************day4****************
-- 1965
-- 考察union拼接
-- FROM 可以将别的查询结果作为一个临时表，这里为t
SELECT * FROM (
SELECT employee_id FROM employees RIGHT OUTER JOIN salaries USING(employee_id) WHERE employees.name is NULL
UNION
SELECT employee_id FROM employees LEFT OUTER JOIN salaries USING(employee_id) WHERE salaries.salary IS NULL) t
ORDER BY t.employee_id

-- 1795
-- 考察列转行/行转列
-- 列转行代码
SELECT product_id, 'store1' store, store1 price FROM products WHERE store1 IS NOT NULL
UNION
SELECT product_id, 'store2' store, store2 price FROM products WHERE store2 IS NOT NULL
UNION
SELECT product_id, 'store3' store, store3 price FROM products WHERE store3 IS NOT NULL

-- 行转列代码
-- 不是很懂这个SUM的意义
SELECT product_id,
SUM(IF(store = 'store1', price, NULL)) 'store1',  
SUM(IF(store = 'store2', price, NULL)) 'store2',
SUM(IF(store = 'store3', price, NULL)) 'store3' 
FROM products 
GROUP BY product_id ;

-- 608
SELECT id, 'Root' type FROM tree WHERE p_id IS NULL 
UNION
-- 这句会出错
-- SELECT id, 'Leaf' type FROM tree WHERE id NOT IN(SELECT p_id from tree)
SELECT id, 'Leaf' type FROM tree WHERE id NOT IN(SELECT p_id from tree WHERE p_id IS NOT NULL) AND p_id IS NOT NULL
UNION
SELECT id, 'Inner' type FROM tree WHERE id IN(SELECT DISTINCT p_id from tree WHERE p_id IS NOT NULL)  AND p_id IS NOT NULL

-- 176
-- 查找数据库中的第K大元素
SELECT IFNULL(
    (SELECT DISTINCT salary
    FROM employee
    ORDER BY salary DESC
    LIMIT 1, 1
    ), NULL 
) AS SecondHighestSalary 

-- ****************day5****************
-- 175
SELECT 
    firstName, lastName, city, state
FROM
    Person LEFT OUTER JOIN Address
USING(personId)

-- 1581
SELECT customer_id,
COUNT(*) AS count_no_trans
FROM visits LEFT OUTER JOIN transactions
USING(visit_id)
WHERE transaction_id IS NULL
GROUP BY customer_id

-- 1148
SELECT DISTINCT author_id AS id
FROM views 
WHERE author_id = viewer_id
ORDER BY author_id

-- day6
-- 197
-- DATEDIFF函数的使用
-- w1.date - w2.date = 1
-- 会在月初和月末的时候出错，不会自动跨月
SELECT 
    id
FROM 
    weather w1, weather w2
WHERE 
    DATEDIFF(w1.date, w2.date) = 1
AND 
    w1.temperature > w2.temperature

-- 607
SELECT name
FROM salesperson
WHERE sales_id NOT IN (
    SELECT sales_id
    FROM 
        (SELECT com_id
        FROM company
        WHERE name = 'RED') t, orders 
    WHERE t.com_id = orders.com_id
)

-- day7
-- 1141
SELECT activity_date AS day, COUNT(DISTINCT user_id) AS active_users 
FROM activity
WHERE DATEDIFF('2019-07-27', activity_date) < 30 AND DATEDIFF('2019-07-27', activity_date) >= 0
GROUP BY activity_date 

SELECT activity_date AS day, COUNT(DISTINCT user_id) AS active_users 
FROM activity
WHERE activity_date BETWEEN  '2019-06-28' AND '2019-07-27'
GROUP BY activity_date 

-- 1693
SELECT date_id, make_name, COUNT(DISTINCT lead_id) AS unique_leads, COUNT(DISTINCT partner_id) unique_partners
FROM dailysales
GROUP BY date_id, make_name

-- 1729
SELECT user_id, COUNT(DISTINCT follower_id) AS followers_count
FROM followers
GROUP BY user_id

-- day8
-- 586
Select customer_number 
FROM orders
GROUP BY customer_number 
ORDER BY COUNT(*) DESC
LIMIT 1

-- 511
SELECT player_id, MIN(event_date) AS first_login
FROM activity
GROUP BY player_id

-- 1890
SELECT user_id, MAX(time_stamp) AS last_stamp
FROM logins
WHERE time_stamp > '2019-12-31 23:59:59' AND time_stamp < '2020-12-31 23:59:59'
GROUP BY user_id

-- 1741
SELECT event_day AS day, emp_id, SUM(out_time-in_time) AS total_time
FROM employees
GROUP BY emp_id, event_day

-- day9
-- 1393
SELECT stock_name, sell_price-buy_price AS capital_gain_loss
FROM 
    (SELECT stock_name, SUM(price) AS buy_price
    FROM stocks
    WHERE operation = 'Sell'
    GROUP BY stock_name) t1, 
    (SELECT stock_name, operation, SUM(price) AS sell_price
    FROM stocks
    WHERE operation = 'Buy'
    GROUP BY stock_name) t2
WHERE
    t1.stock_name = t2.stock_name

-- 1407
SELECT name, IFNULL(total_distance, 0) AS travelled_distance
FROM
    users LEFT OUTER JOIN
    (SELECT users.id AS id, SUM(distance) AS total_distance
    FROM users, rides 
    WHERE users.id = rides.user_id
    GROUP BY users.name) t
USING(id)
ORDER BY travelled_distance DESC, name

-- 1158
SELECT user_id, join_date, IFNULL(num, 0) orders_in_2019
FROM
    users LEFT OUTER JOIN
    (SELECT buyerid AS user_id , COUNT(*) AS num
    FROM orders
    WHERE order_date BETWEEN '2019-01-01' AND '2019-12-31'
    GROUP BY buyer_id) t
    USING(user_id)

-- day10
-- 182
SELECT email
FROM person
GROUP BY email
HAVING COUNT(*) >= 2

-- 1050
SELECT actor_id, director_id
FROM ActorDirector
GROUP BY actor_id, director_id
HAVING COUNT(*) >= 3 

-- 1587
SELECT name, balance
FROM users, 
    (SELECT account, SUM(amount) AS balance
    FROM transactions
    GROUP BY account
    HAVING SUM(amount) >= 10000) t
WHERE users.account = t.account

-- 1084
SELECT product_id, product_name 
FROM product
WHERE product_id IN (
    SELECT product_id
    FROM sales
    GROUP BY product_id
    HAVING MIN(sale_date) >= '2019-01-01' AND  MAX(sale_date) <= '2019-03-31')