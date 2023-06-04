build:
	echo "Building HelloAPI..."
	cd ./api/hello && make outputdir=${workdir}/bin
	cd ../../
	echo "Build Success HelloAPI..."