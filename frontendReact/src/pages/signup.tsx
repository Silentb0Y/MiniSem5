
import "../all.css"

export function Signup() {
    return (
           <div>
            <div className="container ">
                <input type="checkbox" id="check" />
                <div className="login form">
                    <h3>Signup</h3>
                    <form action="#">
                        <input type="text" placeholder="Enter your email" id="email1" />
                        <input
                            type="password"
                            placeholder="Enter your password"
                            id="password"
                        />
                        <input type="button" className="button" value="Login" />
                    </form>
                    <div className="signup">
                        <span className="signup"
                        >Already have an account?
                            <label>Login</label>
                        </span>
                    </div>
                </div>
            </div>
        </div> 
    )
}