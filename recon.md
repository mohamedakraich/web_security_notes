## Reconnaissance Random Notes

- dig -t SOA z.hackycorp.com
- dig -t NS z.hackycorp.com
- dig AXFR int @z.hackycorp.com
- dig @z.hackycorp.com version.bind chaos txt

- [github/commit/path].patch to get email address of the one who commited
- git clone https://github.com/hackycorp/repo7.git
- cd repo7/
- git log
- tig
- git log --diff-filter=D --summary

- curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
- unzip awscliv2.zip
- sudo ./aws/install
- aws configure
- aws s3 cp s3://assets.hackycorp.com/key2.txt ./key2.txt

- robots.txt file
- 404 Error page
- .well-known/security.txt file
- Look for directory listing...
- /admin/

#### Directory Bruteforcing

- wfuzz, ffuf, patator
- docker run -it python /bin/bash
- mkdir /code
- cd code
- git clone https://github.com/xmendez/wfuzz.git
- cd wfuzz /
- python setup install
- wfuzz -c -z file,wordlist/general/common.txt --sc 200 http://hackycorp.com/FUZZ/
- wfuzz -c -z file,wordlist/general/common.txt --hc 404 http://hackycorp.com/FUZZ

- Default vhost => dig hackycorp.com
- curl http://51.158.147.132/ -v
- curl http://hackycorp.com/ -v -H "Host: test"

- Default vhost over TLS => curl https://51.158.147.132/ --insecure -v
- curl https://51.158.147.132/ --insecure -v -H "Host: test"

- Changing the Host header may allow you to get a different version of the website

#### Certificate SAN Alternative Names

- openssl s_client -connect hackycorp.com:443 </dev/null 2>/dev/null | openssl x509 -noout -text | grep DNS:
- echo | openssl s_client -connect hackycorp.com:443 | openssl x509 -noout -text | grep DNS:
- openssl x509 -noout -text -in MyCertificate.crt | grep DNS:

#### HEADER INSPECTION

- curl https://hackycorp.com/ --dump-header - -o /dev/null -s

#### Visual Reconnaissance

- for i in range(0,256): print("0x{:02x}".format(i)+".a.hackycorp.com")
- cat targets.txt | ./aquatone
- for i in `seq 0 255`; do printf "0x%02x.a.hackycorp.com\n" $i; done > dump.data
- for i in `cat hosts.txt` do curl $i/logo.png -o $i.png done
- open .

#### VIRTUAL HOST BRUTE FORCING

- docker run -it golang
- go install github.com/ffuf/ffuf@latest
- git clone https://github.com/ffuf/ffuf ; cd ffuf ; go get ; go build
- https://github.com/xmendez/wfuzz
- curl http://hackycorp.com -H "Host: random123.hackycorp.com"
- ffuf -w wfuzz/wordlist/general/common.txt -H "Host: FUZZ.hackycorp.com" -u http://hackycorp.com -fr recon_06 => Admin, admin, www
- curl -H "Host: admin.hackycorp.com"  http://hackycorp.com
- ffuf -w ~/wordlists/subdomains.txt -H "Host: FUZZ.hackycorp.com" -u http://hackycorp.com -fs 1495

#### LOAD BALANCING

- dig -t TXT key.z.hackycorp.com

#### Zone transfer

- dig axfr z.hackycorp.com
- dig -t SOA z.hackycorp.com
- dig axfr z.hackycorp.com @z.hackycorp.com





