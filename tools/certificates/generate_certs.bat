REM use prebuilt binary for generating certificates as compiling is a problem on windows

if exist "ca-key.pem" (
    echo "ca-key.pem exists"
) else (
    echo "generating certs:"
    cfssl.exe gencert -initca ca-csr.json | cfssljson -bare ca
    cfssl.exe gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json ^
     -hostname=localhost,127.0.0.1,192.168.28.11,192.168.28.12,192.168.28.13,15.206.75.71 ^
      -profile=styx styxnode-csr.json | cfssljson.exe -bare styxnode   
)
