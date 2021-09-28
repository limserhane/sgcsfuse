
gcsfuseurl='https://github.com/limserhane/sgcsfuse/releases/download/v0.1.1/sgcsfuse'
blockchainurl='https://github.com/limserhane/js-blokchain/archive/refs/tags/v1.2.tar.gz'

tmpdir='/tmp/sgcsfuse'
installationdir="$HOME/sgcsfuse"

function check_installed {
	if ! type $1 > /dev/null; then
	    echo "$1 required"
	    return 1
	else
	    echo "$1 found"
	    return 0
	fi
}


#########################################
#########################################

# check curl installation
if ! check_installed curl; then
	return 1
fi

# check go installation
if ! check_installed go; then
	return 1
fi

# check npm installation
if ! check_installed npm; then
	return 1
fi

# create folders
if ! mkdir --parents $tmpdir
then
	echo "error seting up directories"
	return 2
fi
echo 'temporary directory created'
if [ -d "$installationdir" ]; then
	if ! rm -r "$installationdir"
	then
		echo "error seting up directories"
		return 2
	fi
fi
if ! mkdir --parents $installationdir
then
	echo "error seting up directories"
	return 2
fi
echo 'installation directory created'
	
# download sgcsfuse
echo -ne 'sgcsfuse fetching ... '
if ! curl $gcsfuseurl --output $installationdir/sgcsfuse --silent --location
then
	echo "error downloading"
	return 3
fi
if ! chmod +x $installationdir/sgcsfuse
then
	echo "error granting execution permissions"
	return 4
fi

echo 'complete'

# download blockchain
echo -ne 'blockchain fetching ... '
if ! curl $blockchainurl --output $tmpdir/sgcsfuse-blockchain.tar.gz --silent --location
then
	echo "error downloading"
	return 3
fi
if ! tar -xf $tmpdir/sgcsfuse-blockchain.tar.gz --directory $tmpdir
then
	echo "error untaring"
	return 5
fi
if ! mv $tmpdir/js-blokchain-1.2 $installationdir/blockchain
then
	echo "error moving"
	return 6
fi
echo 'complete'

# seting up blockchain
echo -ne 'blockchain dependencies installing ... '
if ! (cd $installationdir/blockchain && npm install --silent &>/dev/null)
then
	echo "error"
	return 7
fi
echo 'complete'


# set environment variables
echo 'updating environment variables'
echo "" >>~/.bashrc
echo "# sgcsfuse installation" >>~/.bashrc
echo "export SGCSFUSEPATH=$installationdir" >>~/.bashrc
echo "export PATH=\$PATH:\$SGCSFUSEPATH" >>~/.bashrc










