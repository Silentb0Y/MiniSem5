
import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import "../all.css"

interface LoginPageProps {
    setLogIn: React.Dispatch<React.SetStateAction<boolean>>
    isLogin: boolean
}
export const Login: React.FC<LoginPageProps> = ({ setLogIn , isLogin}) => {
  const [data, setData] = useState("")
  const [username, setUsername] = useState("")
  const [password, setpassword] = useState("")

  const navigate = useNavigate();

  useEffect(() => {
    if (isLogin){
      navigate('/home')
    }
  },[isLogin, navigate])

  function sendLoginRequest() {
    console.log("done")
  fetch("http://localhost:7001/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json", // Set Content-Type header here
      },
      body: JSON.stringify({
        username: username,
        password: password,
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
        console.log(res)
        if (res.ok) {
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
                        <input type="text" placeholder="Enter your email" id="email1" value={username} onChange={e => setUsername(e.target.value)} />
                        <input type="password" placeholder="Enter your password" id="password" value={password} onChange={e => setpassword(e.target.value)} />
                        <button className="btn btn-success bg-black" type="button" onClick={sendLoginRequest}>Login</button>
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