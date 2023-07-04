# Toblerone
a simple encrypt and decrypt program written in Go<br>
has:<br>
keygen<br>
encrypt<br>
decrypt<br>

none of that pgp and age nonsense, just nice and easy for you simpletons
## Install
you really should know by now<br>
[click](https://github.com/keef1212/toblerone/releases/download/v1.0.0/toblerone.sh) that to install<br>
then, in order, run these lines of code where you just installed that:<br>
````
chmod +x toblerone.sh
./toblerone.sh
````
## Usage
```
keygen              gen a new encryption key
encrypt <input-file.txt> <output-file.tobl> <encryption-key>   encrypt a file
decrypt <input-file.tobl> <output-file.txt> <sender-decryption-key>   decrypt a file
Your sender-decryption-key should be the text string sent along with the file
sender-decryption-key is essentially the senders encryption key
```
