## LDAP Injection Random Notes

- Some LDAP servers authorise NULL Bind:This is an important check to perform on all login forms that you will test in the future, even if the backend is not LDAP-based.

- A boolean OR using |: (|(cn=[INPUT1])(cn=[INPUT2]))

- A boolean AND using &: (&(cn=[INPUT1])(userPassword=[INPUT2]))

- (&(cn=[INPUT1])(userPassword=HASH[INPUT2]))

- LDAP supports several formats: {CLEARTEXT}, {MD5}, {SMD5} (salted MD5), {SHA}, {SSHA} (salted SHA1), {CRYPT} for storing passwords.

- /?name=a*)(cn=*))%00&password=admin

- Testing for error: " ' )