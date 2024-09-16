#!/bin/sh

export $(grep -v '^#' .env | xargs)

PROJECT=$(eval echo $PROJECT)
PROJECT_DATABASE_CONTAINER=$(eval echo $PROJECT_DATABASE_CONTAINER)
PROJECT_DATABASE_IMAGE=$(eval echo $PROJECT_DATABASE_IMAGE)
PROJECT_API_CONTAINER=$(eval echo $PROJECT_API_CONTAINER)
PROJECT_API_IMAGE=$(eval echo $PROJECT_API_IMAGE)
# PROJECT_APP_CONTAINER=$(eval echo $PROJECT_APP_CONTAINER)
PROJECT_APP_IMAGE=$(eval echo $PROJECT_APP_IMAGE)

# Check postgresql database ready or not
check_postgresql_for_doing() {
	# Maximum number of attempts to check PostgreSQL readiness
	max_attempts=2

	# Counter for tracking the number of successful checks
	success_count=0

	# Check PostgreSQL readiness with a loop
	while [ $success_count -lt $max_attempts ]; do
		if docker exec ${PROJECT_DATABASE_CONTAINER} pg_isready; then
			echo "PostgreSQL is ready"
			success_count=$((success_count + 1))
			echo "Successful checks: $success_count"
		else
			echo "Waiting for PostgreSQL to become ready..."
		fi
		sleep 2
	done

	# Run the test if PostgreSQL is ready for the specified number of times
	if [ $success_count -eq $max_attempts ]; then
		echo "PostgreSQL is ready for $max_attempts checks. Running the test..."

		# Do testing
		case "$1" in
			"documentation" )
				docker exec ${PROJECT_DATABASE_CONTAINER} mkdir /app/documentation/
				docker exec ${PROJECT_DATABASE_CONTAINER} java -jar /app/schemaspy/schemaspy-6.2.4.jar -t pgsql -dp /app/schemaspy/postgresql-42.7.3.jar -db planetv -host localhost -port 5432 -u postgres -p postgres -o /app/documentation

				# Check documentation files in container
				FILES=$(docker exec ${PROJECT_DATABASE_CONTAINER} ls /app/documentation)
				if [ -z "$FILES" ]; then
					echo "No files found in container documentation directory."
					exit 1
				else
					echo "Files found in container documentation directory."
				fi

				# Copy paste documentation to host
				docker cp ${PROJECT_DATABASE_CONTAINER}:/app/documentation ./documentation
				FILES=$(ls ./documentation)
				if [ -z "$FILES" ]; then
					echo "No files found in host documentation directory."
					exit 1
				else
					echo "Files found in host documentation directory."
				fi
			;;

			"test-service-blogcategory" )
				docker exec ${PROJECT_API_CONTAINER} go test -v \
				/api-fiber/services/blogcategory.go /api-fiber/services/blogcategory_test.go
			;;

			"test-service-blogfile" )
				docker exec ${PROJECT_API_CONTAINER} go test -v \
				/api-fiber/services/blogfile.go /api-fiber/services/blogfile_test.go
			;;

			"test-service-blogtag" )
				docker exec ${PROJECT_API_CONTAINER} go test -v \
				/api-fiber/services/blogtag.go /api-fiber/services/blogtag_test.go
			;;

			"test-service-blogtagfile" )
				docker exec ${PROJECT_API_CONTAINER} go test -v \
				/api-fiber/services/blogtagfile.go /api-fiber/services/blogtagfile_test.go
			;;
		esac
	else
		echo "PostgreSQL is not ready after $max_attempts attempts"
		exit 1
	fi
}


build_start_container() {
	docker compose up -d
}
stop_container() {
	docker compose stop
}
remove_container() {
	docker compose down
}
remove_image() {
	docker rmi ${PROJECT_DATABASE_IMAGE}
	docker rmi ${PROJECT_API_IMAGE}
	docker rmi ${PROJECT_APP_IMAGE}
}
remove_all() {
	remove_container
	remove_image
}

# All in one command
rebuild_all() {
	remove_all
	build_start_container
}

# Print list of options
print_list() {
	echo "Pass wrong arguments! Here is list of arguments for docker script"
	echo -e "\tbuild-start                   : build image and start container"
	echo -e "\tstop                          : stop container"
	echo -e "\trebuild-all                   : rebuild all"
	echo -e "\tremove-all                    : remove all (container, network, image)"
	echo -e "\tremove-container              : remove container"
	echo -e "\tremove-image                  : remove all image"
	echo -e "\ttest-api-service-auth         : test api service auth"
	echo -e "\ttest-api-service-blogcategory : test api service blogcategory"
	echo -e "\ttest-api-service-blogtag      : test api service blogtag"
	echo -e "\ttest-api-service-blogfile     : test api service blogfile"
	echo -e "\ttest-api-service-blogtagfile  : test api service blogtagfile"
	echo -e "\ttest-api-controller-auth      : test api controller auth"
}

# Main script
if [ $# -eq 1 ]; then
	case "$1" in
		"build-start" )
			build_start_container ;;
		"rebuild-all" )
			rebuild_all ;;
		"remove-all" )
			remove_all ;;
		"remove-container" )
			remove_container ;;
		"remove-image" )
			remove_image ;;
		"generate-database-doc" )
			check_postgresql_for_doing documentation ;;

		# Test service
		"test-api-service-auth" )
			docker exec ${PROJECT_API_CONTAINER} go test -v \
			/api-fiber/services/auth.go /api-fiber/services/auth_test.go ;;
		"test-api-service-blogcategory" )
			check_postgresql_for_doing test-service-blogcategory ;;
		"test-api-service-blogtag" )
			check_postgresql_for_doing test-service-blogtag ;;
		"test-api-service-blogfile" )
			check_postgresql_for_doing test-service-blogfile ;;
		"test-api-service-blogtagfile" )
			check_postgresql_for_doing test-service-blogtagfile ;;

		# Test controller
		"test-api-controller-auth" )
			docker exec ${PROJECT_API_CONTAINER} go test -v \
			/api-fiber/controllers/auth.go /api-fiber/controllers/auth_test.go ;;
		* )
			print_list ;;
	esac
else
	print_list
fi
