# SSL and TLS
A certificate is used to guarantee trust between two devices between a transaction. TLS certificates ensures that communication between the user and the server are encrypted and provides server identity verification. 

Without secure connectivity, any credentials and data will be sent by the user to the server in plain text which can be intercepted by malicious third parties. Encryption keys are used to secure credentials and data by encrypting the data sent to the server. The data will then be unreadable for humans. The hacker intercepting this data cannot do anything with this data as they do not have the encryption key. The server will have the same issue and therefore, in this case, a copy of the key must also be sent to the server. If the hacker gets access to this, they are able to decrypt the encypted data. This is known as symmetric encryption. 

Asymmetric encryption is a more secure way of encrypting data and is less likely to be decrypted by hackers. Instead of using a single key to encrypt and decrypt data, a pair of keys are used, a private key and public key. The public key can be shared with others but the private key should remain a secret as this is what will "unlock" the public key. 

An example of asymmetric encrption is SSH keys. You wish to access a server but don't want to use passwords. You decide to generate SSH keys using the `ssh-keygen` command. It creates the private key in file `id_rsa` and public key in file `id.rsa.pub`. You then assign the public key to the server by adding the public key as an entry to the `~/.ssh/authorized_keys` file. As long as no one gains access to the private key, the server will be secure. You will now be able to ssh into the server using the command `ssh -i id_rsa user1@server1` command. You can create copies of the public key and place them on as many servers as you want. The same private key will be used to gain access to these servers.

If other users want access to the same servers, they can generate their own keypairs. They would then ask you to add their public key to the `~/.ssh/authorized_keys` file and they will then be able to access the servers. 

In the case of a bank website, asymmetric encyption can be achieved by generating private and public keys using the below commands;

`openssl genrsa -out my-bank.key 1024`

`openssl rsa -in my-bank.key -pubout > mybank.pem`

When the user accesses the web server using HTTPS, they get the public key from the server. The hacker is able to obtain the public key. The users browser is able to encrypt the symmetric key using the public key provided by the server. The user then sends this to the server. The server then uses the private key to decrypt the message and retrieve the symmetric key from it. The hacker is not able to do this as they don't have the private key to decrypt the symmetric key. Now the user and server only have access to the symmetric key. 

## SSL Certificates

Other ways that hackers can obtain your data is by creating a website similar to the one you are trying to access. To avoid having your data hacked, you should check the certificate of the website you are accessing. The certificate contains details such as the issuer and who the certificate is issued to. Most browsers are smart enough to display if the website has been validated, usually in the form of a green lock symbol.

There are two types of certificates, a self-signed certificate which is signed by the owner of the web application (not secure) and CA (certificate authority) signed certificate (more secure). CAs are authorities which are well know and trust worthy. They include organisations which include Symantec and Commodo. To get your certificates signed, you would need to generate a certificate signing request (CSR) using the key generated earlier and your website. You can use the below command

`openssl req -new -key my-bank.key -out my-bank.csr -subj "/C=US/ST=CA/O=MyOrg, Inc./CN=my-bank.com"`

The CAs then validate the details and if all checks out, they sign the certificate and send it back to you.

Browsers have built in public keys which are provided by the CAs. The private keys are used to sign the certificates. This is how the browsers know if the certificate is valid or not.

Most CAs also offer certificates for privately hosted sites such as one within an organisation. The private keys will be built in to the employees browsers to validate they are accessing the correct internal applications.