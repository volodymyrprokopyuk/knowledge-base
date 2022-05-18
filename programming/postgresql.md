# PostgreSQL

- Query
    - `WITH _ AS` named subquery (CTE = Common Table Expression) is evaluated only once
      and  can be used in other named subqueries
    - Recursive query can refer to its own output

      ```sql
      WITH RECURSIVE t(n) AS (
        -- non-recursive term is evaluated at the beginning -> create result and create
        -- working table
        VALUES (1)
        UNION ALL
        -- recursive term is evalutaed as long as the working table is not empty
        SELECT n + 1
        -- self-reference to the current content of the working table -> append to result
        -- overwrite the working table
        FROM t
        WHERE n < 5) -- eventually returns the empty working table -> stops iteration
      SELECT n FROM t;
      ```

    - `SELECT DISTINCT ON (...)` computed value, column alias `c`, `agg_func(DISTINCT
      ... ORDER BY ...)`
    - `FROM` table alias `t`, `t(c)`
        - `JOIN` restrict both relations
        - `LEFT JOIN`, `RIGHT JOIN` keep one relation and enrich with the other if
          matches, otherwise `NULL`
        - `FULL JOIN` keep both relations
        - `CROSS JOIN`=`t1, t2` Cartesian product
        - Join conditions `ON expression`, `USING (list)`
        - `JOIN LATERAL (SELECT ... WHERE ...)` restrict subquery to the current row in
          `FROM` `ON true`. Run the subquery in a loop for each row in `FROM`. Push down
          the join condition into the subquery
    - `WHERE` no alias, `AND`, `OR`, `NOT`, `IN (...)`, `EXISTS (SELECT 1 ...)`
        - Anti-join `WHERE NOT EXISTS (SELECT 1 ...)`
        - Row comparator/assignment `(a, b) = (c, d)`
    - `GROUP BY` alias
        - `GROUPING SETS ((...), ())`=`GROUP` data separately `BY` each `GROUPING SET`
          and then `UNION` with appropriate `NULL`. Aggregate over more than one group
          at the same time in a single query scan
        - `ROLLUP` all prefixes for hierarchical data analysis
        - `CUBE` all subsets (power set)
    - `HAVING` no alias
    - `WINDOW _ AS (PARTITION BY ... ORDER BY ... ROWS PRECEDING | CURRENT | FOLLOWING
      EXCLUDE)` for each input row a frame of peer rows sharing a common property is
      available. Different `OVER` definitions can be used in the same query even with
      `GROUP BY`
    - `ORDER BY` alias `ASC | DESC`
    - `LIMIT n` never use `OFFSET m` use `FETCH` cursor instead
- Set operations `UNION [ALL]`, `INTERSECT`, `EXCEPT` combine query result sets
- Three-valuded logic `TRUE`, `FALSE`, `NULL` + `=`, `<>`
    - `IS DISTINCT FROM` two-valued logic with `NULL`
- Conditional `CASE _ WHEN _ THEN _ ELSE _ END`
- Aggregate function `count(*) FILTER (WHERE ...)` (conditional aggregation)

```sql
CREATE AGGREGATE <agg> (type, ...) (
  SFUNC = state transition function (previous state, current row arguments) new state
  STYPE = state value data type
  INITCOND = state initial value
  FINALFUNC = transform the final state value into the aggregate output value)
```

- `WITH ... INSERT INTO ... VALUES | SELECT (JOIN ...) ON CONFLICT ... DO NOTHING | DO UPDATE SET ... EXCLUDED RETURNING *`
- `WITH ... UPDATE ... SET ... FROM (JOIN ...) WHERE ... RETURNING *`
- `WITH ... DELETE FROM ... USING (JOIN ...) WHERE ... RETURNING *`
- `TRUNCATE ...`
- `EXPLAIN (ANALYZE, COSTS OFF)`
    - An index is used if only a small subset of data is retrieved using an operator
      supported by the index on a table with a large enough number of rows + up-to-date
      statistics `ANALYZE`
    - Scan types
        - Sequential scan = reads all table from a disk, fast for small tables
        - Index scan = scans index, randomly reads some records from a disk, fast on
          small subset of records from a large table
        - Index only scan = scans index, does not read from disk as all queried data is
          stored in the index
    - Join types
        - Nested loop = scans the innter table for each row in the outer table, fast
          for small tables
        - Merge join = merges sorted large tables, requires initial sort
        - Hash join = scans the outer table for matches in a hash of inner table,
          requires initial hashing, only for equality conditions
- JSONB + JSONPath operators and indexes
    - Not all index types support all operator classes, so create indexes based on the
      type of operators in queries
    - Operator types
        - Existence `?`, `?|`, `?&` only top-level object key or array element matching
        - Containment `@>`, `<@` nested structure + content matching
        - JSONPath `@@`, `@?`
    - GIN (generalized inverted index) serching for individual elements in composite
      structures (`jsonb`)
        - `jsonb_ops` operator class (default) creates independent indexes for each key
          and value (large index size), supports existence (use GIN expression index to
          check existence of nested keys), containment and JSONPath
        - `jsonb_path_ops` operator class creates indexes only for values (smaller index
          size), supports containment and JSONPath
    - BTREE (binary tree) comparison `=`, `<` `<=`, `>`, `>=` of whole objects
    - HASH (hash) equality `=` of whole objects (smaller index size, than BTREE)
- **Normalization**. Reduces data redundacy, improves data consistency, allows to extend
  the data model without changing existing tables (DDL)
    - **Anomalies**
        - **Update anomaly**. Multiple rows has to be updated with the same data
        - **Insertion anomaly**. More than necessary data has to be inserted
        - **Deletion anomaly**. More than necessary data has to be deleted
    - **Normalization forms** (split tables using identity PK, many-to-one FK and
      many-to-many pivot table)
        - **1NF**. Every attribute value must be atomic
        - **2NF**. Every non-candidate key attribute must depend on the whole candidate
          key
        - **3NF**. Transtivie dependencies between attributes must be removed
- **Transactions**. Lower isolaiton, fewer locks, more concurrency, more
  phenomena. Highter isolation, more locks, less concurrency, less phenomena (TCL, DML)
    - **Phenomena**
        - **Dirty read**. A transaction reads data written by a concurrent uncommitted
          transaction
        - **Nonrepeatable read**. A transaction re-reads data and finds that data has
          been updated by another transaction that commited after the initial read
          (`UPDATE` + `COMMIT`)
        - **Phantom read**. A transaction re-reads data and finds that data has been
          inserted or deleted by another transaction that commited after the initial
          read (`INSERT`, `DELETE` + `COMMIT`)
        - **Serialization anomaly**
    - **Isolation levels**
        - **Read committed**
        - **Repeatable read**
        - **Serializable**
- Random string of length

  ```sql
  CREATE FUNCTION random_string(length integer) RETURNS text LANGUAGE SQL AS $$
  SELECT string_agg(substring(
      'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789',
      (random() * 62 + 1)::integer, 1), '')
  FROM generate_series(1, length);
  $$;
  ```
