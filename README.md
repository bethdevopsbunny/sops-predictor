# sops-predictor
sops doesnt hide the length of the encrypted fields very well, I noticed this and found that mozilla are aware of this [bug](https://github.com/mozilla/sops/issues/815),
but that doesnt mean i cant have fun with it! 

for all data types it can return you the legnth of the encrypted value.

for boolean however, as it can only be true or false having the length of the value allows you to assume
its encrypted value as true has 4 characters and false has 5.

after re-reading the bug linked above, it shows that  



here is the age encrypted file in sopsdata/secrets.enc.yml being predicted. 
![image](https://user-images.githubusercontent.com/80027170/196878553-e9782c52-c507-4b59-90f8-277ff002b661.png)

