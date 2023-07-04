# Toblerone
an on-the-fly encryption program written in Go<br>

## Install (Unix based operating systems)
[click](https://github.com/keef1212/toblerone/releases/download/v1.0.0/toblerone.sh) that to install<br>
then, in order, run these lines of code where you just installed that:<br>
````
chmod +x toblerone.sh
./toblerone.sh
````
### Compile from source
#### Unix based systems
```
git clone https://github.com/keef1212/toblerone.git
cd toblerone
go build toblerone.go
chmod +x toblerone
mv toblerone /usr/local/bin
```
#### Windows
```
git clone https://github.com/keef1212/toblerone.git
cd toblerone
go build toblerone.go
mv toblerone C:\Windows\System32
```
## Usage
```
keygen              generate a new encryption key
encrypt <input-file.txt> <output-file.tobl> <encryption-key>   encrypt a file
decrypt <input-file.tobl> <output-file.txt> <sender-decryption-key>   decrypt a file
```
Your sender-decryption-key should be the text string sent along with the file<br>
sender-decryption-key is essentially the senders encryption key

