OpenSSL Cheatsheet

# Commands


## Print x509 Certificate Infos:

<details><summary>show</summary>
<p>

        openssl x509 -text -in my.pem
        openssl ca -text -in my_ca.pem
        openssl req -text -in csr.pem

</p>
</details>

## Checking Files

<details><summary>show</summary>
<p>


        openssl req -text -noout -verify -in csr.pem
        openssl rsa -in my.key -check
        openssl pkcs12 -info -in keystore.p12

</p>
</details>

## Check for expiration

<details><summary>show</summary>
<p>

        openssl x509 -enddate -noout -in file.pem            # prints something like 'notAfter=Nov  3 22:23:50 2014 GMT'
        openssl x509 -checkend 86400 -noout -in file.pem     # gives exitcode 0 if not expired

</p>
</details>

## Testing SSL webserver

<details><summary>show</summary>
<p>


        openssl s_client -connect example.com:443

        # With advanced TLS and OSCP debugging:
        openssl s_client -connect example.com:443 -tls1 -tlsextdebug -status

</p>
</details>

## OpenSSL Version and Certificates directory:

<details><summary>show</summary>
<p>

        openssl version -a

</p>
</details>

## Rehash OpenSSL certificates

<details><summary>show</summary>
<p>


        c_rehash <directory>

</p>
</details>

## Verifying certificates

<details><summary>show</summary>
<p>


        Certificate: openssl x509 -noout -modulus -in server.crt | openssl md5
        Private Key: openssl rsa -noout -modulus -in server.key | openssl md5
        CSR: openssl req -noout -modulus -in server.csr | openssl md5

</p>
</details>

## Stripping password from private keys

<details><summary>show</summary>
<p>

        openssl rsa -in key-with-pwd.pem -out key-without-pwd.pem

</p>
</details>


