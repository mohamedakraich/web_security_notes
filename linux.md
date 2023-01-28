## Linux Commands Random Notes

- find /home -name .bash_history

- find /home -name .bashrc -exec grep [PATTERN] {} \;

- find /home -name .bashrc -exec grep 'check' {} \;

- find /home -name .bashrc -exec grep -H 'check' {} \;

- .bash_profile

- echo admin | md5sum

- find /home -name .bashrc -exec grep 'export' {} \; | grep -v GCC_COLORS

- .zsh_history

- find /home -name .\*sh_history

- find /home -name .zsh_history

- grep passwd .bash_history
- grep -A 1 passwd .bash_history
- find /home -name .bash_history -exec grep -A 1 passwd {} \;
- find . -name .bash_history -exec grep -A 1 '^passwd' {} \;

- ls -ld /tmp
- file archive.tgz
- tar -xvf archive.tar.gz

- file /var/tmp/backup.tbz 
- tar -jxvf /var/tmp/backup.tbz

- rm -r home/

- bunzip2 backup.bz2
- cpio -t < backup
- cpio -idv --no-absolute-filename < backup
- strings backup

- ls -al /etc/cron.daily/
- openssl enc -d -aes256 -k P3NT35T3RL48 -in /tmp/backup.tgz.enc  -out /tmp/backup.tgz