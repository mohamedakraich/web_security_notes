## Random Notes On App Analysis

##### Testing Layers

- Open Ports and Services.
- Web hosting software.
- Application framework.
- Application: Custom code or COTS.
- Application libraries (usually Javascript).
- Integrations.

##### Tech Profiling

- Browser extensions:  Whatruns,  Wappalyzer
- CMD: webanalyze -host http://example.com -crawl 2

##### Finding CVE's and misconfigs

- Nessus (Enterprise level)
- Nuclei
- Others: Jaeles Scanner, Gofingerprint, Sn1per, Vulners, intrigue Core, Retire.js

##### Port Scanning

- naabu -host hackerone.com

### Content Discovery

###### Tools

- turbo intruder
- gobuster
- ffuf
- feroxbuster
- dirsearch
- wfuzz

###### Lists

- wordlists.assetnote.com
- Source2Url
- Burp Suite extension Scavenger
- Historical: echo bugcrowd.com | gau | wordlistgen | sort -u

###### Technology TIPS

- Pay special attention to config files for DB connections
- Pay special attention to where the admin login and routes/endpoints are 

### Application Analysis

###### Big Questions

1. How does the app pass data?
2. How/where does the app talk about users?
3. Does the site have multi-tenancy or user levels?
4. Does the site have a unique threat model?
5. Has there been past security research & vulns?
6. How does the app handle: XSS, CSRF, Code Injection (SQL, Template, ++) ?

###### Spidering

- Tools: ZAP and BURP suite
- CMD: hakrawler, gospider

###### Javascript Analysis

- Linkfinder
- xnLinkFinder
- Javascript parsing in Burp => GAP extension
- Look at Javascript libs and dependencies

###### Hot Areas (Heat Mapping)

- Upload functions
- Content Types
- APIs
- Profile Section
- Errors

###### Parameter Analysis

- https://github.com/bugcrowd/HUNT
- https://github.com/1ndianl33t/Gf-Patterns => cat urls.txt | gf sqli