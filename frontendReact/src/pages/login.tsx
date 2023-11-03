
import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import "../all.css"

interface LoginPageProps {
    setLogIn: React.Dispatch<React.SetStateAction<boolean>>
    isLogin: boolean
}
export const Login: React.FC<LoginPageProps> = ({ setLogIn , isLogin}) => {
  const [data, setData] = useState("")

  const navigate = useNavigate();

  useEffect(() => {
    if (isLogin){
      navigate('/home')
    }
  },[isLogin, navigate])

  function sendLoginRequest() {
  fetch("http://localhost:9000/login", {
      method: "GET",
      headers: {
        "Content-Type": "application/json", // Set Content-Type header here
      },
      body: JSON.stringify({
        username: "yash",
        password: "yash",
        courseId: 1,
        course: {
          courseId: 1,
          courseMetaData: "s",
          courseName: "s",
          courseImageLink: "w",
          courseLink: "w",
          ytCourseLinks: ["", "", "", ""],
        },
        courseList: [],
      }),
    })
      .then((res: any) => {
        if (res.headers == 200) {
            setLogIn(true)
        }
      })
      .catch((error) => console.log("ERROR:", error));
  }
  function heeh(){
    setLogIn(true)
  }
    return (
        <div>
            <div className="container">
                <input type="checkbox" id="check" />
                <div className="login form">
                    <h3>Login</h3>
                    <form action="#">
                        <input type="text" placeholder="Enter your email" id="email1" />
                        <input
                            type="password"
                            placeholder="Enter your password"
                            id="password"
                        />
                        {/* <a href="#">Forgot password?</a> */}
                        <input type="button" className="button" value="Login" onClick={heeh}/>
                    </form>
                    <div className="signup">
                        <span className="signup"
                        >Don't have an account?
                            <label >Signup</label>
                        </span>
                    </div>
                </div>
            </div>
        </div>
    )
}