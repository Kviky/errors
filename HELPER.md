## How to add new tag
Creating tags:

vscode:<br />
    1) choose source control on activity bar<br />
    2) commit all changes that were done<br />
    3) choose tab tags and press the plus button<br />
    4) when the tag is created, run command pallete and choose 'Git: Push (Follow Tags)'<br />
    
git:<br />
    1) commit all changes<br />
    2) exec cmd in the terminal or shell `git tag -a v1.4 -m 'my version 1.4'`<br />
    3) git push origin --tags

github:<br />
    1) choose Code tab<br />
    2) on the right side choose Releases<br />
    3) choose Draft new release<br />
    4) follow the instruction<br />