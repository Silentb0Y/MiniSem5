apitest:
	curl -X POST \
		-H "Content-Type: application/json" \
		-d '{ \
			"username":"user123", \
			"password":"password123", \
			"courseId":1, \
			"courseName":"Sample Course", \
		 	"courseMetaData":"Sample Metadata", \
		 	"courseVideoLinks":["video1.mp4", "video2.mp4"]}' \
		http://localhost:8088/
userLogin2:
	curl -X POST \
		-H "Content-Type: application/json" \
		-d '{ \
				"username": "yash",\
				"password": "yash",\
				"courseId": 1,\
				"course": {\
					"courseId": 1,\
					"courseMetaData": "s",\
					"courseName": "s",\
					"courseImageLink": "w",\
					"courseLink": "w",\
					"ytCourseLinks": ["", "", "", ""]\
				},
				"courseList": []
			}' \
		http://localhost:8088/login

userLogin:
	curl -X POST \
		-H "Content-Type: application/json" \
		-d '{ \
			"username":"yash", \
			"password":"yash", \
			"courseId":1, \
			"course":{\
				"courseId": 1,\
				"courseMetaData":"s", \
				"courseName":"s",\
				"courseImageLink":"w",\
				"courseLink": "w",\
				"ytCourseLink":["","","",""],\
			}, \
			"courseList"[] \ 
		}' \
		http://localhost:8088/login

getCourses:
	curl -X GET \
		-H "Content-Type: application/json" \
		http://localhost:8098/getCourses
