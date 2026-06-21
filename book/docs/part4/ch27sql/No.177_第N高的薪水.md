## 177 第N高的薪水-中等

题目：

表：`Employee`

| Column Name | Type |
|-------------|------|
| id          | int  |
| salary      | int  |

`id` 是该表的主键（列中的值互不相同）。该表的每一行都包含有关员工工资的信息。

编写一个解决方案查询 `Employee` 表中第 `n` 高的**不同**工资。如果少于 `n` 个不同工资，查询结果应该为 `null`。



**解题思路**

核心难点有三个：① **去重**（`DISTINCT`，相同工资算同一名次）；② 当不足 `n` 个不同工资时要返回 **null**；③ MySQL 的 `LIMIT` 不能直接写 `N-1` 表达式，需要用变量中转。

- LIMIT + OFFSET。对不同工资降序排列，取第 `n` 个（即跳过 `n-1` 个）。最经典、面试常考的写法，详见解法1。
- DENSE_RANK 窗口函数。用 `DENSE_RANK()` 给不同工资编号（相同的并列、序号不跳号），取 `rnk = n`。写法直观，详见解法2。

为什么外层再套一层 `SELECT`：函数 `RETURNS INT` 要返回单值；内层 `LIMIT 1 OFFSET M` 取不到行时结果为 `NULL`，外层 `SELECT` 把它变成标量返回，正好满足「不足 n 个返回 null」。

```sql
-- date 2026/06/21
-- 解法1
-- LIMIT + OFFSET（变量中转 N-1）
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  DECLARE M INT;
  SET M = N - 1;          -- MySQL 的 LIMIT/OFFSET 不能直接用表达式，必须用变量
  RETURN (
      SELECT (
          SELECT DISTINCT salary
          FROM Employee
          ORDER BY salary DESC
          LIMIT 1 OFFSET M
      )
  );
END
```

```sql
-- date 2026/06/21
-- 解法2
-- DENSE_RANK() 窗口函数
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      SELECT DISTINCT salary
      FROM (
          SELECT salary,
                 DENSE_RANK() OVER (ORDER BY salary DESC) AS rnk
          FROM Employee
      ) t
      WHERE rnk = N
  );
END
```

补充一个**不用 function 的普通查询**写法（参数 `n` 直接代入）：

```sql
-- 给定 n=2 时，查询写法（不必定义函数）
SELECT (
    SELECT DISTINCT salary
    FROM Employee
    ORDER BY salary DESC
    LIMIT 1 OFFSET 1          -- OFFSET = n-1
) AS getNthHighestSalary;
```

以 `Employee` 工资为 `100, 200, 300` 为例：

| n | 结果 | 说明 |
|---|------|------|
| 2 | `200` | 跳过第1高(300)，取第2高 |
| 1 | `300` | 最高 |
| 4 | `null` | 只有3个不同工资，不足4个 → null |
