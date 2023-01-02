## Introduce
Open source [mtproto](https://core.telegram.org/mtproto) server implementation written in golang, support private deployment.

## Installing elloapp 
`elloapp` relies on open source high-performance components: 

- **mysql5.7**
- [redis](https://redis.io/)
- [etcd](https://etcd.io/)
- [kafka](https://kafka.apache.org/quickstart)
- [minio](https://docs.min.io/docs/minio-quickstart-guide.html#GNU/Linux)
- [ffmpeg](https://www.johnvansickle.com/ffmpeg/)

Privatization deployment Before `elloapp`, please make sure that the above five components have been installed. If your server does not have the above components, you must first install Missing components. 

- [Centos9 Stream Build and Install](docs/install-centos-9.md) [@A Feel]
- [CentOS7 elloapp_tg_backend](docs/install-centos-7.md) [@saeipi]

If you have the above components, it is recommended to use them directly. If not, it is recommended to use `docker-compose-env.yaml`.


### Source code deployment
#### Install [Go environment](https://go.dev/doc/install). Make sure Go version is at least 1.17.


#### Get source code　

```
git clone https://gitlab.com/merehead/elloapp/backend/elloapp_tg_backend.git
cd elloapp_tg_backend
```

#### Init data
- init database

	```
	1. create database elloapp
	2. init elloapp database
	   mysql -uroot elloapp < elloappd/sql/elloapp2.sql
	   mysql -uroot elloapp < elloappd/sql/migrate-*.sql
	```

- init minio buckets
	- bucket names
	  - `documents`
	  - `encryptedfiles`
	  - `photos`
	  - `videos`
	- Access `http://ip:xxxxx` and create


#### Build
	
```
make
```

#### Run

```
cd elloappd/bin
./runall2.sh
```

### Docker deployment
#### Install [Docker](https://docs.docker.com/get-docker/)

#### Install [Docker Compose](https://docs.docker.com/compose/install/)

#### Get source code

```
git clone https://github.com/elloapp/elloapp-server.git
cd elloapp_tg_backend
```

#### Install depends
- **change `192.168.1.150` to your ip in `docker-compose-env.yaml`**
- install depends

  ```
  # pull docker images
  docker-compose -f docker-compose-env.yaml pull
  
  # run docker-compose
  docker-compose -f docker-compose-env.yaml up -d
  ```
  
#### Init data
- init database

  ```

  # Copy some files to container
  docker cp ./elloappd/sql/ mysql:/elloappd/sql/

  # get mysql
  docker exec -it mysql /bin/bash
  
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/elloapp2.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20220321.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20220326.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20220328.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20220401.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20220412.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20220419.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20220423.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20220504.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20220721.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20220826.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20220919.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20221008.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20221011.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20221016.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20221023.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/migrate-20221101.sql
  mysql -uelloapp -h127.0.0.1 -pelloapp elloapp < elloappd/sql/init.sql
  
  # quit docker mysql
  exit
  ```

- init minio buckets
	- bucket names:
	  - `documents`
	  - `encryptedfiles`
	  - `photos`
	  - `videos`
	- create buckets
		
		```
		# get mc
		docker run -it --entrypoint=/bin/bash minio/mc
		   
		# change 192.168.1.150 to your ip    
		mc alias set minio http://192.168.1.150:9000 minio miniostorage
		
		# create buckets
		mc mb minio/documents
		mc mb minio/encryptedfiles
		mc mb minio/photos
		mc mb minio/videos
  
		# quit docker minio/mc
		exit
		```

#### Run

```  
# run docker-compose
docker-compose up -d
```
