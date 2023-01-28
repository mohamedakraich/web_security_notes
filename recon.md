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