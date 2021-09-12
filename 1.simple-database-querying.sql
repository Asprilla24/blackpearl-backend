SELECT u.ID, u.UserName, p.UserName AS ParentUserName
FROM USER as u 
  LEFT JOIN USER as p
    ON u.Parent = p.ID