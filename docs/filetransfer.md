
# FileTransfer

pkg/filetransfer

The filetransfer package uses the nodeconfig package (pkg/nodeconfig) to connect to other styxnodes and transfer files.

### Notes:

testdata commands:

touch -d 2020-03-03 testdata/conditions/time/2020.txt

touch -d 2019-03-03 testdata/conditions/time/2019.txt

touch -d 2018-03-03 testdata/conditions/time/2018.txt


styxnode            styxnode
                    electon app / web client / cli
linux srv    ->     windows srv
sftp                100
sz 1                sz 4
                    ^
config {glob path destsrv sourcesrv}
