# SGCSFUSE 

## Installation

The project can be installed by running the `install.sh` shell script. It does not require root privileges.

The installation script will do the following :

1. Ensure the required packages are installed. If not, the installation will stop and the user will be notified.

2. Create the folder where SGCSFUSE will be installed. It is located at : `~/sgcsfuse`

3. Fetch the project and set it up

4. Set the environment variables

## Usage

Once GCSFUSE is installed (by default at `~/sgcsfuse`), it can run as a command since its location is included in the `PATH`.

SGCSFUSE needs several arguments :

1. The password used to encrypt the files stored in the cloud

2. The absolute path to your GCS service account key file.

    To get one, use the GCS console. In the *IAM & Admin* menu, reach to the *Service Accouts* tab. Create a service account, and access it by clicking on its e-mail address. Inside the *Keys* tab, you can create a JSON key, which will be used by the program to access the GCS storage. Make sure you download it when you are proposed to.

3. The name of the bucket synced to the mounted folder.

4. The path to the directory you want to mount the file system on.

To start the project, run `start.sh` with the above arguments. For example :

```
source start.sh "thisismy_P@$sPHras3!" "~/gcskeyfile.json" "my-bucket" "synceddir/"
```

This will start the blockchain and then start SGCS with the provided arguments.