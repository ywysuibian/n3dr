#!/bin/bash -e

NEXUS_VERSION="${1:-3.21.1}"
NEXUS_API_VERSION="${2:-v1}"
TOOL="${3:-./n3dr}"

readonly DOWNLOAD_LOCATION=/tmp/n3dr*

validate(){
    if [ -z "${TOOL}" ]; then
        echo "No deliverable defined. Assuming that 'go run main.go' 
should be run."
        TOOL="go run main.go"
    fi

    if [ -z "${NEXUS_VERSION}" ] || [ -z "${NEXUS_API_VERSION}" ]; then
        echo "NEXUS_VERSION and NEXUS_API_VERSION should be specified."
        exit 1
    fi
}

build(){
  source ./build.sh
}

nexus(){
    docker run --rm -d -p 9999:8081 --name nexus "sonatype/nexus3:${NEXUS_VERSION}"
}

readiness(){
    until docker logs nexus | grep 'Started Sonatype Nexus OSS'
    do
        echo "Nexus unavailable"
        sleep 3
    done
}

# Since nexus 3.17.0 the default 'admin123' was changed by an autogenerated
# one. This function retrieves the autogenerated password and if this file
# is unavailable, the default 'admin123' is returned.
password(){
    pw=$(docker exec -it nexus cat /nexus-data/admin.password || true)
    if [[ "${pw}" =~ [0-9a-z-]{36} ]]; then
        export PASSWORD=$pw
    else
        export PASSWORD="admin123"
    fi
}

artifact(){
    mkdir -p "maven-releases/some/group${1}/file${1}/1.0.0"
    echo someContent > "maven-releases/some/group${1}/file${1}/1.0.0/file${1}-1.0.0.jar"
    echo someContentZIP > "maven-releases/some/group${1}/file${1}/1.0.0/file${1}-1.0.0.zip"
    echo -e "<project>\n<modelVersion>4.0.0</modelVersion>\n<groupId>some.group${1}</groupId>\n<artifactId>file${1}</artifactId>\n<version>1.0.0</version>\n</project>" > "maven-releases/some/group${1}/file${1}/1.0.0/file${1}-1.0.0.pom"
}

files(){
    for a in $(seq 100); do artifact "${a}"; done
}

upload(){
    echo "Testing upload..."
    $TOOL upload -u admin -p $PASSWORD -r maven-releases -n http://localhost:9999 -v "${NEXUS_API_VERSION}"
    echo
}

backupHelper(){
    if [ "${NEXUS_VERSION}" == "3.9.0" ]; then
        count_downloads 300
        test_zip 148
    else
    ls $DOWNLOAD_LOCATION
        count_downloads 400
        test_zip 192
    fi

    cleanup_downloads
}

anonymous(){
    echo "Testing backup by anonymous user..."
    $TOOL backup -n http://localhost:9999 -r maven-releases -v "${NEXUS_API_VERSION}" -z --anonymous
    backupHelper
}

backup(){
    echo "Testing backup..."
    $TOOL backup -n http://localhost:9999 -u admin -p $PASSWORD -r maven-releases -v "${NEXUS_API_VERSION}" -z
    backupHelper
}

regex(){
    echo "Testing backup regex..."
    $TOOL backup -n http://localhost:9999 -u admin -p $PASSWORD -r maven-releases -v "${NEXUS_API_VERSION}" -x 'some/group42' -z
    if [ "${NEXUS_VERSION}" == "3.9.0" ]; then
        count_downloads 3
        test_zip 4
    else
        count_downloads 4
        test_zip 4
    fi
    cleanup_downloads

    echo -e "\nTesting repositories regex..."
    $TOOL repositories -n http://localhost:9999 -u admin -p $PASSWORD -v "${NEXUS_API_VERSION}" -b -x 'some/group42' -z
    if [ "${NEXUS_VERSION}" == "3.9.0" ]; then
        count_downloads 3
        test_zip 4
    else
        count_downloads 4
        test_zip 4
    fi
    cleanup_downloads
}

repositories(){
    local cmd="$TOOL repositories -n http://localhost:9999 -u admin -p $PASSWORD -v ${NEXUS_API_VERSION}"

    echo "Testing repositories..."
    $cmd -a | grep maven-releases
    $cmd -c | grep 5
    $cmd -b -z

    if [ "${NEXUS_VERSION}" == "3.9.0" ]; then
        count_downloads 300
        test_zip 148
    else
        count_downloads 400
        test_zip 192
    fi

    cleanup_downloads
}

zipName(){
    echo "Testing zipName..."
    $TOOL backup -n=http://localhost:9999 -u=admin -p=$PASSWORD -r=maven-releases -v="${NEXUS_API_VERSION}" -z -i=helloZipFile.zip
    $TOOL repositories -n http://localhost:9999 -u admin -p $PASSWORD -v ${NEXUS_API_VERSION} -b -z -i=helloZipRepositoriesFile.zip
    ls helloZip* | wc -l | grep 2
}

cleanup(){
    cleanup_downloads
    docker stop nexus
}

count_downloads(){
    local actual
    actual=$(find $DOWNLOAD_LOCATION -type f | wc -l)
    echo "Expected: ${1}"
    echo "Actual: ${actual}"
    echo "${actual}" | grep "${1}"
}

test_zip(){
    local size
    size=$(du n3dr-backup-*zip)
    echo "Actual: ${size}"
    echo "Expected: ${1}"
    echo "${size}" | grep "^${1}"
}

cleanup_downloads(){
    rm -rf maven-releases
    rm -rf $DOWNLOAD_LOCATION
    rm -f n3dr-backup-*zip
    rm -f helloZip*zip
}

main(){
    validate
    build
    nexus
    readiness
    password
    files
    upload
    anonymous
    backup
    repositories
    regex
    zipName
    bats --tap tests.bats
}

trap cleanup EXIT
main
