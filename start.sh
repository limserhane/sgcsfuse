
if [ $1 -eq 0 ]
then
    echo "No password supplied"
    return -1
else
    password=$1
fi

if [ $2 -eq 0 ]
then
    echo "No key file supplied"
    return -1
else
    keyfile=$2
fi

if [ $3 -eq 0 ]
then
    echo "No bucket supplied"
    return -1
else
    bucket=$3
fi

if [ $4 -eq 0 ]
then
    echo "No bucket supplied"
    return -1
else
    mounteddir=$4
fi

source $SGCSFUSEPATH/blockchain/close.sh
setsid npm start --prefix $SGCSFUSEPATH/blockchain # >/dev/null

sleep 3s

umount fs/
if ! sgcsfuse --password $password --key-file $keyfile $bucket $mounteddir
then
	source $SGCSFUSEPATH/blockchain/close.sh
	return 1
fi


